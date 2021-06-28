package serializer_test

import (
	"testing"

	"github.com/azcov/grpc-techschool/pb"
	"github.com/azcov/grpc-techschool/sample"
	"github.com/azcov/grpc-techschool/serializer"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../temp/laptop.bin"
	jsonFile := "../temp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))
}
