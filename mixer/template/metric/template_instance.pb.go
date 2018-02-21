// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/template/metric/template_instance.proto

/*
	Package metric is a generated protocol buffer package.

	It is generated from these files:
		mixer/template/metric/template_instance.proto

	It has these top-level messages:
		Type
		InstanceParam
*/
package metric

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "istio.io/api/mixer/v1/template"
import istio_policy_v1beta1 "istio.io/api/policy/v1beta1"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The `metric` template is designed to let you describe runtime metric to dispatch to
// monitoring backends.
// The `metric` template represents a single piece of data to report.
//
// When writing the configuration, the value for the fields associated with this template can either be a
// literal or an [expression](https://istio.io/docs/reference/config/mixer/expression-language.html). Please note that if the datatype of a field is not istio.mixer.adapter.model.v1beta1.Value,
// then the expression's [inferred type](https://istio.io/docs/reference/config/mixer/expression-language.html#type-checking) must match the datatype of the field.
//
// Example config:
// ```yaml
// apiVersion: "config.istio.io/v1alpha2"
// kind: metric
// metadata:
//   name: requestsize
//   namespace: istio-system
// spec:
//   value: request.size | 0
//   dimensions:
//     source_service: source.service | "unknown"
//     source_version: source.labels["version"] | "unknown"
//     destination_service: destination.service | "unknown"
//     destination_version: destination.labels["version"] | "unknown"
//     response_code: response.code | 200
//   monitored_resource_type: '"UNSPECIFIED"'
// ```
type Type struct {
	// The value being reported.
	Value istio_policy_v1beta1.ValueType `protobuf:"varint,1,opt,name=value,proto3,enum=istio.policy.v1beta1.ValueType" json:"value,omitempty"`
	// The unique identity of the particular metric to report.
	Dimensions map[string]istio_policy_v1beta1.ValueType `protobuf:"bytes,2,rep,name=dimensions" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=istio.policy.v1beta1.ValueType"`
	// Optional. A set of expressions that will form the dimensions of the monitored resource this metric is being reported on.
	// If the metric backend supports monitored resources, these fields are used to populate that resource. Otherwise
	// these fields will be ignored by the adapter.
	MonitoredResourceDimensions map[string]istio_policy_v1beta1.ValueType `protobuf:"bytes,4,rep,name=monitored_resource_dimensions,json=monitoredResourceDimensions" json:"monitored_resource_dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=istio.policy.v1beta1.ValueType"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorTemplateInstance, []int{0} }

func (m *Type) GetValue() istio_policy_v1beta1.ValueType {
	if m != nil {
		return m.Value
	}
	return istio_policy_v1beta1.VALUE_TYPE_UNSPECIFIED
}

func (m *Type) GetDimensions() map[string]istio_policy_v1beta1.ValueType {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Type) GetMonitoredResourceDimensions() map[string]istio_policy_v1beta1.ValueType {
	if m != nil {
		return m.MonitoredResourceDimensions
	}
	return nil
}

type InstanceParam struct {
	Value                       string            `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Dimensions                  map[string]string `protobuf:"bytes,2,rep,name=dimensions" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MonitoredResourceType       string            `protobuf:"bytes,3,opt,name=monitored_resource_type,json=monitoredResourceType,proto3" json:"monitored_resource_type,omitempty"`
	MonitoredResourceDimensions map[string]string `protobuf:"bytes,4,rep,name=monitored_resource_dimensions,json=monitoredResourceDimensions" json:"monitored_resource_dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *InstanceParam) Reset()                    { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage()               {}
func (*InstanceParam) Descriptor() ([]byte, []int) { return fileDescriptorTemplateInstance, []int{1} }

func (m *InstanceParam) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *InstanceParam) GetDimensions() map[string]string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *InstanceParam) GetMonitoredResourceType() string {
	if m != nil {
		return m.MonitoredResourceType
	}
	return ""
}

func (m *InstanceParam) GetMonitoredResourceDimensions() map[string]string {
	if m != nil {
		return m.MonitoredResourceDimensions
	}
	return nil
}

func init() {
	proto.RegisterType((*Type)(nil), "metric.Type")
	proto.RegisterType((*InstanceParam)(nil), "metric.InstanceParam")
}
func (this *Type) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Type)
	if !ok {
		that2, ok := that.(Type)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if len(this.Dimensions) != len(that1.Dimensions) {
		return false
	}
	for i := range this.Dimensions {
		if this.Dimensions[i] != that1.Dimensions[i] {
			return false
		}
	}
	if len(this.MonitoredResourceDimensions) != len(that1.MonitoredResourceDimensions) {
		return false
	}
	for i := range this.MonitoredResourceDimensions {
		if this.MonitoredResourceDimensions[i] != that1.MonitoredResourceDimensions[i] {
			return false
		}
	}
	return true
}
func (this *InstanceParam) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*InstanceParam)
	if !ok {
		that2, ok := that.(InstanceParam)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if len(this.Dimensions) != len(that1.Dimensions) {
		return false
	}
	for i := range this.Dimensions {
		if this.Dimensions[i] != that1.Dimensions[i] {
			return false
		}
	}
	if this.MonitoredResourceType != that1.MonitoredResourceType {
		return false
	}
	if len(this.MonitoredResourceDimensions) != len(that1.MonitoredResourceDimensions) {
		return false
	}
	for i := range this.MonitoredResourceDimensions {
		if this.MonitoredResourceDimensions[i] != that1.MonitoredResourceDimensions[i] {
			return false
		}
	}
	return true
}
func (this *Type) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&metric.Type{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]istio_policy_v1beta1.ValueType{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%#v: %#v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	if this.Dimensions != nil {
		s = append(s, "Dimensions: "+mapStringForDimensions+",\n")
	}
	keysForMonitoredResourceDimensions := make([]string, 0, len(this.MonitoredResourceDimensions))
	for k, _ := range this.MonitoredResourceDimensions {
		keysForMonitoredResourceDimensions = append(keysForMonitoredResourceDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMonitoredResourceDimensions)
	mapStringForMonitoredResourceDimensions := "map[string]istio_policy_v1beta1.ValueType{"
	for _, k := range keysForMonitoredResourceDimensions {
		mapStringForMonitoredResourceDimensions += fmt.Sprintf("%#v: %#v,", k, this.MonitoredResourceDimensions[k])
	}
	mapStringForMonitoredResourceDimensions += "}"
	if this.MonitoredResourceDimensions != nil {
		s = append(s, "MonitoredResourceDimensions: "+mapStringForMonitoredResourceDimensions+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InstanceParam) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&metric.InstanceParam{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]string{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%#v: %#v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	if this.Dimensions != nil {
		s = append(s, "Dimensions: "+mapStringForDimensions+",\n")
	}
	s = append(s, "MonitoredResourceType: "+fmt.Sprintf("%#v", this.MonitoredResourceType)+",\n")
	keysForMonitoredResourceDimensions := make([]string, 0, len(this.MonitoredResourceDimensions))
	for k, _ := range this.MonitoredResourceDimensions {
		keysForMonitoredResourceDimensions = append(keysForMonitoredResourceDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMonitoredResourceDimensions)
	mapStringForMonitoredResourceDimensions := "map[string]string{"
	for _, k := range keysForMonitoredResourceDimensions {
		mapStringForMonitoredResourceDimensions += fmt.Sprintf("%#v: %#v,", k, this.MonitoredResourceDimensions[k])
	}
	mapStringForMonitoredResourceDimensions += "}"
	if this.MonitoredResourceDimensions != nil {
		s = append(s, "MonitoredResourceDimensions: "+mapStringForMonitoredResourceDimensions+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTemplateInstance(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTemplateInstance(dAtA, i, uint64(m.Value))
	}
	if len(m.Dimensions) > 0 {
		for k, _ := range m.Dimensions {
			dAtA[i] = 0x12
			i++
			v := m.Dimensions[k]
			mapSize := 1 + len(k) + sovTemplateInstance(uint64(len(k))) + 1 + sovTemplateInstance(uint64(v))
			i = encodeVarintTemplateInstance(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintTemplateInstance(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintTemplateInstance(dAtA, i, uint64(v))
		}
	}
	if len(m.MonitoredResourceDimensions) > 0 {
		for k, _ := range m.MonitoredResourceDimensions {
			dAtA[i] = 0x22
			i++
			v := m.MonitoredResourceDimensions[k]
			mapSize := 1 + len(k) + sovTemplateInstance(uint64(len(k))) + 1 + sovTemplateInstance(uint64(v))
			i = encodeVarintTemplateInstance(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintTemplateInstance(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintTemplateInstance(dAtA, i, uint64(v))
		}
	}
	return i, nil
}

func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateInstance(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if len(m.Dimensions) > 0 {
		for k, _ := range m.Dimensions {
			dAtA[i] = 0x12
			i++
			v := m.Dimensions[k]
			mapSize := 1 + len(k) + sovTemplateInstance(uint64(len(k))) + 1 + len(v) + sovTemplateInstance(uint64(len(v)))
			i = encodeVarintTemplateInstance(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintTemplateInstance(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintTemplateInstance(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.MonitoredResourceType) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTemplateInstance(dAtA, i, uint64(len(m.MonitoredResourceType)))
		i += copy(dAtA[i:], m.MonitoredResourceType)
	}
	if len(m.MonitoredResourceDimensions) > 0 {
		for k, _ := range m.MonitoredResourceDimensions {
			dAtA[i] = 0x22
			i++
			v := m.MonitoredResourceDimensions[k]
			mapSize := 1 + len(k) + sovTemplateInstance(uint64(len(k))) + 1 + len(v) + sovTemplateInstance(uint64(len(v)))
			i = encodeVarintTemplateInstance(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintTemplateInstance(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintTemplateInstance(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintTemplateInstance(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Type) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovTemplateInstance(uint64(m.Value))
	}
	if len(m.Dimensions) > 0 {
		for k, v := range m.Dimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTemplateInstance(uint64(len(k))) + 1 + sovTemplateInstance(uint64(v))
			n += mapEntrySize + 1 + sovTemplateInstance(uint64(mapEntrySize))
		}
	}
	if len(m.MonitoredResourceDimensions) > 0 {
		for k, v := range m.MonitoredResourceDimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTemplateInstance(uint64(len(k))) + 1 + sovTemplateInstance(uint64(v))
			n += mapEntrySize + 1 + sovTemplateInstance(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *InstanceParam) Size() (n int) {
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTemplateInstance(uint64(l))
	}
	if len(m.Dimensions) > 0 {
		for k, v := range m.Dimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTemplateInstance(uint64(len(k))) + 1 + len(v) + sovTemplateInstance(uint64(len(v)))
			n += mapEntrySize + 1 + sovTemplateInstance(uint64(mapEntrySize))
		}
	}
	l = len(m.MonitoredResourceType)
	if l > 0 {
		n += 1 + l + sovTemplateInstance(uint64(l))
	}
	if len(m.MonitoredResourceDimensions) > 0 {
		for k, v := range m.MonitoredResourceDimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTemplateInstance(uint64(len(k))) + 1 + len(v) + sovTemplateInstance(uint64(len(v)))
			n += mapEntrySize + 1 + sovTemplateInstance(uint64(mapEntrySize))
		}
	}
	return n
}

func sovTemplateInstance(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTemplateInstance(x uint64) (n int) {
	return sovTemplateInstance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]istio_policy_v1beta1.ValueType{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%v: %v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	keysForMonitoredResourceDimensions := make([]string, 0, len(this.MonitoredResourceDimensions))
	for k, _ := range this.MonitoredResourceDimensions {
		keysForMonitoredResourceDimensions = append(keysForMonitoredResourceDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMonitoredResourceDimensions)
	mapStringForMonitoredResourceDimensions := "map[string]istio_policy_v1beta1.ValueType{"
	for _, k := range keysForMonitoredResourceDimensions {
		mapStringForMonitoredResourceDimensions += fmt.Sprintf("%v: %v,", k, this.MonitoredResourceDimensions[k])
	}
	mapStringForMonitoredResourceDimensions += "}"
	s := strings.Join([]string{`&Type{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Dimensions:` + mapStringForDimensions + `,`,
		`MonitoredResourceDimensions:` + mapStringForMonitoredResourceDimensions + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]string{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%v: %v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	keysForMonitoredResourceDimensions := make([]string, 0, len(this.MonitoredResourceDimensions))
	for k, _ := range this.MonitoredResourceDimensions {
		keysForMonitoredResourceDimensions = append(keysForMonitoredResourceDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMonitoredResourceDimensions)
	mapStringForMonitoredResourceDimensions := "map[string]string{"
	for _, k := range keysForMonitoredResourceDimensions {
		mapStringForMonitoredResourceDimensions += fmt.Sprintf("%v: %v,", k, this.MonitoredResourceDimensions[k])
	}
	mapStringForMonitoredResourceDimensions += "}"
	s := strings.Join([]string{`&InstanceParam{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Dimensions:` + mapStringForDimensions + `,`,
		`MonitoredResourceType:` + fmt.Sprintf("%v", this.MonitoredResourceType) + `,`,
		`MonitoredResourceDimensions:` + mapStringForMonitoredResourceDimensions + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTemplateInstance(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateInstance
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Type: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Type: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateInstance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= (istio_policy_v1beta1.ValueType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateInstance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateInstance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dimensions == nil {
				m.Dimensions = make(map[string]istio_policy_v1beta1.ValueType)
			}
			var mapkey string
			var mapvalue istio_policy_v1beta1.ValueType
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTemplateInstance
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTemplateInstance
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTemplateInstance
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTemplateInstance
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (istio_policy_v1beta1.ValueType(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTemplateInstance(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTemplateInstance
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Dimensions[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonitoredResourceDimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateInstance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateInstance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MonitoredResourceDimensions == nil {
				m.MonitoredResourceDimensions = make(map[string]istio_policy_v1beta1.ValueType)
			}
			var mapkey string
			var mapvalue istio_policy_v1beta1.ValueType
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTemplateInstance
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTemplateInstance
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTemplateInstance
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTemplateInstance
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (istio_policy_v1beta1.ValueType(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTemplateInstance(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTemplateInstance
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MonitoredResourceDimensions[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateInstance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateInstance
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InstanceParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateInstance
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstanceParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateInstance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateInstance
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateInstance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateInstance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dimensions == nil {
				m.Dimensions = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTemplateInstance
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTemplateInstance
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTemplateInstance
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTemplateInstance
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTemplateInstance
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTemplateInstance(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTemplateInstance
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Dimensions[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonitoredResourceType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateInstance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateInstance
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MonitoredResourceType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonitoredResourceDimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateInstance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateInstance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MonitoredResourceDimensions == nil {
				m.MonitoredResourceDimensions = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTemplateInstance
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTemplateInstance
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTemplateInstance
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTemplateInstance
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTemplateInstance
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTemplateInstance(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTemplateInstance
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MonitoredResourceDimensions[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateInstance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateInstance
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTemplateInstance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTemplateInstance
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTemplateInstance
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTemplateInstance
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTemplateInstance
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTemplateInstance
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTemplateInstance(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTemplateInstance = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTemplateInstance   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/template/metric/template_instance.proto", fileDescriptorTemplateInstance)
}

var fileDescriptorTemplateInstance = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcd, 0xcd, 0xac, 0x48,
	0x2d, 0xd2, 0x2f, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0xd5, 0xcf, 0x4d, 0x2d, 0x29, 0xca,
	0x4c, 0x86, 0xf3, 0xe3, 0x33, 0xf3, 0x8a, 0x4b, 0x12, 0xf3, 0x92, 0x53, 0xf5, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0xd8, 0x20, 0xf2, 0x52, 0x3a, 0x10, 0x6d, 0x89, 0x29, 0x89, 0x05, 0x25, 0xa9,
	0x45, 0xfa, 0xb9, 0xf9, 0x29, 0xa9, 0x39, 0xfa, 0x65, 0x86, 0x49, 0xa9, 0x25, 0x89, 0x86, 0xfa,
	0xa9, 0x15, 0x25, 0xa9, 0x79, 0xc5, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x5d, 0x52, 0xf2, 0x05, 0xf9,
	0x39, 0x99, 0xc9, 0x95, 0x70, 0x05, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0xf1, 0x25, 0x95, 0x05, 0x50,
	0x63, 0x95, 0xf6, 0x33, 0x73, 0xb1, 0x84, 0x54, 0x16, 0xa4, 0x0a, 0x99, 0x72, 0xb1, 0x82, 0x25,
	0x25, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0xe4, 0xf5, 0x32, 0x8b, 0x4b, 0x32, 0xf3, 0xf5, 0x20,
	0xfa, 0xf5, 0xa0, 0xfa, 0xf5, 0xc2, 0x40, 0x4a, 0x40, 0xea, 0x83, 0x20, 0xaa, 0x85, 0x6c, 0xb8,
	0xb8, 0x52, 0x32, 0x73, 0xa1, 0x96, 0x4a, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x1b, 0xc9, 0xe8, 0x41,
	0xdc, 0xaa, 0x07, 0x52, 0xa8, 0xe7, 0x02, 0x97, 0x76, 0xcd, 0x2b, 0x29, 0xaa, 0x0c, 0x42, 0x52,
	0x2f, 0x54, 0xc8, 0x25, 0x9b, 0x9b, 0x9f, 0x97, 0x59, 0x92, 0x5f, 0x94, 0x9a, 0x12, 0x5f, 0x94,
	0x5a, 0x9c, 0x5f, 0x5a, 0x94, 0x9c, 0x1a, 0x8f, 0x64, 0x20, 0x0b, 0xd8, 0x40, 0x5d, 0x14, 0x03,
	0x7d, 0x61, 0x3a, 0x82, 0xa0, 0x1a, 0xd0, 0x6d, 0x90, 0xce, 0xc5, 0xad, 0x42, 0x2a, 0x8e, 0x8b,
	0x1f, 0x4d, 0xbd, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0x25, 0xd8, 0xe3, 0x9c, 0x41, 0x20, 0x26,
	0x22, 0x30, 0x98, 0x48, 0x09, 0x0c, 0x2b, 0x26, 0x0b, 0x46, 0xa9, 0x7c, 0x2e, 0x05, 0x42, 0x0e,
	0xa4, 0xaa, 0x85, 0x4a, 0xeb, 0x99, 0xb9, 0x78, 0x3d, 0xa1, 0x69, 0x25, 0x20, 0xb1, 0x28, 0x31,
	0x57, 0x48, 0x04, 0x39, 0x2a, 0x39, 0x61, 0x31, 0xe5, 0x8a, 0x25, 0xa6, 0x54, 0x61, 0x01, 0x8b,
	0x62, 0x00, 0xde, 0x28, 0x33, 0xe3, 0x12, 0xc7, 0x12, 0x65, 0xa0, 0x14, 0x25, 0xc1, 0x0c, 0xb6,
	0x4e, 0x14, 0x23, 0xf4, 0xc1, 0xe9, 0xab, 0x8a, 0xb8, 0xa8, 0x36, 0xc3, 0xee, 0x22, 0xca, 0xe2,
	0xdc, 0x96, 0x98, 0x38, 0x17, 0x41, 0x8e, 0x02, 0x4e, 0xe4, 0x28, 0xf5, 0x23, 0x2b, 0x4a, 0x71,
	0x9a, 0xe7, 0x64, 0x74, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31,
	0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x7e, 0x5c, 0x7a, 0x32, 0x99, 0x89, 0x31, 0x89, 0x0d, 0x9c, 0x5d, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x10, 0xe2, 0xde, 0x20, 0x36, 0x04, 0x00, 0x00,
}
