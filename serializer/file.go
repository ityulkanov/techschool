package serializer

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {

	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary %w", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("can not write binary data to file: %w", err)
	}
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("can not read binary data from file %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("can not unmarshall binary to proto message: %w", err)
	}

	return nil
}

func WriteProtobufToJsonFormat(message proto.Message, filename string) error {
	data, err := protobufToJson(message)
	if err != nil {
		return fmt.Errorf("cannot marhsall data from message %w", err)
	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("unable to write file %w", err)
	}
	return nil
}
