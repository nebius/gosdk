// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package nebius

import "github.com/nebius/gosdk/services/nebius/audit"

func (s Services) Audit() audit.Services {
	return audit.New(s.sdk)
}
