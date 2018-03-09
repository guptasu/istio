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

package list

import (
	"testing"

	adapter2 "istio.io/istio/mixer/pkg/adapter"
	adapter_integration "istio.io/istio/mixer/pkg/adapter/test"
	"istio.io/istio/mixer/template"
)

const (
	h1_override_src1src2 = `
apiVersion: "config.istio.io/v1alpha2"
kind: listchecker
metadata:
  name: staticversion
  namespace: istio-system
spec:
  overrides: ["src1", "src2"]
  blacklist: false
`
	i1_val_src_name_attr = `
apiVersion: "config.istio.io/v1alpha2"
kind: listentry
metadata:
  name: appversion
  namespace: istio-system
spec:
  value: source.name | ""
`
	r1_h1_i1 = `
apiVersion: "config.istio.io/v1alpha2"
kind: rule
metadata:
  name: checkwl
  namespace: istio-system
spec:
  actions:
  - handler: staticversion.listchecker
    instances:
    - appversion.listentry
`
)

func TestReport(t *testing.T) {
	adapter_integration.RunTest(
		t,
		[]adapter2.InfoFn{GetInfo},
		template.SupportedTmplInfo,
		nil, /*no setup*/
		nil, /*no teardown*/
		nil, /*no adapter specific state to test*/
		adapter_integration.TestCase{
			ParallelCalls: []adapter_integration.Call{
				{
					CallKind: adapter_integration.CHECK,
					Attrs:    map[string]interface{}{"source.name": "src1"},
				},
				{
					CallKind: adapter_integration.CHECK,
				},
			},
			Cfgs: []string{
				h1_override_src1src2,
				r1_h1_i1,
				i1_val_src_name_attr,
			},
			Want: `{
            "AdapterState": null,
            "Returns": [
             {
              "Check": {
                "Status": {},
                "ValidDuration": 300000000000,
                "ValidUseCount": 10000
              }
             },
             {
              "Check": {
               "Status": {
                "code": 5,
                "message": "staticversion.listchecker.istio-system: is not whitelisted"
               },
               "ValidDuration": 300000000000,
               "ValidUseCount": 10000
              }
             }
            ]
            }`,
		},
	)
}
