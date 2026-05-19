package gosdk

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	DefaultKeepaliveTime                = 20 * time.Second
	DefaultKeepaliveTimeout             = 10 * time.Second
	DefaultKeepalivePermitWithoutStream = true

	EnvGRPCKeepaliveTime                = "NEBIUS_GRPC_KEEPALIVE_TIME"
	EnvGRPCKeepaliveTimeout             = "NEBIUS_GRPC_KEEPALIVE_TIMEOUT"
	EnvGRPCKeepalivePermitWithoutStream = "NEBIUS_GRPC_KEEPALIVE_PERMIT_WITHOUT_STREAM"
)

type keepaliveConfig struct {
	params  keepalive.ClientParameters
	enabled bool
}

func defaultKeepaliveConfig() keepaliveConfig {
	return keepaliveConfig{
		params: keepalive.ClientParameters{
			Time:                DefaultKeepaliveTime,
			Timeout:             DefaultKeepaliveTimeout,
			PermitWithoutStream: DefaultKeepalivePermitWithoutStream,
		},
		enabled: true,
	}
}

func keepaliveConfigFromEnv() (keepaliveConfig, error) {
	cfg := defaultKeepaliveConfig()

	if value, ok := lookupKeepaliveEnv(EnvGRPCKeepaliveTime); ok {
		parsed, err := parseKeepaliveDurationEnv(EnvGRPCKeepaliveTime, value)
		if err != nil {
			return keepaliveConfig{}, err
		}
		cfg.params.Time = parsed
		cfg.enabled = parsed != 0
	}

	if value, ok := lookupKeepaliveEnv(EnvGRPCKeepaliveTimeout); ok {
		parsed, err := parseKeepaliveDurationEnv(EnvGRPCKeepaliveTimeout, value)
		if err != nil {
			return keepaliveConfig{}, err
		}
		cfg.params.Timeout = parsed
	}

	if value, ok := lookupKeepaliveEnv(EnvGRPCKeepalivePermitWithoutStream); ok {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			return keepaliveConfig{}, fmt.Errorf("parse %s: %w", EnvGRPCKeepalivePermitWithoutStream, err)
		}
		cfg.params.PermitWithoutStream = parsed
	}

	if cfg.enabled && cfg.params.Timeout <= 0 {
		return keepaliveConfig{}, fmt.Errorf("%s must be positive when keepalive is enabled", EnvGRPCKeepaliveTimeout)
	}

	return cfg, nil
}

func lookupKeepaliveEnv(name string) (string, bool) {
	value, ok := os.LookupEnv(name)
	if !ok {
		return "", false
	}
	value = strings.TrimSpace(value)
	if value == "" {
		return "", false
	}
	return value, true
}

func parseKeepaliveDurationEnv(name string, value string) (time.Duration, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, fmt.Errorf("%s is empty", name)
	}
	if value == "0" {
		return 0, nil
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("parse %s: %w", name, err)
	}
	if parsed < 0 {
		return 0, fmt.Errorf("%s must not be negative", name)
	}
	return parsed, nil
}

func (cfg keepaliveConfig) dialOption() grpc.DialOption {
	return grpc.WithKeepaliveParams(cfg.params)
}
