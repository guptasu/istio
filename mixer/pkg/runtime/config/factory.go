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
	"context"
	"errors"
	"fmt"

	"github.com/gogo/protobuf/proto"

	istio_mixer_v1_config_descriptor "istio.io/api/policy/v1beta1"
	"istio.io/istio/mixer/pkg/adapter"
	"istio.io/istio/mixer/pkg/lang/checker"
	"istio.io/istio/mixer/pkg/runtime/safecall"
	"istio.io/istio/pkg/log"
)

// Factory is used to instantiate handlers.
type Factory struct {
	snapshot *Snapshot

	checker checker.TypeChecker

	// Map of instance name to inferred type (proto.Message)
	inferredTypesCache map[string]proto.Message
}

func NewFactory(snapshot *Snapshot) *Factory {
	return &Factory{
		snapshot: snapshot,
		checker:  checker.NewTypeChecker(),

		inferredTypesCache: make(map[string]proto.Message),
	}
}


// build instantiates a handler object using the passed in handler and instances configuration.
func (f *Factory) ValidateBuilder(
	handler *HandlerLegacy,
	instances []*InstanceLegacy) (hb adapter.HandlerBuilder, err error) {

	// Do not assign the error to err directly, as this would overwrite the err returned by the inner function.
	panicErr := safecall.Execute("factory.build", func() {
		var inferredTypesByTemplates map[string]InferredTypesMap
		if inferredTypesByTemplates, err = f.inferTypes(instances); err != nil {
			return
		}

		// Adapter should always be present for a valid configuration (reference integrity should already be checked).
		info := handler.Adapter

		hb := info.NewBuilder()
		if hb == nil {
			err = errors.New("nil HandlerBuilder")
			return
		}
		// validate and only construct if the validation passes.
		if err = ValidateBuilder(hb, f.snapshot.Templates, inferredTypesByTemplates, handler); err != nil {
			err = fmt.Errorf("adapter validation failed : %v", err)
			hb = nil
			return
		}
	})

	if panicErr != nil {
		err = panicErr
		hb = nil
		return
	}

	return
}

// build instantiates a handler object using the passed in handler and instances configuration.
func (f *Factory) Build(
	handler *HandlerLegacy,
	instances []*InstanceLegacy,
	env adapter.Env) (h adapter.Handler, err error) {

	// Do not assign the error to err directly, as this would overwrite the err returned by the inner function.
	panicErr := safecall.Execute("factory.build", func() {
		var inferredTypesByTemplates map[string]InferredTypesMap
		if inferredTypesByTemplates, err = f.inferTypes(instances); err != nil {
			return
		}

		// Adapter should always be present for a valid configuration (reference integrity should already be checked).
		info := handler.Adapter

		builder := info.NewBuilder()
		if builder == nil {
			err = errors.New("nil HandlerBuilder")
			return
		}
		// validate and only construct if the validation passes.
		if err = ValidateBuilder(builder, f.snapshot.Templates, inferredTypesByTemplates, handler); err != nil {
			h = nil
			err = fmt.Errorf("adapter validation failed : %v", err)
			return
		}
		h, err = f.buildHandler(builder, env)
		if err != nil {
			h = nil
			err = fmt.Errorf("adapter instantiation error: %v", err)
			return
		}

		// validate if the handlerConfig supports all the necessary interfaces
		for _, tmplName := range info.SupportedTemplates {
			// ti should be there for a valid configuration.
			ti, found := f.snapshot.Templates[tmplName]
			if !found {
				// This is similar to the condition check in the previous loop. That already does logging, so
				// simply skip.
				continue
			}

			if supports := ti.HandlerSupportsTemplate(h); !supports {
				if err = h.Close(); err != nil {
					h = nil
					// log this error, but return the one below. That is likely to be the more important one.
					log.Errorf("error during adapter close: '%v'", err)
					return
				}

				h = nil
				// adapter is bad since it does not support the necessary interface
				err = fmt.Errorf("builder for adapter does not actually support template: template='%s', interface='%s'",
					tmplName, ti.HndlrInterfaceName)
				return
			}
		}
	})

	if panicErr != nil {
		h = nil
		err = panicErr
	}

	return
}

func (f *Factory) buildHandler(builder adapter.HandlerBuilder, env adapter.Env) (handler adapter.Handler, err error) {
	return builder.Build(context.Background(), env)
}

func (f *Factory) inferTypes(instances []*InstanceLegacy) (map[string]InferredTypesMap, error) {

	typesByTemplate := make(map[string]InferredTypesMap)
	for _, instance := range instances {

		inferredType, err := f.inferType(instance)
		if err != nil {
			return nil, err
		}

		if _, exists := typesByTemplate[instance.Template.Name]; !exists {
			typesByTemplate[instance.Template.Name] = make(InferredTypesMap)
		}

		typesByTemplate[instance.Template.Name][instance.Name] = inferredType
	}
	return typesByTemplate, nil
}

func (f *Factory) inferType(instance *InstanceLegacy) (proto.Message, error) {

	var inferredType proto.Message
	var err error
	var found bool

	if inferredType, found = f.inferredTypesCache[instance.Name]; found {
		return inferredType, nil
	}

	// t should be there since the config is already validated
	t := instance.Template

	inferredType, err = t.InferType(instance.Params, func(expr string) (istio_mixer_v1_config_descriptor.ValueType, error) {
		return f.checker.EvalType(expr, f.snapshot.Attributes)
	})
	if err != nil {
		return nil, fmt.Errorf("cannot infer type information from params in instanceConfig '%s': %v", instance.Name, err)
	}

	f.inferredTypesCache[instance.Name] = inferredType

	return inferredType, nil
}
