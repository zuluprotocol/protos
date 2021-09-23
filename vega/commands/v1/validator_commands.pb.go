// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vega/commands/v1/validator_commands.proto

package v1

import (
	vega "code.vegaprotocol.io/protos/vega"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

// The kind of the signature created by a node, for example, allow-listing a new asset, withdrawal etc
type NodeSignatureKind int32

const (
	// Represents an unspecified or missing value from the input
	NodeSignatureKind_NODE_SIGNATURE_KIND_UNSPECIFIED NodeSignatureKind = 0
	// Represents a signature for a new asset allow-listing
	NodeSignatureKind_NODE_SIGNATURE_KIND_ASSET_NEW NodeSignatureKind = 1
	// Represents a signature for an asset withdrawal
	NodeSignatureKind_NODE_SIGNATURE_KIND_ASSET_WITHDRAWAL NodeSignatureKind = 2
)

var NodeSignatureKind_name = map[int32]string{
	0: "NODE_SIGNATURE_KIND_UNSPECIFIED",
	1: "NODE_SIGNATURE_KIND_ASSET_NEW",
	2: "NODE_SIGNATURE_KIND_ASSET_WITHDRAWAL",
}

var NodeSignatureKind_value = map[string]int32{
	"NODE_SIGNATURE_KIND_UNSPECIFIED":      0,
	"NODE_SIGNATURE_KIND_ASSET_NEW":        1,
	"NODE_SIGNATURE_KIND_ASSET_WITHDRAWAL": 2,
}

func (x NodeSignatureKind) String() string {
	return proto.EnumName(NodeSignatureKind_name, int32(x))
}

func (NodeSignatureKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52dee7aaa07d334d, []int{0}
}

// Used to Register a node as a validator during network start-up
type NodeRegistration struct {
	// Vega public key, required field
	VegaPubKey string `protobuf:"bytes,1,opt,name=vega_pub_key,json=vegaPubKey,proto3" json:"vega_pub_key,omitempty"`
	// Ethereum public key, required field
	EthereumAddress string `protobuf:"bytes,2,opt,name=ethereum_address,json=ethereumAddress,proto3" json:"ethereum_address,omitempty"`
	// Public key for the blockchain, required field
	ChainPubKey string `protobuf:"bytes,3,opt,name=chain_pub_key,json=chainPubKey,proto3" json:"chain_pub_key,omitempty"`
	// URL with more info on the node
	InfoUrl string `protobuf:"bytes,4,opt,name=info_url,json=infoUrl,proto3" json:"info_url,omitempty"`
	// Country code (ISO 3166-1 alpha-2) for the location of the node
	Country string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	// ID of the validator, (public master key)
	Id string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the validator
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	// AvatarURL of the validator
	AvatarUrl            string   `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRegistration) Reset()         { *m = NodeRegistration{} }
func (m *NodeRegistration) String() string { return proto.CompactTextString(m) }
func (*NodeRegistration) ProtoMessage()    {}
func (*NodeRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_52dee7aaa07d334d, []int{0}
}

func (m *NodeRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRegistration.Unmarshal(m, b)
}
func (m *NodeRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRegistration.Marshal(b, m, deterministic)
}
func (m *NodeRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRegistration.Merge(m, src)
}
func (m *NodeRegistration) XXX_Size() int {
	return xxx_messageInfo_NodeRegistration.Size(m)
}
func (m *NodeRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRegistration proto.InternalMessageInfo

func (m *NodeRegistration) GetVegaPubKey() string {
	if m != nil {
		return m.VegaPubKey
	}
	return ""
}

func (m *NodeRegistration) GetEthereumAddress() string {
	if m != nil {
		return m.EthereumAddress
	}
	return ""
}

func (m *NodeRegistration) GetChainPubKey() string {
	if m != nil {
		return m.ChainPubKey
	}
	return ""
}

func (m *NodeRegistration) GetInfoUrl() string {
	if m != nil {
		return m.InfoUrl
	}
	return ""
}

func (m *NodeRegistration) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *NodeRegistration) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeRegistration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeRegistration) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

// Used when a node votes for validating a given resource exists or is valid,
// for example, an ERC20 deposit is valid and exists on ethereum
type NodeVote struct {
	// Public key, required field
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	// Reference, required field
	Reference            string   `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeVote) Reset()         { *m = NodeVote{} }
func (m *NodeVote) String() string { return proto.CompactTextString(m) }
func (*NodeVote) ProtoMessage()    {}
func (*NodeVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_52dee7aaa07d334d, []int{1}
}

func (m *NodeVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeVote.Unmarshal(m, b)
}
func (m *NodeVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeVote.Marshal(b, m, deterministic)
}
func (m *NodeVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeVote.Merge(m, src)
}
func (m *NodeVote) XXX_Size() int {
	return xxx_messageInfo_NodeVote.Size(m)
}
func (m *NodeVote) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeVote.DiscardUnknown(m)
}

var xxx_messageInfo_NodeVote proto.InternalMessageInfo

func (m *NodeVote) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *NodeVote) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

// Represents a signature from a validator, to be used by a foreign chain in order to recognise a decision taken by the Vega network
type NodeSignature struct {
	// The identifier of the resource being signed
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The signature
	Sig []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	// The kind of resource being signed
	Kind                 NodeSignatureKind `protobuf:"varint,3,opt,name=kind,proto3,enum=vega.commands.v1.NodeSignatureKind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NodeSignature) Reset()         { *m = NodeSignature{} }
func (m *NodeSignature) String() string { return proto.CompactTextString(m) }
func (*NodeSignature) ProtoMessage()    {}
func (*NodeSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_52dee7aaa07d334d, []int{2}
}

func (m *NodeSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeSignature.Unmarshal(m, b)
}
func (m *NodeSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeSignature.Marshal(b, m, deterministic)
}
func (m *NodeSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeSignature.Merge(m, src)
}
func (m *NodeSignature) XXX_Size() int {
	return xxx_messageInfo_NodeSignature.Size(m)
}
func (m *NodeSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeSignature.DiscardUnknown(m)
}

var xxx_messageInfo_NodeSignature proto.InternalMessageInfo

func (m *NodeSignature) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeSignature) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *NodeSignature) GetKind() NodeSignatureKind {
	if m != nil {
		return m.Kind
	}
	return NodeSignatureKind_NODE_SIGNATURE_KIND_UNSPECIFIED
}

// An event forwarded to the Vega network to provide information on events happening on other networks
type ChainEvent struct {
	// The identifier of the transaction in which the events happened, usually a hash
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// Arbitrary one-time integer used to prevent replay attacks
	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The event
	//
	// Types that are valid to be assigned to Event:
	//	*ChainEvent_Builtin
	//	*ChainEvent_Erc20
	//	*ChainEvent_Btc
	//	*ChainEvent_Validator
	//	*ChainEvent_StakingEvent
	Event                isChainEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ChainEvent) Reset()         { *m = ChainEvent{} }
func (m *ChainEvent) String() string { return proto.CompactTextString(m) }
func (*ChainEvent) ProtoMessage()    {}
func (*ChainEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_52dee7aaa07d334d, []int{3}
}

func (m *ChainEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainEvent.Unmarshal(m, b)
}
func (m *ChainEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainEvent.Marshal(b, m, deterministic)
}
func (m *ChainEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainEvent.Merge(m, src)
}
func (m *ChainEvent) XXX_Size() int {
	return xxx_messageInfo_ChainEvent.Size(m)
}
func (m *ChainEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ChainEvent proto.InternalMessageInfo

func (m *ChainEvent) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *ChainEvent) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type isChainEvent_Event interface {
	isChainEvent_Event()
}

type ChainEvent_Builtin struct {
	Builtin *vega.BuiltinAssetEvent `protobuf:"bytes,1001,opt,name=builtin,proto3,oneof"`
}

type ChainEvent_Erc20 struct {
	Erc20 *vega.ERC20Event `protobuf:"bytes,1002,opt,name=erc20,proto3,oneof"`
}

type ChainEvent_Btc struct {
	Btc *vega.BTCEvent `protobuf:"bytes,1003,opt,name=btc,proto3,oneof"`
}

type ChainEvent_Validator struct {
	Validator *vega.ValidatorEvent `protobuf:"bytes,1004,opt,name=validator,proto3,oneof"`
}

type ChainEvent_StakingEvent struct {
	StakingEvent *vega.StakingEvent `protobuf:"bytes,1005,opt,name=staking_event,json=stakingEvent,proto3,oneof"`
}

func (*ChainEvent_Builtin) isChainEvent_Event() {}

func (*ChainEvent_Erc20) isChainEvent_Event() {}

func (*ChainEvent_Btc) isChainEvent_Event() {}

func (*ChainEvent_Validator) isChainEvent_Event() {}

func (*ChainEvent_StakingEvent) isChainEvent_Event() {}

func (m *ChainEvent) GetEvent() isChainEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ChainEvent) GetBuiltin() *vega.BuiltinAssetEvent {
	if x, ok := m.GetEvent().(*ChainEvent_Builtin); ok {
		return x.Builtin
	}
	return nil
}

func (m *ChainEvent) GetErc20() *vega.ERC20Event {
	if x, ok := m.GetEvent().(*ChainEvent_Erc20); ok {
		return x.Erc20
	}
	return nil
}

func (m *ChainEvent) GetBtc() *vega.BTCEvent {
	if x, ok := m.GetEvent().(*ChainEvent_Btc); ok {
		return x.Btc
	}
	return nil
}

func (m *ChainEvent) GetValidator() *vega.ValidatorEvent {
	if x, ok := m.GetEvent().(*ChainEvent_Validator); ok {
		return x.Validator
	}
	return nil
}

func (m *ChainEvent) GetStakingEvent() *vega.StakingEvent {
	if x, ok := m.GetEvent().(*ChainEvent_StakingEvent); ok {
		return x.StakingEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ChainEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ChainEvent_Builtin)(nil),
		(*ChainEvent_Erc20)(nil),
		(*ChainEvent_Btc)(nil),
		(*ChainEvent_Validator)(nil),
		(*ChainEvent_StakingEvent)(nil),
	}
}

// A transaction to allow validator to rotate their vega keys
type KeyRotateSubmission struct {
	KeyNumber            uint64   `protobuf:"varint,1,opt,name=key_number,json=keyNumber,proto3" json:"key_number,omitempty"`
	TargetBlock          uint64   `protobuf:"varint,2,opt,name=target_block,json=targetBlock,proto3" json:"target_block,omitempty"`
	Time                 int64    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	NewPubKeyHash        string   `protobuf:"bytes,4,opt,name=new_pub_key_hash,json=newPubKeyHash,proto3" json:"new_pub_key_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyRotateSubmission) Reset()         { *m = KeyRotateSubmission{} }
func (m *KeyRotateSubmission) String() string { return proto.CompactTextString(m) }
func (*KeyRotateSubmission) ProtoMessage()    {}
func (*KeyRotateSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_52dee7aaa07d334d, []int{4}
}

func (m *KeyRotateSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRotateSubmission.Unmarshal(m, b)
}
func (m *KeyRotateSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRotateSubmission.Marshal(b, m, deterministic)
}
func (m *KeyRotateSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRotateSubmission.Merge(m, src)
}
func (m *KeyRotateSubmission) XXX_Size() int {
	return xxx_messageInfo_KeyRotateSubmission.Size(m)
}
func (m *KeyRotateSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRotateSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRotateSubmission proto.InternalMessageInfo

func (m *KeyRotateSubmission) GetKeyNumber() uint64 {
	if m != nil {
		return m.KeyNumber
	}
	return 0
}

func (m *KeyRotateSubmission) GetTargetBlock() uint64 {
	if m != nil {
		return m.TargetBlock
	}
	return 0
}

func (m *KeyRotateSubmission) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *KeyRotateSubmission) GetNewPubKeyHash() string {
	if m != nil {
		return m.NewPubKeyHash
	}
	return ""
}

func init() {
	proto.RegisterEnum("vega.commands.v1.NodeSignatureKind", NodeSignatureKind_name, NodeSignatureKind_value)
	proto.RegisterType((*NodeRegistration)(nil), "vega.commands.v1.NodeRegistration")
	proto.RegisterType((*NodeVote)(nil), "vega.commands.v1.NodeVote")
	proto.RegisterType((*NodeSignature)(nil), "vega.commands.v1.NodeSignature")
	proto.RegisterType((*ChainEvent)(nil), "vega.commands.v1.ChainEvent")
	proto.RegisterType((*KeyRotateSubmission)(nil), "vega.commands.v1.KeyRotateSubmission")
}

func init() {
	proto.RegisterFile("vega/commands/v1/validator_commands.proto", fileDescriptor_52dee7aaa07d334d)
}

var fileDescriptor_52dee7aaa07d334d = []byte{
	// 734 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5b, 0x6e, 0xe3, 0x36,
	0x14, 0x1d, 0xbf, 0xe3, 0x1b, 0x3b, 0x55, 0x39, 0x03, 0x8c, 0x30, 0x40, 0x90, 0x89, 0x33, 0x40,
	0x33, 0x83, 0xc6, 0x4e, 0xdc, 0x17, 0xd0, 0x3f, 0x3f, 0xd4, 0xc6, 0x70, 0xe1, 0xa6, 0xb2, 0x9d,
	0x14, 0xfd, 0x11, 0x28, 0x89, 0x91, 0x59, 0x5b, 0x64, 0x40, 0x51, 0x76, 0xbc, 0x80, 0xee, 0xa0,
	0x3b, 0xe8, 0x2a, 0xba, 0x9a, 0x02, 0xfd, 0xeb, 0x6b, 0x0f, 0x05, 0x49, 0xcb, 0x89, 0xd3, 0xce,
	0x1f, 0x75, 0x1e, 0xf7, 0x5e, 0x1c, 0x5d, 0x12, 0xde, 0x2e, 0x49, 0x84, 0x5b, 0x01, 0x8f, 0x63,
	0xcc, 0xc2, 0xa4, 0xb5, 0xbc, 0x68, 0x2d, 0xf1, 0x82, 0x86, 0x58, 0x72, 0xe1, 0x65, 0x68, 0xf3,
	0x4e, 0x70, 0xc9, 0x91, 0xa5, 0xa4, 0xcd, 0x2d, 0xb8, 0xbc, 0x78, 0xf5, 0xd2, 0x98, 0x67, 0x98,
	0x32, 0x8f, 0x2c, 0x09, 0x93, 0x1b, 0xe9, 0xab, 0xcf, 0x23, 0x2a, 0x67, 0xa9, 0xaf, 0xc4, 0xad,
	0x78, 0x45, 0xe5, 0x9c, 0xaf, 0x5a, 0x11, 0x3f, 0xd3, 0xe4, 0xd9, 0xb6, 0x41, 0xf2, 0xd0, 0xcb,
	0xf8, 0x1a, 0xbf, 0xe4, 0xc1, 0x1a, 0xf1, 0x90, 0xb8, 0x24, 0xa2, 0x89, 0x14, 0x58, 0x52, 0xce,
	0xd0, 0x29, 0xd4, 0x54, 0x1f, 0xef, 0x2e, 0xf5, 0xbd, 0x39, 0x59, 0xdb, 0xb9, 0xd7, 0xb9, 0xd3,
	0x6a, 0xb7, 0xfc, 0xfb, 0x6f, 0x47, 0xf9, 0xef, 0x73, 0x2e, 0x28, 0xee, 0x2a, 0xf5, 0x87, 0x64,
	0x8d, 0x2e, 0xc0, 0x22, 0x72, 0x46, 0x04, 0x49, 0x63, 0x0f, 0x87, 0xa1, 0x20, 0x49, 0x62, 0xe7,
	0x77, 0xd4, 0x1f, 0x64, 0x7c, 0xc7, 0xd0, 0xe8, 0x1d, 0xd4, 0xcd, 0xfc, 0x59, 0xf5, 0xc2, 0x8e,
	0x7e, 0x5f, 0x93, 0x9b, 0xf2, 0xc7, 0xb0, 0x47, 0xd9, 0x2d, 0xf7, 0x52, 0xb1, 0xb0, 0x8b, 0x3b,
	0xb2, 0x8a, 0xc2, 0xa7, 0x62, 0x81, 0x5e, 0x43, 0x25, 0xe0, 0x29, 0x93, 0x62, 0x6d, 0x97, 0x76,
	0x15, 0x1b, 0x18, 0x1d, 0x40, 0x9e, 0x86, 0x76, 0x59, 0x91, 0x6e, 0x9e, 0x86, 0x08, 0x41, 0x91,
	0xe1, 0x98, 0xd8, 0x15, 0x8d, 0xe8, 0x33, 0x3a, 0x04, 0xc0, 0x4b, 0x2c, 0xb1, 0xd0, 0xad, 0xf6,
	0x34, 0x53, 0x35, 0xc8, 0x54, 0x2c, 0x1a, 0xdf, 0xc1, 0x9e, 0x0a, 0xe9, 0x9a, 0x4b, 0x82, 0x8e,
	0xa0, 0xf2, 0x38, 0x97, 0xda, 0xb6, 0x61, 0xf9, 0xce, 0x0c, 0xfd, 0x06, 0xaa, 0x82, 0xdc, 0x12,
	0x41, 0x58, 0x40, 0x9e, 0x84, 0xf1, 0x40, 0x34, 0x7e, 0x84, 0xba, 0x2a, 0x39, 0xa6, 0x11, 0xc3,
	0x32, 0x15, 0x64, 0x33, 0x66, 0x6e, 0x3b, 0xa6, 0x05, 0x85, 0x84, 0x46, 0xba, 0x40, 0xcd, 0x55,
	0x47, 0xf4, 0x05, 0x14, 0xe7, 0x94, 0x85, 0x3a, 0xb0, 0x83, 0xf6, 0x49, 0xf3, 0xe9, 0x76, 0x34,
	0x77, 0x0a, 0x0e, 0x29, 0x0b, 0x5d, 0x6d, 0x68, 0xfc, 0x9a, 0x07, 0xe8, 0xa9, 0x58, 0x1d, 0xb5,
	0x32, 0xe8, 0x39, 0x94, 0xe4, 0xbd, 0xb7, 0x6d, 0x56, 0x94, 0xf7, 0x83, 0x10, 0xbd, 0x80, 0x12,
	0xe3, 0xd9, 0xc4, 0x45, 0xd7, 0x7c, 0xa0, 0x4f, 0xa1, 0xe2, 0xa7, 0x74, 0x21, 0x29, 0xb3, 0xff,
	0x50, 0x79, 0xed, 0xb7, 0x5f, 0x9a, 0xb6, 0x5d, 0x83, 0x76, 0x92, 0x84, 0x48, 0x5d, 0xf5, 0xf2,
	0x99, 0x9b, 0x49, 0xd1, 0x5b, 0x28, 0x11, 0x11, 0xb4, 0xcf, 0xed, 0x3f, 0x8d, 0xc7, 0x32, 0x1e,
	0xc7, 0xed, 0xb5, 0xcf, 0x33, 0xb1, 0x51, 0xa0, 0x13, 0x28, 0xf8, 0x32, 0xb0, 0xff, 0x32, 0xc2,
	0x83, 0x4d, 0xf1, 0x49, 0x2f, 0x93, 0x29, 0x16, 0x7d, 0x06, 0xd5, 0xed, 0xde, 0xda, 0x7f, 0x1b,
	0xe9, 0x0b, 0x23, 0xbd, 0xce, 0xf0, 0xcc, 0xf0, 0xa0, 0x44, 0x5f, 0x42, 0x3d, 0x91, 0x78, 0x4e,
	0x59, 0x64, 0xee, 0x8a, 0xfd, 0x8f, 0xb1, 0x22, 0x63, 0x1d, 0x1b, 0x2e, 0x33, 0xd6, 0x92, 0x47,
	0xdf, 0xdd, 0x0a, 0x94, 0xb4, 0xa7, 0xf1, 0x73, 0x0e, 0x9e, 0x0f, 0xc9, 0xda, 0xe5, 0x12, 0x4b,
	0x32, 0x4e, 0xfd, 0x98, 0x26, 0x89, 0xba, 0x23, 0x87, 0x00, 0x73, 0xb2, 0xf6, 0x58, 0x1a, 0xfb,
	0x44, 0xe8, 0x24, 0x8b, 0x6e, 0x75, 0x4e, 0xd6, 0x23, 0x0d, 0xa0, 0x63, 0xa8, 0x49, 0x2c, 0x22,
	0x22, 0x3d, 0x7f, 0xc1, 0x83, 0xf9, 0x26, 0xd5, 0x7d, 0x83, 0x75, 0x15, 0xa4, 0xf6, 0x50, 0xd2,
	0x98, 0xe8, 0xdf, 0x59, 0x70, 0xf5, 0x19, 0x7d, 0x04, 0x16, 0x23, 0xab, 0xec, 0x6a, 0x78, 0x33,
	0x9c, 0xcc, 0xcc, 0xe2, 0xbb, 0x75, 0x46, 0x56, 0xe6, 0x56, 0x5c, 0xe2, 0x64, 0xf6, 0xee, 0xa7,
	0x1c, 0x7c, 0xf8, 0x9f, 0xdf, 0x8d, 0x4e, 0xe0, 0x68, 0xf4, 0x6d, 0xdf, 0xf1, 0xc6, 0x83, 0xaf,
	0x47, 0x9d, 0xc9, 0xd4, 0x75, 0xbc, 0xe1, 0x60, 0xd4, 0xf7, 0xa6, 0xa3, 0xf1, 0x95, 0xd3, 0x1b,
	0x7c, 0x35, 0x70, 0xfa, 0xd6, 0x33, 0x74, 0x0c, 0x87, 0xff, 0x27, 0xea, 0x8c, 0xc7, 0xce, 0xc4,
	0x1b, 0x39, 0x37, 0x56, 0x0e, 0x9d, 0xc2, 0x9b, 0xf7, 0x4b, 0x6e, 0x06, 0x93, 0xcb, 0xbe, 0xdb,
	0xb9, 0xe9, 0x7c, 0x63, 0xe5, 0xbb, 0xcd, 0x1f, 0x3e, 0x0e, 0x78, 0x48, 0x74, 0xa2, 0xfa, 0x45,
	0x09, 0xf8, 0xa2, 0x49, 0x79, 0x4b, 0x9f, 0x93, 0xd6, 0xd3, 0xb7, 0xce, 0x2f, 0x6b, 0xe2, 0x93,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x5f, 0x41, 0x1e, 0x06, 0x05, 0x00, 0x00,
}
