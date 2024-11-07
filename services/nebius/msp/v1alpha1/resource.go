package v1alpha1

import "github.com/nebius/gosdk/services/nebius/msp/v1alpha1/resource"

func (s Services) Resource() resource.Services {
	return resource.New(s.sdk)
}
