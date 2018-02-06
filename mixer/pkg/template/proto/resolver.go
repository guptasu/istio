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
	protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type resolver struct {
	fd *protobuf.FileDescriptorProto
}

func newResolver(fd *protobuf.FileDescriptorProto) resolver {
	return resolver{fd: fd}
}

func (r resolver) resolve(name string) *protobuf.DescriptorProto {
	// TODO: Precalculate these and avoid allocation in the loop.
	pkgName := ""
	if r.fd.Package != nil {
		pkgName = *r.fd.Package
	}

	return r.resolveInContext("."+pkgName, name, r.fd.MessageType)
}

func (r resolver) resolveInContext(prefix string, name string, descriptors []*protobuf.DescriptorProto) *protobuf.DescriptorProto {

	for _, t := range descriptors {
		if t.Name == nil {
			continue
		}

		if *t.Name == name {
			return t
		}

		// TODO: avoid this allocation
		fqn := prefix + "." + *t.Name

		if name == fqn {
			return t
		}

		nd := r.resolveInContext(fqn, name, t.NestedType)
		if nd != nil {
			return nd
		}
	}

	return nil
}

func findField(descriptor *protobuf.DescriptorProto, index int) *protobuf.FieldDescriptorProto {
	for _, f := range descriptor.Field {
		// TODO: Handle missing number
		if *f.Number == int32(index) {
			return f
		}
	}
	return nil
}
