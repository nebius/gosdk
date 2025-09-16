package config

import (
	"context"

	"github.com/nebius/gosdk/auth"
)

type Option interface {
	option()
}

type ConfigInterface interface {
	ParentID() string
	Endpoint() string
	GetCredentials(context.Context) (auth.BearerTokener, error)
	GetAuthType() AuthType
	SetOptions(options ...Option)
	GetConfig() *Config
	GetConfigPath() (string, error)
	Load(context.Context) error
	CurrentProfileName() string
}

type StubOption struct{}

func (StubOption) option() {}
