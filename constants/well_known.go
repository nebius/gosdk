package constants

import "google.golang.org/protobuf/reflect/protoreflect"

const (
	MethodGet       = "Get"
	MethodGetByName = "GetByName"
	MethodList      = "List"
	MethodCreate    = "Create"
	MethodDelete    = "Delete"
	MethodUpdate    = "Update"
)

const (
	FieldID              = "id"
	FieldName            = "name"
	FieldParentID        = "parent_id"
	FieldResourceVersion = "resource_version"
	FieldCreatedAt       = "created_at"
	FieldUpdatedAt       = "updated_at"
	FieldLabels          = "labels"

	FieldSpec        = "spec"
	FieldStatus      = "status"
	FieldMetadata    = "metadata"
	FieldWellKnownID = "well_known_id"
)

const (
	AlphaOperationMessageFilePath = "nebius/common/v1alpha1/operation.proto"
	AlphaOperationMessageFullName = "nebius.common.v1alpha1.Operation"
	OperationMessageFilePath      = "nebius/common/v1/operation.proto"
	OperationMessageFullName      = "nebius.common.v1.Operation"
	MetadataMessageFullName       = "nebius.common.v1.ResourceMetadata"
	GetByNameRequestFullName      = "nebius.common.v1.GetByNameRequest"
	ServiceSuffix                 = "Service"
)

var MetadataUnwrapped = []protoreflect.Name{
	FieldID, FieldName, FieldParentID, FieldResourceVersion, FieldCreatedAt,
	FieldUpdatedAt, FieldLabels,
}

var ReservedFields = append([]protoreflect.Name{
	FieldMetadata, FieldStatus, FieldWellKnownID,
}, MetadataUnwrapped...)
