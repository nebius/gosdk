package checknid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"

	checknid "github.com/nebius/gosdk/check-nid"
	"github.com/nebius/gosdk/proto/nebius"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

func TestCheckMetadataParentNID(t *testing.T) {
	t.Parallel()

	metadata := &common.ResourceMetadata{
		ParentId: "project-e0tabc",
	}

	assert.Empty(t, checknid.CheckMetadataParentNID(metadata, []string{"project"}))
	assert.Contains(t, checknid.CheckMetadataParentNID(metadata, []string{"folder"}), "folder")

	metadata.ParentId = ""
	assert.Empty(t, checknid.CheckMetadataParentNID(metadata, []string{"project"}))
}

func TestCheckMetadataParentNIDWithWildcardAny(t *testing.T) {
	t.Parallel()

	metadata := &common.ResourceMetadata{
		ParentId: "project-e0tabc",
	}

	assert.Empty(t, checknid.CheckMetadataParentNID(metadata, []string{"*"}))
	assert.Contains(
		t,
		checknid.CheckMetadataParentNID(&common.ResourceMetadata{ParentId: "bad"}, []string{"*"}),
		"not a valid Nebius ID",
	)
}

func TestCheckMetadataParentNIDWithContext(t *testing.T) {
	t.Parallel()

	metadata := &common.ResourceMetadata{
		ParentId: "project-e0tabc",
	}

	checkCtx := checknid.NewNIDCheckContext([]*checknid.SubfieldSettings{
		nidSubfield("parent_id", []string{"folder"}),
	})

	assert.Contains(
		t,
		checknid.CheckMetadataParentNID(metadata, []string{"project"}, checkCtx),
		"folder",
	)
}

func TestCheckMetadataParentNIDWithWildcardContext(t *testing.T) {
	t.Parallel()

	metadata := &common.ResourceMetadata{
		ParentId: "project-e0tabc",
	}

	checkCtx := checknid.NewNIDCheckContext([]*checknid.SubfieldSettings{
		nidSubfield("parent_id", []string{"*"}),
	})

	assert.Empty(t, checknid.CheckMetadataParentNID(metadata, []string{"folder"}, checkCtx))
}

func TestCheckNIDFieldsParentResourceContextIgnoresNonMetadataField(t *testing.T) {
	t.Parallel()

	request := buildCustomParentIDRequest(t, "project-e0tabc")
	checkCtx := checknid.NewNIDCheckContext([]*checknid.SubfieldSettings{
		parentResourceSubfield("custom", []string{"folder"}),
	})

	warnings := checknid.CheckMessageFields(request, checkCtx)
	require.Empty(t, warnings)
}

func TestMetadataParentAutoFillTypes(t *testing.T) {
	t.Parallel()

	settings := []*checknid.SubfieldSettings{
		parentResourceSubfield("spec", []string{"folder"}),
		parentResourceSubfield("metadata", []string{"project"}),
		nidSubfield("metadata.parent_id.extra", []string{"folder"}),
		nidSubfield("metadata.parent_id", []string{"organization"}),
	}

	require.Equal(
		t,
		[]string{"project"},
		checknid.MetadataParentAutoFillTypes(settings, "metadata"),
	)
}

func TestSubfieldSettingMatchesPath(t *testing.T) {
	t.Parallel()

	require.True(
		t,
		checknid.SubfieldSettingMatchesPath("metadata.parent_id", []string{"metadata", "parent_id"}),
	)
	require.False(
		t,
		checknid.SubfieldSettingMatchesPath("metadata.parent_id.extra", []string{"metadata", "parent_id"}),
	)
}

func TestCheckNIDFieldsMethodContextOverridesFieldSubfieldSettingsForNestedMapAndList(t *testing.T) {
	t.Parallel()

	request := buildNestedMapListRequest(t, "methodlist-e0tabc", "methodmap-e0tabc")
	checkCtx := checknid.NewNIDCheckContext([]*checknid.SubfieldSettings{
		nidSubfield("wrap.items.*.id", []string{"methodlist"}),
		nidSubfield("wrap.by_key.*.id", []string{"methodmap"}),
	})

	warnings := checknid.CheckMessageFields(request, checkCtx)
	require.Empty(t, warnings)
}

func TestCheckNIDFieldsMethodContextFallsBackToFieldSubfieldSettingsForNestedMapAndList(t *testing.T) {
	t.Parallel()

	request := buildNestedMapListRequest(t, "fieldlist-e0tabc", "fieldmap-e0tabc")
	checkCtx := checknid.NewNIDCheckContext([]*checknid.SubfieldSettings{
		nidSubfield("wrap.items.*.id.extra", []string{"methodlist"}),
		nidSubfield("wrap.by_key.*.id.extra", []string{"methodmap"}),
	})

	warnings := checknid.CheckMessageFields(request, checkCtx)
	require.Empty(t, warnings)
}

func buildNestedMapListRequest(t *testing.T, listID, mapID string) proto.Message {
	t.Helper()

	wrapFieldOptions := &descriptorpb.FieldOptions{}
	mustSetExtension(
		t,
		wrapFieldOptions,
		nebius.E_SubfieldSettings,
		[]*nebius.SubfieldSettings{
			{
				FieldPath: "items.*.id",
				Nid: &nebius.NIDFieldSettings{
					Resource: []string{"fieldlist"},
				},
			},
			{
				FieldPath: "by_key.*.id",
				Nid: &nebius.NIDFieldSettings{
					Resource: []string{"fieldmap"},
				},
			},
		},
	)

	file := &descriptorpb.FileDescriptorProto{
		Syntax:  proto.String("proto3"),
		Name:    proto.String("check_nid_nested_map_list_test.proto"),
		Package: proto.String("test"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: proto.String("Leaf"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   proto.String("id"),
						Number: proto.Int32(1),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					},
				},
			},
			{
				Name: proto.String("Wrapper"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:     proto.String("items"),
						Number:   proto.Int32(1),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
						TypeName: proto.String(".test.Leaf"),
					},
					{
						Name:     proto.String("by_key"),
						Number:   proto.Int32(2),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
						TypeName: proto.String(".test.Wrapper.ByKeyEntry"),
					},
				},
				NestedType: []*descriptorpb.DescriptorProto{
					{
						Name: proto.String("ByKeyEntry"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:   proto.String("key"),
								Number: proto.Int32(1),
								Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
							},
							{
								Name:     proto.String("value"),
								Number:   proto.Int32(2),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
								TypeName: proto.String(".test.Leaf"),
							},
						},
						Options: &descriptorpb.MessageOptions{
							MapEntry: proto.Bool(true),
						},
					},
				},
			},
			{
				Name: proto.String("Request"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:     proto.String("wrap"),
						Number:   proto.Int32(1),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
						TypeName: proto.String(".test.Wrapper"),
						Options:  wrapFieldOptions,
					},
				},
			},
		},
	}

	files, err := protodesc.NewFiles(&descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{file},
	})
	require.NoError(t, err)

	requestDesc := mustFindMessageDescriptor(t, files, "test.Request")
	requestMsg := dynamicpb.NewMessage(requestDesc)

	wrapField := requestDesc.Fields().ByName("wrap")
	require.NotNil(t, wrapField)
	wrapMsg := dynamicpb.NewMessage(wrapField.Message())

	itemsField := wrapField.Message().Fields().ByName("items")
	require.NotNil(t, itemsField)
	itemsList := wrapMsg.Mutable(itemsField).List()
	itemsList.Append(protoreflect.ValueOfMessage(newLeafMessage(t, itemsField.Message(), listID)))

	byKeyField := wrapField.Message().Fields().ByName("by_key")
	require.NotNil(t, byKeyField)
	byKeyMap := wrapMsg.Mutable(byKeyField).Map()
	byKeyMap.Set(
		protoreflect.ValueOfString("k").MapKey(),
		protoreflect.ValueOfMessage(newLeafMessage(t, byKeyField.MapValue().Message(), mapID)),
	)

	requestMsg.Set(wrapField, protoreflect.ValueOfMessage(wrapMsg))
	return requestMsg
}

func nidSubfield(fieldPath string, resources []string) *checknid.SubfieldSettings {
	return &checknid.SubfieldSettings{
		FieldPath: fieldPath,
		Nid: &checknid.NIDFieldSettings{
			Resource: resources,
		},
	}
}

func parentResourceSubfield(fieldPath string, resources []string) *checknid.SubfieldSettings {
	return &checknid.SubfieldSettings{
		FieldPath: fieldPath,
		Nid: &checknid.NIDFieldSettings{
			ParentResource: resources,
		},
	}
}

func buildCustomParentIDRequest(t *testing.T, parentID string) proto.Message {
	t.Helper()

	file := &descriptorpb.FileDescriptorProto{
		Syntax:  proto.String("proto3"),
		Name:    proto.String("check_nid_custom_parent_id_test.proto"),
		Package: proto.String("test"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: proto.String("CustomMetadata"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   proto.String("parent_id"),
						Number: proto.Int32(1),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					},
				},
			},
			{
				Name: proto.String("Request"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:     proto.String("custom"),
						Number:   proto.Int32(1),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
						TypeName: proto.String(".test.CustomMetadata"),
					},
				},
			},
		},
	}

	files, err := protodesc.NewFiles(&descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{file},
	})
	require.NoError(t, err)

	requestDesc := mustFindMessageDescriptor(t, files, "test.Request")
	request := dynamicpb.NewMessage(requestDesc)

	customField := requestDesc.Fields().ByName("custom")
	require.NotNil(t, customField)
	custom := dynamicpb.NewMessage(customField.Message())
	parentIDField := customField.Message().Fields().ByName("parent_id")
	require.NotNil(t, parentIDField)
	custom.Set(parentIDField, protoreflect.ValueOfString(parentID))
	request.Set(customField, protoreflect.ValueOfMessage(custom))

	return request
}

func newLeafMessage(
	t *testing.T,
	leafDesc protoreflect.MessageDescriptor,
	id string,
) protoreflect.Message {
	t.Helper()

	leaf := dynamicpb.NewMessage(leafDesc)
	idField := leafDesc.Fields().ByName("id")
	require.NotNil(t, idField)
	leaf.Set(idField, protoreflect.ValueOfString(id))
	return leaf
}

func mustFindMessageDescriptor(
	t *testing.T,
	files *protoregistry.Files,
	name protoreflect.FullName,
) protoreflect.MessageDescriptor {
	t.Helper()

	desc, err := files.FindDescriptorByName(name)
	require.NoError(t, err)
	msg, ok := desc.(protoreflect.MessageDescriptor)
	require.True(t, ok)
	return msg
}

func mustSetExtension(
	t *testing.T,
	opts proto.Message,
	ext protoreflect.ExtensionType,
	val any,
) {
	t.Helper()

	type setterWithErr func(proto.Message, protoreflect.ExtensionType, any) error
	type setterNoErr func(proto.Message, protoreflect.ExtensionType, any)

	switch fn := any(proto.SetExtension).(type) {
	case setterWithErr:
		require.NoError(t, fn(opts, ext, val))
	case setterNoErr:
		fn(opts, ext, val)
	default:
		proto.SetExtension(opts, ext, val)
	}
}
