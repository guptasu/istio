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

package validator

import (
	"time"

	"github.com/gogo/protobuf/proto"

	cpb "istio.io/api/policy/v1beta1"
	"istio.io/istio/mixer/pkg/adapter"
	"istio.io/istio/mixer/pkg/config/store"
	"istio.io/istio/mixer/pkg/lang/checker"
	"istio.io/istio/mixer/pkg/runtime/config"
	"istio.io/istio/mixer/pkg/template"
	"istio.io/istio/pkg/cache"
)

// Validator offers semantic validation of the config changes.
type Validator struct {
	handlerBuilders map[string]adapter.HandlerBuilder
	templates       map[string]*template.Info
	tc              checker.TypeChecker
	c               *validatorCache
	donec           chan struct{}
	e               *config.Ephemeral
}

// New creates a new store.Validator instance which validates runtime semantics of
// the configs.
func New(tc checker.TypeChecker, identityAttribute string, s store.Store,
	adapterInfo map[string]*adapter.Info, templateInfo map[string]*template.Info) (store.Validator, error) {
	kinds := config.KindMap(adapterInfo, templateInfo)
	data, ch, err := store.StartWatch(s, kinds)
	if err != nil {
		return nil, err
	}
	hb := make(map[string]adapter.HandlerBuilder, len(adapterInfo))
	for k, ai := range adapterInfo {
		hb[k] = ai.NewBuilder()
	}
	configData := make(map[store.Key]proto.Message, len(data))
	manifests := map[store.Key]*cpb.AttributeManifest{}
	for k, obj := range data {
		if k.Kind == config.AttributeManifestKind {
			manifests[k] = obj.Spec.(*cpb.AttributeManifest)
		}
		configData[k] = obj.Spec
	}
	e, err := config.NewEphemeral(templateInfo, adapterInfo)
	if err != nil {
		return nil, err
	}
	v := &Validator{
		handlerBuilders: hb,
		templates:       templateInfo,
		tc:              tc,
		c: &validatorCache{
			c:          cache.NewTTL(validatedDataExpiration, validatedDataEviction),
			configData: configData,
		},
		donec: make(chan struct{}),
		e:     e,
	}
	v.e.SetState(data)
	go store.WatchChanges(ch, v.donec, time.Second, v.c.applyChanges)
	return v, nil
}

// Stop stops the validator.
func (v *Validator) Stop() {
	close(v.donec)
}

// Validate implements store.Validator interface.
func (v *Validator) Validate(ev *store.Event) error {
	// get old state so we can revert in case of validation error.
	oldEntryVal, exists := v.e.GetEntry(ev)
	oldAttrs := v.e.GetProcessedAttributes()

	v.e.ApplyEvent(ev)
	_, err := v.e.BuildSnapshot()

	if err != nil {
		v.e.SetProcessedAttributes(oldAttrs)
		reverseEvent := *ev
		if exists {
			reverseEvent.Value = oldEntryVal
			reverseEvent.Type = store.Update
		} else if ev.Type == store.Update { // didn't existed before.
			reverseEvent.Type = store.Delete
		}
		v.e.ApplyEvent(&reverseEvent)
	}

	return err
}
