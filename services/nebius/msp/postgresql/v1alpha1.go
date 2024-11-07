package postgresql

import "github.com/nebius/gosdk/services/nebius/msp/postgresql/v1alpha1"

func (s Services) V1Alpha1() v1alpha1.Services {
	return v1alpha1.New(s.sdk)
}
