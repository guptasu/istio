// Copyright 2016 Istio Authors
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

package template

// InterfaceTemplate defines the template used to generate the adapter
// interfaces for Mixer for a given aspect.
var InterfaceTemplate = `// Copyright 2017 Istio Authors
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

// THIS FILE IS AUTOMATICALLY GENERATED.

package {{.PkgName}}

import (
	"github.com/gogo/protobuf/proto"
	"fmt"
	"context"
	"strings"
	"net"
	"istio.io/istio/mixer/pkg/attribute"
	"istio.io/istio/mixer/pkg/expr"
	"istio.io/istio/mixer/pkg/adapter"
	"istio.io/api/mixer/v1/config/descriptor"
	"istio.io/istio/mixer/pkg/template"
	"github.com/golang/glog"
	adptTmpl "istio.io/api/mixer/v1/template"
	"istio.io/istio/mixer/pkg/config/proto"
	"errors"
	{{range .TemplateModels}}
		"{{.PackageImportPath}}"
	{{end}}
	$$additional_imports$$
)

var _ net.IP
var _ istio_mixer_v1_config.AttributeManifest
var _ = strings.Reader{}

const emptyQuotes = "\"\""

type wrapperAttr struct {
	get         getFn
	names       namesFn
	done        doneFn
	debugString debugStringFn
}

type (
	getFn         func(name string) (value interface{}, found bool)
	namesFn       func() []string
	doneFn        func()
	debugStringFn func() string
)

func newWrapperAttrBag(get getFn, names namesFn, done doneFn, debugString debugStringFn) attribute.Bag {
	return &wrapperAttr{
		debugString: debugString,
		done:        done,
		get:         get,
		names:       names,
	}
}

// Get returns an attribute value.
func (w *wrapperAttr) Get(name string) (value interface{}, found bool) {
	return w.get(name)
}

// Names returns the names of all the attributes known to this bag.
func (w *wrapperAttr) Names() []string {
	return w.names()
}

// Done indicates the bag can be reclaimed.
func (w *wrapperAttr) Done() {
	w.done()
}

// DebugString provides a dump of an attribute Bag that avoids affecting the
// calculation of referenced attributes.
func (w *wrapperAttr) DebugString() string {
	return w.debugString()
}

var (
	SupportedTmplInfo = map[string]template.Info {
	{{range .TemplateModels}}
		{{.GoPackageName}}.TemplateName: {
			Name: {{.GoPackageName}}.TemplateName,
			Impl: "{{.PackageName}}",
			CtrCfg:  &{{.GoPackageName}}.InstanceParam{},
			Variety:   adptTmpl.{{.VarietyName}},
			BldrInterfaceName:  {{.GoPackageName}}.TemplateName + "." + "HandlerBuilder",
			HndlrInterfaceName: {{.GoPackageName}}.TemplateName + "." + "Handler",
			BuilderSupportsTemplate: func(hndlrBuilder adapter.HandlerBuilder) bool {
				_, ok := hndlrBuilder.({{.GoPackageName}}.HandlerBuilder)
				return ok
			},
			HandlerSupportsTemplate: func(hndlr adapter.Handler) bool {
				_, ok := hndlr.({{.GoPackageName}}.Handler)
				return ok
			},
			InferType: func(cp proto.Message, tEvalFn template.TypeEvalFn) (proto.Message, error) {
				{{$goPkgName := .GoPackageName}}
				{{$varietyName := .VarietyName}}
				{{range getAllMsgs .}}
				{{with $msg := .}}

				{{if ne $varietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}
				var {{getBuildFnName $msg.Name}} func(param *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}},
					path string) (*{{$goPkgName}}.{{getResourcMessageTypeName $msg.Name}}, error)
				{{else}}
				var {{getBuildFnName $msg.Name}} func(param *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}},
					path string) (proto.Message, error)
				{{end}}
				_ = {{getBuildFnName $msg.Name}}
				{{end}}
				{{end}}
				{{range getAllMsgs .}}
				{{with $msg := .}}

				{{if ne $varietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}
				{{getBuildFnName $msg.Name}} = func(param *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}},
					path string) (*{{$goPkgName}}.{{getResourcMessageTypeName $msg.Name}}, error) {
				{{else}}
				{{getBuildFnName $msg.Name}} = func(param *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}},
					path string) (proto.Message, error) {
				{{end}}

				if param == nil {
					return nil, nil
				}
				{{if ne $varietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}
				infrdType := &{{$goPkgName}}.{{getResourcMessageTypeName $msg.Name}}{}
				{{end}}
				var err error = nil

				{{range $msg.Fields}}
					{{if containsValueTypeOrResMsg .GoType}}
						{{if ne $varietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}
						{{if .GoType.IsMap}}
							{{$typeName := getTypeName .GoType.MapValue}}
							{{if .GoType.MapValue.IsResourceMessage}}
								infrdType.{{.GoName}} = make(map[{{.GoType.MapKey.Name}}]*{{$goPkgName}}.{{getResourcMessageTypeName $typeName}}, len(param.{{.GoName}}))
							{{else}}
								infrdType.{{.GoName}} = make(map[{{.GoType.MapKey.Name}}]{{$typeName}}, len(param.{{.GoName}}))
							{{end}}
							for k, v := range param.{{.GoName}} {
							{{if .GoType.MapValue.IsResourceMessage}}
								if infrdType.{{.GoName}}[k], err = {{getBuildFnName $typeName}}(v, path + "{{.GoName}}[" + k + "]."); err != nil {
							{{else}}
								if infrdType.{{.GoName}}[k], err = tEvalFn(v); err != nil {
							{{end}}
									return nil, fmt.Errorf("failed to evaluate expression for field '%s'; %v", path + "{{.GoName}}", err)
								}
							}
						{{else}}
							{{if .GoType.IsResourceMessage}}
								if param.{{.GoName}} != nil {
									{{$typeName := getTypeName .GoType}}
									if infrdType.{{.GoName}}, err = {{getBuildFnName $typeName}}(param.{{.GoName}}, path + "{{.GoName}}."); err != nil {
										return nil, fmt.Errorf("failed to evaluate expression for field '%s'; %v", path + "{{.GoName}}", err)
									}
								}
							{{else}}
								if param.{{.GoName}} == "" || param.{{.GoName}} == emptyQuotes {
									return nil, fmt.Errorf("expression for field '%s' cannot be empty", path + "{{.GoName}}")
								}
								if infrdType.{{.GoName}}, err = tEvalFn(param.{{.GoName}}); err != nil {
									return nil, fmt.Errorf("failed to evaluate expression for field '%s'; %v", path + "{{.GoName}}", err)
								}
							{{end}}
						{{end}}
						{{end}}
					{{else}}
						{{if .GoType.IsMap}}
							for _, v := range param.{{.GoName}} {
								if t, e := tEvalFn(v); e != nil || t != {{getValueType .GoType.MapValue}} {
									if e != nil {
										return nil, fmt.Errorf("failed to evaluate expression for field '%s'; %v", path + "{{.GoName}}", e)
									}
									return nil, fmt.Errorf(
										"error type checking for field '%s': Evaluated expression type %v want %v", path + "{{.GoName}}", t, {{getValueType .GoType.MapValue}})
								}
							}
						{{else}}
							if param.{{.GoName}} == "" || param.{{.GoName}} == emptyQuotes {
								return nil, fmt.Errorf("expression for field '%s' cannot be empty", path + "{{.GoName}}")
							}
							if t, e := tEvalFn(param.{{.GoName}}); e != nil || t != {{getValueType .GoType}} {
								if e != nil {
									return nil, fmt.Errorf("failed to evaluate expression for field '%s': %v", path + "{{.GoName}}", e)
								}
								return nil, fmt.Errorf("error type checking for field '%s': Evaluated expression type %v want %v", path + "{{.GoName}}", t, {{getValueType .GoType}})
							}
						{{end}}
					{{end}}
				{{end}}
				{{if ne $varietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}
				return infrdType, err
				{{else}}
				return nil, err
				{{end}}
				}
				{{end}}
				{{end}}

				instParam := cp.(*{{.GoPackageName}}.InstanceParam)
				{{if eq $varietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}
				const fullOutName = "{{.GoPackageName}}.output."
				for attr, exp := range instParam.AttributeBindings {
					expr := strings.Replace(exp, "$out.", fullOutName, -1)
					t1, err := tEvalFn(expr)
					if err != nil {
						return nil, fmt.Errorf("error evaluating AttributeBinding expression '%s' for attribute '%s': %v", expr, attr, err)
					}
					t2, err := tEvalFn(attr)
					if err != nil {
						return nil, fmt.Errorf("error evaluating AttributeBinding expression for attribute key '%s': %v", attr, err)
					}
					if t1 != t2 {
						return nil, fmt.Errorf("type '%v' for attribute '%s' does not match type '%s' for expression '%s'", t2, attr, t1, expr)
					}
				}
				{{end}}
				return BuildTemplate(instParam, "")
			},
			{{if ne $varietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}
			SetType: func(types map[string]proto.Message, builder adapter.HandlerBuilder) {
				// Mixer framework should have ensured the type safety.
				castedBuilder := builder.({{.GoPackageName}}.HandlerBuilder)
				castedTypes := make(map[string]*{{.GoPackageName}}.Type, len(types))
				for k, v := range types {
					// Mixer framework should have ensured the type safety.
					v1 := v.(*{{.GoPackageName}}.Type)
					castedTypes[k] = v1
				}
				castedBuilder.Set{{.InterfaceName}}Types(castedTypes)
			},
			{{end}}
			{{if eq $varietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}
			{{$goPkgName := .GoPackageName}}
			AttributeManifests: []*istio_mixer_v1_config.AttributeManifest{
				{
					Attributes: map[string]*istio_mixer_v1_config.AttributeManifest_AttributeInfo{
						{{range .OutputTemplateMessage.Fields}}
						"{{$goPkgName}}.output.{{tolower .ProtoName}}": {
							ValueType: {{getValueType .GoType}},
						},
						{{end}}
					},
				},
			},
			{{end}}
			{{if eq .VarietyName "TEMPLATE_VARIETY_REPORT"}}
				ProcessReport: func(ctx context.Context, insts map[string]proto.Message, attrs attribute.Bag, mapper expr.Evaluator, handler adapter.Handler) error {
			{{end}}
			{{if eq .VarietyName "TEMPLATE_VARIETY_CHECK"}}
				ProcessCheck: func(ctx context.Context, instName string, inst proto.Message, attrs attribute.Bag,
				mapper expr.Evaluator, handler adapter.Handler) (adapter.CheckResult, error) {
			{{end}}
			{{if eq .VarietyName "TEMPLATE_VARIETY_QUOTA"}}
				ProcessQuota: func(ctx context.Context, instName string, inst proto.Message, attrs attribute.Bag,
				 mapper expr.Evaluator, handler adapter.Handler, args adapter.QuotaArgs) (adapter.QuotaResult, error) {
			{{end}}
				{{if eq .VarietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}
				ProcessGenAttrs: func(ctx context.Context, instName string, inst proto.Message, attrs attribute.Bag,
				mapper expr.Evaluator, handler adapter.Handler) (*attribute.MutableBag, error) {
			{{end}}
			{{$varietyName := .VarietyName}}
			{{$goPkgName := .GoPackageName}}
			{{range getAllMsgs .}}
			{{with $msg := .}}
			var {{getBuildFnName $msg.Name}} func(instName string,
				param *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}}, path string) (
					*{{$goPkgName}}.{{getResourcMessageInstanceName $msg.Name}}, error)
			_ = {{getBuildFnName $msg.Name}}
			{{end}}
			{{end}}
			{{range getAllMsgs .}}
			{{with $msg := .}}
			{{getBuildFnName $msg.Name}} = func(instName string,
				param *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}}, path string) (
					*{{$goPkgName}}.{{getResourcMessageInstanceName $msg.Name}}, error) {
				if param == nil {
					return nil, nil
				}
				var err error
				_ = err
			{{range $msg.Fields}}
				{{if .GoType.IsMap}}
					{{if .GoType.MapValue.IsResourceMessage}}
						{{$typeName := getTypeName .GoType.MapValue}}
						{{.GoName}} := make(map[{{.GoType.MapKey.Name}}]*{{$goPkgName}}.{{$typeName}}, len(param.{{.GoName}}))
						for k, v := range param.{{.GoName}} {
							if {{.GoName}}[k], err = {{getBuildFnName $typeName}}(instName, v, path + "{{.GoName}}[" + k + "]."); err != nil {
								return nil, fmt.Errorf("failed to evaluate field '%s' for instance '%s': %v", path + "{{.GoName}}", instName, err)
							}
						}
					{{else}}
						{{.GoName}}, err := template.EvalAll(param.{{.GoName}}, attrs, mapper)
					{{end}}
				{{else}}
					{{if .GoType.IsResourceMessage}}
					{{$typeName := getTypeName .GoType}}
					{{.GoName}}, err := {{getBuildFnName $typeName}}(instName, param.{{.GoName}}, path + "{{.GoName}}.")
					{{else }}
					{{.GoName}}, err := mapper.Eval(param.{{.GoName}}, attrs)
					{{end}}
				{{end}}
					if err != nil {
						msg := fmt.Sprintf("failed to evaluate field '%s' for instance '%s': %v", path + "{{.GoName}}", instName, err)
						glog.Error(msg)
						return nil, errors.New(msg)
					}
			{{end}}
				_ = param
				return &{{$goPkgName}}.{{getResourcMessageInstanceName $msg.Name}}{
					{{if eq $msg.Name "Template"}}
					Name:       instName,
					{{end}}
					{{range $msg.Fields}}
						{{if containsValueTypeOrResMsg .GoType}}
							{{.GoName}}: {{.GoName}},
						{{else}}
							{{if .GoType.IsMap}}
								{{.GoName}}: func(m map[string]interface{}) map[string]{{.GoType.MapValue.Name}} {
									res := make(map[string]{{.GoType.MapValue.Name}}, len(m))
									for k, v := range m {
										res[k] = v.({{.GoType.MapValue.Name}})
									}
									return res
								}({{.GoName}}),
							{{else}}
								{{.GoName}}: {{.GoName}}.({{.GoType.Name}}),{{reportTypeUsed .GoType}}
							{{end}}
						{{end}}
					{{end}}
				}, nil
			}
			{{end}}
			{{end}}
			{{if eq .VarietyName "TEMPLATE_VARIETY_REPORT"}}
					var instances []*{{.GoPackageName}}.Instance
					for instName, inst := range insts {
						instance, err := BuildTemplate(instName, inst.(*{{.GoPackageName}}.InstanceParam), "")
						if err != nil {
							return err
						}
						instances = append(instances, instance)
					}

					if err := handler.({{.GoPackageName}}.Handler).Handle{{.InterfaceName}}(ctx, instances); err != nil {
						return fmt.Errorf("failed to report all values: %v", err)
					}
					return nil
				},
			{{else}}
					instParam := inst.(*{{.GoPackageName}}.InstanceParam)
					instance, err := BuildTemplate(instName, instParam, "")
					if err != nil {
						{{if eq $varietyName "TEMPLATE_VARIETY_CHECK"}}
						return adapter.CheckResult{}, err
						{{else if eq $varietyName "TEMPLATE_VARIETY_QUOTA"}}return adapter.QuotaResult{}, err
						{{else if eq $varietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}return nil, err
						{{end}}
					}
					{{if eq $varietyName "TEMPLATE_VARIETY_CHECK"}}return handler.({{.GoPackageName}}.Handler).Handle{{.InterfaceName}}(ctx, instance)
					{{else if eq $varietyName "TEMPLATE_VARIETY_QUOTA"}}return handler.({{.GoPackageName}}.Handler).Handle{{.InterfaceName}}(ctx, instance, args)
					{{else if eq $varietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}
					out, err := handler.({{.GoPackageName}}.Handler).Generate{{.InterfaceName}}Attributes(ctx, instance)
					if err != nil {
						return nil, err
					}
					const fullOutName = "{{.GoPackageName}}.output."
					abag := newWrapperAttrBag(
						func(name string) (value interface{}, found bool) {
							field := strings.TrimPrefix(name, fullOutName)
							if len(field) != len(name) {
								switch field {
									{{range .OutputTemplateMessage.Fields}}
									case "{{tolower .ProtoName}}":
										return out.{{.GoName}}, true
									{{end}}
									default:
									// FIXME : any fields in output (or its references) that are of type
									// map<string, non string fields> are not supported yet
									return nil, false
								}

							}
							return attrs.Get(name)
						},
						func() []string {
							return attrs.Names()
						},
						func() {
							attrs.Done()
						},
						func() string {
							return attrs.DebugString()
						},
					)

					resultBag := attribute.GetMutableBag(nil)
					// TODO validate the content of AttributeBindings during inferType function.
					for attrName, outExpr := range instParam.AttributeBindings {
						ex := strings.Replace(outExpr, "$out.", fullOutName, -1)
						val, err := mapper.Eval(ex, abag)
						if err != nil {
							return nil, err
						}
						switch v := val.(type) {
						case net.IP:
							// conversion to []byte necessary based on current IP_ADDRESS handling within Mixer
							// TODO: remove
							glog.V(4).Info("converting net.IP to []byte")
							if v4 := v.To4(); v4 != nil {
								resultBag.Set(attrName, []byte(v4))
								continue
							}
							resultBag.Set(attrName, []byte(v.To16()))
						default:
							resultBag.Set(attrName, val)
						}
					}
					return resultBag, nil
					{{end}}
				},
			{{end}}
		},
	{{end}}
	}
)

`
