// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package nebius

import "github.com/nebius/gosdk/services/nebius/compute"

func (s Services) Compute() compute.Services {
	return compute.New(s.sdk)
}
