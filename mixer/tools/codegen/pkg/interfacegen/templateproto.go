package interfacegen

// augmentedProtoTmpl defines the modified template proto with Type and InstanceParams
var augmentedProtoTmpl = `// Copyright 2017 Istio Authors
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

syntax = "proto3";

{{.Comment}}{{if ne .TemplateMessage.Comment ""}}
//
{{.TemplateMessage.Comment}}{{end}}
package {{.PackageName}};

import "mixer/adapter/model/v1beta1/extensions.proto";
import "google/protobuf/any.proto";
import "google/rpc/status.proto";
import "mixer/adapter/model/v1beta1/type.proto";

$$additional_imports$$

option (istio.mixer.adapter.model.v1beta1.template_variety) = {{.VarietyName}};

{{if ne .VarietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}

// Handle{{.InterfaceName}}Service is implemented by backends that wants to handle request-time '{{.TemplateName}}' instances.
service Handle{{.InterfaceName}}Service {
    // Handle{{.InterfaceName}} is called by Mixer at request-time to deliver '{{.TemplateName}}' instances to the backend.
    rpc Handle{{.InterfaceName}}(Handle{{.InterfaceName}}Request) returns (Handle{{.InterfaceName}}Response);
}

// Request message for Handle{{.InterfaceName}} method.
message Handle{{.InterfaceName}}Request {

    // '{{.TemplateName}}' instances.
    repeated Type instances = 1;

    // Adapter specific handler configuration.
    //
    // Note: Backends can also implement [InfrastructureBackend][https://istio.io/docs/reference/config/mixer/istio.mixer.adapter.model.v1beta1.html#InfrastructureBackend]
    // service and therefore opt to receive handler configuration during session creation through [InfrastructureBackend.CreateSession][TODO: Link to this fragment]
    // call. In that case, adapter_config will have type_url as 'google.protobuf.Any.type_url' and would contain string
    // value of session_id (returned from InfrastructureBackend.CreateSession).
    google.protobuf.Any adapter_config = 2;

    // Id to dedupe identical requests from Mixer.
    string dedup_id = 3;
}

// Response message for Handle{{.InterfaceName}} method.
message Handle{{.InterfaceName}}Response {
    // The success/failure status of Handle{{.InterfaceName}} call.
    google.rpc.Status status = 1;
}

// Contains instance payload for '{{.TemplateName}}' template. This is passed to infrastructure backends during request-time
// through Handle{{.InterfaceName}}Service.Handle{{.InterfaceName}}.
message InstanceMsg {
  // Name of the instance as specified in configuration.
  string name = 72295727;
  {{range .TemplateMessage.Fields}}
  {{.Comment}}
  {{typeName .ProtoType}} {{.ProtoName}} = {{.Number}};{{reportTypeUsed .ProtoType}}
  {{end}}
}

{{range .ResourceMessages}}
{{.Comment}}
message {{.Name}}Msg {
  {{range .Fields}}
  {{.Comment}}
  {{typeName .ProtoType}} {{.ProtoName}} = {{.Number}};{{reportTypeUsed .ProtoType}}
  {{end}}
}
{{end}}

{{end}}

{{if ne .VarietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}

// Contains inferred type information about specific instance of '{{.TemplateName}}' template. This is passed to
// infrastructure backends during configuration-time through [InfrastructureBackend.CreateSession][TODO: Link to this fragment].
message Type {
  {{range .TemplateMessage.Fields -}}
  {{- if valueTypeOrResMsg .ProtoType}}
  {{.Comment}}
  {{valueTypeOrResMsgFieldTypeName .ProtoType}} {{.ProtoName}} = {{.Number}};{{reportTypeUsed .ProtoType}}
  {{- end}}
  {{- end}}
}

{{range .ResourceMessages}}
{{.Comment}}
message {{getResourcMessageTypeName .Name}} {
  {{range .Fields}}
  {{- if valueTypeOrResMsg .ProtoType}}
  {{.Comment}}
  {{valueTypeOrResMsgFieldTypeName .ProtoType}} {{.ProtoName}} = {{.Number}};{{reportTypeUsed .ProtoType}}
  {{- end}}
  {{- end}}
}
{{end}}

{{end}}

// Represents instance configuration schema for '{{.TemplateName}}' template.
message InstanceParam {
  {{range .TemplateMessage.Fields}}
  {{.Comment}}
  {{stringify .ProtoType}} {{.ProtoName}} = {{.Number}};
  {{end}}
  {{if eq .VarietyName "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR"}}
  // Attribute names to expression mapping. These expressions can use the fields from the output object
  // returned by the attribute producing adapters using $out.<fieldName> notation. For example:
  // source.ip : $out.source_pod_ip
  // In the above example, source.ip attribute will be added to the existing attribute list and its value will be set to
  // the value of source_pod_ip field of the output returned by the adapter.
  map<string, string> attribute_bindings = 72295728;
  {{end}}
}

{{range .ResourceMessages}}
message {{getResourcMessageInterfaceParamTypeName  .Name}} {
  {{range .Fields}}
  {{.Comment}}
  {{stringify .ProtoType}} {{.ProtoName}} = {{.Number}};
  {{end}}
}
{{end}}
`
