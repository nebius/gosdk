package checknid

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/nid"
	"github.com/nebius/gosdk/proto/fieldmask/mask"
	"github.com/nebius/gosdk/proto/nebius"
)

// SubfieldSettings is an alias for nebius.SubfieldSettings to simplify usage
// from generated SDK code.
type SubfieldSettings = nebius.SubfieldSettings
type NIDFieldSettings = nebius.NIDFieldSettings

// SubfieldPath is a compiled SelectMask-style path used by subfield settings.
type SubfieldPath struct {
	subMask *mask.Mask
}

const anyNIDResourceType = "*"

// ParseSubfieldPath parses a SelectMask-style field path used by subfield
// settings.
func ParseSubfieldPath(fieldPath string) (*SubfieldPath, error) {
	fieldPath = strings.TrimSpace(fieldPath)
	if fieldPath == "" {
		return nil, errors.New("path must not be empty")
	}

	subMask, err := mask.Parse(fieldPath)
	if err != nil || subMask == nil {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("path must not be empty")
	}

	return &SubfieldPath{subMask: subMask}, nil
}

// SubfieldSettingMatchesPath reports whether fieldPath matches a concrete target
// path in a message. Scalar targets must match terminal paths only.
func SubfieldSettingMatchesPath(fieldPath string, targetPath []string) bool {
	if len(targetPath) == 0 {
		return false
	}
	if strings.TrimSpace(fieldPath) == "" {
		return false
	}

	subMask, err := mask.Parse(strings.TrimSpace(fieldPath))
	if err != nil || subMask == nil {
		return false
	}

	for _, part := range targetPath {
		if subMask.IsEmpty() {
			return true
		}

		next, subMaskErr := subMask.GetSubMask(mask.FieldKey(part))
		if subMaskErr != nil || next == nil {
			return false
		}
		subMask = next
	}

	return subMask.IsEmpty()
}

// MatchField reports whether the path can match the provided field and returns
// the remaining path after this field.
func (p *SubfieldPath) MatchField(fd protoreflect.FieldDescriptor) (*SubfieldPath, bool) {
	if p == nil || p.subMask == nil || fd == nil {
		return nil, false
	}

	if p.subMask.IsEmpty() {
		return p, true
	}

	subMask, err := p.subMask.GetSubMask(mask.FieldKey(fd.Name()))
	if err != nil || subMask == nil {
		return nil, false
	}

	if subMask.IsEmpty() {
		return &SubfieldPath{subMask: subMask}, true
	}

	switch {
	case fd.IsMap():
		nextPath := newCollectionElementPath(subMask)
		if fd.MapValue().Kind() == protoreflect.MessageKind {
			return nextPath, true
		}
		if isCollectionElementMask(subMask) {
			return nextPath, true
		}
		return nil, false
	case fd.IsList():
		nextPath := newCollectionElementPath(subMask)
		if fd.Kind() == protoreflect.MessageKind {
			return nextPath, true
		}
		if fd.Kind() == protoreflect.StringKind && isCollectionElementMask(subMask) {
			return nextPath, true
		}
		return nil, false
	case fd.Kind() == protoreflect.MessageKind:
		return &SubfieldPath{subMask: subMask}, true
	default:
		// Scalar fields (including string) must match terminal paths only.
		return nil, false
	}
}

// IsEmpty reports whether the path is fully consumed.
func (p *SubfieldPath) IsEmpty() bool {
	return p != nil && p.subMask != nil && p.subMask.IsEmpty()
}

// Marshal converts the remaining path back into SelectMask format.
func (p *SubfieldPath) Marshal() (string, error) {
	if p == nil || p.subMask == nil || p.subMask.IsEmpty() {
		return "*", nil
	}

	serialized, err := p.subMask.Marshal()
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(serialized) == "" {
		return "*", nil
	}
	return serialized, nil
}

// ValidateSubfieldPathForMessage validates a subfield path against a message.
func ValidateSubfieldPathForMessage(fieldPath string, msg protoreflect.MessageDescriptor) error {
	subPath, err := ParseSubfieldPath(fieldPath)
	if err != nil {
		return err
	}
	if msg == nil {
		return errors.New("message must not be nil")
	}
	if subfieldPathMatchesMessage(subPath, msg) {
		return nil
	}
	return fmt.Errorf("path %q does not match any subfield of %s", fieldPath, msg.FullName())
}

// ValidateSubfieldPathForField validates a subfield path against a field's
// contained value.
func ValidateSubfieldPathForField(fieldPath string, fd protoreflect.FieldDescriptor) error {
	subPath, err := ParseSubfieldPath(fieldPath)
	if err != nil {
		return err
	}
	if fd == nil {
		return errors.New("field must not be nil")
	}
	if subfieldPathMatchesFieldContainer(subPath, fd) {
		return nil
	}
	return fmt.Errorf("path %q does not match any subfield of %s", fieldPath, fd.FullName())
}

// MetadataParentAutoFillTypes returns the first matching allowed types for
// metadata.parent_id inferred from subfield settings. Nil means no restriction
// (or no matching setting).
func MetadataParentAutoFillTypes(
	settings []*SubfieldSettings,
	metadataFieldName string,
) []string {
	if strings.TrimSpace(metadataFieldName) == "" {
		return nil
	}

	targetPath := []string{metadataFieldName, "parent_id"}
	for _, setting := range settings {
		if setting == nil {
			continue
		}
		nidSettings := setting.GetNid()

		if hasNIDResourceSetting(nidSettings) &&
			SubfieldSettingMatchesPath(setting.GetFieldPath(), targetPath) {
			return getNIDFieldResources(nidSettings)
		}

		if hasNIDParentResourceSetting(nidSettings) &&
			SubfieldSettingMatchesPath(setting.GetFieldPath(), []string{metadataFieldName}) {
			return getNIDParentFieldResources(nidSettings)
		}
	}

	return nil
}

func getNIDFieldResources(settings *nebius.NIDFieldSettings) []string {
	if settings == nil {
		return nil
	}

	resources := append([]string{}, settings.GetResource()...)

	return deduplicate(resources)
}

func getNIDParentFieldResources(settings *nebius.NIDFieldSettings) []string {
	if settings == nil {
		return nil
	}

	resources := append([]string{}, settings.GetParentResource()...)

	return deduplicate(resources)
}

func hasNIDResourceSetting(settings *nebius.NIDFieldSettings) bool {
	if settings == nil {
		return false
	}

	hasResourceSetting := settings.Resource != nil
	return hasResourceSetting
}

func hasNIDParentResourceSetting(settings *nebius.NIDFieldSettings) bool {
	if settings == nil {
		return false
	}

	hasParentResourceSetting := settings.ParentResource != nil
	return hasParentResourceSetting
}

func isCollectionElementMask(subMask *mask.Mask) bool {
	if subMask == nil || subMask.IsEmpty() {
		return false
	}

	hasElements := false
	if subMask.Any != nil {
		if !subMask.Any.IsEmpty() {
			return false
		}
		hasElements = true
	}

	for _, child := range subMask.FieldParts {
		if child != nil && !child.IsEmpty() {
			return false
		}
		hasElements = true
	}

	return hasElements
}

func newCollectionElementPath(subMask *mask.Mask) *SubfieldPath {
	if subMask == nil {
		return nil
	}
	if subMask.Any != nil {
		return &SubfieldPath{subMask: subMask.Any}
	}
	return &SubfieldPath{subMask: subMask}
}

func subfieldPathMatchesMessage(subPath *SubfieldPath, msg protoreflect.MessageDescriptor) bool {
	if subPath == nil || msg == nil {
		return false
	}

	fields := msg.Fields()
	for i := range fields.Len() {
		field := fields.Get(i)
		remaining, ok := subPath.MatchField(field)
		if !ok {
			continue
		}
		if remaining.IsEmpty() {
			return true
		}
		if subfieldPathMatchesFieldContainer(remaining, field) {
			return true
		}
	}

	return false
}

func subfieldPathMatchesFieldContainer(subPath *SubfieldPath, fd protoreflect.FieldDescriptor) bool {
	if subPath == nil || fd == nil {
		return false
	}

	switch {
	case fd.IsMap():
		value := fd.MapValue()
		if value.Kind() == protoreflect.MessageKind {
			return subfieldPathMatchesMessage(subPath, value.Message())
		}
		return subPathMatchesCollectionScalar(subPath)
	case fd.IsList():
		if fd.Kind() == protoreflect.MessageKind {
			return subfieldPathMatchesMessage(subPath, fd.Message())
		}
		return subPathMatchesCollectionScalar(subPath)
	case fd.Kind() == protoreflect.MessageKind:
		return subfieldPathMatchesMessage(subPath, fd.Message())
	default:
		return false
	}
}

func subPathMatchesCollectionScalar(subPath *SubfieldPath) bool {
	if subPath == nil || subPath.subMask == nil {
		return false
	}
	return isCollectionElementMask(subPath.subMask)
}

// IsNIDAllowedForAutoFill reports whether value is a valid Nebius ID whose
// type can be used for parent autofill. It does not validate user requests.
func IsNIDAllowedForAutoFill(value string, allowedTypes []string) bool {
	parsed, err := nid.Parse(value)
	if err != nil {
		return false
	}
	return len(allowedTypes) == 0 ||
		(len(allowedTypes) == 1 && allowedTypes[0] == anyNIDResourceType) ||
		slices.Contains(allowedTypes, string(parsed.Type))
}

func deduplicate(values []string) []string {
	if len(values) == 0 {
		return nil
	}

	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))

	for _, value := range values {
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}

	return result
}
