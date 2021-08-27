// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vega/snapshot/v1/snapshot.proto

package v1

import (
	vega "code.vegaprotocol.io/protos/vega"
	fmt "fmt"
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

// Snapshot is the entire checkpoint serialised (basically serialised the Checkpoint message + hash)
type Snapshot struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	State                []byte   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee6f703c00de307, []int{0}
}

func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snapshot.Unmarshal(m, b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return xxx_messageInfo_Snapshot.Size(m)
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Snapshot) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

// Checkpoint aggregates the various engine snapshots
type Checkpoint struct {
	Governance           []byte   `protobuf:"bytes,1,opt,name=governance,proto3" json:"governance,omitempty"`
	Assets               []byte   `protobuf:"bytes,2,opt,name=assets,proto3" json:"assets,omitempty"`
	Collateral           []byte   `protobuf:"bytes,3,opt,name=collateral,proto3" json:"collateral,omitempty"`
	NetworkParameters    []byte   `protobuf:"bytes,4,opt,name=network_parameters,json=networkParameters,proto3" json:"network_parameters,omitempty"`
	Delegation           []byte   `protobuf:"bytes,5,opt,name=delegation,proto3" json:"delegation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee6f703c00de307, []int{1}
}

func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetGovernance() []byte {
	if m != nil {
		return m.Governance
	}
	return nil
}

func (m *Checkpoint) GetAssets() []byte {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *Checkpoint) GetCollateral() []byte {
	if m != nil {
		return m.Collateral
	}
	return nil
}

func (m *Checkpoint) GetNetworkParameters() []byte {
	if m != nil {
		return m.NetworkParameters
	}
	return nil
}

func (m *Checkpoint) GetDelegation() []byte {
	if m != nil {
		return m.Delegation
	}
	return nil
}

// AssetEntrty is a single (enabled) asset
type AssetEntry struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AssetDetails         *vega.AssetDetails `protobuf:"bytes,2,opt,name=asset_details,json=assetDetails,proto3" json:"asset_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AssetEntry) Reset()         { *m = AssetEntry{} }
func (m *AssetEntry) String() string { return proto.CompactTextString(m) }
func (*AssetEntry) ProtoMessage()    {}
func (*AssetEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee6f703c00de307, []int{2}
}

func (m *AssetEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetEntry.Unmarshal(m, b)
}
func (m *AssetEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetEntry.Marshal(b, m, deterministic)
}
func (m *AssetEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetEntry.Merge(m, src)
}
func (m *AssetEntry) XXX_Size() int {
	return xxx_messageInfo_AssetEntry.Size(m)
}
func (m *AssetEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AssetEntry proto.InternalMessageInfo

func (m *AssetEntry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AssetEntry) GetAssetDetails() *vega.AssetDetails {
	if m != nil {
		return m.AssetDetails
	}
	return nil
}

// Assets contains all the enabled assets as AssetEntries
type Assets struct {
	Assets               []*AssetEntry `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Assets) Reset()         { *m = Assets{} }
func (m *Assets) String() string { return proto.CompactTextString(m) }
func (*Assets) ProtoMessage()    {}
func (*Assets) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee6f703c00de307, []int{3}
}

func (m *Assets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assets.Unmarshal(m, b)
}
func (m *Assets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assets.Marshal(b, m, deterministic)
}
func (m *Assets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assets.Merge(m, src)
}
func (m *Assets) XXX_Size() int {
	return xxx_messageInfo_Assets.Size(m)
}
func (m *Assets) XXX_DiscardUnknown() {
	xxx_messageInfo_Assets.DiscardUnknown(m)
}

var xxx_messageInfo_Assets proto.InternalMessageInfo

func (m *Assets) GetAssets() []*AssetEntry {
	if m != nil {
		return m.Assets
	}
	return nil
}

// AssetBalance represents the total balance of a given asset for a party
type AssetBalance struct {
	Party                string   `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
	Asset                string   `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Balance              string   `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetBalance) Reset()         { *m = AssetBalance{} }
func (m *AssetBalance) String() string { return proto.CompactTextString(m) }
func (*AssetBalance) ProtoMessage()    {}
func (*AssetBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee6f703c00de307, []int{4}
}

func (m *AssetBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetBalance.Unmarshal(m, b)
}
func (m *AssetBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetBalance.Marshal(b, m, deterministic)
}
func (m *AssetBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetBalance.Merge(m, src)
}
func (m *AssetBalance) XXX_Size() int {
	return xxx_messageInfo_AssetBalance.Size(m)
}
func (m *AssetBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetBalance.DiscardUnknown(m)
}

var xxx_messageInfo_AssetBalance proto.InternalMessageInfo

func (m *AssetBalance) GetParty() string {
	if m != nil {
		return m.Party
	}
	return ""
}

func (m *AssetBalance) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *AssetBalance) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

// Collateral contains the balances per party
type Collateral struct {
	Balances             []*AssetBalance `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Collateral) Reset()         { *m = Collateral{} }
func (m *Collateral) String() string { return proto.CompactTextString(m) }
func (*Collateral) ProtoMessage()    {}
func (*Collateral) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee6f703c00de307, []int{5}
}

func (m *Collateral) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collateral.Unmarshal(m, b)
}
func (m *Collateral) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collateral.Marshal(b, m, deterministic)
}
func (m *Collateral) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collateral.Merge(m, src)
}
func (m *Collateral) XXX_Size() int {
	return xxx_messageInfo_Collateral.Size(m)
}
func (m *Collateral) XXX_DiscardUnknown() {
	xxx_messageInfo_Collateral.DiscardUnknown(m)
}

var xxx_messageInfo_Collateral proto.InternalMessageInfo

func (m *Collateral) GetBalances() []*AssetBalance {
	if m != nil {
		return m.Balances
	}
	return nil
}

// NetParams contains all network parameters
type NetParams struct {
	Params               []*vega.NetworkParameter `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NetParams) Reset()         { *m = NetParams{} }
func (m *NetParams) String() string { return proto.CompactTextString(m) }
func (*NetParams) ProtoMessage()    {}
func (*NetParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee6f703c00de307, []int{6}
}

func (m *NetParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetParams.Unmarshal(m, b)
}
func (m *NetParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetParams.Marshal(b, m, deterministic)
}
func (m *NetParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetParams.Merge(m, src)
}
func (m *NetParams) XXX_Size() int {
	return xxx_messageInfo_NetParams.Size(m)
}
func (m *NetParams) XXX_DiscardUnknown() {
	xxx_messageInfo_NetParams.DiscardUnknown(m)
}

var xxx_messageInfo_NetParams proto.InternalMessageInfo

func (m *NetParams) GetParams() []*vega.NetworkParameter {
	if m != nil {
		return m.Params
	}
	return nil
}

// Proposals will contain all accepted proposals
type Proposals struct {
	Proposals            []*vega.Proposal `protobuf:"bytes,1,rep,name=proposals,proto3" json:"proposals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Proposals) Reset()         { *m = Proposals{} }
func (m *Proposals) String() string { return proto.CompactTextString(m) }
func (*Proposals) ProtoMessage()    {}
func (*Proposals) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee6f703c00de307, []int{7}
}

func (m *Proposals) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposals.Unmarshal(m, b)
}
func (m *Proposals) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposals.Marshal(b, m, deterministic)
}
func (m *Proposals) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposals.Merge(m, src)
}
func (m *Proposals) XXX_Size() int {
	return xxx_messageInfo_Proposals.Size(m)
}
func (m *Proposals) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposals.DiscardUnknown(m)
}

var xxx_messageInfo_Proposals proto.InternalMessageInfo

func (m *Proposals) GetProposals() []*vega.Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

// Delegated amounts for party/node
// undelegate and epoch seq are only relevant for pending entries
type DelegateEntry struct {
	Party                string   `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
	Node                 string   `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Amount               string   `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Undelegate           bool     `protobuf:"varint,4,opt,name=undelegate,proto3" json:"undelegate,omitempty"`
	EpochSeq             uint64   `protobuf:"varint,5,opt,name=epoch_seq,json=epochSeq,proto3" json:"epoch_seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelegateEntry) Reset()         { *m = DelegateEntry{} }
func (m *DelegateEntry) String() string { return proto.CompactTextString(m) }
func (*DelegateEntry) ProtoMessage()    {}
func (*DelegateEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee6f703c00de307, []int{8}
}

func (m *DelegateEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelegateEntry.Unmarshal(m, b)
}
func (m *DelegateEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelegateEntry.Marshal(b, m, deterministic)
}
func (m *DelegateEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegateEntry.Merge(m, src)
}
func (m *DelegateEntry) XXX_Size() int {
	return xxx_messageInfo_DelegateEntry.Size(m)
}
func (m *DelegateEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegateEntry.DiscardUnknown(m)
}

var xxx_messageInfo_DelegateEntry proto.InternalMessageInfo

func (m *DelegateEntry) GetParty() string {
	if m != nil {
		return m.Party
	}
	return ""
}

func (m *DelegateEntry) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *DelegateEntry) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *DelegateEntry) GetUndelegate() bool {
	if m != nil {
		return m.Undelegate
	}
	return false
}

func (m *DelegateEntry) GetEpochSeq() uint64 {
	if m != nil {
		return m.EpochSeq
	}
	return 0
}

// Delegate contains all entries for a checkpoint
type Delegate struct {
	Active               []*DelegateEntry `protobuf:"bytes,1,rep,name=active,proto3" json:"active,omitempty"`
	Pending              []*DelegateEntry `protobuf:"bytes,2,rep,name=pending,proto3" json:"pending,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Delegate) Reset()         { *m = Delegate{} }
func (m *Delegate) String() string { return proto.CompactTextString(m) }
func (*Delegate) ProtoMessage()    {}
func (*Delegate) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee6f703c00de307, []int{9}
}

func (m *Delegate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Delegate.Unmarshal(m, b)
}
func (m *Delegate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Delegate.Marshal(b, m, deterministic)
}
func (m *Delegate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delegate.Merge(m, src)
}
func (m *Delegate) XXX_Size() int {
	return xxx_messageInfo_Delegate.Size(m)
}
func (m *Delegate) XXX_DiscardUnknown() {
	xxx_messageInfo_Delegate.DiscardUnknown(m)
}

var xxx_messageInfo_Delegate proto.InternalMessageInfo

func (m *Delegate) GetActive() []*DelegateEntry {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Delegate) GetPending() []*DelegateEntry {
	if m != nil {
		return m.Pending
	}
	return nil
}

func init() {
	proto.RegisterType((*Snapshot)(nil), "vega.snapshot.v1.Snapshot")
	proto.RegisterType((*Checkpoint)(nil), "vega.snapshot.v1.Checkpoint")
	proto.RegisterType((*AssetEntry)(nil), "vega.snapshot.v1.AssetEntry")
	proto.RegisterType((*Assets)(nil), "vega.snapshot.v1.Assets")
	proto.RegisterType((*AssetBalance)(nil), "vega.snapshot.v1.AssetBalance")
	proto.RegisterType((*Collateral)(nil), "vega.snapshot.v1.Collateral")
	proto.RegisterType((*NetParams)(nil), "vega.snapshot.v1.NetParams")
	proto.RegisterType((*Proposals)(nil), "vega.snapshot.v1.Proposals")
	proto.RegisterType((*DelegateEntry)(nil), "vega.snapshot.v1.DelegateEntry")
	proto.RegisterType((*Delegate)(nil), "vega.snapshot.v1.Delegate")
}

func init() { proto.RegisterFile("vega/snapshot/v1/snapshot.proto", fileDescriptor_cee6f703c00de307) }

var fileDescriptor_cee6f703c00de307 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x4f, 0xdb, 0x4e,
	0x10, 0x55, 0x42, 0x12, 0xe2, 0x21, 0xf0, 0xfb, 0xb1, 0xa2, 0xc8, 0xa2, 0x15, 0xa0, 0x3d, 0x71,
	0xa0, 0x8e, 0xa0, 0x48, 0x88, 0x56, 0xaa, 0xc4, 0x9f, 0x4a, 0x3d, 0x21, 0xb4, 0xb4, 0x97, 0x5e,
	0xa2, 0xc5, 0x1e, 0xc5, 0x16, 0x66, 0xd7, 0x78, 0x17, 0x57, 0x5c, 0xfa, 0x11, 0xfa, 0x65, 0xfa,
	0x05, 0x2b, 0xcf, 0xae, 0xed, 0x34, 0x6d, 0xa5, 0xde, 0x76, 0xde, 0xbe, 0x37, 0x79, 0xfb, 0x66,
	0x62, 0xd8, 0xab, 0x70, 0x2e, 0xa7, 0x46, 0xc9, 0xc2, 0xa4, 0xda, 0x4e, 0xab, 0xa3, 0xf6, 0x1c,
	0x15, 0xa5, 0xb6, 0x9a, 0xfd, 0x5f, 0x13, 0xa2, 0x16, 0xac, 0x8e, 0x76, 0xfe, 0x23, 0x09, 0xc1,
	0x44, 0xd9, 0xd9, 0x24, 0x40, 0x1a, 0x83, 0xd6, 0x78, 0xe8, 0x05, 0x41, 0x73, 0x5d, 0x61, 0xa9,
	0xa4, 0x8a, 0xd1, 0xc1, 0xfc, 0x04, 0xc6, 0xb7, 0xbe, 0x13, 0x63, 0x30, 0x48, 0xa5, 0x49, 0xc3,
	0xde, 0x7e, 0xef, 0x60, 0x22, 0xe8, 0xcc, 0xb6, 0x60, 0x68, 0xac, 0xb4, 0x18, 0xf6, 0x09, 0x74,
	0x05, 0xff, 0xd1, 0x03, 0xb8, 0x4c, 0x31, 0xbe, 0x2f, 0x74, 0xa6, 0x2c, 0xdb, 0x05, 0xe8, 0x1a,
	0x7b, 0xf9, 0x02, 0xc2, 0xb6, 0x61, 0xe4, 0xbc, 0xf8, 0x2e, 0xbe, 0xaa, 0x75, 0xb1, 0xce, 0x73,
	0x69, 0xb1, 0x94, 0x79, 0xb8, 0xe2, 0x74, 0x1d, 0xc2, 0x5e, 0x03, 0x53, 0x68, 0xbf, 0xea, 0xf2,
	0x7e, 0x56, 0xc8, 0x52, 0x3e, 0xa0, 0xc5, 0xd2, 0x84, 0x03, 0xe2, 0x6d, 0xfa, 0x9b, 0x9b, 0xf6,
	0xa2, 0x6e, 0x97, 0x60, 0x8e, 0x73, 0x69, 0x33, 0xad, 0xc2, 0xa1, 0x6b, 0xd7, 0x21, 0xfc, 0x33,
	0xc0, 0x79, 0xfd, 0xc3, 0x1f, 0x94, 0x2d, 0x9f, 0xd9, 0x06, 0xf4, 0xb3, 0x84, 0xcc, 0x06, 0xa2,
	0x9f, 0x25, 0xec, 0x14, 0xd6, 0xc9, 0xd6, 0x2c, 0x41, 0x2b, 0xb3, 0xdc, 0x79, 0x5d, 0x3b, 0x66,
	0x11, 0xe5, 0x4a, 0xc2, 0x2b, 0x77, 0x23, 0x26, 0x72, 0xa1, 0xe2, 0xef, 0x61, 0x74, 0xee, 0xde,
	0x73, 0xd2, 0xbe, 0xb3, 0xb7, 0xbf, 0x72, 0xb0, 0x76, 0xfc, 0x2a, 0x5a, 0x1e, 0x55, 0xd4, 0x19,
	0x68, 0x52, 0xe0, 0x9f, 0x60, 0x42, 0xe8, 0x85, 0xcc, 0x29, 0xad, 0x2d, 0x18, 0x16, 0xb2, 0xb4,
	0xcf, 0xde, 0x9b, 0x2b, 0x6a, 0x94, 0xf8, 0x64, 0x2b, 0x10, 0xae, 0x60, 0x21, 0xac, 0xde, 0x39,
	0x19, 0xc5, 0x17, 0x88, 0xa6, 0xe4, 0x1f, 0x01, 0x2e, 0xbb, 0x24, 0xdf, 0xc2, 0xd8, 0x5f, 0x34,
	0xde, 0x76, 0xff, 0xe2, 0xcd, 0xbb, 0x10, 0x2d, 0x9f, 0xbf, 0x83, 0xe0, 0x1a, 0x2d, 0xe5, 0x6c,
	0x58, 0x04, 0x23, 0x1a, 0x45, 0xd3, 0x66, 0xdb, 0xb5, 0xb9, 0x5e, 0x1a, 0x86, 0xf0, 0x2c, 0x7e,
	0x06, 0xc1, 0x4d, 0xa9, 0x0b, 0x6d, 0x64, 0x6e, 0xd8, 0x21, 0x04, 0x45, 0x53, 0x78, 0xfd, 0x86,
	0xd3, 0x37, 0x1c, 0xd1, 0x11, 0xf8, 0xf7, 0x1e, 0xac, 0x5f, 0xb9, 0xe9, 0xa1, 0x1b, 0xd9, 0x9f,
	0x93, 0x61, 0x30, 0x50, 0x3a, 0x41, 0x1f, 0x0c, 0x9d, 0x69, 0xe3, 0x1e, 0xf4, 0x93, 0xb2, 0x3e,
	0x16, 0x5f, 0xd5, 0x2b, 0xf2, 0xa4, 0xfc, 0x4a, 0x20, 0x6d, 0xd2, 0x58, 0x2c, 0x20, 0xec, 0x25,
	0x04, 0x58, 0xe8, 0x38, 0x9d, 0x19, 0x7c, 0xa4, 0x0d, 0x1a, 0x88, 0x31, 0x01, 0xb7, 0xf8, 0xc8,
	0xbf, 0xc1, 0xb8, 0xf1, 0xc3, 0x4e, 0x61, 0x24, 0x63, 0x9b, 0x55, 0xe8, 0xdf, 0xb1, 0xf7, 0x7b,
	0x9c, 0xbf, 0x78, 0x17, 0x9e, 0xce, 0xce, 0x60, 0xb5, 0x40, 0x95, 0x64, 0x6a, 0x1e, 0xf6, 0xff,
	0x4d, 0xd9, 0xf0, 0x2f, 0xa2, 0x2f, 0x87, 0xb1, 0x4e, 0x90, 0xf8, 0xf4, 0xef, 0x8d, 0x75, 0x1e,
	0x65, 0x7a, 0x4a, 0x67, 0x33, 0x5d, 0xfe, 0x6e, 0xdc, 0x8d, 0xe8, 0xe2, 0xcd, 0xcf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xae, 0x36, 0x7d, 0xdd, 0x52, 0x04, 0x00, 0x00,
}
