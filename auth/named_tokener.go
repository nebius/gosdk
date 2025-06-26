package auth

import "context"

type NamedTokener interface {
	BearerTokener
	Name() string
}

type Wrapper interface {
	Unwrap() BearerTokener
}

// NameWrapper is a wrapper around a [BearerTokener] that adds a name to the tokener
// if the tokener does not have a name.
type NameWrapper struct {
	name    string
	tokener BearerTokener
}

var _ NamedTokener = (*NameWrapper)(nil)

func NewNameWrapper(name string, tokener BearerTokener) *NameWrapper {
	return &NameWrapper{
		name:    name,
		tokener: tokener,
	}
}

func (n *NameWrapper) Name() string {
	return n.name
}

func (n *NameWrapper) BearerToken(ctx context.Context) (BearerToken, error) {
	return n.tokener.BearerToken(ctx)
}

func (n *NameWrapper) HandleError(ctx context.Context, token BearerToken, err error) error {
	return n.tokener.HandleError(ctx, token, err)
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
