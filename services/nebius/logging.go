package nebius

import "github.com/nebius/gosdk/services/nebius/logging"

func (s Services) Logging() logging.Services {
	return logging.New(s.sdk)
}
