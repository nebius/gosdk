package nebius

import "github.com/nebius/gosdk/services/nebius/iam"

func (s Services) IAM() iam.Services {
	return iam.New(s.sdk)
}
