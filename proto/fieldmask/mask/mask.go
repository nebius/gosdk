package mask

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

const recursionTooDeep = 1000

var ErrRecursionTooDeep = errors.New("recursion too deep")

// Mask represents a mask structure used for pattern matching.
type Mask struct {
	Any        *Mask              // Pointer to a wildcard mask.
	FieldParts map[FieldKey]*Mask // Map of field keys to corresponding masks.
}

// New creates a new [Mask] instance initialized with empty field parts map.
func New() *Mask {
	return &Mask{
		FieldParts: map[FieldKey]*Mask{},
	}
}

// IsEmpty checks if the [Mask] is empty, i.e., has no wildcard or field parts.
func (m *Mask) IsEmpty() bool {
	if m.Any != nil {
		return false
	}
	for _, p := range m.FieldParts {
		if p != nil {
			return false
		}
	}
	return true
}

func (m *Mask) toFieldPathRecursive(recursion int) (FieldPath, error) {
	if recursion >= recursionTooDeep {
		return nil, ErrRecursionTooDeep
	}
	recursion++
	if m.Any != nil {
		return nil, errors.New("wildcard in the mask")
	}
	if len(m.FieldParts) == 0 {
		return nil, nil
	}
	seen := false
	ret := FieldPath{}
	for k, v := range m.FieldParts {
		if v != nil {
			if seen {
				return nil, errors.New("multiple paths in the mask")
			}
			seen = true
			inner, err := v.toFieldPathRecursive(recursion)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", k, err)
			}
			ret = append(ret, k)
			ret = append(ret, inner...)
		}
	}
	return ret, nil
}

// ToFieldPath converts the [Mask] into a [FieldPath] if it represents a valid
// path to a single field. This method recursively transforms the Mask into a
// [FieldPath] if the mask represents only one field, ie it has no wildcards and
// has only one field at each level.
//
// Returns:
//   - *FieldPath: A [FieldPath] representing the path to a single field.
//   - error: An error if the [Mask] does not meet the criteria for a valid
//     [FieldPath], such as having multiple parts at any level or containing
//     wildcards.
func (m *Mask) ToFieldPath() (FieldPath, error) {
	return m.toFieldPathRecursive(0)
}

// IsFieldPath checks if the [Mask] represents a valid [FieldPath].
// A [Mask] is considered a valid [FieldPath] if it represents a path to a
// single field, meaning it has at any level exactly one FieldPart and does not
// contain any wildcards.
//
// This function is a wrapper around [Mask.ToFieldPath] and requires deep
// traversal.
//
// Returns:
//   - bool: True if the [Mask] corresponds to a valid [FieldPath] (path to a
//     single field), otherwise false.
func (m *Mask) IsFieldPath() bool {
	_, err := m.ToFieldPath()
	return err == nil
}

func appendMarshaler(
	ret []string,
	kMask string,
	mask *Mask,
	recursion int,
) ([]string, error) {
	counter, vMask, err := mask.marshal(recursion + 1)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", kMask, err)
	}
	switch {
	case vMask == "":
		ret = append(ret, kMask)
	case counter == 1:
		ret = append(ret, kMask+"."+vMask)
	default:
		ret = append(ret, kMask+".("+vMask+")")
	}
	return ret, nil
}

func (m *Mask) marshal(recursion int) (int, string, error) {
	if recursion >= recursionTooDeep {
		return 0, "", ErrRecursionTooDeep
	}
	if m.Any == nil && len(m.FieldParts) == 0 {
		return 0, "", nil
	}
	ret := []string{}
	var err error
	if m.Any != nil {
		ret, err = appendMarshaler(ret, "*", m.Any, recursion)
		if err != nil {
			return 0, "", err
		}
	}
	for k, v := range m.FieldParts {
		if v == nil {
			continue
		}
		kMask, err := k.Marshal()
		if err != nil {
			return 0, "", fmt.Errorf("%s: %w", k, err)
		}

		ret, err = appendMarshaler(ret, kMask, v, recursion)
		if err != nil {
			return 0, "", err
		}
	}
	sort.Strings(ret)
	return len(ret), strings.Join(ret, ","), nil
}

// Marshal generates a parseable string representation of the [Mask].
// It recursively constructs a string that represents the structure of the
// [Mask] suitable for parsing back into a [Mask].
func (m *Mask) Marshal() (string, error) {
	_, ret, err := m.marshal(0)
	return ret, err
}

// MarshalText runs [Mask.Marshal] and converts string to []byte.
func (m *Mask) MarshalText() ([]byte, error) {
	ret, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	return []byte(ret), nil
}

// UnmarshalText runs [Parse] against input text as string and then replaces
// current mask with the results. If the result is nil, it will clear the mask.
func (m *Mask) UnmarshalText(text []byte) error {
	if m == nil {
		return errors.New("mask not initialized")
	}
	ret, err := Parse(string(text))
	if err != nil {
		return err
	}
	if ret == nil {
		m.Any = nil
		m.FieldParts = map[FieldKey]*Mask{}
	} else {
		m.Any = ret.Any
		m.FieldParts = ret.FieldParts
	}
	return nil
}

// String returns a string representation of the [Mask] for visualization
// purposes. It generates a human-readable representation of the [Mask]
// structure, providing a meaningful display of the [Mask]'s content.
func (m *Mask) String() string {
	ret, err := m.Marshal()
	if err != nil {
		return "Mask<not-marshalable " + err.Error() + ">"
	}
	return "Mask<" + ret + ">"
}

func (m *Mask) copyRecursive(recursion int) (*Mask, error) {
	if m == nil {
		return nil, nil
	}
	if recursion >= recursionTooDeep {
		return nil, ErrRecursionTooDeep
	}
	recursion++
	ret := New()
	if m.Any != nil {
		mask, err := m.Any.copyRecursive(recursion)
		if err != nil {
			return nil, fmt.Errorf("*: %w", err)
		}
		ret.Any = mask
	}
	if len(m.FieldParts) > 0 {
		for k, v := range m.FieldParts {
			mask, err := v.copyRecursive(recursion)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", k, err)
			}
			if mask != nil {
				ret.FieldParts[k] = mask
			}
		}
	}
	return ret, nil
}

// Copy creates a deep copy of the [Mask], including all its wildcard and field
// parts. It recursively copies the structure of the Mask to produce an
// independent replica.
func (m *Mask) Copy() (*Mask, error) {
	return m.copyRecursive(0)
}

func (m *Mask) mergeRecursive(with *Mask, recursion int) error { //nolint:gocognit
	if recursion >= recursionTooDeep {
		return ErrRecursionTooDeep
	}
	recursion++
	if with == nil || with.IsEmpty() {
		return nil
	}
	if m.Any != nil {
		err := m.Any.mergeRecursive(with.Any, recursion)
		if err != nil {
			return fmt.Errorf("*:M %w", err)
		}
	} else if with.Any != nil {
		mask, err := with.Any.copyRecursive(recursion)
		if err != nil {
			return fmt.Errorf("*:C %w", err)
		}
		m.Any = mask
	}
	if m.FieldParts == nil {
		m.FieldParts = map[FieldKey]*Mask{}
	}
	if len(with.FieldParts) > 0 {
		for k, v := range m.FieldParts {
			if part, ok := with.FieldParts[k]; ok && v != nil {
				err := v.mergeRecursive(part, recursion)
				if err != nil {
					return fmt.Errorf("%s:M %w", k, err)
				}
			}
		}
		for k, v := range with.FieldParts {
			if m.FieldParts[k] == nil {
				mask, err := v.copyRecursive(recursion)
				if err != nil {
					return fmt.Errorf("%s:C %w", k, err)
				}
				if mask != nil {
					m.FieldParts[k] = mask
				} else {
					delete(m.FieldParts, k)
				}
			}
		}
	}
	return nil
}

// Merge combines the current [Mask] with another [Mask] by merging their
// structures. It recursively merges the [Mask] with another [Mask], combining
// their field parts and wildcard.
func (m *Mask) Merge(with *Mask) error {
	return m.mergeRecursive(with, 0)
}

func (m *Mask) equalRecursive(to *Mask, recursion int) bool {
	if recursion >= recursionTooDeep {
		return false
	}
	recursion++
	if m == nil || to == nil {
		return to == m
	}
	if !m.Any.equalRecursive(to.Any, recursion) {
		return false
	}
	mfp := m.FieldParts
	tfp := to.FieldParts
	for k, v := range mfp {
		if !v.equalRecursive(tfp[k], recursion) {
			return false
		}
	}
	for k, v := range tfp {
		if _, ok := mfp[k]; !ok && v != nil {
			return false
		}
	}
	return true
}

// GetSubMask retrieves the full field mask for the specified [FieldKey]
// from the [Mask], optionally merged with the wildcard mask (m.Any) if
// available.
//
// If a mask associated with the [FieldKey] is found in the [Mask]'s FieldParts
// map, it is returned. If no mask is found for the [FieldKey], the method
// returns the wildcard mask (m.Any) if available.
//
// If both a specific mask and the wildcard mask (m.Any) are present, the method
// creates a copy of the specific mask and merges it with the wildcard mask
// using the Merge method.
//
// Note: may return direct link instead of copying.
//
// Parameters:
//   - key: The [FieldKey] for which the full field mask is requested.
//
// Returns:
//   - *Mask: The full field mask associated with the specified FieldKey,
//     possibly merged with m.Any.
//   - error: An error if there was any issue copying or merging the masks.
func (m *Mask) GetSubMask(key FieldKey) (*Mask, error) {
	if m == nil {
		return nil, nil
	}
	ret := m.FieldParts[key]
	if ret == nil {
		return m.Any, nil
	}
	if m.Any != nil {
		cp, err := ret.Copy()
		if err != nil {
			return nil, fmt.Errorf("failed to copy key mask: %w", err)
		}
		if err := cp.Merge(m.Any); err != nil {
			return nil, fmt.Errorf(
				"failed to merge key mask with wildcard mask: %w",
				err,
			)
		}
		ret = cp
	}
	return ret, nil
}

// GetSubMaskByPath retrieves the sub-mask at the specified [FieldPath] from the
// [Mask].
//
// The GetSubMaskByPath method traverses the [Mask] structure based on the given
// [FieldPath] to retrieve the sub-mask located at that path. It combines
// wildcard masks (if present, using Any) matching the path with direct masks to
// ensure correct retrieval of nested masks.
//
// Example:
//
//	mask := ParseMust("user.*.telegram,user.profile.email")
//
//	path := NewFieldPath("user", "profile")
//	subMask, err := mask.GetSubMaskByPath(path)
//	if err != nil {
//	    // Handle error
//	}
//
// The subMask variable in the example above would represent the Mask structure
// corresponding to the "user.profile" path within the original Mask with both
// email and telegram fields.
//
// Note: may return direct link instead of copying.
//
// Parameters:
//   - path: The [FieldPath] indicating the path to the desired sub-mask.
//
// Returns:
//   - *Mask: The sub-mask located at the specified [FieldPath] within the
//     [Mask].
//   - error: An error, if any, encountered during the retrieval process.
//     This could occur if the specified path is not found or if there
//     are any issues accessing nested masks along the path.
func (m *Mask) GetSubMaskByPath(path FieldPath) (*Mask, error) {
	sub := m
	var err error
	for i, k := range path {
		sub, err = sub.GetSubMask(k)
		if err != nil {
			return nil, fmt.Errorf("failed to get field mask at path %s: %w", path[0:i], err)
		}
		if sub == nil {
			return nil, nil
		}
	}
	return sub, nil
}

// Equal checks if the current [Mask] is equal to another [Mask] by comparing
// their structures. It recursively compares the [Mask] with another Mask to
// determine if they are identical.
func (m *Mask) Equal(to *Mask) bool {
	return m.equalRecursive(to, 0)
}

func (m *Mask) intersectRMRecursive( //nolint:gocognit // TODO: simplify?
	other *Mask,
	recursion int,
) (*Mask, error) {
	if recursion >= recursionTooDeep {
		return nil, ErrRecursionTooDeep
	}
	recursion++
	if m == nil || other == nil {
		return nil, nil
	}
	ret := &Mask{
		FieldParts: map[FieldKey]*Mask{},
	}
	inner, err := m.Any.intersectRMRecursive(other.Any, recursion)
	if err != nil {
		return nil, fmt.Errorf("*×*: %w", err)
	}
	ret.Any = inner

	if other.Any != nil {
		for k, v := range m.FieldParts {
			inner, err := v.intersectRMRecursive(other.Any, recursion)
			if err != nil {
				return nil, fmt.Errorf("%s×*: %w", k, err)
			}
			if inner != nil {
				ret.FieldParts[k] = inner
			}
		}
	}

	if m.Any != nil {
		for k, v := range other.FieldParts {
			inner, err := m.Any.intersectRMRecursive(v, recursion)
			if err != nil {
				return nil, fmt.Errorf("*×%s: %w", k, err)
			}
			if inner != nil {
				if err := inner.Merge(ret.FieldParts[k]); err != nil {
					return nil, fmt.Errorf("*×%s: Merge: %w", k, err)
				}
				ret.FieldParts[k] = inner
			}
		}
	}

	for k, v := range m.FieldParts {
		inner, err := v.intersectRMRecursive(other.FieldParts[k], recursion)
		if err != nil {
			return nil, fmt.Errorf("%s×%s: %w", k, k, err)
		}
		if inner != nil {
			if err := inner.Merge(ret.FieldParts[k]); err != nil {
				return nil, fmt.Errorf("%s×%s: Merge: %w", k, k, err)
			}
			ret.FieldParts[k] = inner
		}
	}

	return ret, nil
}

// IntersectResetMask creates arithmetic intersection of two reset masks.
//
// This function recursively intersects all the fields and wildcards in masks
// and returns the new mask containing only fields included in both.
// It also checks for wildcards when working with named fields.
//
// Intersection of two masks is essentially their applications composition.
//
// Parameters:
//   - other *[Mask]: the mask to intersect with.
//
// Returns:
//   - *[Mask]: the new mask representing intersection between two masks.
//   - error: An error that may appear during intersection process.
func (m *Mask) IntersectResetMask(other *Mask) (*Mask, error) {
	return m.intersectRMRecursive(other, 0)
}

func (m *Mask) intersectDumbRecursive(
	other *Mask,
	recursion int,
) (*Mask, error) {
	if recursion >= recursionTooDeep {
		return nil, ErrRecursionTooDeep
	}
	recursion++
	if m == nil || other == nil {
		return nil, nil
	}
	ret := &Mask{
		FieldParts: map[FieldKey]*Mask{},
	}
	inner, err := m.Any.intersectDumbRecursive(other.Any, recursion)
	if err != nil {
		return nil, fmt.Errorf("*: %w", err)
	}
	ret.Any = inner

	for k, v := range m.FieldParts {
		inner, err := v.intersectDumbRecursive(other.FieldParts[k], recursion)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", k, err)
		}
		if inner != nil {
			ret.FieldParts[k] = inner
		}
	}

	return ret, nil
}

// IntersectDumb creates arithmetic intersection of two masks.
//
// This function recursively intersects all the fields and wildcards in masks
// and returns the new mask containing only fields included in both.
//
// Unlike [Mask.IntersectResetMask], this function doesn't cross-check wildcards
// while intersecting named fields.
//
// Parameters:
//   - other *[Mask]: the mask to intersect with.
//
// Returns:
//   - *[Mask]: the new mask representing intersection between two masks.
//   - error: An error that may appear during intersection process.
func (m *Mask) IntersectDumb(other *Mask) (*Mask, error) {
	return m.intersectDumbRecursive(other, 0)
}

func (m *Mask) subtractRecursive(other *Mask, recursion int) error {
	if recursion >= recursionTooDeep {
		return ErrRecursionTooDeep
	}
	recursion++

	if m == nil || other == nil {
		return nil
	}
	if m.Any != nil && other.Any != nil {
		if err := m.Any.subtractRecursive(other.Any, recursion); err != nil {
			return fmt.Errorf("*: %w", err)
		}
		if m.Any.IsEmpty() {
			m.Any = nil
		}
	}
	for k, v := range m.FieldParts {
		if v == nil {
			delete(m.FieldParts, k)
		}
		if o := other.FieldParts[k]; o != nil {
			if err := v.subtractRecursive(o, recursion); err != nil {
				return fmt.Errorf("%s: %w", k, err)
			}
			if v.IsEmpty() {
				delete(m.FieldParts, k)
			}
		}
	}

	return nil
}

// SubtractDumb subtracts other [Mask] from this one.
//
// This function recursively removes fields existing in both masks. If the field
// exists in both and is empty after recursion it is being removed.
//
// Wild cards are subtracted separately without cross-action on named fields.
//
// Note: subtraction is done in-place.
//
// Parameters:
//   - other *[Mask]: the mask being subtracted.
//
// Returns:
//   - error: An error that may appear during intersection process.
func (m *Mask) SubtractDumb(other *Mask) error {
	return m.subtractRecursive(other, 0)
}

func (m *Mask) subtractResetRecursive(other *Mask, recursion int) error { //nolint:gocognit
	if recursion >= recursionTooDeep {
		return ErrRecursionTooDeep
	}
	recursion++

	if m == nil || other == nil {
		return nil
	}

	if m.Any != nil && other.Any != nil {
		err := m.Any.subtractResetRecursive(other.Any, recursion)
		if err != nil {
			return fmt.Errorf("*\\*: %w", err)
		}
		if m.Any.IsEmpty() {
			m.Any = nil
		}
	}
	for k, v := range m.FieldParts {
		if v == nil {
			delete(m.FieldParts, k)
		}
		if other.Any == nil && other.FieldParts[k] == nil {
			continue
		}
		if other.Any != nil {
			err := v.subtractResetRecursive(other.Any, recursion)
			if err != nil {
				return fmt.Errorf("%s\\*: %w", k, err)
			}
		}
		if other.FieldParts[k] != nil {
			err := v.subtractResetRecursive(other.FieldParts[k], recursion)
			if err != nil {
				return fmt.Errorf("%s\\%s: %w", k, k, err)
			}
		}
		if v.IsEmpty() {
			delete(m.FieldParts, k)
		}
	}

	return nil
}

// SubtractResetMask subtracts other [Mask] from this one according to
// ResetMask logic.
//
// This function recursively removes fields existing in both masks. If the field
// exists in both and is empty after recursive subtraction, it is being
// removed.
//
// Wild cards are subtracted both from the original wildcard and from each of
// the named fields.
//
// Note: subtraction is done in-place.
//
// Parameters:
//   - other *[Mask]: the mask being subtracted.
//
// Returns:
//   - error: An error that may appear during intersection process.
func (m *Mask) SubtractResetMask(other *Mask) error {
	return m.subtractResetRecursive(other, 0)
}
