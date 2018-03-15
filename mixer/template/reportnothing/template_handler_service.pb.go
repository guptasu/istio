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
	HandleReportNothingResponse
	InstanceMsg
	Type
	InstanceParam
*/
package reportnothing

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "istio.io/api/mixer/adapter/model/v1beta1"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"
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

// Response message for HandleReportNothing method.
type HandleReportNothingResponse struct {
	// The success/failure status of HandleReportNothing call.
	Status *google_rpc.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *HandleReportNothingResponse) Reset()      { *m = HandleReportNothingResponse{} }
func (*HandleReportNothingResponse) ProtoMessage() {}
func (*HandleReportNothingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{1}
}

func (m *HandleReportNothingResponse) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
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
	return fileDescriptorTemplateHandlerService, []int{2}
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
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorTemplateHandlerService, []int{3} }

// Represents instance configuration schema for 'reportnothing' template.
type InstanceParam struct {
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{4}
}

func init() {
	proto.RegisterType((*HandleReportNothingRequest)(nil), "reportnothing.HandleReportNothingRequest")
	proto.RegisterType((*HandleReportNothingResponse)(nil), "reportnothing.HandleReportNothingResponse")
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
func (this *HandleReportNothingResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*HandleReportNothingResponse)
	if !ok {
		that2, ok := that.(HandleReportNothingResponse)
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
	if !this.Status.Equal(that1.Status) {
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
func (this *HandleReportNothingResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&reportnothing.HandleReportNothingResponse{")
	if this.Status != nil {
		s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	}
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

func (m *HandleReportNothingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandleReportNothingResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.Status.Size()))
		n2, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
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

func (m *HandleReportNothingResponse) Size() (n int) {
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
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
func (this *HandleReportNothingResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HandleReportNothingResponse{`,
		`Status:` + strings.Replace(fmt.Sprintf("%v", this.Status), "Status", "google_rpc.Status", 1) + `,`,
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
func (m *HandleReportNothingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: HandleReportNothingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandleReportNothingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
			if m.Status == nil {
				m.Status = &google_rpc.Status{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x3f, 0x6f, 0xd4, 0x30,
	0x14, 0x3f, 0xb7, 0xd5, 0x41, 0x7d, 0x3a, 0x90, 0x5c, 0x10, 0xe9, 0x21, 0x59, 0xa7, 0x08, 0xa1,
	0xa3, 0x42, 0xb1, 0x2e, 0x8c, 0x9d, 0x80, 0x85, 0x1b, 0x40, 0x28, 0x65, 0x8f, 0x7c, 0xc9, 0x6b,
	0x1a, 0x29, 0xb1, 0x8d, 0xed, 0x54, 0xcd, 0xc6, 0xcc, 0x84, 0xc4, 0x87, 0x80, 0x8d, 0x2f, 0xc0,
	0x07, 0x40, 0x4c, 0x15, 0x13, 0x23, 0x17, 0x18, 0x18, 0x3b, 0x32, 0x22, 0x9c, 0x04, 0x54, 0x74,
	0xa0, 0x8e, 0xcf, 0xbf, 0xdf, 0xf3, 0xef, 0xcf, 0xc3, 0xfb, 0x65, 0x7e, 0x02, 0x9a, 0x59, 0x28,
	0x55, 0xc1, 0x2d, 0x30, 0x0d, 0x4a, 0x6a, 0x2b, 0xa4, 0x3d, 0xca, 0x45, 0xf6, 0xfb, 0x39, 0x3e,
	0xe2, 0x22, 0x2d, 0x40, 0xc7, 0x06, 0xf4, 0x71, 0x9e, 0x40, 0xa0, 0xb4, 0xb4, 0x92, 0x8c, 0xcf,
	0xb1, 0x27, 0x77, 0xdb, 0xbf, 0x78, 0xca, 0x95, 0x05, 0xcd, 0x4a, 0x99, 0x42, 0xc1, 0x8e, 0xe7,
	0x4b, 0xb0, 0x7c, 0xce, 0xe0, 0xc4, 0x82, 0x30, 0xb9, 0x14, 0xa6, 0x5d, 0x9e, 0xec, 0x66, 0x52,
	0x66, 0x05, 0x30, 0x37, 0x2d, 0xab, 0x43, 0xc6, 0x45, 0xdd, 0x41, 0x37, 0x3a, 0x48, 0xab, 0x84,
	0x19, 0xcb, 0x6d, 0xd5, 0xef, 0xdc, 0xfe, 0x9f, 0x82, 0xad, 0x55, 0x67, 0xcc, 0x7f, 0x83, 0xf0,
	0xe4, 0x91, 0xb3, 0x1c, 0x39, 0x87, 0x4f, 0x5a, 0x87, 0x11, 0x3c, 0xaf, 0xc0, 0x58, 0x32, 0xc7,
	0xdb, 0xb9, 0x30, 0x96, 0x8b, 0x04, 0x8c, 0x87, 0xa6, 0x9b, 0xb3, 0x51, 0xb8, 0x13, 0x9c, 0xcb,
	0x12, 0x3c, 0xab, 0x15, 0x44, 0x7f, 0x58, 0x64, 0x1f, 0x5f, 0xe9, 0x54, 0xe3, 0x44, 0x8a, 0xc3,
	0x3c, 0xf3, 0x36, 0xa6, 0x68, 0x36, 0x0a, 0xaf, 0x05, 0xad, 0xd7, 0xa0, 0x8f, 0x11, 0xdc, 0x17,
	0x75, 0x34, 0xee, 0xb8, 0x0f, 0x1d, 0x95, 0xec, 0xe2, 0xcb, 0x29, 0xa4, 0x95, 0x8a, 0xf3, 0xd4,
	0xdb, 0x9c, 0xa2, 0xd9, 0x76, 0x74, 0xc9, 0xcd, 0x8b, 0xd4, 0x5f, 0xe0, 0x9b, 0x6b, 0x8d, 0x1a,
	0x25, 0x85, 0x01, 0xb2, 0x87, 0x87, 0x6d, 0x01, 0x1e, 0x72, 0x72, 0xa4, 0x97, 0xd3, 0x2a, 0x09,
	0x0e, 0x1c, 0x12, 0x75, 0x0c, 0xff, 0x16, 0x1e, 0x2d, 0x3a, 0xbf, 0x8f, 0x4d, 0x46, 0xae, 0xe3,
	0x2d, 0xc1, 0x4b, 0xf0, 0xde, 0x7d, 0x7c, 0xef, 0x3b, 0x4d, 0x37, 0xfa, 0x43, 0xbc, 0xf5, 0x2b,
	0x9b, 0x7f, 0x15, 0x8f, 0x7b, 0xf6, 0x53, 0xae, 0x79, 0x19, 0xbe, 0x5c, 0xdf, 0xd9, 0x41, 0x7b,
	0x71, 0x52, 0xe0, 0x9d, 0x35, 0x28, 0xb9, 0xf3, 0x57, 0x6f, 0xff, 0x6e, 0x7d, 0xb2, 0x77, 0x11,
	0x6a, 0x9b, 0xfb, 0x41, 0x78, 0xba, 0xa2, 0x83, 0xcf, 0x2b, 0x3a, 0x38, 0x5b, 0x51, 0xf4, 0xa2,
	0xa1, 0xe8, 0x6d, 0x43, 0xd1, 0x87, 0x86, 0xa2, 0xd3, 0x86, 0xa2, 0x2f, 0x0d, 0x45, 0xdf, 0x1b,
	0x3a, 0x38, 0x6b, 0x28, 0x7a, 0xf5, 0x95, 0x0e, 0x7e, 0x7c, 0xfa, 0xf6, 0x7a, 0x03, 0x2d, 0x87,
	0xee, 0x04, 0xf7, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x72, 0x52, 0xd5, 0x93, 0xd3, 0x02, 0x00,
	0x00,
}
