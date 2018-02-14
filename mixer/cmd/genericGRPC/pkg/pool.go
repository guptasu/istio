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

package pkg

import (
	"sync"

	"github.com/golang/protobuf/proto"
)

var bufferPool = sync.Pool{New: func() interface{} { return proto.NewBuffer(make([]byte, 0, 256)) }}

func GetBuffer() *proto.Buffer {
	return bufferPool.Get().(*proto.Buffer)
}

func putBuffer(b *proto.Buffer) {
	b.Reset()
	bufferPool.Put(b)
}
