package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nebius/gosdk/config/paths"
)

// fileTokener is a [BearerTokener] that serves token from file, reading it every time.
type fileTokener string

var _ BearerTokener = (*fileTokener)(nil)

// newFileTokener returns a [BearerTokener] that serves token from file, reading it every time.
func newFileTokener(path string) (*fileTokener, error) {
	path, err := paths.ExpandHomeDir(path)
	if err != nil {
		return nil, fmt.Errorf("expand home: %w", err)
	}
	t := fileTokener(path)
	return &t, nil
}

func (t *fileTokener) BearerToken(context.Context) (BearerToken, error) {
	file := string(*t)
	data, err := os.ReadFile(file)
	if err != nil {
		return BearerToken{}, fmt.Errorf("read file: %w", err)
	}
	token := strings.TrimSpace(string(data))
	if strings.Contains(token, "\n") {
		return BearerToken{}, fmt.Errorf("invalid token file: %s contains newline", file)
	}
	return BearerToken{
		Token:     token,
		ExpiresAt: time.Time{},
	}, nil
}

func (t *fileTokener) HandleError(ctx context.Context, token BearerToken, err error) error {
	rt, tokErr := t.BearerToken(ctx)
	if tokErr != nil {
		return errors.Join(tokErr, err)
	}
	if rt.Token != token.Token {
		return nil // token has changed
	}

	return err // Same token, but it is invalid, fail
}

type FileTokener struct {
	tokener       *fileTokener
	token         *BearerToken
	refreshPeriod time.Duration
	expiration    time.Time
}

var _ BearerTokener = (*FileTokener)(nil)

func NewFileTokener(
	path string,
	refreshPeriod time.Duration,
) (*FileTokener, error) {
	tokener, err := newFileTokener(path)
	if err != nil {
		return nil, err
	}
	return &FileTokener{
		tokener:       tokener,
		refreshPeriod: refreshPeriod,
	}, nil
}

func (c *FileTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	if c.token == nil || (c.refreshPeriod != 0 && time.Now().After(c.expiration)) {
		token, err := c.tokener.BearerToken(ctx)
		if err != nil {
			return BearerToken{}, err
		}
		c.token = &token
		c.expiration = time.Now().Add(c.refreshPeriod)
	}
	return *c.token, nil
}

func (c *FileTokener) HandleError(ctx context.Context, token BearerToken, err error) error {
	c.token = nil // force refresh on next call
	c.expiration = time.Time{}
	return c.tokener.HandleError(ctx, token, err)
}
