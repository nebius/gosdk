package grpcheader

import (
	"context"

	"github.com/nebius/gosdk/fieldmask/mask"
)

// SelectMask is the gRPC header name where the select mask will be added.
const SelectMask = "X-SelectMask"

// IsSelectMaskInOutgoingContext checks if the select mask is present in the
// outgoing context metadata (gRPC headers).
//
// Parameters:
//   - ctx [context.Context]: The context to check for the presence of the
//     select mask.
//
// Returns:
//   - bool: True if the select mask is present in the outgoing context
//     metadata, false otherwise.
func IsSelectMaskInOutgoingContext(ctx context.Context) bool {
	return IsInOutgoingContext(ctx, SelectMask)
}

// AddSelectMaskToOutgoingContext adds the select [mask.Mask] to the outgoing
// context metadata (gRPC headers).
//
// Parameters:
//   - ctx [context.Context]: The context to which the select mask will be
//     added.
//   - selectMask *[mask.Mask]: The select mask to be added.
//
// Returns:
//   - [context.Context]: The context with the select mask added to the
//     metadata.
//   - error: An error if there was any issue marshalling the select mask or
//     appending it to the context metadata.
func AddSelectMaskToOutgoingContext(
	ctx context.Context,
	selectMask *mask.Mask,
) (context.Context, error) {
	return AddMaskToOutgoingContext(ctx, SelectMask, selectMask)
}

// GetSelectMaskFromIncomingContext retrieves the select [mask.Mask] from the
// incoming context metadata (grpc headers) and parses it.
// If several masks were present, they are merged together.
// If no mask is present, an empty mask will be returned.
//
// Parameters:
//   - ctx [context.Context]: The context from which to extract the select mask.
//
// Returns:
//   - *[mask.Mask]: The select mask retrieved from the context metadata.
//   - error: An error if there was any issue parsing or merging the select
//     masks.
func GetSelectMaskFromIncomingContext(ctx context.Context) (*mask.Mask, error) {
	return GetMaskFromIncomingContext(ctx, SelectMask)
}
