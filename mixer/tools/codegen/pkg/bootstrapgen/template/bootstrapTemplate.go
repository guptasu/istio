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
	"istio.io/istio/mixer/pkg/attribute"
	"istio.io/istio/mixer/pkg/expr"
	"istio.io/istio/mixer/pkg/adapter"
	"istio.io/api/mixer/v1/config/descriptor"
	"istio.io/istio/mixer/pkg/template"
	"github.com/golang/glog"
	adptTmpl "istio.io/api/mixer/v1/template"
	"errors"
	{{range .TemplateModels}}
		"{{.PackageImportPath}}"
	{{end}}
	$$additional_imports$$
)


const emptyQuotes = "\"\""
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
				{{range getAllMsgs .}}
				{{with $msg := .}}
				var Build{{$msg.Name}} func(cpb *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}}) (*{{$goPkgName}}.{{getResourcMessageTypeName $msg.Name}}, error)
                _ = Build{{$msg.Name}}
				{{end}}
				{{end}}

				{{range getAllMsgs .}}
				{{with $msg := .}}
				Build{{$msg.Name}} = func(cpb *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}}) (*{{$goPkgName}}.{{getResourcMessageTypeName $msg.Name}}, error) {
				infrdType := &{{$goPkgName}}.{{getResourcMessageTypeName $msg.Name}}{}
				var err error = nil

				{{range $msg.Fields}}
					{{if containsValueTypeOrResMsg .GoType}}
						{{if .GoType.IsMap}}
							infrdType.{{.GoName}} = make(map[{{.GoType.MapKey.Name}}]istio_mixer_v1_config_descriptor.ValueType, len(cpb.{{.GoName}}))
							for k, v := range cpb.{{.GoName}} {
							{{if .GoType.MapValue.IsResourceMessage}}
								if infrdType.{{.GoName}}[k], err = Build{{getBuildTypeFnName .GoType.MapValue}}(v); err != nil {
									return nil, err
								}
							{{else}}
								if infrdType.{{.GoName}}[k], err = tEvalFn(v); err != nil {
									return nil, err
								}
							{{end}}
							}
						{{else}}
							{{if .GoType.IsResourceMessage}}
								if cpb.{{.GoName}} != nil {
									if infrdType.{{.GoName}}, err = Build{{getBuildTypeFnName .GoType}}(cpb.{{.GoName}}); err != nil {
										return nil, err
									}
								}
							{{else}}
								if cpb.{{.GoName}} == "" || cpb.{{.GoName}} == emptyQuotes {
									return nil, errors.New("expression for field {{.GoName}} cannot be empty")
								}
								if infrdType.{{.GoName}}, err = tEvalFn(cpb.{{.GoName}}); err != nil {
									return nil, err
								}
							{{end}}
						{{end}}
					{{else}}
						{{if .GoType.IsMap}}
							for _, v := range cpb.{{.GoName}} {
								if t, e := tEvalFn(v); e != nil || t != {{getValueType .GoType.MapValue}} {
									if e != nil {
										return nil, fmt.Errorf("failed to evaluate expression for field {{.GoName}}: %v", e)
									}
									return nil, fmt.Errorf("error type checking for field {{.GoName}}: Evaluated expression type %v want %v", t, {{getValueType .GoType.MapValue}})
								}
							}
						{{else}}
							if cpb.{{.GoName}} == "" || cpb.{{.GoName}} == emptyQuotes {
								return nil, errors.New("expression for field {{.GoName}} cannot be empty")
							}
							if t, e := tEvalFn(cpb.{{.GoName}}); e != nil || t != {{getValueType .GoType}} {
								if e != nil {
									return nil, fmt.Errorf("failed to evaluate expression for field {{.GoName}}: %v", e)
								}
								return nil, fmt.Errorf("error type checking for field {{.GoName}}: Evaluated expression type %v want %v", t, {{getValueType .GoType}})
							}
						{{end}}
					{{end}}
				{{end}}
                return infrdType, err
				}
				{{end}}
				{{end}}

                return BuildTemplate(cp.(*{{.GoPackageName}}.InstanceParam))
			},
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
			{{if eq .VarietyName "TEMPLATE_VARIETY_REPORT"}}
				ProcessReport: func(ctx context.Context, insts map[string]proto.Message, attrs attribute.Bag, mapper expr.Evaluator, handler adapter.Handler) error {
					{{$goPkgName := .GoPackageName}}
					{{range getAllMsgs .}}
					{{with $msg := .}}
					var Build{{$msg.Name}} func(instName string, md *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}}) (*{{$goPkgName}}.{{getResourcMessageInstanceName $msg.Name}}, error)
    	            _ = Build{{$msg.Name}}
					{{end}}
					{{end}}

					{{range getAllMsgs .}}
					{{with $msg := .}}
					Build{{$msg.Name}} = func(instName string, md *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}}) (*{{$goPkgName}}.{{getResourcMessageInstanceName $msg.Name}}, error) {
						if md == nil {
							return nil, nil
						}
						var err error
						_ = err
						{{range $msg.Fields}}
							{{if .GoType.IsMap}}
								{{if .GoType.MapValue.IsResourceMessage}}
									{{.GoName}} = make(map[{{.GoType.MapKey.Name}}]{{.GoType.MapValue.Name}}, len(md.{{.GoName}}))
									for k, v := range md.{{.GoName}} {
										if {{.GoName}}[k], err = Build{{getBuildTypeFnName .GoType.MapValue}}(v); err != nil {
											return nil, err
										}
									}
								{{else}}
									{{.GoName}}, err := template.EvalAll(md.{{.GoName}}, attrs, mapper)
								{{end}}
							{{else}}
								{{if .GoType.IsResourceMessage}}
								{{.GoName}}, err := Build{{getBuildTypeFnName .GoType}}(instName, md.{{.GoName}})
								{{else }}
								{{.GoName}}, err := mapper.Eval(md.{{.GoName}}, attrs)
								{{end}}
							{{end}}
								if err != nil {
									msg := fmt.Sprintf("failed to eval {{.GoName}} for instance '%s': %v", instName, err)
									glog.Error(msg)
									return nil, errors.New(msg)
								}
						{{end}}

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

					var instances []*{{.GoPackageName}}.Instance
					for instName, inst := range insts {
						instance, err := BuildTemplate(instName, inst.(*{{.GoPackageName}}.InstanceParam))
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
				{{$varietyName := .VarietyName}}
				{{if eq $varietyName "TEMPLATE_VARIETY_CHECK"}}
				ProcessCheck: func(ctx context.Context, instName string, inst proto.Message, attrs attribute.Bag,
				mapper expr.Evaluator, handler adapter.Handler) (adapter.CheckResult, error) {
				{{else}}
				ProcessQuota: func(ctx context.Context, instName string, inst proto.Message, attrs attribute.Bag,
				 mapper expr.Evaluator, handler adapter.Handler, args adapter.QuotaArgs) (adapter.QuotaResult, error) {
				{{end}}

					{{$goPkgName := .GoPackageName}}
					{{range getAllMsgs .}}
					{{with $msg := .}}
					var Build{{$msg.Name}} func(instName string, md *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}}) (*{{$goPkgName}}.{{getResourcMessageInstanceName $msg.Name}}, error)
    	            _ = Build{{$msg.Name}}
					{{end}}
					{{end}}

					{{range getAllMsgs .}}
					{{with $msg := .}}
					Build{{$msg.Name}} = func(instName string, md *{{$goPkgName}}.{{getResourcMessageInterfaceParamTypeName $msg.Name}}) (*{{$goPkgName}}.{{getResourcMessageInstanceName $msg.Name}}, error) {
						if md == nil {
							return nil, nil
						}
						var err error
						_ = err
					{{range $msg.Fields}}
						{{if .GoType.IsMap}}
							{{if .GoType.MapValue.IsResourceMessage}}
								{{.GoName}} = make(map[{{.GoType.MapKey.Name}}]{{.GoType.MapValue.Name}}, len(md.{{.GoName}}))
								for k, v := range md.{{.GoName}} {
									if {{.GoName}}[k], err = Build{{getBuildTypeFnName .GoType.MapValue}}(v); err != nil {
										return nil, err
									}
								}
							{{else}}
								{{.GoName}}, err := template.EvalAll(md.{{.GoName}}, attrs, mapper)
							{{end}}
						{{else}}
							{{if .GoType.IsResourceMessage}}
							{{.GoName}}, err := Build{{getBuildTypeFnName .GoType}}(instName, md.{{.GoName}})
							{{else }}
							{{.GoName}}, err := mapper.Eval(md.{{.GoName}}, attrs)
							{{end}}
						{{end}}
							if err != nil {
								msg := fmt.Sprintf("failed to eval {{.GoName}} for instance '%s': %v", instName, err)
								glog.Error(msg)
								return nil, errors.New(msg)
							}
					{{end}}
						_ = md
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

					instance, err := BuildTemplate(instName, inst.(*{{.GoPackageName}}.InstanceParam))
					if err != nil {
						{{if eq $varietyName "TEMPLATE_VARIETY_CHECK"}}
						return adapter.CheckResult{}, err
						{{else}}return adapter.QuotaResult{}, err
						{{end}}
					}
					{{if eq $varietyName "TEMPLATE_VARIETY_CHECK"}}return handler.({{.GoPackageName}}.Handler).Handle{{.InterfaceName}}(ctx, instance)
					{{else}}return handler.({{.GoPackageName}}.Handler).Handle{{.InterfaceName}}(ctx, instance, args)
					{{end}}
				},
			{{end}}

		},
	{{end}}
	}
)

`
