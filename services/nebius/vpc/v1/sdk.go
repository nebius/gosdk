package v1

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
