package v1

import "github.com/nebius/gosdk/services/nebius/logging/v1/agentmanager"

func (s Services) AgentManager() agentmanager.Services {
	return agentmanager.New(s.sdk)
}
