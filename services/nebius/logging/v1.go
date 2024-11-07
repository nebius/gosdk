package logging

import "github.com/nebius/gosdk/services/nebius/logging/v1"

func (s Services) V1() v1.Services {
	return v1.New(s.sdk)
}
