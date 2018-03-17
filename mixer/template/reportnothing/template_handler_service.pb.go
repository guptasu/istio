// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/template/reportnothing/template_handler_service.proto

/*
	Package reportnothing is a generated protocol buffer package.

	The `reportnothing` template represents an empty block of data, which can useful
	in different testing scenarios.

	ReportNothing represents an empty block of data that is used for Report-capable
	adapters which don't require any parameters. This is primarily intended for testing
	scenarios.

	Example config:
	```yaml
	apiVersion: "config.istio.io/v1alpha2"
	kind: reportnothing
	metadata:
	  name: reportrequest
	  namespace: istio-system
	spec:
	```

	It is generated from these files:
		mixer/template/reportnothing/template_handler_service.proto

	It has these top-level messages:
		HandleReportNothingRequest
		InstanceMsg
		Type
		InstanceParam
*/
package reportnothing

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "istio.io/api/mixer/adapter/model/v1beta1"
import google_protobuf1 "google/protobuf"
import _ "google/rpc"
import _ "istio.io/api/mixer/adapter/model/v1beta1"
import _ "istio.io/api/mixer/adapter/model/v1beta1"
import _ "istio.io/api/mixer/adapter/model/v1beta1"
import _ "istio.io/api/mixer/adapter/model/v1beta1"

import strings "strings"
import reflect "reflect"

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

// Request message for HandleReportNothing method.
type HandleReportNothingRequest struct {
	// 'reportnothing' instances.
	Instances []*Type `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
	// Adapter specific handler configuration.
	//
	// Note: Backends can also implement [InfrastructureBackend][https://istio.io/docs/reference/config/mixer/istio.mixer.adapter.model.v1beta1.html#InfrastructureBackend]
	// service and therefore opt to receive handler configuration during session creation through [InfrastructureBackend.CreateSession][TODO: Link to this fragment]
	// call. In that case, adapter_config will have type_url as 'google.protobuf.Any.type_url' and would contain string
	// value of session_id (returned from InfrastructureBackend.CreateSession).
	AdapterConfig *google_protobuf1.Any `protobuf:"bytes,2,opt,name=adapter_config,json=adapterConfig" json:"adapter_config,omitempty"`
	// Id to dedupe identical requests from Mixer.
	DedupId string `protobuf:"bytes,3,opt,name=dedup_id,json=dedupId,proto3" json:"dedup_id,omitempty"`
}

func (m *HandleReportNothingRequest) Reset()      { *m = HandleReportNothingRequest{} }
func (*HandleReportNothingRequest) ProtoMessage() {}
func (*HandleReportNothingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{0}
}

func (m *HandleReportNothingRequest) GetInstances() []*Type {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *HandleReportNothingRequest) GetAdapterConfig() *google_protobuf1.Any {
	if m != nil {
		return m.AdapterConfig
	}
	return nil
}

func (m *HandleReportNothingRequest) GetDedupId() string {
	if m != nil {
		return m.DedupId
	}
	return ""
}

// Contains instance payload for 'reportnothing' template. This is passed to infrastructure backends during request-time
// through HandleReportNothingService.HandleReportNothing.
type InstanceMsg struct {
	// Name of the instance as specified in configuration.
	Name string `protobuf:"bytes,72295727,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *InstanceMsg) Reset()      { *m = InstanceMsg{} }
func (*InstanceMsg) ProtoMessage() {}
func (*InstanceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{1}
}

func (m *InstanceMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Contains inferred type information about specific instance of 'reportnothing' template. This is passed to
// infrastructure backends during configuration-time through [InfrastructureBackend.CreateSession][TODO: Link to this fragment].
type Type struct {
}

func (m *Type) Reset()                    { *m = Type{} }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorTemplateHandlerService, []int{2} }

// Represents instance configuration schema for 'reportnothing' template.
type InstanceParam struct {
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{3}
}

func init() {
	proto.RegisterType((*HandleReportNothingRequest)(nil), "reportnothing.HandleReportNothingRequest")
	proto.RegisterType((*InstanceMsg)(nil), "reportnothing.InstanceMsg")
	proto.RegisterType((*Type)(nil), "reportnothing.Type")
	proto.RegisterType((*InstanceParam)(nil), "reportnothing.InstanceParam")
}
func (this *HandleReportNothingRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*HandleReportNothingRequest)
	if !ok {
		that2, ok := that.(HandleReportNothingRequest)
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
	if len(this.Instances) != len(that1.Instances) {
		return false
	}
	for i := range this.Instances {
		if !this.Instances[i].Equal(that1.Instances[i]) {
			return false
		}
	}
	if !this.AdapterConfig.Equal(that1.AdapterConfig) {
		return false
	}
	if this.DedupId != that1.DedupId {
		return false
	}
	return true
}
func (this *InstanceMsg) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*InstanceMsg)
	if !ok {
		that2, ok := that.(InstanceMsg)
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
	if this.Name != that1.Name {
		return false
	}
	return true
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
	return true
}
func (this *HandleReportNothingRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&reportnothing.HandleReportNothingRequest{")
	if this.Instances != nil {
		s = append(s, "Instances: "+fmt.Sprintf("%#v", this.Instances)+",\n")
	}
	if this.AdapterConfig != nil {
		s = append(s, "AdapterConfig: "+fmt.Sprintf("%#v", this.AdapterConfig)+",\n")
	}
	s = append(s, "DedupId: "+fmt.Sprintf("%#v", this.DedupId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InstanceMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&reportnothing.InstanceMsg{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Type) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&reportnothing.Type{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InstanceParam) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&reportnothing.InstanceParam{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTemplateHandlerService(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *HandleReportNothingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandleReportNothingRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Instances) > 0 {
		for _, msg := range m.Instances {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTemplateHandlerService(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.AdapterConfig != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.AdapterConfig.Size()))
		n1, err := m.AdapterConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.DedupId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.DedupId)))
		i += copy(dAtA[i:], m.DedupId)
	}
	return i, nil
}

func (m *InstanceMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xfa
		i++
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0xe4
		i++
		dAtA[i] = 0x93
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
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
	return i, nil
}

func encodeVarintTemplateHandlerService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HandleReportNothingRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Instances) > 0 {
		for _, e := range m.Instances {
			l = e.Size()
			n += 1 + l + sovTemplateHandlerService(uint64(l))
		}
	}
	if m.AdapterConfig != nil {
		l = m.AdapterConfig.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.DedupId)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *InstanceMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 5 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *Type) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *InstanceParam) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovTemplateHandlerService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTemplateHandlerService(x uint64) (n int) {
	return sovTemplateHandlerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HandleReportNothingRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HandleReportNothingRequest{`,
		`Instances:` + strings.Replace(fmt.Sprintf("%v", this.Instances), "Type", "Type", 1) + `,`,
		`AdapterConfig:` + strings.Replace(fmt.Sprintf("%v", this.AdapterConfig), "Any", "google_protobuf1.Any", 1) + `,`,
		`DedupId:` + fmt.Sprintf("%v", this.DedupId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceMsg{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Type{`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceParam{`,
		`}`,
	}, "")
	return s
}
func valueToStringTemplateHandlerService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HandleReportNothingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
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
			return fmt.Errorf("proto: HandleReportNothingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandleReportNothingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instances = append(m.Instances, &Type{})
			if err := m.Instances[len(m.Instances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdapterConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AdapterConfig == nil {
				m.AdapterConfig = &google_protobuf1.Any{}
			}
			if err := m.AdapterConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DedupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DedupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func (m *InstanceMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
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
			return fmt.Errorf("proto: InstanceMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 72295727:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
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
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
				return ErrIntOverflowTemplateHandlerService
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
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func skipTemplateHandlerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTemplateHandlerService
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
					return 0, ErrIntOverflowTemplateHandlerService
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
					return 0, ErrIntOverflowTemplateHandlerService
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
				return 0, ErrInvalidLengthTemplateHandlerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTemplateHandlerService
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
				next, err := skipTemplateHandlerService(dAtA[start:])
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
	ErrInvalidLengthTemplateHandlerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTemplateHandlerService   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/template/reportnothing/template_handler_service.proto", fileDescriptorTemplateHandlerService)
}

var fileDescriptorTemplateHandlerService = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0xe3, 0xb6, 0x2a, 0xd4, 0x51, 0x40, 0xba, 0x82, 0x48, 0x6f, 0xb0, 0xa2, 0x13, 0x82,
	0x20, 0x21, 0x5b, 0x09, 0x63, 0x27, 0x60, 0xa1, 0x03, 0x08, 0x1d, 0xec, 0x91, 0x73, 0xf7, 0x7a,
	0xb1, 0xb8, 0xb3, 0x2f, 0xb6, 0xaf, 0x6a, 0x36, 0xbe, 0x01, 0x48, 0x7c, 0x08, 0xd8, 0xf8, 0x02,
	0x7c, 0x00, 0xc4, 0x54, 0x31, 0x31, 0x92, 0x83, 0x81, 0xb1, 0x23, 0x23, 0xc2, 0x76, 0x40, 0x45,
	0x25, 0xea, 0xf8, 0xee, 0xfd, 0x9e, 0xdf, 0xfb, 0xff, 0x74, 0x78, 0xbf, 0x12, 0xc7, 0xa0, 0x99,
	0x85, 0xaa, 0x2e, 0xb9, 0x05, 0xa6, 0xa1, 0x56, 0xda, 0x4a, 0x65, 0x67, 0x42, 0x16, 0x7f, 0x3e,
	0x4f, 0x66, 0x5c, 0xe6, 0x25, 0xe8, 0x89, 0x01, 0x7d, 0x24, 0x32, 0xa0, 0xb5, 0x56, 0x56, 0x45,
	0xbd, 0x33, 0x74, 0x7c, 0xd7, 0xbf, 0xc5, 0x73, 0x5e, 0x5b, 0xd0, 0xac, 0x52, 0x39, 0x94, 0xec,
	0x68, 0x34, 0x05, 0xcb, 0x47, 0x0c, 0x8e, 0x2d, 0x48, 0x23, 0x94, 0x34, 0x7e, 0x38, 0xde, 0x2b,
	0x94, 0x2a, 0x4a, 0x60, 0xae, 0x9a, 0x36, 0x87, 0x8c, 0xcb, 0x45, 0x68, 0xdd, 0x08, 0x2d, 0x5d,
	0x67, 0xcc, 0x58, 0x6e, 0x9b, 0xd5, 0xcc, 0xad, 0x75, 0x1b, 0xec, 0xa2, 0x0e, 0x87, 0xc5, 0xb7,
	0xd7, 0x71, 0xd9, 0x0c, 0xb2, 0x17, 0x01, 0x1c, 0xae, 0x03, 0x7d, 0xba, 0x8b, 0x3c, 0x39, 0x6f,
	0x94, 0xe5, 0x1e, 0x4c, 0xde, 0x22, 0x1c, 0x3f, 0x72, 0xba, 0x52, 0x37, 0xff, 0xc4, 0xdb, 0x49,
	0x61, 0xde, 0x80, 0xb1, 0xd1, 0x08, 0xef, 0x08, 0x69, 0x2c, 0x97, 0x19, 0x98, 0x3e, 0x1a, 0x6c,
	0x0e, 0xbb, 0xe3, 0x5d, 0x7a, 0xc6, 0x23, 0x7d, 0xbe, 0xa8, 0x21, 0xfd, 0x4b, 0x45, 0xfb, 0xf8,
	0x4a, 0x58, 0x3b, 0xc9, 0x94, 0x3c, 0x14, 0x45, 0x7f, 0x63, 0x80, 0x86, 0xdd, 0xf1, 0x35, 0xea,
	0x3d, 0xd1, 0x95, 0x42, 0x7a, 0x5f, 0x2e, 0xd2, 0x5e, 0x60, 0x1f, 0x3a, 0x34, 0xda, 0xc3, 0x97,
	0x73, 0xc8, 0x9b, 0x7a, 0x22, 0xf2, 0xfe, 0xe6, 0x00, 0x0d, 0x77, 0xd2, 0x4b, 0xae, 0x3e, 0xc8,
	0x93, 0x9b, 0xb8, 0x7b, 0x10, 0x96, 0x3c, 0x36, 0x45, 0x74, 0x1d, 0x6f, 0x49, 0x5e, 0x41, 0xff,
	0xfd, 0xa7, 0x0f, 0x89, 0x03, 0x5d, 0x99, 0x6c, 0xe3, 0xad, 0xdf, 0x07, 0x25, 0x57, 0x71, 0x6f,
	0x45, 0x3f, 0xe5, 0x9a, 0x57, 0xe3, 0x57, 0xe7, 0x07, 0x7d, 0xe6, 0x7f, 0x91, 0x68, 0x8e, 0x77,
	0xcf, 0xe9, 0x46, 0x77, 0xfe, 0x09, 0xfb, 0x7f, 0x55, 0x31, 0xa3, 0xc2, 0x58, 0xa1, 0xa8, 0x33,
	0x4f, 0x43, 0x2c, 0xea, 0xcc, 0xd3, 0x60, 0x9e, 0xfa, 0xc1, 0x14, 0x4c, 0x53, 0xda, 0x07, 0xe3,
	0x93, 0x25, 0xe9, 0x7c, 0x59, 0x92, 0xce, 0xe9, 0x92, 0xa0, 0x97, 0x2d, 0x41, 0xef, 0x5a, 0x82,
	0x3e, 0xb6, 0x04, 0x9d, 0xb4, 0x04, 0x7d, 0x6d, 0x09, 0xfa, 0xd1, 0x92, 0xce, 0x69, 0x4b, 0xd0,
	0xeb, 0x6f, 0xa4, 0xf3, 0xf3, 0xf3, 0xf7, 0x37, 0x1b, 0x68, 0xba, 0xed, 0xe4, 0xdd, 0xfb, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x55, 0xe8, 0x04, 0xc6, 0x09, 0x03, 0x00, 0x00,
}
