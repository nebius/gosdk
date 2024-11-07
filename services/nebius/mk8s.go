package nebius

import "github.com/nebius/gosdk/services/nebius/mk8s"

func (s Services) MK8S() mk8s.Services {
	return mk8s.New(s.sdk)
}
