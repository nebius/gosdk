package msp

import "github.com/nebius/gosdk/services/nebius/msp/spark"

func (s Services) Spark() spark.Services {
	return spark.New(s.sdk)
}
