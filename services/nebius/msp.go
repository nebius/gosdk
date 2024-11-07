package nebius

import "github.com/nebius/gosdk/services/nebius/msp"

func (s Services) MSP() msp.Services {
	return msp.New(s.sdk)
}
