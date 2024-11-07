package mask

import (
	"errors"
)

// FieldPath represents a path to a specific field within a message structure.
// It is a sequence of [FieldKey]s that identifies the location of a field.
type FieldPath []FieldKey

func NewFieldPath(keys ...FieldKey) FieldPath {
	ret := make(FieldPath, 0, len(keys))
	ret = append(ret, keys...)
	return ret
}

// Join concatenates the current [FieldPath] with additional [FieldKey]s,
// returning a new extended [FieldPath].
//
// The Join method appends the specified [FieldKey]s (keys...) to the end of the
// current [FieldPath] (fp), producing a new [FieldPath] that represents a
// longer path. This operation does not modify the original [FieldPath].
//
// Example:
//
//	fp := NewFieldPath("user", "profile")
//	extendedPath := fp.Join("email") // Result: ["user", "profile", "email"]
//
// If no additional [FieldKey]s are provided (keys...), the method will return
// a copy of the current [FieldPath].
//
// Note that the new [FieldPath] is created as a separate instance from the
// original, ensuring that immutability is maintained for the current
// [FieldPath].
//
// Parameters:
//   - keys ([FieldKey]...): Additional [FieldKey]s to append to the current
//     [FieldPath].
//
// Returns:
//   - FieldPath: A new [FieldPath] instance representing the extended path.
func (fp FieldPath) Join(keys ...FieldKey) FieldPath {
	ret := make(FieldPath, 0, len(keys)+len(fp))
	ret = append(ret, fp...)
	ret = append(ret, keys...)
	return ret
}

// Parent returns the parent [FieldPath] by excluding the last [FieldKey],
// representing the path to the parent field of the current field.
//
// If the current [FieldPath] is empty (nil or has zero length), the Parent
// method returns nil, indicating that there is no parent path.
//
// Example:
//
//	fp := NewFieldPath("user", "profile", "email")
//	parentPath := fp.Parent() // Result: ["user", "profile"]
//
// If the [FieldPath] contains only one [FieldKey], calling Parent on it
// will return an empty [FieldPath], as it represents the
// top-level path with no parent.
//
// Note that the Parent method does not modify the original [FieldPath];
// it returns a new [FieldPath] representing the parent path.
func (fp FieldPath) Parent() FieldPath {
	if len(fp) == 0 {
		return nil
	}
	return fp[0 : len(fp)-1]
}

// Copy creates and returns a copy of the current [FieldPath].
//
// The Copy method duplicates the current [FieldPath], producing an identical
// [FieldPath] instance with the same sequence of [FieldKey]s. This operation
// ensures that modifications to the copied [FieldPath] do not affect the
// original [FieldPath].
//
// Example:
//
//	fp := NewFieldPath("user", "profile", "email")
//	copiedPath := fp.Copy() // Result: ["user", "profile", "email"]
//
// Note that the Copy method is equivalent to calling [FieldPath.Join]() with no
// additional [FieldKey]s, resulting in a new [FieldPath] that is a clone of the
// original.
func (fp FieldPath) Copy() FieldPath {
	return fp.Join()
}

// ToMask converts the [FieldPath] into a [Mask] structure, representing a path
// mask.
//
// The ToMask method constructs a [Mask] structure based on the [FieldPath],
// where each [FieldKey] in the [FieldPath] corresponds to a nested [Mask]
// within the resulting [Mask]. This method is useful for generating a [Mask]
// that can be used to specify which parts of a data structure are accessed or
// modified.
//
// Example:
//
//	fp := NewFieldPath("user", "profile", "email")
//	mask := fp.ToMask()
//
// The resulting [Mask] from the example above would be: `user.profile.email`
//
// Returns:
//   - *Mask: A [Mask] structure representing the path mask derived from the
//     [FieldPath].
func (fp FieldPath) ToMask() *Mask {
	ret := New()
	cur := ret
	for _, k := range fp {
		next := New()
		cur.FieldParts[k] = next
		cur = next
	}
	return ret
}

// Equal checks if two [FieldPath]s are identical in sequence and length.
//
// The Equal method compares the current [FieldPath] with another [FieldPath]
// (other) to determine if they have the same sequence of [FieldKey]s and are of
// the same length. It returns true if both [FieldPath]s are identical, and
// false otherwise.
//
// Example:
//
//	fp1 := NewFieldPath("user", "profile", "email")
//	fp2 := NewFieldPath("user", "profile", "email")
//	isEqual := fp1.Equal(fp2) // Result: true
//
// Note that the Equal method performs a deep comparison of the [FieldPath]s.
func (fp FieldPath) Equal(other FieldPath) bool {
	if len(fp) != len(other) {
		return false
	}
	for i := range fp {
		if fp[i] != other[i] {
			return false
		}
	}
	return true
}

// IsPrefixOf checks if the current [FieldPath] is a prefix of another
// [FieldPath].
//
// The IsPrefixOf method determines whether the current [FieldPath] is a prefix
// of another [FieldPath] (other). It returns true if every [FieldKey] in the
// current [FieldPath] matches the corresponding [FieldKey] in the beginning of
// the other [FieldPath], and false otherwise.
//
// Example:
//
//	prefix := NewFieldPath("user", "profile")
//	fullPath := NewFieldPath("user", "profile", "email")
//	isPrefix := prefix.IsPrefixOf(fullPath) // Result: true
//
// If the current [FieldPath] is longer than the other [FieldPath], IsPrefixOf
// returns false.
func (fp FieldPath) IsPrefixOf(other FieldPath) bool {
	if len(fp) >= len(other) {
		return false
	}
	for i := range fp {
		if fp[i] != other[i] {
			return false
		}
	}
	return true
}

func fieldKeyMatchResetMaskRecursive(
	slice FieldPath,
	mask *Mask,
) (bool, bool) {
	if mask == nil {
		return false, false
	}
	if len(slice) == 0 {
		return true, mask.IsEmpty()
	}
	key, rest := slice[0], slice[1:]
	hasMatch := false
	isFinal := false
	if mask.Any != nil {
		hasMatch, isFinal = fieldKeyMatchResetMaskRecursive(rest, mask.Any)
	}
	if km, ok := mask.FieldParts[key]; ok && km != nil {
		kMatch, kFinal := fieldKeyMatchResetMaskRecursive(rest, km)
		hasMatch = hasMatch || kMatch
		if kMatch {
			isFinal = isFinal || kFinal
		}
	}
	return hasMatch, isFinal
}

// MatchesResetMask checks if the [FieldPath] matches the provided [Mask] in
// terms of `X-ResetMask` behavior by [Design].
//
// The MatchesResetMask method recursively matches each [FieldKey] in the
// current [FieldPath] against the provided Mask. It returns true if there is a
// complete match according to the Mask structure, and false otherwise.
//
// Example:
//
//		fp := NewFieldPath("user", "profile", "email")
//		mask := ParseMust("user.profile.email")
//		doesMatch := fp.MatchesResetMask(mask) // Result: true
//	    doesMatchParent := fp.Parent().MatchesResetMask(mask) // Result: true
//
// If the length of the [FieldPath] exceeds a defined recursion limit
// [recursionTooDeep], MatchesResetMask returns false to prevent deep recursion.
//
// [Design]: https://nebius.atlassian.net/wiki/spaces/NEWBIUS/pages/131367768/ResetMask+tech+design
func (fp FieldPath) MatchesResetMask(mask *Mask) bool {
	if len(fp) >= recursionTooDeep {
		return false
	}
	ret, _ := fieldKeyMatchResetMaskRecursive(fp, mask)
	return ret
}

// MatchesResetMaskFinal checks if the [FieldPath] matches the provided [Mask]
// in terms of `X-ResetMask` behavior by [Design], and if the path ends exactly
// where the [Mask] structure ends (is fully consumed).
//
// The MatchesResetMaskFinal method recursively matches each [FieldKey] in the
// current [FieldPath] against the provided [Mask]. It returns true if there is
// a complete match according to the [Mask] structure and the [FieldPath]
// exactly corresponds to the last node (non-submask) in the [Mask], ensuring
// that no submasks are present beyond the last [FieldKey] in the [FieldPath].
//
// Example:
//
//	fp := NewFieldPath("user", "profile", "email")
//	mask := ParseMust("user.profile.email")
//	doesMatch := fp.MatchesResetMaskFinal(mask) // Result: true
//	doesMatchParent := fp.Parent().MatchesResetMaskFinal(mask) // Result: false
//
// If the length of the [FieldPath] exceeds a defined recursion limit
// [recursionTooDeep], MatchesResetMaskFinal returns false to prevent deep
// recursion.
//
// [Design]: https://nebius.atlassian.net/wiki/spaces/NEWBIUS/pages/131367768/ResetMask+tech+design
func (fp FieldPath) MatchesResetMaskFinal(mask *Mask) bool {
	if len(fp) >= recursionTooDeep {
		return false
	}
	ret, emp := fieldKeyMatchResetMaskRecursive(fp, mask)
	return ret && emp
}

func fieldKeyMatchSelectMaskRecursive(
	slice []FieldKey,
	mask *Mask,
) (bool, bool) {
	if mask == nil || mask.IsEmpty() {
		return true, len(slice) != 0
	}
	if len(slice) == 0 {
		return true, false
	}
	key, rest := slice[0], slice[1:]
	hasMatch := false
	isInner := false
	if mask.Any != nil {
		hasMatch, isInner = fieldKeyMatchSelectMaskRecursive(rest, mask.Any)
	}
	if km, ok := mask.FieldParts[key]; ok && km != nil {
		kMatch, kInner := fieldKeyMatchSelectMaskRecursive(rest, km)
		hasMatch = hasMatch || kMatch
		if kMatch {
			isInner = isInner || kInner
		}
	}
	return hasMatch, isInner
}

// MatchesSelectMask checks if the [FieldPath] matches the provided [Mask]
// according to `X-SelectMask` behavior by [Design].
//
// This method recursively matches each [FieldKey] in the current [FieldPath]
// against the provided [Mask]. It returns true if there is a complete match
// according to the [Mask] structure, allowing any submasks within the provided
// [Mask].
//
// For example:
//
//	mask := ParseMust("a.b")
//	doesMatch, isInner := NewFieldPath("a", "x").MatchesSelectMask(mask) // Result: false
//	doesMatch, isInner := NewFieldPath("a", "b").MatchesSelectMask(mask) // Result: true
//	doesMatch, isInner := NewFieldPath("a", "b", "c").MatchesSelectMask(mask) // Result: true
//
// If the length of the [FieldPath] exceeds a defined recursion limit
// [recursionTooDeep], MatchesSelectMask returns false to prevent deep recursion.
//
// [Design]: https://nebius.atlassian.net/wiki/spaces/NEWBIUS/pages/131367768/ResetMask+tech+design
func (fp FieldPath) MatchesSelectMask(mask *Mask) bool {
	ret, _ := fp.MatchesSelectMaskInner(mask)
	return ret
}

// MatchesSelectMaskInner checks if the [FieldPath] matches the provided [Mask]
// according to `X-SelectMask` behavior by [Design], and also indicates whether
// the match was not the final (inner match).
//
// This method recursively matches each [FieldKey] in the current [FieldPath]
// against the provided [Mask]. It returns true if there is a complete match
// according to the [Mask] structure. Additionally, it returns true if the match
// was not the final match in the [Mask] structure, indicating that submasks
// were present beyond the matched [FieldKey].
//
// For example:
//
//	mask := ParseMust("a.b")
//	doesMatch, isInner := NewFieldPath("a", "x").MatchesSelectMaskInner(mask) // Result: (false, false)
//	doesMatch, isInner := NewFieldPath("a", "b").MatchesSelectMaskInner(mask) // Result: (true, false)
//	doesMatch, isInner := NewFieldPath("a", "b", "c").MatchesSelectMaskInner(mask) // Result: (true, true)
//
// If the length of the [FieldPath] exceeds a defined recursion limit
// [recursionTooDeep], MatchesSelectMaskInner returns (false, false) to prevent
// deep recursion.
//
// [Design]: https://nebius.atlassian.net/wiki/spaces/NEWBIUS/pages/131367768/ResetMask+tech+design
func (fp FieldPath) MatchesSelectMaskInner(mask *Mask) (bool, bool) {
	if len(fp) >= recursionTooDeep {
		return false, false
	}
	return fieldKeyMatchSelectMaskRecursive(fp, mask)
}

// Marshal converts the [FieldPath] into its string representation in a [Mask]
// form.
//
// The Marshal method first converts the [FieldPath] into a [Mask] using the
// [FieldPath.ToMask] method. It then obtains a string representation of the
// resulting mask using the [Mask.Marshal] method.
//
// Example:
//
//	fp := NewFieldPath("user", "profile", "email")
//	str, err := fp.Marshal()
//
// The resulting string (str) from the example above represents the [Mask]
// derived from the [FieldPath] in its marshaled form.
//
// The resulting [Mask] represents the field path "user.profile.email".
//
// To parse the string representation back into a [FieldPath], use the
// [Parse] function followed by [Mask.ToFieldPath] method:
//
//	m, err := mask.Parse(str)
//	if err != nil {
//	   // Handle error
//	}
//	fieldPath, err := m.ToFieldPath()
//	if err != nil {
//	   // Handle error
//	}
//
// Returns:
//   - string: A string representation of the path mask derived from the
//     [FieldPath].
//   - error: An error, if any, encountered during the marshaling process.
func (fp FieldPath) Marshal() (string, error) {
	m := fp.ToMask()
	return m.Marshal()
}

func (fp FieldPath) String() string {
	ret, err := fp.Marshal()
	if err != nil {
		return "FieldPath<not-marshalable " + err.Error() + ">"
	}
	return ret
}

// MarshalText wraps [FieldPath.Marshal] to produce []byte instead of string.
func (fp FieldPath) MarshalText() ([]byte, error) {
	ret, err := fp.Marshal()
	if err != nil {
		return nil, err
	}
	return []byte(ret), nil
}

// UnmarshalText parses text as [Mask] and converts it to [FieldPath].
func (fp *FieldPath) UnmarshalText(text []byte) error {
	if fp == nil {
		return errors.New("passed nil into unmarshal")
	}
	mask, err := Parse(string(text))
	if err != nil {
		return err
	}
	*fp, err = mask.ToFieldPath()
	return err
}
