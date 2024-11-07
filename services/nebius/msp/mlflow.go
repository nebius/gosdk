package msp

import "github.com/nebius/gosdk/services/nebius/msp/mlflow"

func (s Services) MLFlow() mlflow.Services {
	return mlflow.New(s.sdk)
}
