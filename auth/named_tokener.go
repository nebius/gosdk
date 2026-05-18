package auth

import "context"

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
