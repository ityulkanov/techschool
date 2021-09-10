package serializer_test

import (
	"fmt"
	"github.com/ityulkanov/techschool/pb"
	"github.com/ityulkanov/techschool/sample"
	"github.com/ityulkanov/techschool/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	fmt.Println("serialization completed")
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}

	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	fmt.Println("read file completed ")
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJsonFormat(laptop1, jsonFile)
	require.NoError(t, err)
}
