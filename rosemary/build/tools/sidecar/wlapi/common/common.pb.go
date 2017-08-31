// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common.proto

It has these top-level messages:
	Empty
	AttestedData
	Selector
	Selectors
	RegistrationEntry
	RegistrationEntries
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// @exclude Represents a message with no fields
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// *A type which contains attestation data for specific platform.
type AttestedData struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *AttestedData) Reset()                    { *m = AttestedData{} }
func (m *AttestedData) String() string            { return proto.CompactTextString(m) }
func (*AttestedData) ProtoMessage()               {}
func (*AttestedData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AttestedData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AttestedData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// *A type which describes the conditions under which a registration entry is matched.
type Selector struct {
	Type  string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Selector) Reset()                    { *m = Selector{} }
func (m *Selector) String() string            { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()               {}
func (*Selector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Selector) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Selector) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// *Represents a type with a list of NodeResolution.
type Selectors struct {
	Entries []*Selector `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *Selectors) Reset()                    { *m = Selectors{} }
func (m *Selectors) String() string            { return proto.CompactTextString(m) }
func (*Selectors) ProtoMessage()               {}
func (*Selectors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Selectors) GetEntries() []*Selector {
	if m != nil {
		return m.Entries
	}
	return nil
}

// *This is a curated record that the Control Plane uses to set up and manage the various registered nodes and workloads that are controlled by it.
type RegistrationEntry struct {
	Selectors   []*Selector `protobuf:"bytes,1,rep,name=selectors" json:"selectors,omitempty"`
	ParentId    string      `protobuf:"bytes,2,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	SpiffeId    string      `protobuf:"bytes,3,opt,name=spiffe_id,json=spiffeId" json:"spiffe_id,omitempty"`
	Ttl         int32       `protobuf:"varint,4,opt,name=ttl" json:"ttl,omitempty"`
	FbSpiffeIds []string    `protobuf:"bytes,5,rep,name=fb_spiffe_ids,json=fbSpiffeIds" json:"fb_spiffe_ids,omitempty"`
}

func (m *RegistrationEntry) Reset()                    { *m = RegistrationEntry{} }
func (m *RegistrationEntry) String() string            { return proto.CompactTextString(m) }
func (*RegistrationEntry) ProtoMessage()               {}
func (*RegistrationEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegistrationEntry) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *RegistrationEntry) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *RegistrationEntry) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *RegistrationEntry) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *RegistrationEntry) GetFbSpiffeIds() []string {
	if m != nil {
		return m.FbSpiffeIds
	}
	return nil
}

// *A list of registration entries.
type RegistrationEntries struct {
	Entries []*RegistrationEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *RegistrationEntries) Reset()                    { *m = RegistrationEntries{} }
func (m *RegistrationEntries) String() string            { return proto.CompactTextString(m) }
func (*RegistrationEntries) ProtoMessage()               {}
func (*RegistrationEntries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RegistrationEntries) GetEntries() []*RegistrationEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "common.Empty")
	proto.RegisterType((*AttestedData)(nil), "common.AttestedData")
	proto.RegisterType((*Selector)(nil), "common.Selector")
	proto.RegisterType((*Selectors)(nil), "common.Selectors")
	proto.RegisterType((*RegistrationEntry)(nil), "common.RegistrationEntry")
	proto.RegisterType((*RegistrationEntries)(nil), "common.RegistrationEntries")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x96, 0x49, 0xd3, 0xd6, 0xd7, 0x22, 0x15, 0xc3, 0x10, 0xc4, 0x12, 0x79, 0x8a, 0x18, 0x3a,
	0x50, 0x04, 0x33, 0x12, 0x1d, 0xca, 0xe8, 0x3e, 0x40, 0xe5, 0x34, 0x17, 0x64, 0x29, 0x89, 0x23,
	0xfb, 0x40, 0xca, 0x73, 0xf1, 0x82, 0x28, 0x7f, 0x45, 0x2a, 0xb0, 0x7d, 0xf7, 0xfd, 0x9c, 0xbe,
	0xb3, 0x61, 0x79, 0xb4, 0x65, 0x69, 0xab, 0x75, 0xed, 0x2c, 0x59, 0x31, 0xed, 0x27, 0x39, 0x83,
	0x70, 0x5b, 0xd6, 0xd4, 0xc8, 0x27, 0x58, 0xbe, 0x10, 0xa1, 0x27, 0xcc, 0x5e, 0x35, 0x69, 0x21,
	0x60, 0x42, 0x4d, 0x8d, 0x11, 0x8b, 0x59, 0xc2, 0x55, 0x87, 0x5b, 0x2e, 0xd3, 0xa4, 0xa3, 0x8b,
	0x98, 0x25, 0x4b, 0xd5, 0x61, 0xf9, 0x08, 0xf3, 0x3d, 0x16, 0x78, 0x24, 0xeb, 0xfe, 0xcc, 0xdc,
	0x40, 0xf8, 0xa9, 0x8b, 0x0f, 0xec, 0x42, 0x5c, 0xf5, 0x83, 0x7c, 0x06, 0x3e, 0xa6, 0xbc, 0xb8,
	0x87, 0x19, 0x56, 0xe4, 0x0c, 0xfa, 0x88, 0xc5, 0x41, 0xb2, 0x78, 0x58, 0xad, 0x87, 0xae, 0xa3,
	0x47, 0x8d, 0x06, 0xf9, 0xc5, 0xe0, 0x4a, 0xe1, 0xbb, 0xf1, 0xe4, 0x34, 0x19, 0x5b, 0x6d, 0x2b,
	0x72, 0x8d, 0x58, 0x03, 0xf7, 0xe3, 0xba, 0x7f, 0x77, 0xfc, 0x58, 0xc4, 0x1d, 0xf0, 0x5a, 0x3b,
	0xac, 0xe8, 0x60, 0xb2, 0xa1, 0xd8, 0xbc, 0x27, 0x76, 0x59, 0x2b, 0xfa, 0xda, 0xe4, 0x39, 0xb6,
	0x62, 0xd0, 0x8b, 0x3d, 0xb1, 0xcb, 0xc4, 0x0a, 0x02, 0xa2, 0x22, 0x9a, 0xc4, 0x2c, 0x09, 0x55,
	0x0b, 0x85, 0x84, 0xcb, 0x3c, 0x3d, 0x9c, 0x12, 0x3e, 0x0a, 0xe3, 0x20, 0xe1, 0x6a, 0x91, 0xa7,
	0xfb, 0x21, 0xe4, 0xe5, 0x1b, 0x5c, 0x9f, 0x97, 0x36, 0xe8, 0xc5, 0xe6, 0xfc, 0xf0, 0xdb, 0xb1,
	0xf4, 0xaf, 0x13, 0x4f, 0x2f, 0x90, 0x4e, 0xbb, 0x0f, 0xdc, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x25, 0x30, 0x9e, 0x52, 0xd0, 0x01, 0x00, 0x00,
}
