package protovalidate

import (
	"fmt"
	"strconv"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"

	"github.com/nebius/gosdk/proto/fieldmask/mask"
)

func FieldPathFromProto(protoPath *validate.FieldPath) (*mask.FieldPath, error) {
	if protoPath == nil {
		return nil, nil
	}

	ret := make(mask.FieldPath, 0, len(protoPath.GetElements())*2)
	for i, element := range protoPath.GetElements() {
		if element == nil {
			return nil, fmt.Errorf("field path element %d is nil", i)
		}

		fieldName := element.GetFieldName()
		if fieldName == "" {
			if !element.HasFieldNumber() {
				return nil, fmt.Errorf("field path element %d has no field name or field number", i)
			}
			fieldName = strconv.FormatInt(int64(element.GetFieldNumber()), 10)
		}
		ret = append(ret, mask.FieldKey(fieldName))

		switch subscript := element.GetSubscript().(type) {
		case nil:
		case *validate.FieldPathElement_Index:
			ret = append(ret, mask.FieldKey(strconv.FormatUint(subscript.Index, 10)))
		case *validate.FieldPathElement_BoolKey:
			ret = append(ret, mask.FieldKey(strconv.FormatBool(subscript.BoolKey)))
		case *validate.FieldPathElement_IntKey:
			ret = append(ret, mask.FieldKey(strconv.FormatInt(subscript.IntKey, 10)))
		case *validate.FieldPathElement_UintKey:
			ret = append(ret, mask.FieldKey(strconv.FormatUint(subscript.UintKey, 10)))
		case *validate.FieldPathElement_StringKey:
			ret = append(ret, mask.FieldKey(subscript.StringKey))
		default:
			return nil, fmt.Errorf("field path element %d has unsupported subscript type %T", i, subscript)
		}
	}

	return &ret, nil
}
