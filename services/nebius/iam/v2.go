// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package iam

import "github.com/nebius/gosdk/services/nebius/iam/v2"

func (s Services) V2() v2.Services {
	return v2.New(s.sdk)
}
