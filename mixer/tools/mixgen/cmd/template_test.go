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

//go:generate protoc testdata/foo.proto -otestdata/foo.descriptor -I$GOPATH/src/istio.io/istio/vendor/istio.io/api -I.
package cmd

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type templateCmdTestdata struct {
	descriptorContent string
	wantCfg           string
}

func TestTemplateCmd(t *testing.T) {
	for i, td := range []templateCmdTestdata{
		{
			descriptorContent: "some text",
			wantCfg: `# this config is created through command
# mixgen template -d tempDescriptorFile -n myTemplateResourceName --namespace mynamespace
apiVersion: "config.istio.io/v1alpha2"
kind: template
metadata:
  name: myTemplateResourceName
  namespace: mynamespace
spec:
  descriptor: "c29tZSB0ZXh0"
---
`,
		},
	} {
		file, _ := os.Create("tempDescriptorFile")
		defer func() {
			if removeErr := os.Remove(file.Name()); removeErr != nil {
				t.Logf("could not remove temporary file %s: %v", file.Name(), removeErr)
			}
		}()

		t.Run(fmt.Sprintf("%d", i), func(tt *testing.T) {
			var args []string

			_, _ = file.WriteString(td.descriptorContent)
			args = []string{"template", "-d", file.Name(), "-n", "myTemplateResourceName", "--namespace", "mynamespace"}

			gotCfg := ""
			root := GetRootCmd(args,
				func(format string, a ...interface{}) {
					gotCfg = fmt.Sprintf(format, a...)
				},
				func(format string, a ...interface{}) {
					gotError := fmt.Sprintf(format, a...)
					tt.Fatalf("want error 'nil'; got '%s'", gotError)
				})

			_ = root.Execute()

			if gotCfg != td.wantCfg {
				tt.Errorf("want :\n%v\ngot :\n%v", td.wantCfg, gotCfg)
			}
		})
	}
}

func TestTemplateCmd_NoInputFile(t *testing.T) {
	var gotError string
	cmd := GetRootCmd([]string{"template"},
		func(format string, a ...interface{}) {},
		func(format string, a ...interface{}) {
			gotError = fmt.Sprintf(format, a...)
			if !strings.Contains(gotError, "unable to read") {
				t.Fatalf("want error 'unable to read'; got '%s'", gotError)
			}
		})
	_ = cmd.Execute()
	if gotError == "" {
		t.Errorf("want error; got nil")
	}
}
