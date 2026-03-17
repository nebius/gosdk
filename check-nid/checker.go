package checknid

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/nebius/gosdk/constants"
	"github.com/nebius/gosdk/nid"
	"github.com/nebius/gosdk/proto/fieldmask/mask"
	"github.com/nebius/gosdk/proto/nebius"
)

// SubfieldSettings is an alias for nebius.SubfieldSettings to simplify usage
// from generated SDK code.
type SubfieldSettings = nebius.SubfieldSettings
type NIDFieldSettings = nebius.NIDFieldSettings

// NIDCheckContext carries pre-parsed NID-related subfield settings for the
// current message level.
type NIDCheckContext struct {
	subfields []*compiledSubfieldSetting
}

// SubfieldPath is a compiled SelectMask-style path used by subfield settings.
type SubfieldPath struct {
	subMask *mask.Mask
}

type compiledSubfieldSetting struct {
	subPath        *SubfieldPath
	resourceTypes  []string
	metadataParent bool
}

const anyNIDResourceType = "*"

// TODO: DO NOT MERGE
var Abc = "is global to trigger linter"

// NewNIDCheckContext compiles method-level subfield settings used by NID
// checks. Non-NID subfield settings are ignored.
func NewNIDCheckContext(settings []*SubfieldSettings) *NIDCheckContext {
	return newNIDCheckContextFromCompiled(compileNIDSubfieldSettings(settings))
}

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

		if direct := directResourceSubfieldSetting(setting); direct != nil &&
			SubfieldSettingMatchesPath(direct.GetFieldPath(), targetPath) {
			resourceTypes, hasResourceSetting := getSubfieldResourceTypes(direct)
			if hasResourceSetting {
				return resourceTypes
			}
		}

		if parent := metadataParentResourceSubfieldSetting(setting); parent != nil &&
			SubfieldSettingMatchesPath(parent.GetFieldPath(), []string{metadataFieldName}) {
			resourceTypes, hasResourceSetting := getSubfieldResourceTypes(parent)
			if hasResourceSetting {
				return resourceTypes
			}
		}
	}

	return nil
}

// CheckMessageFields walks through the message recursively, validates all fields
// annotated with (nid) and fields matched by context subfield settings, and
// returns a path -> warning map for all failed preconditions.
func CheckMessageFields(msg proto.Message, checkCtx ...*NIDCheckContext) map[string]string {
	if isNilProtoMessage(msg) {
		return nil
	}

	warnings := map[string]string{}
	checkMessageNIDs(msg.ProtoReflect(), "", warnings, firstNIDCheckContext(checkCtx))

	if len(warnings) == 0 {
		return nil
	}

	return warnings
}

// CheckMetadataParentNID validates metadata.parent_id against allowed NID types.
// allowedTypes empty means any NID type is accepted.
// Returns a single failed precondition message (empty string if everything is valid).
func CheckMetadataParentNID(
	metadata proto.Message,
	allowedTypes []string,
	checkCtx ...*NIDCheckContext,
) string {
	if isNilProtoMessage(metadata) {
		return ""
	}

	md := metadata.ProtoReflect()
	parentIDField := md.Descriptor().Fields().ByName("parent_id")
	if parentIDField == nil || parentIDField.Kind() != protoreflect.StringKind {
		return ""
	}

	parentID := md.Get(parentIDField).String()
	if parentID == "" {
		return ""
	}

	if overrideTypes, ok := resolveFieldResources(parentIDField, firstNIDCheckContext(checkCtx)); ok {
		allowedTypes = overrideTypes
	}

	return ValidateNIDString(parentID, allowedTypes)
}

func checkMessageNIDs(
	msg protoreflect.Message,
	path string,
	warnings map[string]string,
	checkCtx *NIDCheckContext,
) {
	msg.Range(func(fd protoreflect.FieldDescriptor, val protoreflect.Value) bool {
		fieldPath := joinPath(path, string(fd.Name()))
		fieldCtx := nextNIDCheckContext(checkCtx, fd)

		if permittedTypes, ok := resolveFieldResources(fd, checkCtx); ok {
			validateNIDField(fd, val, fieldPath, permittedTypes, warnings)
		}

		switch {
		case fd.IsMap():
			valueDesc := fd.MapValue()
			if valueDesc.Kind() == protoreflect.MessageKind {
				val.Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
					checkMessageNIDs(v.Message(), mapPath(fieldPath, k.String()), warnings, fieldCtx)
					return true
				})
			}
		case fd.IsList():
			if fd.Kind() == protoreflect.MessageKind {
				list := val.List()
				for i := range list.Len() {
					checkMessageNIDs(list.Get(i).Message(), indexPath(fieldPath, i), warnings, fieldCtx)
				}
			}
		case fd.Kind() == protoreflect.MessageKind:
			if msg.Has(fd) {
				checkMessageNIDs(val.Message(), fieldPath, warnings, fieldCtx)
			}
		default:
			// nothing to do, not validatable type or recursive message
		}

		return true
	})
}

func validateNIDField(
	fd protoreflect.FieldDescriptor,
	val protoreflect.Value,
	path string,
	permittedTypes []string,
	warnings map[string]string,
) {
	switch {
	case fd.Kind() == protoreflect.StringKind && !fd.IsList():
		if warning := ValidateNIDString(val.String(), permittedTypes); warning != "" {
			warnings[path] = warning
		}
	case fd.IsList() && fd.Kind() == protoreflect.StringKind:
		list := val.List()
		for i := range list.Len() {
			if warning := ValidateNIDString(list.Get(i).String(), permittedTypes); warning != "" {
				warnings[indexPath(path, i)] = warning
			}
		}
	case fd.IsMap():
		valueDesc := fd.MapValue()
		if valueDesc.Kind() == protoreflect.StringKind {
			val.Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
				if warning := ValidateNIDString(v.String(), permittedTypes); warning != "" {
					warnings[mapPath(path, k.String())] = warning
				}
				return true
			})
		}
	}
}

func resolveFieldResources(fd protoreflect.FieldDescriptor, checkCtx *NIDCheckContext) ([]string, bool) {
	if checkCtx != nil {
		for _, subfield := range checkCtx.subfields {
			if subfield == nil || subfield.metadataParent {
				continue
			}
			_, ok := subfield.matchField(fd)
			if !ok {
				continue
			}
			return subfield.resourceTypes, true
		}
	}

	settings, ok := getNIDSettings(fd)
	if !ok {
		return nil, false
	}

	return getNIDFieldResources(settings), true
}

func getNIDFieldResources(settings *nebius.NIDFieldSettings) []string {
	if settings == nil {
		return nil
	}

	resources := settings.GetResource()

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

func parentSubfieldSetting(settings *nebius.NIDFieldSettings, fieldPath string) *SubfieldSettings {
	if !hasNIDParentResourceSetting(settings) || strings.TrimSpace(fieldPath) == "" {
		return nil
	}

	nidSettings := &SubfieldSettings{
		FieldPath: fieldPath,
		Nid: &NIDFieldSettings{
			Resource: cloneSliceIfPresent(
				settings.GetParentResource(),
				settings.ParentResource != nil,
			),
		},
	}

	return nidSettings
}

func directResourceSubfieldSetting(setting *SubfieldSettings) *SubfieldSettings {
	if setting == nil || !hasNIDResourceSetting(setting.GetNid()) {
		return nil
	}

	nidSettings := &SubfieldSettings{
		FieldPath: setting.GetFieldPath(),
		Nid: &NIDFieldSettings{
			Resource: cloneSliceIfPresent(
				setting.GetNid().GetResource(),
				setting.GetNid().Resource != nil,
			),
		},
	}

	return nidSettings
}

func metadataParentResourceSubfieldSetting(setting *SubfieldSettings) *SubfieldSettings {
	if setting == nil || !hasNIDParentResourceSetting(setting.GetNid()) {
		return nil
	}

	nidSettings := &SubfieldSettings{
		FieldPath: setting.GetFieldPath(),
		Nid: &NIDFieldSettings{
			Resource: cloneSliceIfPresent(
				setting.GetNid().GetParentResource(),
				setting.GetNid().ParentResource != nil,
			),
		},
	}

	return nidSettings
}

func cloneSliceIfPresent(values []string, present bool) []string {
	if !present {
		return nil
	}
	if len(values) == 0 {
		return []string{}
	}
	return append([]string{}, values...)
}

func nextNIDCheckContext(checkCtx *NIDCheckContext, fd protoreflect.FieldDescriptor) *NIDCheckContext {
	nextSubfields := inheritedSubfieldSettings(checkCtx, fd)

	// Field-level parent_resource rules are normalized into a synthetic
	// field_path=parent_id subfield setting. It must have higher priority than
	// message-level subfield_settings, but lower priority than method-level
	// context.
	if seeded := getSyntheticFieldSubfieldSettings(fd); len(seeded) > 0 {
		nextSubfields = append(nextSubfields, seeded...)
	}

	fieldSubfields := getFieldSubfieldSettings(fd)
	if len(fieldSubfields) > 0 {
		nextSubfields = append(nextSubfields, fieldSubfields...)
	}

	return newNIDCheckContextFromCompiled(nextSubfields)
}

func inheritedSubfieldSettings(
	checkCtx *NIDCheckContext,
	fd protoreflect.FieldDescriptor,
) []*compiledSubfieldSetting {
	if checkCtx == nil || len(checkCtx.subfields) == 0 {
		return nil
	}

	nextSubfields := make([]*compiledSubfieldSetting, 0, len(checkCtx.subfields))
	for _, subfield := range checkCtx.subfields {
		nextSubfields = append(nextSubfields, advanceSubfieldSetting(fd, subfield)...)
	}

	return nextSubfields
}

func advanceSubfieldSetting(
	fd protoreflect.FieldDescriptor,
	subfield *compiledSubfieldSetting,
) []*compiledSubfieldSetting {
	subPath, ok := subfield.matchField(fd)
	if !ok {
		return nil
	}

	if subfield.metadataParent {
		return advanceMetadataParentSubfieldSetting(fd, subfield, subPath)
	}

	return []*compiledSubfieldSetting{{
		subPath:       subPath,
		resourceTypes: subfield.resourceTypes,
	}}
}

func advanceMetadataParentSubfieldSetting(
	fd protoreflect.FieldDescriptor,
	subfield *compiledSubfieldSetting,
	subPath *SubfieldPath,
) []*compiledSubfieldSetting {
	if subPath != nil && subPath.IsEmpty() {
		if seeded := metadataParentContextSetting(fd, subfield.resourceTypes); seeded != nil {
			return []*compiledSubfieldSetting{seeded}
		}
		return nil
	}

	return []*compiledSubfieldSetting{{
		subPath:        subPath,
		resourceTypes:  subfield.resourceTypes,
		metadataParent: true,
	}}
}

func newNIDCheckContextFromCompiled(subfields []*compiledSubfieldSetting) *NIDCheckContext {
	if len(subfields) == 0 {
		return nil
	}
	return &NIDCheckContext{
		subfields: subfields,
	}
}

func getFieldSubfieldSettings(fd protoreflect.FieldDescriptor) []*compiledSubfieldSetting {
	settings, ok := getSubfieldSettings(fd)
	if !ok {
		return nil
	}

	return compileNIDSubfieldSettings(settings)
}

func getSyntheticFieldSubfieldSettings(fd protoreflect.FieldDescriptor) []*compiledSubfieldSetting {
	if !isResourceMetadataFieldDescriptor(fd) {
		return nil
	}

	settings, ok := getNIDSettings(fd)
	if !ok || settings == nil {
		return nil
	}

	parentSetting := parentSubfieldSetting(settings, "parent_id")
	if parentSetting == nil {
		return nil
	}

	return compileNIDSubfieldSettings([]*SubfieldSettings{parentSetting})
}

func getSubfieldSettings(fd protoreflect.FieldDescriptor) ([]*SubfieldSettings, bool) {
	opts, ok := fd.Options().(*descriptorpb.FieldOptions)
	if !ok {
		return nil, false
	}

	if !proto.HasExtension(opts, nebius.E_SubfieldSettings) {
		return nil, false
	}

	ext, ok := proto.GetExtension(opts, nebius.E_SubfieldSettings).([]*SubfieldSettings)
	if !ok || len(ext) == 0 {
		return nil, false
	}

	return ext, true
}

func compileNIDSubfieldSettings(settings []*SubfieldSettings) []*compiledSubfieldSetting {
	compiled := make([]*compiledSubfieldSetting, 0, len(settings))

	for _, raw := range settings {
		compiled = appendCompiledSubfieldSetting(compiled, directResourceSubfieldSetting(raw), false)
		compiled = appendCompiledSubfieldSetting(compiled, metadataParentResourceSubfieldSetting(raw), true)
	}

	return compiled
}

func appendCompiledSubfieldSetting(
	compiled []*compiledSubfieldSetting,
	setting *SubfieldSettings,
	metadataParent bool,
) []*compiledSubfieldSetting {
	if setting == nil {
		return compiled
	}

	resourceTypes, hasResourceSetting := getSubfieldResourceTypes(setting)
	if !hasResourceSetting {
		return compiled
	}

	fieldPath := strings.TrimSpace(setting.GetFieldPath())
	if fieldPath == "" {
		return compiled
	}

	subPath, err := ParseSubfieldPath(fieldPath)
	if err != nil {
		return compiled
	}

	return append(compiled, &compiledSubfieldSetting{
		subPath:        subPath,
		resourceTypes:  resourceTypes,
		metadataParent: metadataParent,
	})
}

func getSubfieldResourceTypes(setting *SubfieldSettings) ([]string, bool) {
	if setting == nil {
		return nil, false
	}

	nidSettings := setting.GetNid()
	return getNIDFieldResources(nidSettings), hasNIDResourceSetting(nidSettings)
}

func mustSingleFieldSubfieldPath(fieldName string) *SubfieldPath {
	ret, err := ParseSubfieldPath(fieldName)
	if err != nil {
		panic(err)
	}
	return ret
}

func (c *compiledSubfieldSetting) matchField(fd protoreflect.FieldDescriptor) (*SubfieldPath, bool) {
	if c == nil || c.subPath == nil || fd == nil {
		return nil, false
	}
	return c.subPath.MatchField(fd)
}

func metadataParentContextSetting(
	fd protoreflect.FieldDescriptor,
	resourceTypes []string,
) *compiledSubfieldSetting {
	if !isResourceMetadataFieldDescriptor(fd) {
		return nil
	}

	return &compiledSubfieldSetting{
		subPath:       mustSingleFieldSubfieldPath(constants.FieldParentID),
		resourceTypes: resourceTypes,
	}
}

func isResourceMetadataFieldDescriptor(fd protoreflect.FieldDescriptor) bool {
	if fd == nil || fd.IsList() || fd.IsMap() || fd.Kind() != protoreflect.MessageKind {
		return false
	}
	if fd.Name() != protoreflect.Name(constants.FieldMetadata) {
		return false
	}
	msg := fd.Message()
	return msg != nil && string(msg.FullName()) == constants.MetadataMessageFullName
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

func firstNIDCheckContext(checkCtx []*NIDCheckContext) *NIDCheckContext {
	for _, ctx := range checkCtx {
		if ctx != nil {
			return ctx
		}
	}
	return nil
}

func getNIDSettings(fd protoreflect.FieldDescriptor) (*nebius.NIDFieldSettings, bool) {
	opts, ok := fd.Options().(*descriptorpb.FieldOptions)
	if !ok {
		return nil, false
	}

	if !proto.HasExtension(opts, nebius.E_Nid) {
		return nil, false
	}

	ext, ok := proto.GetExtension(opts, nebius.E_Nid).(*nebius.NIDFieldSettings)
	if !ok || ext == nil {
		return nil, false
	}

	return ext, true
}

func ValidateNIDString(value string, allowedTypes []string) string {
	if value == "" {
		return ""
	}
	if len(allowedTypes) == 1 && allowedTypes[0] == anyNIDResourceType {
		allowedTypes = nil
	}

	parsed, err := nid.Parse(value)
	if err != nil {
		return fmt.Sprintf("value %q is not a valid Nebius ID: %v", value, err)
	}

	if len(allowedTypes) == 0 {
		return ""
	}

	if slices.Contains(allowedTypes, string(parsed.Type)) {
		return ""
	}

	allowedNids := make([]string, len(allowedTypes))
	for i, t := range allowedTypes {
		allowed, errX := nid.NewWithRoutingCode(
			nid.Type(t),
			parsed.RoutingCode,
			"xxx",
		)
		if errX != nil {
			allowedNids[i] = t + strings.TrimPrefix(t, parsed.Type.String())
		}
		allowedNids[i] = allowed.String()
	}

	return fmt.Sprintf(
		"Nebius ID %q has invalid type %q, allowed IDs: %s",
		value,
		parsed.Type,
		strings.Join(allowedNids, ", "),
	)
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

func joinPath(path, elem string) string {
	if path == "" {
		return elem
	}
	return path + "." + elem
}

func indexPath(path string, index int) string {
	if path == "" {
		return fmt.Sprintf("[%d]", index)
	}
	return fmt.Sprintf("%s[%d]", path, index)
}

func mapPath(path, key string) string {
	if path == "" {
		return fmt.Sprintf("[%q]", key)
	}
	return fmt.Sprintf("%s[%q]", path, key)
}

func isNilProtoMessage(msg proto.Message) bool {
	if msg == nil {
		return true
	}

	rv := reflect.ValueOf(msg)
	return rv.Kind() == reflect.Pointer && rv.IsNil()
}
