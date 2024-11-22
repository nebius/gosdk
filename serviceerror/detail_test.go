package serviceerror_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	"github.com/nebius/gosdk/serviceerror"
)

func TestDetailNoUnknowns(t *testing.T) {
	t.Parallel()

	seProto := (&common.ServiceError{}).ProtoReflect()
	details := seProto.Descriptor().Oneofs().ByName("details")
	require.NotNil(t, details)

	for i := range details.Fields().Len() {
		field := details.Fields().Get(i)
		val := seProto.NewField(field)
		seProto.Set(field, val)
		myAny := &anypb.Any{}
		err := anypb.MarshalFrom(myAny, seProto.Interface(), proto.MarshalOptions{})
		require.NoError(t, err)
		se, err := serviceerror.NewDetailFromAny(myAny)
		require.NoError(t, err)
		_, ok := se.(*serviceerror.Unknown)
		assert.False(t, ok, "Service error type not implemented", field.Name())
	}
}
