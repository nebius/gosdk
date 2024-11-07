package nebius

import "github.com/nebius/gosdk/services/nebius/storage"

func (s Services) Storage() storage.Services {
	return storage.New(s.sdk)
}
