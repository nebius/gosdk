package grpcheader

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/nebius/gosdk/proto/fieldmask/mask"
	"github.com/nebius/gosdk/proto/fieldmask/protobuf"
)

// ResetMask is the gRPC header name where the reset mask will be added.
const ResetMask = "X-ResetMask"

// IsResetMaskInOutgoingContext checks if the reset mask is present in the
// outgoing context metadata (gRPC headers).
//
// Parameters:
//   - ctx [context.Context]: The context to check for the presence of the reset
//     mask.
//
// Returns:
//   - bool: True if the reset mask is present in the outgoing context metadata,
//     false otherwise.
func IsResetMaskInOutgoingContext(ctx context.Context) bool {
	return IsInOutgoingContext(ctx, ResetMask)
}

// AddResetMaskToOutgoingContext adds the reset [mask.Mask] to the outgoing
// context metadata (gRPC headers).
//
// Parameters:
//   - ctx [context.Context]: The context to which the reset mask will be added.
//   - resetMask *[mask.Mask]: The reset mask to be added.
//
// Returns:
//   - [context.Context]: The context with the reset mask added to the metadata.
//   - error: An error if there was any issue marshalling the reset mask or
//     appending it to the context metadata.
func AddResetMaskToOutgoingContext(
	ctx context.Context,
	resetMask *mask.Mask,
) (context.Context, error) {
	return AddMaskToOutgoingContext(ctx, ResetMask, resetMask)
}

// AddMessageResetMaskToOutgoingContext adds the reset [mask.Mask] derived from
// the given [proto.Message] to the outgoing context metadata (gRPC headers).
//
// Parameters:
//   - ctx [context.Context]: The context to which the reset mask will be added.
//   - msg [proto.Message]: The proto message from which the reset mask will be
//     derived.
//
// Returns:
//   - [context.Context]: The context with the reset mask.
//   - error: An error if there was any issue deriving the reset mask from the
//     message or adding it to the context metadata.
func AddMessageResetMaskToOutgoingContext(
	ctx context.Context,
	msg proto.Message,
) (context.Context, error) {
	resetMask, err := protobuf.ResetMaskFromMessage(msg)
	if err != nil {
		return ctx, fmt.Errorf("failed to get mask from message: %w", err)
	}

	return AddResetMaskToOutgoingContext(ctx, resetMask)
}

// EnsureMessageResetMaskInOutgoingContext ensures that the reset [mask.Mask] is
// present in the outgoing context metadata. If the reset mask is already
// present, it returns the context unchanged. Otherwise, it derives the reset
// mask from the given [proto.Message] and adds it to the context metadata.
//
// Parameters:
//   - ctx [context.Context]: The context to check for the presence of the reset
//     mask.
//   - msg [proto.Message]: The proto message from which to derive the reset
//     mask if it's not already present.
//
// Returns:
//   - [context.Context]: The context with the reset mask added to the metadata,
//     if necessary.
//   - error: An error if there was any issue deriving or adding the reset mask
//     to the context metadata.
func EnsureMessageResetMaskInOutgoingContext(
	ctx context.Context,
	msg proto.Message,
) (context.Context, error) {
	if IsResetMaskInOutgoingContext(ctx) {
		return ctx, nil
	}
	return AddMessageResetMaskToOutgoingContext(ctx, msg)
}

// GetResetMaskFromIncomingContext retrieves the reset [mask.Mask] from the
// incoming context metadata (grpc headers) and parses it.
// If several masks were present, they are merged together.
// If no mask is present, an empty mask will be returned.
//
// Parameters:
//   - ctx [context.Context]: The context from which to extract the reset mask.
//
// Returns:
//   - *[mask.Mask]: The reset mask retrieved from the context metadata.
//   - error: An error if there was any issue parsing or merging the reset
//     masks.
func GetResetMaskFromIncomingContext(ctx context.Context) (*mask.Mask, error) {
	return GetMaskFromIncomingContext(ctx, ResetMask)
}
