// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package billing

import (
	"github.com/nebius/gosdk/internal/iface"
)

type Services struct {
	sdk iface.SDK
}

func New(sdk iface.SDK) Services {
	return Services{
		sdk: sdk,
	}
}
