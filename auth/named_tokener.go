package auth

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/nebius/gosdk/internal/logging"
)

type NamedTokener interface {
	BearerTokener
	Name() string
}

type TypedTokener interface {
	Type() string
}

type Wrapper interface {
	Unwrap() BearerTokener
}

// NameWrapper is a wrapper around a [BearerTokener] that adds a name to the tokener
// if the tokener does not have a name.
//
// Formatting uses %v for the wrapped tokener.
// Structured logging uses LogValue, then String, or only the wrapped type.
// The wrapped type should sanitize its own sensitive values for both formatting and structured logging.
type NameWrapper struct {
	name    string
	typ     string
	tokener BearerTokener
}

var _ NamedTokener = (*NameWrapper)(nil)
var _ TypedTokener = (*NameWrapper)(nil)
var _ MetricsSetter = (*NameWrapper)(nil)

func NewNameWrapper(name string, tokener BearerTokener) *NameWrapper {
	return NewTypedNameWrapper(name, "", tokener)
}

func NewTypedNameWrapper(name string, typ string, tokener BearerTokener) *NameWrapper {
	return &NameWrapper{
		name:    name,
		typ:     typ,
		tokener: tokener,
	}
}

func (n *NameWrapper) LogValue() slog.Value {
	if n == nil {
		return slog.AnyValue(nil)
	}
	return slog.GroupValue(
		slog.String("type", "NameWrapper"),
		slog.String("name", n.name),
		slog.String("tokener_type", n.Type()),
		logging.Object("tokener", n.tokener),
	)
}

func (n *NameWrapper) String() string {
	return fmt.Sprintf("NameWrapper(name=%q, type=%q, tokener=%v)", n.name, n.Type(), n.tokener)
}

func (n *NameWrapper) Name() string {
	return n.name
}

func (n *NameWrapper) Type() string {
	if n.typ != "" {
		return n.typ
	}
	if typ, ok := TypeOfTokener(n.tokener); ok {
		return typ
	}
	return "custom"
}

func (n *NameWrapper) BearerToken(ctx context.Context) (BearerToken, error) {
	return n.tokener.BearerToken(ctx)
}

func (n *NameWrapper) HandleError(ctx context.Context, token BearerToken, err error) error {
	return n.tokener.HandleError(ctx, token, err)
}

func (n *NameWrapper) SetMetrics(metrics Metrics) {
	if setter, ok := n.tokener.(MetricsSetter); ok {
		setter.SetMetrics(metrics)
	}
}

func (n *NameWrapper) Unwrap() BearerTokener {
	return n.tokener
}

func NameOfTokener(tokener BearerTokener) (string, bool) {
	if namedTokener, ok := tokener.(NamedTokener); ok {
		return namedTokener.Name(), true
	}
	if wrapper, ok := tokener.(Wrapper); ok {
		return NameOfTokener(wrapper.Unwrap())
	}
	return "", false
}

func TypeOfTokener(tokener BearerTokener) (string, bool) {
	if typedTokener, ok := tokener.(TypedTokener); ok {
		return typedTokener.Type(), true
	}
	if wrapper, ok := tokener.(Wrapper); ok {
		return TypeOfTokener(wrapper.Unwrap())
	}
	return "", false
}
