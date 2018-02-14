// Copyright 2018 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proto

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	yaml "gopkg.in/yaml.v2"
)

func yamlToBytes(configYaml string, fd *protobuf.FileDescriptorProto, msgName string) ([]byte, error) {
	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(configYaml), &m)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	r := newResolver(fd)
	desc := r.resolve(msgName)

	buf := proto.NewBuffer([]byte{})
	for k, v := range m {
		fieldDescriptor := findFieldByName(desc, k.(string))
		if fieldDescriptor == nil {
			return nil, fmt.Errorf("field not found in instance: %s", k)
		}
		encode(*fieldDescriptor, v, buf)
	}

	return buf.Bytes(), nil
}

func encode(fieldDesc protobuf.FieldDescriptorProto, data interface{}, buffer *proto.Buffer) error {
	switch *fieldDesc.Type {
	case protobuf.FieldDescriptorProto_TYPE_STRING:
		v, ok := data.(string)
		if !ok {
			return fmt.Errorf("yaml val %v didn't match field type string", data)
		}

		buffer.EncodeVarint(encodeIndexAndType(int(*fieldDesc.Number), proto.WireBytes))
		buffer.EncodeStringBytes(v)
	default:
		// TODO: Come up with a strategy for mapping various other types (i.e. int32, fixed64, float etc.)
		panic("Unrecognized field type:" + (*fieldDesc.Type).String())
	}

	return nil
}
