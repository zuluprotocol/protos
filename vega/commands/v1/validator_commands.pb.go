// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vega/commands/v1/validator_commands.proto

package v1

import (
	vega "code.vegaprotocol.io/protos/vega"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The kind of the signature created by a node, for example, allow-listing a new asset, withdrawal etc
type NodeSignatureKind int32

const (
	// Represents an unspecified or missing value from the input
	NodeSignatureKind_NODE_SIGNATURE_KIND_UNSPECIFIED NodeSignatureKind = 0
	// Represents a signature for a new asset allow-listing
	NodeSignatureKind_NODE_SIGNATURE_KIND_ASSET_NEW NodeSignatureKind = 1
	// Represents a signature for an asset withdrawal
	NodeSignatureKind_NODE_SIGNATURE_KIND_ASSET_WITHDRAWAL NodeSignatureKind = 2
	// Represents a signature for a new signer added to the erc20 multisig contract
	NodeSignatureKind_NODE_SIGNATURE_KIND_ERC20_MULTISIG_SIGNER_ADDED NodeSignatureKind = 3
	// Represents a signature for a signer removed from the erc20 multisig contract
	NodeSignatureKind_NODE_SIGNATURE_KIND_ERC20_MULTISIG_SIGNER_REMOVED NodeSignatureKind = 4
)

// Enum value maps for NodeSignatureKind.
var (
	NodeSignatureKind_name = map[int32]string{
		0: "NODE_SIGNATURE_KIND_UNSPECIFIED",
		1: "NODE_SIGNATURE_KIND_ASSET_NEW",
		2: "NODE_SIGNATURE_KIND_ASSET_WITHDRAWAL",
		3: "NODE_SIGNATURE_KIND_ERC20_MULTISIG_SIGNER_ADDED",
		4: "NODE_SIGNATURE_KIND_ERC20_MULTISIG_SIGNER_REMOVED",
	}
	NodeSignatureKind_value = map[string]int32{
		"NODE_SIGNATURE_KIND_UNSPECIFIED":                   0,
		"NODE_SIGNATURE_KIND_ASSET_NEW":                     1,
		"NODE_SIGNATURE_KIND_ASSET_WITHDRAWAL":              2,
		"NODE_SIGNATURE_KIND_ERC20_MULTISIG_SIGNER_ADDED":   3,
		"NODE_SIGNATURE_KIND_ERC20_MULTISIG_SIGNER_REMOVED": 4,
	}
)

func (x NodeSignatureKind) Enum() *NodeSignatureKind {
	p := new(NodeSignatureKind)
	*p = x
	return p
}

func (x NodeSignatureKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeSignatureKind) Descriptor() protoreflect.EnumDescriptor {
	return file_vega_commands_v1_validator_commands_proto_enumTypes[0].Descriptor()
}

func (NodeSignatureKind) Type() protoreflect.EnumType {
	return &file_vega_commands_v1_validator_commands_proto_enumTypes[0]
}

func (x NodeSignatureKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeSignatureKind.Descriptor instead.
func (NodeSignatureKind) EnumDescriptor() ([]byte, []int) {
	return file_vega_commands_v1_validator_commands_proto_rawDescGZIP(), []int{0}
}

// A message from a validator signaling they are still online and validating blocks
// or ready to validate block when they are till a potential validator
type ValidatorHeartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the id of the node emitting the heartbeat
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Signature from the validator made using the ethereum wallet
	EthereumSignature *Signature `protobuf:"bytes,2,opt,name=ethereum_signature,json=ethereumSignature,proto3" json:"ethereum_signature,omitempty"`
	// Signature from the validator made using the vega wallet
	VegaSignature *Signature `protobuf:"bytes,3,opt,name=vega_signature,json=vegaSignature,proto3" json:"vega_signature,omitempty"`
}

func (x *ValidatorHeartbeat) Reset() {
	*x = ValidatorHeartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorHeartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorHeartbeat) ProtoMessage() {}

func (x *ValidatorHeartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorHeartbeat.ProtoReflect.Descriptor instead.
func (*ValidatorHeartbeat) Descriptor() ([]byte, []int) {
	return file_vega_commands_v1_validator_commands_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatorHeartbeat) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *ValidatorHeartbeat) GetEthereumSignature() *Signature {
	if x != nil {
		return x.EthereumSignature
	}
	return nil
}

func (x *ValidatorHeartbeat) GetVegaSignature() *Signature {
	if x != nil {
		return x.VegaSignature
	}
	return nil
}

// Used announce a node as a new potential validator
type AnnounceNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	AvatarUrl string `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	// Vega public key derivation index
	VegaPubKeyIndex uint32 `protobuf:"varint,9,opt,name=vega_pub_key_index,json=vegaPubKeyIndex,proto3" json:"vega_pub_key_index,omitempty"`
	// The epoch from which the validator is expected
	// to be ready to validate blocks
	FromEpoch uint64 `protobuf:"varint,10,opt,name=from_epoch,json=fromEpoch,proto3" json:"from_epoch,omitempty"`
	// Signature from the validator made using the ethereum wallet
	EthereumSignature *Signature `protobuf:"bytes,11,opt,name=ethereum_signature,json=ethereumSignature,proto3" json:"ethereum_signature,omitempty"`
	// Signature from the validator made using the vega wallet
	VegaSignature *Signature `protobuf:"bytes,12,opt,name=vega_signature,json=vegaSignature,proto3" json:"vega_signature,omitempty"`
}

func (x *AnnounceNode) Reset() {
	*x = AnnounceNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceNode) ProtoMessage() {}

func (x *AnnounceNode) ProtoReflect() protoreflect.Message {
	mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceNode.ProtoReflect.Descriptor instead.
func (*AnnounceNode) Descriptor() ([]byte, []int) {
	return file_vega_commands_v1_validator_commands_proto_rawDescGZIP(), []int{1}
}

func (x *AnnounceNode) GetVegaPubKey() string {
	if x != nil {
		return x.VegaPubKey
	}
	return ""
}

func (x *AnnounceNode) GetEthereumAddress() string {
	if x != nil {
		return x.EthereumAddress
	}
	return ""
}

func (x *AnnounceNode) GetChainPubKey() string {
	if x != nil {
		return x.ChainPubKey
	}
	return ""
}

func (x *AnnounceNode) GetInfoUrl() string {
	if x != nil {
		return x.InfoUrl
	}
	return ""
}

func (x *AnnounceNode) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *AnnounceNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AnnounceNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnnounceNode) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *AnnounceNode) GetVegaPubKeyIndex() uint32 {
	if x != nil {
		return x.VegaPubKeyIndex
	}
	return 0
}

func (x *AnnounceNode) GetFromEpoch() uint64 {
	if x != nil {
		return x.FromEpoch
	}
	return 0
}

func (x *AnnounceNode) GetEthereumSignature() *Signature {
	if x != nil {
		return x.EthereumSignature
	}
	return nil
}

func (x *AnnounceNode) GetVegaSignature() *Signature {
	if x != nil {
		return x.VegaSignature
	}
	return nil
}

// Used when a node votes for validating a given resource exists or is valid,
// for example, an ERC20 deposit is valid and exists on ethereum
type NodeVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Public key, required field
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	// Reference, required field
	Reference string `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
}

func (x *NodeVote) Reset() {
	*x = NodeVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeVote) ProtoMessage() {}

func (x *NodeVote) ProtoReflect() protoreflect.Message {
	mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeVote.ProtoReflect.Descriptor instead.
func (*NodeVote) Descriptor() ([]byte, []int) {
	return file_vega_commands_v1_validator_commands_proto_rawDescGZIP(), []int{2}
}

func (x *NodeVote) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *NodeVote) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

// Represents a signature from a validator, to be used by a foreign chain in order to recognise a decision taken by the Vega network
type NodeSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the resource being signed
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The signature
	Sig []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	// The kind of resource being signed
	Kind NodeSignatureKind `protobuf:"varint,3,opt,name=kind,proto3,enum=vega.commands.v1.NodeSignatureKind" json:"kind,omitempty"`
}

func (x *NodeSignature) Reset() {
	*x = NodeSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeSignature) ProtoMessage() {}

func (x *NodeSignature) ProtoReflect() protoreflect.Message {
	mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeSignature.ProtoReflect.Descriptor instead.
func (*NodeSignature) Descriptor() ([]byte, []int) {
	return file_vega_commands_v1_validator_commands_proto_rawDescGZIP(), []int{3}
}

func (x *NodeSignature) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NodeSignature) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *NodeSignature) GetKind() NodeSignatureKind {
	if x != nil {
		return x.Kind
	}
	return NodeSignatureKind_NODE_SIGNATURE_KIND_UNSPECIFIED
}

// An event forwarded to the Vega network to provide information on events happening on other networks
type ChainEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the transaction in which the events happened, usually a hash
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// Arbitrary one-time integer used to prevent replay attacks
	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The event
	//
	// Types that are assignable to Event:
	//	*ChainEvent_Builtin
	//	*ChainEvent_Erc20
	//	*ChainEvent_StakingEvent
	//	*ChainEvent_Erc20Multisig
	Event isChainEvent_Event `protobuf_oneof:"event"`
}

func (x *ChainEvent) Reset() {
	*x = ChainEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainEvent) ProtoMessage() {}

func (x *ChainEvent) ProtoReflect() protoreflect.Message {
	mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainEvent.ProtoReflect.Descriptor instead.
func (*ChainEvent) Descriptor() ([]byte, []int) {
	return file_vega_commands_v1_validator_commands_proto_rawDescGZIP(), []int{4}
}

func (x *ChainEvent) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *ChainEvent) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (m *ChainEvent) GetEvent() isChainEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *ChainEvent) GetBuiltin() *vega.BuiltinAssetEvent {
	if x, ok := x.GetEvent().(*ChainEvent_Builtin); ok {
		return x.Builtin
	}
	return nil
}

func (x *ChainEvent) GetErc20() *vega.ERC20Event {
	if x, ok := x.GetEvent().(*ChainEvent_Erc20); ok {
		return x.Erc20
	}
	return nil
}

func (x *ChainEvent) GetStakingEvent() *vega.StakingEvent {
	if x, ok := x.GetEvent().(*ChainEvent_StakingEvent); ok {
		return x.StakingEvent
	}
	return nil
}

func (x *ChainEvent) GetErc20Multisig() *vega.ERC20MultiSigEvent {
	if x, ok := x.GetEvent().(*ChainEvent_Erc20Multisig); ok {
		return x.Erc20Multisig
	}
	return nil
}

type isChainEvent_Event interface {
	isChainEvent_Event()
}

type ChainEvent_Builtin struct {
	// Built-in asset event
	Builtin *vega.BuiltinAssetEvent `protobuf:"bytes,1001,opt,name=builtin,proto3,oneof"`
}

type ChainEvent_Erc20 struct {
	// Ethereum ERC20 event
	Erc20 *vega.ERC20Event `protobuf:"bytes,1002,opt,name=erc20,proto3,oneof"`
}

type ChainEvent_StakingEvent struct {
	// Ethereum Staking event
	StakingEvent *vega.StakingEvent `protobuf:"bytes,1005,opt,name=staking_event,json=stakingEvent,proto3,oneof"`
}

type ChainEvent_Erc20Multisig struct {
	// Ethereum ERC20 multisig event
	Erc20Multisig *vega.ERC20MultiSigEvent `protobuf:"bytes,1006,opt,name=erc20_multisig,json=erc20Multisig,proto3,oneof"`
}

func (*ChainEvent_Builtin) isChainEvent_Event() {}

func (*ChainEvent_Erc20) isChainEvent_Event() {}

func (*ChainEvent_StakingEvent) isChainEvent_Event() {}

func (*ChainEvent_Erc20Multisig) isChainEvent_Event() {}

// A transaction to allow validator to rotate their vega keys
type KeyRotateSubmission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// New Vega public key derivation index
	NewPubKeyIndex uint32 `protobuf:"varint,1,opt,name=new_pub_key_index,json=newPubKeyIndex,proto3" json:"new_pub_key_index,omitempty"`
	// Target block at which the key rotation will take effect on
	TargetBlock uint64 `protobuf:"varint,2,opt,name=target_block,json=targetBlock,proto3" json:"target_block,omitempty"`
	// The new public key to rotate to
	NewPubKey string `protobuf:"bytes,3,opt,name=new_pub_key,json=newPubKey,proto3" json:"new_pub_key,omitempty"`
	// Hash of currently used public key
	CurrentPubKeyHash string `protobuf:"bytes,4,opt,name=current_pub_key_hash,json=currentPubKeyHash,proto3" json:"current_pub_key_hash,omitempty"`
}

func (x *KeyRotateSubmission) Reset() {
	*x = KeyRotateSubmission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRotateSubmission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRotateSubmission) ProtoMessage() {}

func (x *KeyRotateSubmission) ProtoReflect() protoreflect.Message {
	mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRotateSubmission.ProtoReflect.Descriptor instead.
func (*KeyRotateSubmission) Descriptor() ([]byte, []int) {
	return file_vega_commands_v1_validator_commands_proto_rawDescGZIP(), []int{5}
}

func (x *KeyRotateSubmission) GetNewPubKeyIndex() uint32 {
	if x != nil {
		return x.NewPubKeyIndex
	}
	return 0
}

func (x *KeyRotateSubmission) GetTargetBlock() uint64 {
	if x != nil {
		return x.TargetBlock
	}
	return 0
}

func (x *KeyRotateSubmission) GetNewPubKey() string {
	if x != nil {
		return x.NewPubKey
	}
	return ""
}

func (x *KeyRotateSubmission) GetCurrentPubKeyHash() string {
	if x != nil {
		return x.CurrentPubKeyHash
	}
	return ""
}

type StateVariableProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state value proposal details
	Proposal *vega.StateValueProposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
}

func (x *StateVariableProposal) Reset() {
	*x = StateVariableProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateVariableProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateVariableProposal) ProtoMessage() {}

func (x *StateVariableProposal) ProtoReflect() protoreflect.Message {
	mi := &file_vega_commands_v1_validator_commands_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateVariableProposal.ProtoReflect.Descriptor instead.
func (*StateVariableProposal) Descriptor() ([]byte, []int) {
	return file_vega_commands_v1_validator_commands_proto_rawDescGZIP(), []int{6}
}

func (x *StateVariableProposal) GetProposal() *vega.StateValueProposal {
	if x != nil {
		return x.Proposal
	}
	return nil
}

var File_vega_commands_v1_validator_commands_proto protoreflect.FileDescriptor

var file_vega_commands_v1_validator_commands_proto_rawDesc = []byte{
	0x0a, 0x29, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x76, 0x65, 0x67,
	0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x0f, 0x76,
	0x65, 0x67, 0x61, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x65, 0x67, 0x61, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x12, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x12, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x11, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x76, 0x65, 0x67, 0x61, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0d, 0x76, 0x65, 0x67, 0x61,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xd3, 0x03, 0x0a, 0x0c, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x76, 0x65,
	0x67, 0x61, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x76, 0x65, 0x67, 0x61, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x10,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x6e, 0x66, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x55, 0x72, 0x6c, 0x12, 0x2b, 0x0a, 0x12, 0x76, 0x65, 0x67, 0x61, 0x5f, 0x70, 0x75, 0x62, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x76, 0x65, 0x67, 0x61, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12,
	0x4a, 0x0a, 0x12, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x65,
	0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x11, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x76,
	0x65, 0x67, 0x61, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x0d, 0x76, 0x65, 0x67, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0x41, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75,
	0x62, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x22, 0x6a, 0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x73, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0xb1,
	0x02, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a,
	0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c,
	0x74, 0x69, 0x6e, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x65, 0x67,
	0x61, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x12, 0x29,
	0x0a, 0x05, 0x65, 0x72, 0x63, 0x32, 0x30, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x63, 0x32, 0x30, 0x12, 0x3a, 0x0a, 0x0d, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0xed, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x5f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x18, 0xee, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x45, 0x52, 0x43, 0x32, 0x30, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x53, 0x69, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x72, 0x63, 0x32,
	0x30, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x4a, 0x06, 0x08, 0xeb, 0x07, 0x10, 0xec, 0x07, 0x4a, 0x06, 0x08, 0xec, 0x07, 0x10,
	0xed, 0x07, 0x22, 0xb4, 0x01, 0x0a, 0x13, 0x4b, 0x65, 0x79, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x11, 0x6e, 0x65,
	0x77, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f,
	0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x65, 0x77, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x22, 0x4d, 0x0a, 0x15, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2a, 0xf1, 0x01, 0x0a, 0x11, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x23,
	0x0a, 0x1f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45,
	0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e,
	0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54,
	0x5f, 0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x28, 0x0a, 0x24, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x53,
	0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x53,
	0x53, 0x45, 0x54, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x10, 0x02,
	0x12, 0x33, 0x0a, 0x2f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55,
	0x52, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x45, 0x52, 0x43, 0x32, 0x30, 0x5f, 0x4d, 0x55,
	0x4c, 0x54, 0x49, 0x53, 0x49, 0x47, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x52, 0x5f, 0x41, 0x44,
	0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x35, 0x0a, 0x31, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x49,
	0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x45, 0x52, 0x43,
	0x32, 0x30, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x53, 0x49, 0x47, 0x5f, 0x53, 0x49, 0x47, 0x4e,
	0x45, 0x52, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x04, 0x42, 0x2e, 0x5a, 0x2c,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x67, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vega_commands_v1_validator_commands_proto_rawDescOnce sync.Once
	file_vega_commands_v1_validator_commands_proto_rawDescData = file_vega_commands_v1_validator_commands_proto_rawDesc
)

func file_vega_commands_v1_validator_commands_proto_rawDescGZIP() []byte {
	file_vega_commands_v1_validator_commands_proto_rawDescOnce.Do(func() {
		file_vega_commands_v1_validator_commands_proto_rawDescData = protoimpl.X.CompressGZIP(file_vega_commands_v1_validator_commands_proto_rawDescData)
	})
	return file_vega_commands_v1_validator_commands_proto_rawDescData
}

var file_vega_commands_v1_validator_commands_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vega_commands_v1_validator_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_vega_commands_v1_validator_commands_proto_goTypes = []interface{}{
	(NodeSignatureKind)(0),          // 0: vega.commands.v1.NodeSignatureKind
	(*ValidatorHeartbeat)(nil),      // 1: vega.commands.v1.ValidatorHeartbeat
	(*AnnounceNode)(nil),            // 2: vega.commands.v1.AnnounceNode
	(*NodeVote)(nil),                // 3: vega.commands.v1.NodeVote
	(*NodeSignature)(nil),           // 4: vega.commands.v1.NodeSignature
	(*ChainEvent)(nil),              // 5: vega.commands.v1.ChainEvent
	(*KeyRotateSubmission)(nil),     // 6: vega.commands.v1.KeyRotateSubmission
	(*StateVariableProposal)(nil),   // 7: vega.commands.v1.StateVariableProposal
	(*Signature)(nil),               // 8: vega.commands.v1.Signature
	(*vega.BuiltinAssetEvent)(nil),  // 9: vega.BuiltinAssetEvent
	(*vega.ERC20Event)(nil),         // 10: vega.ERC20Event
	(*vega.StakingEvent)(nil),       // 11: vega.StakingEvent
	(*vega.ERC20MultiSigEvent)(nil), // 12: vega.ERC20MultiSigEvent
	(*vega.StateValueProposal)(nil), // 13: vega.StateValueProposal
}
var file_vega_commands_v1_validator_commands_proto_depIdxs = []int32{
	8,  // 0: vega.commands.v1.ValidatorHeartbeat.ethereum_signature:type_name -> vega.commands.v1.Signature
	8,  // 1: vega.commands.v1.ValidatorHeartbeat.vega_signature:type_name -> vega.commands.v1.Signature
	8,  // 2: vega.commands.v1.AnnounceNode.ethereum_signature:type_name -> vega.commands.v1.Signature
	8,  // 3: vega.commands.v1.AnnounceNode.vega_signature:type_name -> vega.commands.v1.Signature
	0,  // 4: vega.commands.v1.NodeSignature.kind:type_name -> vega.commands.v1.NodeSignatureKind
	9,  // 5: vega.commands.v1.ChainEvent.builtin:type_name -> vega.BuiltinAssetEvent
	10, // 6: vega.commands.v1.ChainEvent.erc20:type_name -> vega.ERC20Event
	11, // 7: vega.commands.v1.ChainEvent.staking_event:type_name -> vega.StakingEvent
	12, // 8: vega.commands.v1.ChainEvent.erc20_multisig:type_name -> vega.ERC20MultiSigEvent
	13, // 9: vega.commands.v1.StateVariableProposal.proposal:type_name -> vega.StateValueProposal
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_vega_commands_v1_validator_commands_proto_init() }
func file_vega_commands_v1_validator_commands_proto_init() {
	if File_vega_commands_v1_validator_commands_proto != nil {
		return
	}
	file_vega_commands_v1_signature_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vega_commands_v1_validator_commands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorHeartbeat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_commands_v1_validator_commands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceNode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_commands_v1_validator_commands_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeVote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_commands_v1_validator_commands_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeSignature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_commands_v1_validator_commands_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_commands_v1_validator_commands_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRotateSubmission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_commands_v1_validator_commands_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateVariableProposal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_vega_commands_v1_validator_commands_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ChainEvent_Builtin)(nil),
		(*ChainEvent_Erc20)(nil),
		(*ChainEvent_StakingEvent)(nil),
		(*ChainEvent_Erc20Multisig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vega_commands_v1_validator_commands_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vega_commands_v1_validator_commands_proto_goTypes,
		DependencyIndexes: file_vega_commands_v1_validator_commands_proto_depIdxs,
		EnumInfos:         file_vega_commands_v1_validator_commands_proto_enumTypes,
		MessageInfos:      file_vega_commands_v1_validator_commands_proto_msgTypes,
	}.Build()
	File_vega_commands_v1_validator_commands_proto = out.File
	file_vega_commands_v1_validator_commands_proto_rawDesc = nil
	file_vega_commands_v1_validator_commands_proto_goTypes = nil
	file_vega_commands_v1_validator_commands_proto_depIdxs = nil
}
