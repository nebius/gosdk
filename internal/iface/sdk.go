package iface

import (
	"github.com/nebius/gosdk/conn"
)

type SDK interface {
	conn.Resolver
	conn.Dialer
}

type SDKWithParentID interface {
	SDK
	ParentID() string
}

var _ SDKWithParentID = (*SDKWithNoParentID)(nil)

type SDKWithNoParentID struct {
	SDK
}

func (s *SDKWithNoParentID) ParentID() string {
	return ""
}

func WrapSDK(sdk SDK) SDKWithParentID {
	if withParentID, ok := sdk.(SDKWithParentID); ok {
		return withParentID
	}
	return &SDKWithNoParentID{
		SDK: sdk,
	}
}
