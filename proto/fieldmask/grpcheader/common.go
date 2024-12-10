package grpcheader

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/nebius/gosdk/proto/fieldmask/mask"
)

// IsInOutgoingContext checks if a metadata key (grpc header) exists in the
// outgoing context.
//
// Parameters:
//   - ctx [context.Context]: The context from which to extract metadata.
//   - name string: The name of the metadata key to check.
//
// Returns:
//   - bool: True if the metadata key exists in the outgoing context, false
//     otherwise.
func IsInOutgoingContext(ctx context.Context, name string) bool {
	name = strings.ToLower(name)
	md, _ := metadata.FromOutgoingContext(ctx)
	return len(md[name]) != 0
}

// AddMaskToOutgoingContext adds a [mask.Mask] as a header to the outgoing
// grpc context.
//
// Parameters:
//   - ctx [context.Context]: The context to which the mask will be added.
//   - maskName string: The name of the mask to be added to the metadata.
//   - mask *[mask.Mask]: The mask to be added.
//
// Returns:
//   - [context.Context]: The context with the mask added to the metadata.
//   - error: An error if there was any issue marshalling the mask or appending
//     it to the context metadata.

func AddMaskToOutgoingContext(
	ctx context.Context,
	maskName string,
	mask *mask.Mask,
) (context.Context, error) {
	mstr, err := mask.Marshal()
	if err != nil {
		return ctx, fmt.Errorf("failed to marshal mask: %w", err)
	}
	return metadata.AppendToOutgoingContext(ctx, maskName, mstr), nil
}

// GetMaskFromIncomingContext retrieves a [mask.Mask] from the incoming context
// metadata (grpc headers) by its `maskHeaderName` and parses it.
// If several masks were present, they are merged together.
// If no mask is present, an empty mask will be returned.
//
// Parameters:
//   - ctx [context.Context]: The context from which to extract the mask.
//   - maskHeaderName string: The name of the mask to retrieve from the
//     metadata.
//
// Returns:
//   - *[mask.Mask]: The mask retrieved from the context metadata.
//   - error: An error if there was any issue parsing or merging the masks.
func GetMaskFromIncomingContext(
	ctx context.Context,
	maskHeaderName string,
) (*mask.Mask, error) {
	maskHeaderName = strings.ToLower(maskHeaderName)
	maskStrings := metadata.ValueFromIncomingContext(ctx, maskHeaderName)
	fieldMask := mask.New()
	for i, str := range maskStrings {
		oneMask, err := mask.Parse(str)
		if err != nil {
			return nil, fmt.Errorf("failed parse mask %d: %w", i, err)
		}
		if err := fieldMask.Merge(oneMask); err != nil {
			return nil, fmt.Errorf("failed merge mask %d: %w", i, err)
		}
	}
	return fieldMask, nil
}
