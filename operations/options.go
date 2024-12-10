package operations

import (
	"time"

	"google.golang.org/grpc"
)

const DefaultPollInterval = 1 * time.Second

type pollInterval struct {
	grpc.EmptyCallOption

	interval time.Duration
}

// PollInterval returns [grpc.CallOption] you can pass to [Operation.Wait] to override the [DefaultPollInterval].
//
// If an interval isn't positive, the [DefaultPollInterval] is used.
func PollInterval(interval time.Duration) grpc.CallOption {
	return pollInterval{
		interval: interval,
	}
}
