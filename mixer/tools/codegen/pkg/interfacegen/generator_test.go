// Copyright 2017 Istio Authors
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

package interfacegen

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -t mixer/tools/codegen/pkg/interfacegen/testdata/apa/template.proto
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -t mixer/tools/codegen/pkg/interfacegen/testdata/check/template.proto
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -t mixer/tools/codegen/pkg/interfacegen/testdata/report/template.proto
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -t mixer/tools/codegen/pkg/interfacegen/testdata/quota/template.proto
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -t mixer/tools/codegen/pkg/interfacegen/testdata/error/template.proto

type logFn func(string, ...interface{})

// TestGenerator_Generate uses the outputs file descriptors generated via bazel
// and compares them against the golden files.
func TestGenerator_Generate(t *testing.T) {
	importmap := map[string]string{
		"mixer/v1/config/descriptor/value_type.proto": "istio.io/api/mixer/v1/config/descriptor",
		"mixer/adapter/model/v1beta1/extensions.proto":          "istio.io/api/mixer/v1/template",
		"gogoproto/gogo.proto":                        "github.com/gogo/protobuf/gogoproto",
		"google/protobuf/duration.proto":              "github.com/gogo/protobuf/types",
	}

	tests := []struct {
		name, descriptor, wantIntFace, wantProto string
	}{
		{"Report", "testdata/report/template_proto.descriptor_set",
			"testdata/report/template_handler.gen.go.golden",
			"testdata/report/template_instance.proto.golden"},
		{"Quota", "testdata/quota/template_proto.descriptor_set",
			"testdata/quota/template_handler.gen.go.golden",
			"testdata/quota/template_instance.proto.golden"},
		{"Check", "testdata/check/template_proto.descriptor_set",
			"testdata/check/template_handler.gen.go.golden",
			"testdata/check/template_instance.proto.golden"},
		{"APA", "testdata/apa/template_proto.descriptor_set",
			"testdata/apa/template_handler.gen.go.golden",
			"testdata/apa/template_instance.proto.golden"},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			tmpDir := path.Join(os.TempDir(), v.name)
			_ = os.MkdirAll(tmpDir, os.ModeDir|os.ModePerm)
			oIntface, err := os.Create(path.Join(tmpDir, path.Base(v.wantIntFace)))
			if err != nil {
				t.Fatal(err)
			}
			oTmpl, err := os.Create(path.Join(tmpDir, path.Base(v.wantProto)))
			if err != nil {
				t.Fatal(err)
			}

			defer func() {
				if !t.Failed() {
					if removeErr := os.Remove(oIntface.Name()); removeErr != nil {
						t.Logf("Could not remove temporary file %s: %v", oIntface.Name(), removeErr)
					}
					if removeErr := os.Remove(oTmpl.Name()); removeErr != nil {
						t.Logf("Could not remove temporary file %s: %v", oTmpl.Name(), removeErr)
					}
				}
			}()

			g := Generator{OutInterfacePath: oIntface.Name(), OAugmentedTmplPath: oTmpl.Name(), ImptMap: importmap}

			if err := g.Generate(v.descriptor); err != nil {
				t.Fatalf("Generate(%s) produced an error: %v", v.descriptor, err)
			}

			if same := fileCompare(oIntface.Name(), v.wantIntFace, t.Errorf, false); !same {
				t.Errorf("File %s does not match baseline %s.", oIntface.Name(), v.wantIntFace)
			}

			if same := fileCompare(oTmpl.Name(), v.wantProto, t.Errorf, true); !same {
				t.Errorf("File %s does not match baseline %s.", oTmpl.Name(), v.wantProto)
			}
		})
	}
}

func TestGenerator_GenerateErrors(t *testing.T) {
	file, err := ioutil.TempFile("", "error_file")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if removeErr := os.Remove(file.Name()); removeErr != nil {
			t.Logf("Could not remove temporary file %s: %v", file.Name(), removeErr)
		}
	}()

	g := Generator{OutInterfacePath: file.Name()}
	err = g.Generate("testdata/error/template_proto.descriptor_set")
	if err == nil {
		t.Fatalf("Generate(%s) should have produced an error", "testdata/error/template_proto.descriptor_set")
	}
	b, fileErr := ioutil.ReadFile("testdata/error/template.baseline")
	if fileErr != nil {
		t.Fatalf("Could not read baseline file: %v", err)
	}
	want := fmt.Sprintf("%s", b)
	got := err.Error()
	if got != want {
		t.Fatalf("Generate(%s) => '%s'\nwanted: '%s'", "testdata/error/template_proto.descriptor_set", got, want)
	}
}

const chunkSize = 64000

func fileCompare(actual, want string, logf logFn, skipSpaces bool) bool {
	f1, err := os.Open(actual)
	if err != nil {
		logf("could not open file: %v", err)
		return false
	}

	f2, err := os.Open(want)
	if err != nil {
		logf("could not open file: %v", err)
		return false
	}

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)
		b1 = bytes.Trim(b1, "\x00")
		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)
		b2 = bytes.Trim(b2, "\x00")

		b1ToCmp := b1
		b2ToCmp := b2

		if skipSpaces {
			b1ToCmp = bytes.Replace(b1ToCmp, []byte(" "), []byte(""), -1)
			b1ToCmp = bytes.Replace(b1ToCmp, []byte("\n"), []byte(""), -1)
			b2ToCmp = bytes.Replace(b2ToCmp, []byte(" "), []byte(""), -1)
			b2ToCmp = bytes.Replace(b2ToCmp, []byte("\n"), []byte(""), -1)
		}

		if err1 == io.EOF && err2 == io.EOF {
			return true
		}

		if err1 != nil || err2 != nil {
			return false
		}
		if !bytes.Equal(b1ToCmp, b2ToCmp) {
			logf("bytes don't match (sizes: %d, %d):\n%s\nNOT EQUALS\n%s.\n"+
				"Got file content:\n%s\nWant file content:\n%s\n",
				len(b1ToCmp), len(b2ToCmp), string(b1ToCmp), string(b2ToCmp), string(b1), string(b2))
			return false
		}
	}
}
