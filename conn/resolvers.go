package conn

import (
	"context"
	"errors"
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
		resolver = r.next
	} else {
		r.logger.DebugContext(ctx, "using custom service address resolver", slog.String("service", string(id)))
	}
	return resolver.Resolve(ctx, id)
}

func NewBasicResolver(id, address string) Resolver {
	if strings.HasSuffix(id, "*") {
		return NewPrefixResolver(strings.TrimSuffix(id, "*"), Address(address))
	}
	return NewSingleResolver(ServiceID(id), Address(address))
}

type ConventionResolver struct{}

var _ Resolver = ConventionResolver{}

func NewConventionResolver() ConventionResolver {
	return ConventionResolver{}
}

func (ConventionResolver) Resolve(_ context.Context, id ServiceID) (Address, error) {
	parts := strings.Split(string(id), ".")
	if len(parts) < 3 || parts[0] != "nebius" || !strings.HasSuffix(parts[len(parts)-1], "Service") {
		return "", NewUnknownServiceError(id)
	}
	serviceName := parts[1]
	if name, ok := ConventionResolverServiceIDToNameMap[id]; ok {
		serviceName = name
	}
	return Address(serviceName + ".{domain}"), nil
}

type ConstantResolver string

var _ Resolver = ConstantResolver("")

func NewConstantResolver(address Address) ConstantResolver {
	return ConstantResolver(address)
}

func (r ConstantResolver) Resolve(context.Context, ServiceID) (Address, error) {
	return Address(r), nil
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
