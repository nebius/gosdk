package iface

import (
	"log/slog"

	"github.com/nebius/gosdk/conn"
)

type SDK interface {
	conn.Resolver
	conn.Dialer
}

type WithLogger interface {
	GetLogger() *slog.Logger
}

type SDKWithParentID interface {
	SDK
	WithLogger
	ParentID() string
	TenantID() string
}

var _ SDKWithParentID = (*SDKWithNoParentID)(nil)

type SDKWithNoParentID struct {
	SDK
	logger *slog.Logger
}

func (s *SDKWithNoParentID) ParentID() string {
	return ""
}
func (s *SDKWithNoParentID) TenantID() string {
	return ""
}

func (s *SDKWithNoParentID) GetLogger() *slog.Logger {
	return s.logger
}

func WrapSDK(sdk SDK) SDKWithParentID {
	if withParentID, ok := sdk.(SDKWithParentID); ok {
		return withParentID
	}
	var logger *slog.Logger
	if withLogger, ok := sdk.(WithLogger); ok {
		logger = withLogger.GetLogger()
	}
	return &SDKWithNoParentID{
		SDK:    sdk,
		logger: logger,
	}
}
