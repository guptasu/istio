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
	"io"
	"math"
	"strings"

	"github.com/golang/protobuf/proto"
	protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"istio.io/istio/mixer/pkg/attribute"
	"istio.io/istio/mixer/pkg/il/compiled"
)

// Assembler directly serialized proto into the passed buffer, based on the given attribute bag.
type Assembler interface {
	Assemble(bag attribute.Bag, buffer *proto.Buffer) error
}

type assembler interface {
	assemble(bag attribute.Bag, index int, buffer *proto.Buffer) error
}

type messageAssembler struct {
	fields []fieldEntry
}

type fieldEntry struct {
	index     int
	assembler assembler
}

var _ assembler = &messageAssembler{}
var _ Assembler = &messageAssembler{}

type scalarAssembler struct {
	expression compiled.Expression
	fieldType  protobuf.FieldDescriptorProto_Type
}

var _ assembler = &scalarAssembler{}

type constantAssembler struct {
	value string
}

var _ assembler = &constantAssembler{}

// Assemble implements the Assembler interface.
func (m *messageAssembler) Assemble(bag attribute.Bag, buffer *proto.Buffer) error {
	return m.assemble(bag, -1, buffer)
}

func (m *messageAssembler) assemble(bag attribute.Bag, index int, buffer *proto.Buffer) error {
	original := buffer
	// TODO: We can avoid this double-copy of data by filling up the buffer in reverse order (thus, writing
	// size last, after completing the full serialization). This requires our own implementation of proto.Buffer.
	if index != -1 {
		buffer = getBuffer()
	}

	for _, field := range m.fields {
		err := field.assembler.assemble(bag, field.index, buffer)
		if err != nil {
			return err
		}
	}

	if index != -1 {
		original.EncodeVarint(encodeIndexAndType(index, proto.WireBytes))
		original.EncodeRawBytes(buffer.Bytes())
		putBuffer(buffer)
	}

	return nil
}

func (s *scalarAssembler) assemble(bag attribute.Bag, index int, buffer *proto.Buffer) error {
	switch s.fieldType {
	case protobuf.FieldDescriptorProto_TYPE_BOOL:
		v, err := s.expression.EvaluateBoolean(bag)
		if err != nil {
			return err
		}

		val := uint64(0)
		if v {
			val = uint64(1)
		}
		buffer.EncodeVarint(encodeIndexAndType(index, proto.WireVarint))
		buffer.EncodeVarint(val)

	case protobuf.FieldDescriptorProto_TYPE_INT64:
		v, err := s.expression.EvaluateInteger(bag)
		if err != nil {
			return err
		}

		buffer.EncodeVarint(encodeIndexAndType(index, proto.WireVarint))
		buffer.EncodeVarint(uint64(v)) // TODO: Test edge cases for this cast.

	case protobuf.FieldDescriptorProto_TYPE_DOUBLE:
		v, err := s.expression.EvaluateDouble(bag)
		if err != nil {
			return err
		}

		buffer.EncodeVarint(encodeIndexAndType(index, proto.WireFixed64))

		buffer.EncodeFixed64(math.Float64bits(v))

	case protobuf.FieldDescriptorProto_TYPE_STRING:
		v, err := s.expression.EvaluateString(bag)
		if err != nil {
			return err
		}

		buffer.EncodeVarint(encodeIndexAndType(index, proto.WireBytes))
		buffer.EncodeStringBytes(v)

	// TODO: map

	default:
		// TODO: Come up with a strategy for mapping various other types (i.e. int32, fixed64, float etc.)
		panic("Unrecognized field type:" + s.fieldType.String())
	}

	return nil
}

func encodeIndexAndType(index int, typeid uint64) uint64 {
	return (uint64(index) << 3) | typeid
}

func decodeIndexAndType(u uint64) (index int, typeid uint64) {
	typeid = u & 0x07
	index = int(u >> 3)
	return
}

// NewAssemblerFor returns a new assembler instances based on the given instance and instanceParam descriptor,
// instanceParam proto and the expression builder.
func NewAssemblerFor(
	fd *protobuf.FileDescriptorProto,
	instanceName string,
	paramName string,
	instanceParam []byte,
	builder *compiled.ExpressionBuilder) (Assembler, error) {

	r := newResolver(fd)
	return buildMessageAssembler(r, instanceName, paramName, proto.NewBuffer(instanceParam), builder)
}

func buildMessageAssembler(
	r resolver,
	instanceName string,
	paramName string,
	instParamBuf *proto.Buffer,
	builder *compiled.ExpressionBuilder) (*messageAssembler, error) {

	instanceDescriptor := r.resolve(instanceName)
	if instanceDescriptor == nil {
		return nil, fmt.Errorf("descriptor not found: %s", instanceName)
	}

	parameterDescriptor := r.resolve(paramName)
	if parameterDescriptor == nil {
		return nil, fmt.Errorf("descriptor not found: %s", paramName)
	}

	result := &messageAssembler{
		fields: []fieldEntry{},
	}

	for {
		vint, err := instParamBuf.DecodeVarint()
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				return result, nil
			}
			return nil, err
		}
		i, t := decodeIndexAndType(vint)
		if t != proto.WireBytes {
			return nil, fmt.Errorf("unexpected field type encountered: %d", t)
		}

		fieldDescriptor := findField(instanceDescriptor, i)
		if fieldDescriptor == nil {
			return nil, fmt.Errorf("field not found in instance: %d", i)
		}

		// TODO: Support other nonscalar types as well
		// TODO: support missing type info.
		if *fieldDescriptor.Type == protobuf.FieldDescriptorProto_TYPE_MESSAGE {
			name := fieldDescriptor.GetTypeName()
			idx := strings.Index(name, instanceName)
			if idx == -1 {
				paramName = name + "Param"
			} else {
				// TODO: The right name matching algorithm.
				prefix := name[:idx+len(instanceName)]
				suffix := name[idx+len(instanceName):]
				paramName = prefix + "Param" + suffix
			}
			var mbytes []byte
			if mbytes, err = instParamBuf.DecodeRawBytes(false); err != nil {
				return nil, err
			}
			mbuf := proto.NewBuffer(mbytes)
			var asm *messageAssembler
			if asm, err = buildMessageAssembler(r, name, paramName, mbuf, builder); err != nil {
				return nil, err
			}
			entry := fieldEntry{
				index:     i,
				assembler: asm,
			}
			result.fields = append(result.fields, entry)
			continue
		}
		// Let the map key pass through
		if strings.HasSuffix(instanceName, "MapPrimitiveEntry") && i == 1 {
			mapKey, err := instParamBuf.DecodeStringBytes()
			if err != nil {
				return nil, err
			}

			entry := fieldEntry{
				index: i,
				assembler: &constantAssembler{
					value: mapKey,
				},
			}
			result.fields = append(result.fields, entry)
			continue
		}

		expression, err := instParamBuf.DecodeStringBytes()
		if err != nil {
			return nil, fmt.Errorf("unable to decode expression for field: %d, %v", i, err)
		}

		expr, exprType, err := builder.Compile(expression)
		if err != nil {
			return nil, fmt.Errorf("error compiling expression '%s': %v", expression, err)
		}
		// TODO: Detect expression type mismatch
		_ = exprType

		entry := fieldEntry{
			index: i,
			assembler: &scalarAssembler{
				expression: expr,
				fieldType:  *fieldDescriptor.Type,
			},
		}
		result.fields = append(result.fields, entry)
	}
}


func (m *constantAssembler) assemble(bag attribute.Bag, index int, buffer *proto.Buffer) error {
	buffer.EncodeVarint(encodeIndexAndType(index, proto.WireBytes))
	buffer.EncodeStringBytes(m.value)
	return nil
}
