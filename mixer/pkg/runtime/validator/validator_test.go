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
	"context"
	"errors"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/hashicorp/go-multierror"

	cpb "istio.io/api/policy/v1beta1"
	adapter2 "istio.io/istio/mixer/adapter"
	"istio.io/istio/mixer/pkg/adapter"
	"istio.io/istio/mixer/pkg/config"
	"istio.io/istio/mixer/pkg/config/store"
	"istio.io/istio/mixer/pkg/lang/checker"
	"istio.io/istio/mixer/pkg/template"
	template2 "istio.io/istio/mixer/template"
)

// testAdapterConfig is a types.Struct instance which represents the
// adapter config used in the mixer/testdata/config/listentry.yaml.
// This is commonly used in the test cases and expectations in this file.
var testAdapterConfig = &types.Struct{
	Fields: map[string]*types.Value{
		"overrides": {
			Kind: &types.Value_ListValue{ListValue: &types.ListValue{
				Values: []*types.Value{
					{Kind: &types.Value_StringValue{StringValue: "v1"}},
					{Kind: &types.Value_StringValue{StringValue: "v2"}},
				},
			}},
		},
		"blacklist": {Kind: &types.Value_BoolValue{BoolValue: false}},
	},
}

type dummyHandlerBuilder struct {
	want proto.Message
	got  proto.Message
}

func (d *dummyHandlerBuilder) SetAdapterConfig(cfg adapter.Config) {
	d.got = cfg
}

func (d *dummyHandlerBuilder) Validate() *adapter.ConfigErrors {
	var err *adapter.ConfigErrors
	if !reflect.DeepEqual(d.want, d.got) {
		err = err.Appendf("", "Got %v, Want %v", d.got, d.want)
	}
	return err
}

func (d *dummyHandlerBuilder) Build(ctx context.Context, env adapter.Env) (adapter.Handler, error) {
	return nil, errors.New("dummy can't build")
}

func getValidatorForTest() (*Validator, error) {
	path, err := filepath.Abs("../../../testdata/config")
	if err != nil {
		return nil, err
	}
	s, err := store.NewRegistry(config.StoreInventory()...).NewStore("fs://" + path)
	if err != nil {
		return nil, err
	}
	tc := checker.NewTypeChecker()
	adapterInfo := make(map[string]*adapter.Info)
	for _, y := range adapter2.Inventory() {
		i := y()
		adapterInfo[i.Name] = &i
	}

	adapterInfo["listchecker"] = &adapter.Info{
		DefaultConfig: &types.Struct{},
		NewBuilder: func() adapter.HandlerBuilder {
			return &dummyHandlerBuilder{want: testAdapterConfig}
		},
	}

	templateInfo := make(map[string]*template.Info)
	for x, y := range template2.SupportedTmplInfo {
		tmp := y
		templateInfo[x] = &tmp
	}

	templateInfo["listentry"] = &template.Info{
		CtrCfg: &types.Struct{},
		InferType: func(msg proto.Message, fn template.TypeEvalFn) (proto.Message, error) {
			st := msg.(*types.Struct)
			v, ok := st.Fields["value"]
			if !ok {
				return nil, errors.New("no value field")
			}
			value := v.GetStringValue()
			if value == "" {
				return nil, errors.New("not string value")
			}
			_, ierr := fn(value)
			return nil, ierr
		},
	}

	v, err := New(tc, "destination.service", s, adapterInfo, templateInfo)
	if err != nil {
		return nil, err
	}
	return v.(*Validator), nil
}

func updateEvent(keystr string, spec proto.Message) *store.Event {
	keySegments := strings.Split(keystr, ".")
	key := store.Key{Name: keySegments[0], Kind: keySegments[1], Namespace: keySegments[2]}
	return &store.Event{Type: store.Update, Key: key, Value: &store.Resource{
		Metadata: store.ResourceMeta{Name: key.Name, Namespace: key.Namespace},
		Spec:     spec,
	}}
}

func deleteEvent(keystr string) *store.Event {
	keySegments := strings.Split(keystr, ".")
	return &store.Event{Type: store.Delete, Key: store.Key{Name: keySegments[0], Kind: keySegments[1], Namespace: keySegments[2]}}
}

func TestValidator(t *testing.T) {
	for _, cc := range []struct {
		title   string
		evs     []*store.Event
		ok      bool
		wantErr string
	}{
		{
			"new rule",
			[]*store.Event{updateEvent("test.rule.default", &cpb.Rule{
				Actions: []*cpb.Action{
					{Handler: "staticversion.listchecker.istio-system", Instances: []string{"appversion.listentry.istio-system"}},
				}})},
			true,
			"",
		},
		//		{
		//			"update rule",
		//			[]*store.Event{updateEvent("checkwl.rule.istio-system", &cpb.Rule{
		//				Actions: []*cpb.Action{
		//					{Handler: "staticversion.listchecker", Instances: []string{"appversion.listentry"}},
		//				}})},
		//			true,
		//			"",
		//		},
		//		{
		//			"delete rule",
		//			[]*store.Event{deleteEvent("checkwl.rule.istio-system")},
		//			true,
		//			"",
		//		},
		//		{
		//			"invalid updating rule: match syntax error",
		//			[]*store.Event{updateEvent("test.rule.default", &cpb.Rule{Match: "foo"})},
		//			false,
		//			"",
		//		},
		//		{
		//			"invalid updating rule: match type error",
		//			[]*store.Event{updateEvent("test.rule.default", &cpb.Rule{Match: "1"})},
		//			false,
		//			"",
		//		},
		//		{
		//			"invalid updating rule: reference not found",
		//			[]*store.Event{updateEvent("test.rule.default", &cpb.Rule{Actions: []*cpb.Action{{Handler: "nonexistent.listchecker.istio-system"}}})},
		//			false,
		//			"",
		//		},
		//		{
		//			"adding adapter",
		//			[]*store.Event{updateEvent("test.listchecker.default", testAdapterConfig)},
		//			true,
		//			"",
		//		},
		//		{
		//			"adding instance",
		//			[]*store.Event{updateEvent("test.listentry.default", &types.Struct{Fields: map[string]*types.Value{
		//				"value": {Kind: &types.Value_StringValue{StringValue: "0"}},
		//			}})},
		//			true,
		//			"",
		//		},
		//		{
		//			"adapter validation failure",
		//			[]*store.Event{updateEvent("test.listchecker.default", &types.Struct{})},
		//			false,
		//			"",
		//		},
		//		{
		//			"invalid instance",
		//			[]*store.Event{updateEvent("test.listentry.default", &types.Struct{})},
		//			false,
		//			"",
		//		},
		//		{
		//			"invalid instance syntax",
		//			[]*store.Event{updateEvent("test.listentry.default", &types.Struct{Fields: map[string]*types.Value{
		//				"value": {Kind: &types.Value_StringValue{StringValue: ""}},
		//			}})},
		//			false,
		//			"",
		//		},
		//		{
		//			"invalid delete handler",
		//			[]*store.Event{deleteEvent("staticversion.listchecker.istio-system")},
		//			false,
		//			"",
		//		},
		//		{
		//			"invalid delete instance",
		//			[]*store.Event{deleteEvent("appversion.listentry.istio-system")},
		//			false,
		//			"",
		//		},
		//		{
		//			"invalid removal of attributemanifest",
		//			[]*store.Event{deleteEvent("kubernetes.attributemanifest.istio-system")},
		//			false,
		//			"",
		//		},
	} {
		t.Run(cc.title, func(tt *testing.T) {
			v, err := getValidatorForTest()
			if err != nil {
				tt.Fatal(err)
			}
			defer v.Stop()
			var result *multierror.Error
			for _, ev := range cc.evs {
				e := v.Validate(ev)
				v.refreshTypeChecker()
				result = multierror.Append(result, e)
			}
			ok := result.ErrorOrNil() == nil
			if cc.ok != ok {
				tt.Errorf("Got %v, Want %v", result.ErrorOrNil(), cc.ok)
			}
			if cc.wantErr != "" {
				if !strings.Contains(result.Error(), cc.wantErr) {
					tt.Errorf("Got error %s, Want err %s", result.Error(), cc.wantErr)
				}
			}
		})
	}
}

func TestValidatorToRememberValidation(t *testing.T) {
	t.SkipNow()
	for _, c := range []struct {
		title string
		ev1   *store.Event
		evs   []*store.Event
	}{
		{
			"reference",
			updateEvent("test.rule.default", &cpb.Rule{
				Actions: []*cpb.Action{
					{Handler: "test.listchecker", Instances: []string{"appversion.listentry.istio-system"}},
				},
			}),
			[]*store.Event{updateEvent("test.listchecker.default", testAdapterConfig)},
		},
		{
			"deletion order",
			deleteEvent("checkwl.rule.istio-system"),
			[]*store.Event{
				deleteEvent("staticversion.listchecker.istio-system"),
				deleteEvent("appversion.listentry.istio-system"),
			},
		},
		{
			"deleting attribute manifests",
			deleteEvent("kubernetes.attributemanifest.istio-system"),
			[]*store.Event{
				deleteEvent("staticversion.listchecker.istio-system"),
				deleteEvent("appversion.listentry.istio-system"),
				deleteEvent("checkwl.rule.istio-system"),
			},
		},
	} {
		t.Run(c.title, func(tt *testing.T) {
			v, err := getValidatorForTest()
			if err != nil {
				tt.Fatal(err)
			}
			defer v.Stop()

			if err = v.Validate(c.ev1); err == nil {
				tt.Error("Got nil, Want error")
			}
			for _, ev := range c.evs {
				if err = v.Validate(ev); err != nil {
					tt.Errorf("Got %v, Want nil", err)
				}
			}
			if err = v.Validate(c.ev1); err != nil {
				tt.Errorf("Got %v, Want nil", err)
			}
		})
	}
}
