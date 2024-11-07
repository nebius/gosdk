package iface

import (
	"github.com/nebius/gosdk/conn"
)

type SDK interface {
	conn.Resolver
	conn.Dialer
}
