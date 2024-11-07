package registry

import "github.com/nebius/gosdk/services/nebius/registry/v1"

func (s Services) V1() v1.Services {
	return v1.New(s.sdk)
}
