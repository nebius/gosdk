package conn

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"
)

// ConventionResolverServiceIDToNameMap is a global registry for api_service_name annotations.
// It is filled in by init functions inside services directory.
var ConventionResolverServiceIDToNameMap = map[ServiceID]string{} //nolint:gochecknoglobals // ok

type resolverContextKey struct{}

func ContextWithResolver(ctx context.Context, resolver Resolver) context.Context {
	return context.WithValue(ctx, resolverContextKey{}, resolver)
}

func ResolverFromContext(ctx context.Context) Resolver {
	res, ok := ctx.Value(resolverContextKey{}).(Resolver)
	if !ok {
		return nil
	}
	return res
}

type ContextResolver struct {
	logger *slog.Logger
	next   Resolver
}

var _ Resolver = ContextResolver{}

func NewContextResolver(logger *slog.Logger, next Resolver) ContextResolver {
	return ContextResolver{
		logger: logger,
		next:   next,
	}
}

func (r ContextResolver) Resolve(ctx context.Context, id ServiceID) (Address, error) {
	resolver := ResolverFromContext(ctx)
	if resolver == nil {
		return r.next.Resolve(ctx, id)
	}
	return NewLoggingResolver(r.logger, resolver).Resolve(ctx, id)
}

type loggingResolver struct {
	logger *slog.Logger
	next   Resolver
}

var _ Resolver = loggingResolver{}

type resolverLogAttributer interface {
	resolverLogAttributes(ServiceID, Address) []any
}

// NewLoggingResolver logs the concrete resolver that successfully resolves a service.
// Resolver chains are decorated recursively so the matching resolver is logged instead of the chain.
func NewLoggingResolver(logger *slog.Logger, next Resolver) Resolver {
	switch resolver := next.(type) {
	case loggingResolver:
		return resolver
	case ResolversChain:
		logged := make([]Resolver, 0, len(resolver))
		for _, child := range resolver {
			logged = append(logged, NewLoggingResolver(logger, child))
		}
		return NewResolversChain(logged...)
	default:
		return loggingResolver{logger: logger, next: next}
	}
}

func (r loggingResolver) Resolve(ctx context.Context, id ServiceID) (Address, error) {
	address, err := r.next.Resolve(ctx, id)
	if err != nil {
		return "", err
	}
	if !r.logger.Enabled(ctx, slog.LevelDebug) {
		return address, nil
	}
	attrs := []any{
		slog.String("service", string(id)),
		slog.String("address", string(address)),
		slog.String("resolver_type", fmt.Sprintf("%T", r.next)),
	}
	if attributer, ok := r.next.(resolverLogAttributer); ok {
		attrs = append(attrs, attributer.resolverLogAttributes(id, address)...)
	}
	r.logger.DebugContext(ctx, "service address is resolved by resolver", attrs...)
	return address, nil
}

func NewBasicResolver(id, address string) Resolver {
	if prefix, isWildcard := strings.CutSuffix(id, "*"); isWildcard {
		return NewPrefixResolver(prefix, Address(address))
	}
	return NewSingleResolver(ServiceID(id), Address(address))
}

type ConventionResolver struct{}

var _ Resolver = ConventionResolver{}

func NewConventionResolver() ConventionResolver {
	return ConventionResolver{}
}

func (ConventionResolver) Resolve(_ context.Context, id ServiceID) (Address, error) {
	serviceName, err := conventionServiceName(id)
	if err != nil {
		return "", err
	}
	return Address(serviceName + ".{domain}"), nil
}

func (ConventionResolver) resolverLogAttributes(ServiceID, Address) []any {
	return []any{slog.String("route", "gateway")}
}

func conventionServiceName(id ServiceID) (string, error) {
	parts := strings.Split(string(id), ".")
	if len(parts) < 3 || parts[0] != "nebius" {
		return "", NewUnknownServiceError(id)
	}
	serviceName := parts[1]
	if name, ok := ConventionResolverServiceIDToNameMap[id]; ok {
		serviceName = name
	}
	return serviceName, nil
}

type ConstantResolver string

var _ Resolver = ConstantResolver("")

func NewConstantResolver(address Address) ConstantResolver {
	return ConstantResolver(address)
}

func (r ConstantResolver) Resolve(context.Context, ServiceID) (Address, error) {
	return Address(r), nil
}

func (ConstantResolver) resolverLogAttributes(ServiceID, Address) []any {
	return []any{slog.String("route", "override")}
}

type SingleResolver struct {
	id      ServiceID
	address Address
}

var _ Resolver = SingleResolver{}

func NewSingleResolver(id ServiceID, address Address) SingleResolver {
	return SingleResolver{
		id:      id,
		address: address,
	}
}

func (r SingleResolver) Resolve(_ context.Context, id ServiceID) (Address, error) {
	if id == r.id {
		return r.address, nil
	}
	return "", NewUnknownServiceError(id)
}

func (SingleResolver) resolverLogAttributes(ServiceID, Address) []any {
	return []any{slog.String("route", "override")}
}

type PrefixResolver struct {
	prefix  string
	address Address
}

var _ Resolver = PrefixResolver{}

func NewPrefixResolver(prefix string, address Address) PrefixResolver {
	return PrefixResolver{
		prefix:  prefix,
		address: address,
	}
}

func (r PrefixResolver) Resolve(_ context.Context, id ServiceID) (Address, error) {
	if strings.HasPrefix(string(id), r.prefix) {
		return r.address, nil
	}
	return "", NewUnknownServiceError(id)
}

func (r PrefixResolver) resolverLogAttributes(ServiceID, Address) []any {
	return []any{
		slog.String("route", "override"),
		slog.String("resolver_prefix", r.prefix),
	}
}

type ResolversChain []Resolver

var _ Resolver = ResolversChain{}

func NewResolversChain(resolvers ...Resolver) ResolversChain {
	res := make(ResolversChain, 0, len(resolvers))
	for _, resolver := range resolvers {
		switch r := resolver.(type) {
		case ResolversChain:
			res = append(res, r...)
		default:
			res = append(res, r)
		}
	}
	return res
}

func (r ResolversChain) Resolve(ctx context.Context, id ServiceID) (Address, error) {
	var errs error
	for _, resolver := range r {
		address, err := resolver.Resolve(ctx, id)
		if err == nil {
			return address, nil
		}
		unknownService := &UnknownServiceError{}
		if errors.As(err, &unknownService) && unknownService.ID == id {
			continue
		}
		errs = errors.Join(errs, err)
	}
	if errs != nil {
		return "", errs
	}
	return "", NewUnknownServiceError(id)
}

type CachedResolver struct {
	logger *slog.Logger
	next   Resolver
	cache  sync.Map
}

var _ Resolver = &CachedResolver{}

func NewCachedResolver(logger *slog.Logger, next Resolver) *CachedResolver {
	return &CachedResolver{
		logger: logger,
		next:   next,
		cache:  sync.Map{},
	}
}

func (r *CachedResolver) Resolve(ctx context.Context, id ServiceID) (Address, error) {
	if value, ok := r.cache.Load(id); ok {
		return value.(Address), nil //nolint:errcheck // ok to panic
	}

	log := r.logger.With(slog.String("service", string(id)))
	log.DebugContext(ctx, "resolving service address")
	address, err := r.next.Resolve(ctx, id)
	if err != nil {
		return "", err
	}
	log.DebugContext(ctx, "service address is resolved", slog.String("address", string(address)))

	r.cache.Store(id, address)
	return address, nil
}

type TemplateExpander struct {
	next          Resolver
	substitutions map[string]string
}

var _ Resolver = TemplateExpander{}

func NewTemplateExpander(substitutions map[string]string, next Resolver) TemplateExpander {
	return TemplateExpander{
		next:          next,
		substitutions: substitutions,
	}
}

func (r TemplateExpander) Resolve(ctx context.Context, id ServiceID) (Address, error) {
	address, err := r.next.Resolve(ctx, id)
	if err != nil {
		return "", err
	}
	for find, replace := range r.substitutions {
		address = Address(strings.ReplaceAll(string(address), find, replace))
	}
	return address, nil
}
