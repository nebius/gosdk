package nebius

import "github.com/nebius/gosdk/services/nebius/vpc"

func (s Services) VPC() vpc.Services {
	return vpc.New(s.sdk)
}
