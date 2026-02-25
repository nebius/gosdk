package constants

import "time"

const TokenEnv = "NEBIUS_IAM_TOKEN"
const ProfileEnv = "NEBIUS_PROFILE"
const EndpointEnv = "NEBIUS_ENDPOINT"

const Domain = "api.nebius.cloud:443"

const HTTPMaxAttempts = 3
const HTTPBaseBackoff = 200 * time.Millisecond
