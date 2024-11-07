package nebius

import "github.com/nebius/gosdk/services/nebius/registry"

func (s Services) Registry() registry.Services {
	return registry.New(s.sdk)
}
