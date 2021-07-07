// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bloxapp/ssv/ibft/proto/beacon.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type InputValue struct {
	// Types that are valid to be assigned to Data:
	//	*InputValue_AttestationData
	//	*InputValue_AggregationData
	//	*InputValue_BeaconBlock
	Data isInputValue_Data `protobuf_oneof:"data"`
	// Types that are valid to be assigned to SignedData:
	//	*InputValue_Attestation
	//	*InputValue_Aggregation
	//	*InputValue_Block
	SignedData           isInputValue_SignedData `protobuf_oneof:"signed_data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *InputValue) Reset()         { *m = InputValue{} }
func (m *InputValue) String() string { return proto.CompactTextString(m) }
func (*InputValue) ProtoMessage()    {}
func (*InputValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_933ff678c25e0138, []int{0}
}

func (m *InputValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputValue.Unmarshal(m, b)
}
func (m *InputValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputValue.Marshal(b, m, deterministic)
}
func (m *InputValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputValue.Merge(m, src)
}
func (m *InputValue) XXX_Size() int {
	return xxx_messageInfo_InputValue.Size(m)
}
func (m *InputValue) XXX_DiscardUnknown() {
	xxx_messageInfo_InputValue.DiscardUnknown(m)
}

var xxx_messageInfo_InputValue proto.InternalMessageInfo

type isInputValue_Data interface {
	isInputValue_Data()
}

type InputValue_AttestationData struct {
	AttestationData *v1alpha1.AttestationData `protobuf:"bytes,2,opt,name=attestation_data,json=attestationData,proto3,oneof"`
}

type InputValue_AggregationData struct {
	AggregationData *v1alpha1.AggregateAttestationAndProof `protobuf:"bytes,3,opt,name=aggregation_data,json=aggregationData,proto3,oneof"`
}

type InputValue_BeaconBlock struct {
	BeaconBlock *v1alpha1.BeaconBlock `protobuf:"bytes,4,opt,name=beacon_block,json=beaconBlock,proto3,oneof"`
}

func (*InputValue_AttestationData) isInputValue_Data() {}

func (*InputValue_AggregationData) isInputValue_Data() {}

func (*InputValue_BeaconBlock) isInputValue_Data() {}

func (m *InputValue) GetData() isInputValue_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InputValue) GetAttestationData() *v1alpha1.AttestationData {
	if x, ok := m.GetData().(*InputValue_AttestationData); ok {
		return x.AttestationData
	}
	return nil
}

func (m *InputValue) GetAggregationData() *v1alpha1.AggregateAttestationAndProof {
	if x, ok := m.GetData().(*InputValue_AggregationData); ok {
		return x.AggregationData
	}
	return nil
}

func (m *InputValue) GetBeaconBlock() *v1alpha1.BeaconBlock {
	if x, ok := m.GetData().(*InputValue_BeaconBlock); ok {
		return x.BeaconBlock
	}
	return nil
}

type isInputValue_SignedData interface {
	isInputValue_SignedData()
}

type InputValue_Attestation struct {
	Attestation *v1alpha1.Attestation `protobuf:"bytes,5,opt,name=attestation,proto3,oneof"`
}

type InputValue_Aggregation struct {
	Aggregation *v1alpha1.SignedAggregateAttestationAndProof `protobuf:"bytes,6,opt,name=aggregation,proto3,oneof"`
}

type InputValue_Block struct {
	Block *v1alpha1.SignedBeaconBlock `protobuf:"bytes,7,opt,name=block,proto3,oneof"`
}

func (*InputValue_Attestation) isInputValue_SignedData() {}

func (*InputValue_Aggregation) isInputValue_SignedData() {}

func (*InputValue_Block) isInputValue_SignedData() {}

func (m *InputValue) GetSignedData() isInputValue_SignedData {
	if m != nil {
		return m.SignedData
	}
	return nil
}

func (m *InputValue) GetAttestation() *v1alpha1.Attestation {
	if x, ok := m.GetSignedData().(*InputValue_Attestation); ok {
		return x.Attestation
	}
	return nil
}

func (m *InputValue) GetAggregation() *v1alpha1.SignedAggregateAttestationAndProof {
	if x, ok := m.GetSignedData().(*InputValue_Aggregation); ok {
		return x.Aggregation
	}
	return nil
}

func (m *InputValue) GetBlock() *v1alpha1.SignedBeaconBlock {
	if x, ok := m.GetSignedData().(*InputValue_Block); ok {
		return x.Block
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*InputValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*InputValue_AttestationData)(nil),
		(*InputValue_AggregationData)(nil),
		(*InputValue_BeaconBlock)(nil),
		(*InputValue_Attestation)(nil),
		(*InputValue_Aggregation)(nil),
		(*InputValue_Block)(nil),
	}
}

func init() {
	proto.RegisterType((*InputValue)(nil), "proto.InputValue")
}

func init() {
	proto.RegisterFile("github.com/bloxapp/ssv/ibft/proto/beacon.proto", fileDescriptor_933ff678c25e0138)
}

var fileDescriptor_933ff678c25e0138 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd2, 0xcf, 0x4f, 0xc2, 0x30,
	0x14, 0x07, 0x70, 0x50, 0xc0, 0xa4, 0xd3, 0x48, 0x7a, 0x5a, 0x38, 0xf8, 0x03, 0x13, 0xc3, 0xa9,
	0x0d, 0x72, 0xf2, 0x26, 0x8b, 0x51, 0xbc, 0x99, 0x91, 0x78, 0x30, 0x31, 0xf8, 0x0a, 0x65, 0x5b,
	0x1c, 0xeb, 0xb2, 0xbd, 0x11, 0xff, 0x08, 0xff, 0x68, 0xb3, 0x37, 0xc9, 0xca, 0x81, 0x79, 0xea,
	0xba, 0x7d, 0xfb, 0xe9, 0x7b, 0x79, 0x63, 0x22, 0x88, 0x30, 0x2c, 0x94, 0x58, 0x9a, 0x8d, 0x54,
	0xb1, 0xf9, 0x86, 0x34, 0x95, 0x79, 0xbe, 0x95, 0x91, 0x5a, 0xa3, 0x4c, 0x33, 0x83, 0x46, 0x2a,
	0x0d, 0x4b, 0x93, 0x08, 0xda, 0xf0, 0x2e, 0x2d, 0x83, 0x0b, 0x8d, 0xa1, 0xdc, 0x8e, 0x21, 0x4e,
	0x43, 0x18, 0x4b, 0x40, 0xd4, 0x39, 0x02, 0x46, 0xbb, 0xd8, 0xe0, 0x72, 0xef, 0x7b, 0x25, 0x2c,
	0x54, 0x6c, 0x96, 0x5f, 0x55, 0x60, 0xf8, 0xd3, 0x61, 0xec, 0x25, 0x49, 0x0b, 0x7c, 0x83, 0xb8,
	0xd0, 0x7c, 0xce, 0xfa, 0x16, 0xb2, 0x58, 0x01, 0x82, 0x7b, 0x74, 0xd5, 0x1e, 0x39, 0x77, 0xb7,
	0x42, 0x63, 0xa8, 0x33, 0x5d, 0x6c, 0xca, 0x07, 0xb1, 0x33, 0xc5, 0xb4, 0x8e, 0x3f, 0x02, 0xc2,
	0xac, 0xe5, 0x9f, 0xc3, 0xfe, 0x2b, 0xfe, 0xc9, 0xfa, 0x10, 0x04, 0x99, 0x0e, 0x2c, 0xf4, 0x98,
	0xd0, 0xc9, 0x21, 0xf4, 0x2f, 0xae, 0x2d, 0x7d, 0x9a, 0xac, 0x5e, 0x33, 0x63, 0xd6, 0x74, 0x43,
	0xcd, 0xd1, 0x0d, 0xcf, 0xec, 0xd4, 0xee, 0xcd, 0xed, 0x90, 0x3e, 0x3c, 0xa0, 0x7b, 0x14, 0xf5,
	0xca, 0xe4, 0xac, 0xe5, 0x3b, 0xaa, 0xde, 0xf2, 0x27, 0xe6, 0x58, 0xd5, 0xbb, 0xdd, 0x46, 0xc7,
	0x2a, 0x6e, 0xd6, 0xf6, 0xed, 0x83, 0xfc, 0x83, 0x39, 0x56, 0x8d, 0x6e, 0x8f, 0x9c, 0xfb, 0x03,
	0xce, 0x3c, 0x0a, 0x12, 0xbd, 0x6a, 0xec, 0xb9, 0xe4, 0x6b, 0x8f, 0x3f, 0xb0, 0x6e, 0xd5, 0xe8,
	0x09, 0xc1, 0xa3, 0x46, 0xd8, 0x6e, 0xb7, 0xed, 0x57, 0x07, 0xbd, 0x1e, 0xeb, 0x94, 0x73, 0xf0,
	0xce, 0x98, 0x93, 0x53, 0x8a, 0xc6, 0xe2, 0xdd, 0xbc, 0x5f, 0xff, 0xfb, 0x23, 0xaa, 0x1e, 0x2d,
	0x93, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xc6, 0x61, 0x6c, 0xb4, 0x02, 0x00, 0x00,
}
