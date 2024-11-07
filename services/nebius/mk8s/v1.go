package mk8s

import "github.com/nebius/gosdk/services/nebius/mk8s/v1"

func (s Services) V1() v1.Services {
	return v1.New(s.sdk)
}
