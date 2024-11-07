package iam

import "github.com/nebius/gosdk/services/nebius/iam/v1"

func (s Services) V1() v1.Services {
	return v1.New(s.sdk)
}
