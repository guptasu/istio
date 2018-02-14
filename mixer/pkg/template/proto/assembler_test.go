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
	"bytes"
	"compress/gzip"
	"reflect"
	"testing"

	gogo "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"istio.io/istio/mixer/pkg/attribute"
	"istio.io/istio/mixer/pkg/expr"
	"istio.io/istio/mixer/pkg/il/compiled"

	pbv "istio.io/api/mixer/v1/config/descriptor"
	configpb "istio.io/api/mixer/v1/config"
	tst "istio.io/istio/mixer/pkg/template/proto/testing"
	"fmt"

)

var tests = []struct {
	n string
	i *tst.InstanceParam
	a map[string]interface{}
	e *tst.Instance
}{
	{
		n: "basic",
		i: &tst.InstanceParam{
			StringPrimitive: "as",
			BoolPrimitive:   `as == "foo"`,
			DoublePrimitive: "43.45",
			Int64Primitive:  "23",
		},
		a: map[string]interface{}{
			"as": "baz",
		},
		e: &tst.Instance{
			StringPrimitive: "baz",
			BoolPrimitive:   false,
			DoublePrimitive: float64(43.45),
			Int64Primitive:  int64(23),
		},
	},
}

var manifest = map[string]*configpb.AttributeManifest_AttributeInfo{
	"as": {
		ValueType: pbv.STRING,
	},
	"ai": {
		ValueType: pbv.INT64,
	},
}

func TestDebug(t *testing.T) {
	fd, err := getFileDescriptor()
	if err != nil {
		t.Fatalf("%v", err)
	}
	yaml  := `
int64Primitive: "23"
boolPrimitive: as == "foo"
stringPrimitive: as
doublePrimitive: "43.45"
`
	bytes, err := yamlToBytes(yaml, fd, "InstanceParam")

	instanceParam := tst.InstanceParam{}
	err = proto.Unmarshal(bytes, &instanceParam)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	fmt.Println(bytes, err)
}

func TestAssembler(t *testing.T) {
	for _, ts := range tests {
		t.Run(ts.n, func(tt *testing.T) {
			fd, err := getFileDescriptor()
			if err != nil {
				tt.Fatalf("%v", err)
			}

			paramBytes, err := proto.Marshal(ts.i)
			if err != nil {
				tt.Fatalf("Error marshalling instance bytes: %v", err)
			}

			// Create an expression builder for building an assembler.
			finder := expr.NewFinder(manifest)
			builder := compiled.NewBuilder(finder)

			// Create a new Assembler which, given an instance and instanceParam descriptor (and the bytes for the
			// instanceParam, creates an assembler that can create an instance.
			assembler, err := NewAssemblerFor(fd, "Instance", "InstanceParam", paramBytes, builder)
			if err != nil {
				tt.Fatalf("error from assembler: %v", err)
			}

			// Now with the assembler, try creating an instance directly into a proto buffer.
			bag := attribute.GetFakeMutableBagForTesting(ts.a)
			buf := proto.NewBuffer([]byte{})
			err = assembler.Assemble(bag, buf)
			if err != nil {
				tt.Fatalf("error: %v", err)
			}

			// Deserialize from the created buffer and compare against an expected format.
			instance := tst.Instance{}
			err = proto.Unmarshal(buf.Bytes(), &instance)
			if err != nil {
				tt.Fatalf("error: %v", err)
			}

			if !reflect.DeepEqual(ts.e, &instance) {
				tt.Fatalf("got:\n%+v\nwanted:\n%+v\n", &instance, ts.e)
			}
		})
	}
}

func getFileDescriptor() (*descriptor.FileDescriptorProto, error) {
	// Get the descriptor for instance and param.
	instanceDescriptorBytes := gogo.FileDescriptor("test.proto")
	bytes, err := ungzip(instanceDescriptorBytes)
	if err != nil {
		return nil, err
	}
	fd := &descriptor.FileDescriptorProto{}
	err = proto.Unmarshal(bytes, fd)
	return fd, err
}

func ungzip(input []byte) ([]byte, error) {
	b := bytes.NewBuffer(input)
	r, err := gzip.NewReader(b)

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
