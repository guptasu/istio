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

package config

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	"istio.io/istio/mixer/pkg/adapter"
	"istio.io/istio/mixer/pkg/template"
	"istio.io/istio/pkg/log"
)

// isFQN returns true if the name is fully qualified.
// every resource name is defined by Key.String()
// shortname.kind.namespace
func isFQN(name string) bool {
	c := 0
	for _, ch := range name {
		if ch == '.' {
			c++
		}

		if c > 2 {
			return false
		}
	}

	return c == 2
}

// canonicalize ensures that the name is fully qualified.
func canonicalize(name string, namespace string) string {
	if isFQN(name) {
		return name
	}

	return name + "." + namespace
}

type InferredTypesMap map[string]proto.Message

func ValidateBuilder(
	builder adapter.HandlerBuilder,
	templates map[string]*template.Info,
	inferredTypes map[string]InferredTypesMap,
	handler *HandlerLegacy) (err error) {
	if builder == nil {
		err = fmt.Errorf("nil builder from adapter: adapter='%s'", handler.Adapter.Name)
		return
	}

	// validate if the builder supports all the necessary interfaces
	for _, tmplName := range handler.Adapter.SupportedTemplates {
		ti, found := templates[tmplName]
		if !found {
			// TODO (Issue #2512): This log is unnecessarily spammy. We should test for this during startup
			// and log it once.
			// One of the templates that is supported by the adapter was not found. We should log and simply
			// move on.
			log.Infof("Ignoring unrecognized template, supported by adapter: adapter='%s', template='%s'",
				handler.Adapter.NewBuilder, tmplName)
			continue
		}

		if supports := ti.BuilderSupportsTemplate(builder); !supports {
			err = fmt.Errorf("adapter does not actually support template: template='%s', interface='%s'", tmplName, ti.BldrInterfaceName)
			return
		}
	}

	for tmplName := range inferredTypes {
		types := inferredTypes[tmplName]
		// ti should be there for a valid configuration.
		ti := templates[tmplName]
		if ti.SetType != nil { // for case like APA template that does not have SetType
			ti.SetType(types, builder)
		}
	}

	builder.SetAdapterConfig(handler.Params)

	var ce *adapter.ConfigErrors
	if ce = builder.Validate(); ce != nil {
		err = fmt.Errorf("builder validation failed: '%v'", ce)
		return
	}
	return
}
