package msp

import "github.com/nebius/gosdk/services/nebius/msp/postgresql"

func (s Services) PostgreSQL() postgresql.Services {
	return postgresql.New(s.sdk)
}
