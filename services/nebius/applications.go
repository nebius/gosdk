package nebius

import "github.com/nebius/gosdk/services/nebius/applications"

func (s Services) Applications() applications.Services {
	return applications.New(s.sdk)
}
