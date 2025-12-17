package checknid

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/nebius/gosdk/nid"
	"github.com/nebius/gosdk/proto/nebius"
)

// CheckMessageFields walks through the message recursively, validates all fields annotated
// with (nid), and returns a path -> warning map for all failed preconditions.
func CheckMessageFields(msg proto.Message) map[string]string {
	if isNilProtoMessage(msg) {
		return nil
	}

	warnings := map[string]string{}
	checkMessageNIDs(msg.ProtoReflect(), "", warnings)

	if len(warnings) == 0 {
		return nil
	}

	return warnings
}

// CheckMetadataParentNID validates metadata.parent_id against allowed NID types.
// allowedTypes empty means any NID type is accepted.
// Returns a single failed precondition message (empty string if everything is valid).
func CheckMetadataParentNID(metadata proto.Message, allowedTypes []string) string {
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

	return ValidateNIDString(parentID, allowedTypes)
}

func checkMessageNIDs(msg protoreflect.Message, path string, warnings map[string]string) {
	msg.Range(func(fd protoreflect.FieldDescriptor, val protoreflect.Value) bool {
		fieldPath := joinPath(path, string(fd.Name()))

		if settings, ok := getNIDSettings(fd); ok {
			validateAnnotatedField(fd, val, fieldPath, settings, warnings)
		}

		switch {
		case fd.IsMap():
			valueDesc := fd.MapValue()
			if valueDesc.Kind() == protoreflect.MessageKind {
				val.Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
					checkMessageNIDs(v.Message(), mapPath(fieldPath, k.String()), warnings)
					return true
				})
			}
		case fd.IsList():
			if fd.Kind() == protoreflect.MessageKind {
				list := val.List()
				for i := range list.Len() {
					checkMessageNIDs(list.Get(i).Message(), indexPath(fieldPath, i), warnings)
				}
			}
		case fd.Kind() == protoreflect.MessageKind:
			if msg.Has(fd) {
				checkMessageNIDs(val.Message(), fieldPath, warnings)
			}
		default:
			// nothing to do, not validatable type or recursive message
		}

		return true
	})
}

func validateAnnotatedField(
	fd protoreflect.FieldDescriptor,
	val protoreflect.Value,
	path string,
	settings *nebius.NIDFieldSettings,
	warnings map[string]string,
) {
	switch {
	case fd.Kind() == protoreflect.StringKind && !fd.IsList():
		if warning := ValidateNIDString(val.String(), settings.GetResource()); warning != "" {
			warnings[path] = warning
		}
	case fd.IsList() && fd.Kind() == protoreflect.StringKind:
		list := val.List()
		for i := range list.Len() {
			if warning := ValidateNIDString(list.Get(i).String(), settings.GetResource()); warning != "" {
				warnings[indexPath(path, i)] = warning
			}
		}
	case fd.IsMap():
		valueDesc := fd.MapValue()
		if valueDesc.Kind() == protoreflect.StringKind {
			val.Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
				if warning := ValidateNIDString(v.String(), settings.GetResource()); warning != "" {
					warnings[mapPath(path, k.String())] = warning
				}
				return true
			})
		}
	}
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
	return rv.Kind() == reflect.Ptr && rv.IsNil()
}
