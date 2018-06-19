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

package mygrpcadapter

import (
	"fmt"
	"io/ioutil"
	"testing"

	adapter_integration "istio.io/istio/mixer/pkg/adapter/test"
	"os"
	"strings"
)

const (
	hconfig = `
apiVersion: "config.istio.io/v1alpha2"
kind: handler
metadata:
  name: h1
  namespace: istio-system
spec:
  adapter: mygrpcadapter
  connection:
    address: "%s"
  params:
    file_path: "out.txt"
---
`
	iconfig = `
apiVersion: "config.istio.io/v1alpha2"
kind: instance
metadata:
  name: i1metric
  namespace: istio-system
spec:
  template: metric
  params:
    value: request.size | 0
    dimensions:
      destination_service: "\"myservice\""
      response_code: "200"
---
`

	rconfig = `
apiVersion: "config.istio.io/v1alpha2"
kind: rule
metadata:
  name: r1
  namespace: istio-system
spec:
  actions:
  - handler: h1.istio-system
    instances:
    - i1metric
---
`
)

func TestReport(t *testing.T) {
	adptCrBytes, err := ioutil.ReadFile("config/mygrpcadapter.yaml")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}


	shutdown := make(chan error, 1)

	var outFile *os.File
	outFile, err = os.OpenFile("out.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if removeErr := os.Remove(outFile.Name()); removeErr != nil {
			t.Logf("Could not remove temporary file %s: %v", outFile.Name(), removeErr)
		}
	}()

	adapter_integration.RunTest(
		t,
		nil,
		adapter_integration.Scenario{
			Setup: func() (ctx interface{}, err error) {
				pServer, err := NewNoSessionServer("")
				if err != nil {
					return nil, err
				}
				go func() {
					pServer.Run(shutdown)
					_ = <-shutdown
				}()
				return pServer, nil
			},
			Teardown: func(ctx interface{}) {
				s := ctx.(Server)
				s.Close()
			},
			ParallelCalls: []adapter_integration.Call{
				{
					CallKind: adapter_integration.REPORT,
					Attrs:    map[string]interface{}{"request.size": 555},
				},
			},
			GetState: func(ctx interface{}) (interface{}, error) {
				// validate if the content of "out.txt" is as expected
				bytes, err := ioutil.ReadFile("out.txt")
				if err != nil {
					return nil, err
				}
				s := string(bytes)
				wantStr := `HandleMetric invoked with:
		  Adapter config: &Params{FilePath:out.txt,}
		  Instances: 'i1metric.instance.istio-system':
		  {
				Value = 555
				Dimensions = map[destination_service:myservice response_code:200]
		  }
`
				if normalize(s) != normalize(wantStr) {
					return nil, fmt.Errorf("got adapters state as : '%s'; want '%s'", s, wantStr)
				}
				return nil, nil
			},
			GetConfig: func(ctx interface{}) ([]string, error) {
				s := ctx.(Server)
				return []string{
					// CRs for built-in templates (metric is what we need for this test)
					// are automatically added by the integration test framework.
					string(adptCrBytes),
					fmt.Sprintf(hconfig, s.Addr()),
					iconfig,
					rconfig,
				}, nil
			},
			Want: `
		{
		 "AdapterState": null,
		 "Returns": [
		  {
		   "Check": {
		    "Status": {},
		    "ValidDuration": 0,
		    "ValidUseCount": 0
		   },
		   "Quota": null,
		   "Error": null
		  }
		 ]
		}`,
		},
	)
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, " ", "", -1)
	return s
}