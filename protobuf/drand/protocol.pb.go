// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drand/protocol.proto

package drand

import (
	fmt "fmt"
	dkg "github.com/Giordyfish/drand/protobuf/crypto/dkg"
	proto "github.com/golang/protobuf/proto"
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

type IdentityRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentityRequest) Reset()         { *m = IdentityRequest{} }
func (m *IdentityRequest) String() string { return proto.CompactTextString(m) }
func (*IdentityRequest) ProtoMessage()    {}
func (*IdentityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{0}
}

func (m *IdentityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityRequest.Unmarshal(m, b)
}
func (m *IdentityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityRequest.Marshal(b, m, deterministic)
}
func (m *IdentityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityRequest.Merge(m, src)
}
func (m *IdentityRequest) XXX_Size() int {
	return xxx_messageInfo_IdentityRequest.Size(m)
}
func (m *IdentityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityRequest proto.InternalMessageInfo

// SignalDKGPacket is the packet nodes send to a coordinator that collects all
// keys and setups the group and sends them back to the nodes such that they can
// start the DKG automatically.
type SignalDKGPacket struct {
	Node        *Identity `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	SecretProof []byte    `protobuf:"bytes,2,opt,name=secret_proof,json=secretProof,proto3" json:"secret_proof,omitempty"`
	// In resharing cases, previous_group_hash is the hash of the previous group.
	// It is to make sure the nodes build on top of the correct previous group.
	PreviousGroupHash    []byte   `protobuf:"bytes,3,opt,name=previous_group_hash,json=previousGroupHash,proto3" json:"previous_group_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignalDKGPacket) Reset()         { *m = SignalDKGPacket{} }
func (m *SignalDKGPacket) String() string { return proto.CompactTextString(m) }
func (*SignalDKGPacket) ProtoMessage()    {}
func (*SignalDKGPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{1}
}

func (m *SignalDKGPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalDKGPacket.Unmarshal(m, b)
}
func (m *SignalDKGPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalDKGPacket.Marshal(b, m, deterministic)
}
func (m *SignalDKGPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalDKGPacket.Merge(m, src)
}
func (m *SignalDKGPacket) XXX_Size() int {
	return xxx_messageInfo_SignalDKGPacket.Size(m)
}
func (m *SignalDKGPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalDKGPacket.DiscardUnknown(m)
}

var xxx_messageInfo_SignalDKGPacket proto.InternalMessageInfo

func (m *SignalDKGPacket) GetNode() *Identity {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *SignalDKGPacket) GetSecretProof() []byte {
	if m != nil {
		return m.SecretProof
	}
	return nil
}

func (m *SignalDKGPacket) GetPreviousGroupHash() []byte {
	if m != nil {
		return m.PreviousGroupHash
	}
	return nil
}

// PushDKGInfor is the packet the coordinator sends that contains the group over
// which to run the DKG on, the secret proof (to prove it's he's part of the
// expected group, and it's not a random packet) and as well the time at which
// every node should start the DKG.
type DKGInfoPacket struct {
	NewGroup    *GroupPacket `protobuf:"bytes,1,opt,name=new_group,json=newGroup,proto3" json:"new_group,omitempty"`
	SecretProof []byte       `protobuf:"bytes,2,opt,name=secret_proof,json=secretProof,proto3" json:"secret_proof,omitempty"`
	// timeout in seconds
	DkgTimeout uint32 `protobuf:"varint,3,opt,name=dkg_timeout,json=dkgTimeout,proto3" json:"dkg_timeout,omitempty"`
	// signature from the coordinator to prove he is the one sending that group
	// file.
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DKGInfoPacket) Reset()         { *m = DKGInfoPacket{} }
func (m *DKGInfoPacket) String() string { return proto.CompactTextString(m) }
func (*DKGInfoPacket) ProtoMessage()    {}
func (*DKGInfoPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{2}
}

func (m *DKGInfoPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DKGInfoPacket.Unmarshal(m, b)
}
func (m *DKGInfoPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DKGInfoPacket.Marshal(b, m, deterministic)
}
func (m *DKGInfoPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DKGInfoPacket.Merge(m, src)
}
func (m *DKGInfoPacket) XXX_Size() int {
	return xxx_messageInfo_DKGInfoPacket.Size(m)
}
func (m *DKGInfoPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_DKGInfoPacket.DiscardUnknown(m)
}

var xxx_messageInfo_DKGInfoPacket proto.InternalMessageInfo

func (m *DKGInfoPacket) GetNewGroup() *GroupPacket {
	if m != nil {
		return m.NewGroup
	}
	return nil
}

func (m *DKGInfoPacket) GetSecretProof() []byte {
	if m != nil {
		return m.SecretProof
	}
	return nil
}

func (m *DKGInfoPacket) GetDkgTimeout() uint32 {
	if m != nil {
		return m.DkgTimeout
	}
	return 0
}

func (m *DKGInfoPacket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type PartialBeaconPacket struct {
	// Round is the round for which the beacon will be created from the partial
	// signatures
	Round uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	// signature of the previous round - could be removed at some point but now
	// is used to verify the signature even before accessing the store
	PreviousSig []byte `protobuf:"bytes,2,opt,name=previous_sig,json=previousSig,proto3" json:"previous_sig,omitempty"`
	// partial signature - a threshold of them needs to be aggregated to produce
	// the final beacon at the given round.
	PartialSig           []byte   `protobuf:"bytes,3,opt,name=partial_sig,json=partialSig,proto3" json:"partial_sig,omitempty"`
	Message              []byte   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartialBeaconPacket) Reset()         { *m = PartialBeaconPacket{} }
func (m *PartialBeaconPacket) String() string { return proto.CompactTextString(m) }
func (*PartialBeaconPacket) ProtoMessage()    {}
func (*PartialBeaconPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{3}
}

func (m *PartialBeaconPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartialBeaconPacket.Unmarshal(m, b)
}
func (m *PartialBeaconPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartialBeaconPacket.Marshal(b, m, deterministic)
}
func (m *PartialBeaconPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialBeaconPacket.Merge(m, src)
}
func (m *PartialBeaconPacket) XXX_Size() int {
	return xxx_messageInfo_PartialBeaconPacket.Size(m)
}
func (m *PartialBeaconPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialBeaconPacket.DiscardUnknown(m)
}

var xxx_messageInfo_PartialBeaconPacket proto.InternalMessageInfo

func (m *PartialBeaconPacket) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *PartialBeaconPacket) GetPreviousSig() []byte {
	if m != nil {
		return m.PreviousSig
	}
	return nil
}

func (m *PartialBeaconPacket) GetPartialSig() []byte {
	if m != nil {
		return m.PartialSig
	}
	return nil
}

func (m *PartialBeaconPacket) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

// DKGPacket is the packet that nodes send to others nodes as part of the
// broadcasting protocol.
type DKGPacket struct {
	Dkg                  *dkg.Packet `protobuf:"bytes,1,opt,name=dkg,proto3" json:"dkg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DKGPacket) Reset()         { *m = DKGPacket{} }
func (m *DKGPacket) String() string { return proto.CompactTextString(m) }
func (*DKGPacket) ProtoMessage()    {}
func (*DKGPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{4}
}

func (m *DKGPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DKGPacket.Unmarshal(m, b)
}
func (m *DKGPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DKGPacket.Marshal(b, m, deterministic)
}
func (m *DKGPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DKGPacket.Merge(m, src)
}
func (m *DKGPacket) XXX_Size() int {
	return xxx_messageInfo_DKGPacket.Size(m)
}
func (m *DKGPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_DKGPacket.DiscardUnknown(m)
}

var xxx_messageInfo_DKGPacket proto.InternalMessageInfo

func (m *DKGPacket) GetDkg() *dkg.Packet {
	if m != nil {
		return m.Dkg
	}
	return nil
}

// SyncRequest is from a node that needs to sync up with the current head of the
// chain
type SyncRequest struct {
	FromRound            uint64   `protobuf:"varint,1,opt,name=from_round,json=fromRound,proto3" json:"from_round,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRequest) Reset()         { *m = SyncRequest{} }
func (m *SyncRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()    {}
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{5}
}

func (m *SyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRequest.Unmarshal(m, b)
}
func (m *SyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRequest.Marshal(b, m, deterministic)
}
func (m *SyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRequest.Merge(m, src)
}
func (m *SyncRequest) XXX_Size() int {
	return xxx_messageInfo_SyncRequest.Size(m)
}
func (m *SyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRequest proto.InternalMessageInfo

func (m *SyncRequest) GetFromRound() uint64 {
	if m != nil {
		return m.FromRound
	}
	return 0
}

type BeaconPacket struct {
	PreviousSig          []byte   `protobuf:"bytes,1,opt,name=previous_sig,json=previousSig,proto3" json:"previous_sig,omitempty"`
	Round                uint64   `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Message              []byte   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeaconPacket) Reset()         { *m = BeaconPacket{} }
func (m *BeaconPacket) String() string { return proto.CompactTextString(m) }
func (*BeaconPacket) ProtoMessage()    {}
func (*BeaconPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{6}
}

func (m *BeaconPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeaconPacket.Unmarshal(m, b)
}
func (m *BeaconPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeaconPacket.Marshal(b, m, deterministic)
}
func (m *BeaconPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconPacket.Merge(m, src)
}
func (m *BeaconPacket) XXX_Size() int {
	return xxx_messageInfo_BeaconPacket.Size(m)
}
func (m *BeaconPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconPacket.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconPacket proto.InternalMessageInfo

func (m *BeaconPacket) GetPreviousSig() []byte {
	if m != nil {
		return m.PreviousSig
	}
	return nil
}

func (m *BeaconPacket) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *BeaconPacket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *BeaconPacket) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*IdentityRequest)(nil), "drand.IdentityRequest")
	proto.RegisterType((*SignalDKGPacket)(nil), "drand.SignalDKGPacket")
	proto.RegisterType((*DKGInfoPacket)(nil), "drand.DKGInfoPacket")
	proto.RegisterType((*PartialBeaconPacket)(nil), "drand.PartialBeaconPacket")
	proto.RegisterType((*DKGPacket)(nil), "drand.DKGPacket")
	proto.RegisterType((*SyncRequest)(nil), "drand.SyncRequest")
	proto.RegisterType((*BeaconPacket)(nil), "drand.BeaconPacket")
}

func init() { proto.RegisterFile("drand/protocol.proto", fileDescriptor_e344a98fea1e2f3a) }

var fileDescriptor_e344a98fea1e2f3a = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x95, 0xfb, 0x80, 0xe6, 0x3a, 0x55, 0xe9, 0x34, 0x42, 0x91, 0x45, 0x45, 0x31, 0x9b, 0xaa,
	0xaa, 0x1c, 0x28, 0x52, 0x25, 0x24, 0x56, 0xa5, 0xc8, 0x54, 0xdd, 0x44, 0x0e, 0x2b, 0x36, 0xd1,
	0xc4, 0x9e, 0xd8, 0xa3, 0xc4, 0x33, 0x66, 0x66, 0x4c, 0x95, 0x15, 0x6b, 0x7e, 0x83, 0x8f, 0xe2,
	0x7b, 0xd0, 0x3c, 0x9c, 0x38, 0xa6, 0x48, 0x2c, 0x2c, 0xf9, 0x9e, 0xfb, 0xd0, 0xb9, 0xe7, 0x1e,
	0x1b, 0x06, 0x99, 0xc0, 0x2c, 0x1b, 0x55, 0x82, 0x2b, 0x9e, 0xf2, 0x65, 0x64, 0x5e, 0xd0, 0xbe,
	0x41, 0x83, 0x41, 0x2a, 0x56, 0x95, 0xe2, 0xa3, 0x6c, 0x91, 0xeb, 0xc7, 0x26, 0x03, 0x64, 0x5b,
	0x52, 0x5e, 0x96, 0x9c, 0x59, 0x2c, 0x3c, 0x86, 0xa3, 0xbb, 0x8c, 0x30, 0x45, 0xd5, 0x2a, 0x21,
	0xdf, 0x6a, 0x22, 0x55, 0xf8, 0xd3, 0x83, 0xa3, 0x09, 0xcd, 0x19, 0x5e, 0xde, 0xde, 0xc7, 0x63,
	0x9c, 0x2e, 0x88, 0x42, 0xaf, 0x61, 0x8f, 0xf1, 0x8c, 0x0c, 0xbd, 0x33, 0xef, 0xdc, 0xbf, 0x3a,
	0x8a, 0xcc, 0xa4, 0x68, 0xdd, 0x69, 0x92, 0xe8, 0x15, 0xf4, 0x25, 0x49, 0x05, 0x51, 0xd3, 0x4a,
	0x70, 0x3e, 0x1f, 0xee, 0x9c, 0x79, 0xe7, 0xfd, 0xc4, 0xb7, 0xd8, 0x58, 0x43, 0x28, 0x82, 0x93,
	0x4a, 0x90, 0xef, 0x94, 0xd7, 0x72, 0x9a, 0x0b, 0x5e, 0x57, 0xd3, 0x02, 0xcb, 0x62, 0xb8, 0x6b,
	0x2a, 0x8f, 0x9b, 0x54, 0xac, 0x33, 0x9f, 0xb1, 0x2c, 0xc2, 0x5f, 0x1e, 0x1c, 0xde, 0xde, 0xc7,
	0x77, 0x6c, 0xce, 0x1d, 0x93, 0x11, 0xf4, 0x18, 0x79, 0xb0, 0xcd, 0x8e, 0x0e, 0x72, 0x74, 0x4c,
	0x9b, 0x2d, 0x4b, 0x0e, 0x18, 0x79, 0x30, 0xf1, 0xff, 0xb0, 0x7a, 0x09, 0x7e, 0xb6, 0xc8, 0xa7,
	0x8a, 0x96, 0x84, 0xd7, 0xca, 0xb0, 0x39, 0x4c, 0x20, 0x5b, 0xe4, 0x5f, 0x2c, 0x82, 0x5e, 0x40,
	0x4f, 0x6a, 0x45, 0x54, 0x2d, 0xc8, 0x70, 0xcf, 0x0c, 0xd8, 0x00, 0x5a, 0xb0, 0x93, 0x31, 0x16,
	0x8a, 0xe2, 0xe5, 0x0d, 0xc1, 0x29, 0x67, 0x8e, 0xea, 0x00, 0xf6, 0x05, 0xaf, 0x59, 0x66, 0x68,
	0xee, 0x25, 0x36, 0xd0, 0x7c, 0xd6, 0x12, 0x48, 0x9a, 0x37, 0x7c, 0x1a, 0x6c, 0x42, 0x73, 0xcd,
	0xa7, 0xb2, 0xf3, 0x4c, 0x85, 0x55, 0x07, 0x1c, 0xa4, 0x0b, 0x86, 0xf0, 0xb4, 0x24, 0x52, 0xe2,
	0xbc, 0x61, 0xd3, 0x84, 0xe1, 0x05, 0xf4, 0x36, 0x57, 0x3b, 0x85, 0xdd, 0x6c, 0x91, 0x3b, 0x95,
	0xfc, 0x48, 0x3b, 0xc1, 0xc9, 0xa3, 0xf1, 0xf0, 0x12, 0xfc, 0xc9, 0x8a, 0xa5, 0xee, 0xee, 0xe8,
	0x14, 0x60, 0x2e, 0x78, 0x39, 0x6d, 0x73, 0xee, 0x69, 0x24, 0xd1, 0x40, 0xf8, 0x03, 0xfa, 0x5b,
	0xdb, 0x75, 0xf7, 0xf0, 0xfe, 0xde, 0x63, 0x2d, 0xc0, 0x4e, 0x5b, 0x80, 0x2d, 0x31, 0x77, 0x3b,
	0x62, 0xfe, 0x7b, 0xb5, 0xab, 0xdf, 0x3b, 0x70, 0x30, 0x76, 0x76, 0x47, 0xd7, 0xe0, 0xc7, 0x44,
	0x35, 0x06, 0x44, 0xcf, 0xbb, 0x8e, 0xb4, 0x3b, 0x05, 0x5d, 0xa7, 0xa2, 0x0f, 0x30, 0x68, 0x79,
	0x5b, 0x28, 0x9a, 0xd2, 0x0a, 0x33, 0xb5, 0x1e, 0xd0, 0x31, 0x7e, 0xd0, 0x77, 0xf8, 0xa7, 0xb2,
	0x52, 0x2b, 0xf4, 0x16, 0xfc, 0x71, 0x2d, 0x0b, 0xe7, 0x48, 0x34, 0x70, 0xc9, 0x2d, 0x87, 0x76,
	0x5a, 0x22, 0xe8, 0xdf, 0x08, 0x8e, 0xb3, 0x14, 0x4b, 0x75, 0x7b, 0x1f, 0xa3, 0x67, 0x9b, 0x9e,
	0x47, 0xeb, 0xdf, 0xc3, 0xe1, 0x96, 0x97, 0x50, 0xe0, 0xd2, 0x8f, 0x38, 0xac, 0xd3, 0x7a, 0x0d,
	0x3d, 0x7d, 0xcf, 0x8f, 0x05, 0xa6, 0x0c, 0x35, 0x1f, 0x45, 0xeb, 0xc2, 0xc1, 0x89, 0xc3, 0xda,
	0x33, 0xde, 0x78, 0x37, 0x97, 0x5f, 0x2f, 0x72, 0xaa, 0x8a, 0x7a, 0x16, 0xa5, 0xbc, 0x1c, 0xc5,
	0x94, 0x8b, 0x6c, 0x35, 0xa7, 0xb2, 0x18, 0xb5, 0x7e, 0x31, 0xb3, 0x7a, 0x6e, 0xc3, 0xd9, 0x13,
	0x13, 0xbf, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xf7, 0x66, 0x36, 0x81, 0x04, 0x00, 0x00,
}
