// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet/v1/wallet.proto

package v1

import (
	v1 "code.vegaprotocol.io/protos/vega/commands/v1"
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

type SubmitTransactionRequest struct {
	PubKey    string `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Propagate bool   `protobuf:"varint,2,opt,name=propagate,proto3" json:"propagate,omitempty"`
	// Types that are valid to be assigned to Command:
	//	*SubmitTransactionRequest_OrderSubmission
	//	*SubmitTransactionRequest_OrderCancellation
	//	*SubmitTransactionRequest_OrderAmendment
	//	*SubmitTransactionRequest_WithdrawSubmission
	//	*SubmitTransactionRequest_ProposalSubmission
	//	*SubmitTransactionRequest_VoteSubmission
	//	*SubmitTransactionRequest_LiquidityProvisionSubmission
	//	*SubmitTransactionRequest_NodeRegistration
	//	*SubmitTransactionRequest_NodeVote
	//	*SubmitTransactionRequest_NodeSignature
	//	*SubmitTransactionRequest_ChainEvent
	//	*SubmitTransactionRequest_OracleDataSubmission
	Command              isSubmitTransactionRequest_Command `protobuf_oneof:"command"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *SubmitTransactionRequest) Reset()         { *m = SubmitTransactionRequest{} }
func (m *SubmitTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionRequest) ProtoMessage()    {}
func (*SubmitTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a45fb539ff6b1e, []int{0}
}

func (m *SubmitTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionRequest.Unmarshal(m, b)
}
func (m *SubmitTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionRequest.Marshal(b, m, deterministic)
}
func (m *SubmitTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionRequest.Merge(m, src)
}
func (m *SubmitTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionRequest.Size(m)
}
func (m *SubmitTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionRequest proto.InternalMessageInfo

func (m *SubmitTransactionRequest) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *SubmitTransactionRequest) GetPropagate() bool {
	if m != nil {
		return m.Propagate
	}
	return false
}

type isSubmitTransactionRequest_Command interface {
	isSubmitTransactionRequest_Command()
}

type SubmitTransactionRequest_OrderSubmission struct {
	OrderSubmission *v1.OrderSubmission `protobuf:"bytes,1001,opt,name=order_submission,json=orderSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_OrderCancellation struct {
	OrderCancellation *v1.OrderCancellation `protobuf:"bytes,1002,opt,name=order_cancellation,json=orderCancellation,proto3,oneof"`
}

type SubmitTransactionRequest_OrderAmendment struct {
	OrderAmendment *v1.OrderAmendment `protobuf:"bytes,1003,opt,name=order_amendment,json=orderAmendment,proto3,oneof"`
}

type SubmitTransactionRequest_WithdrawSubmission struct {
	WithdrawSubmission *v1.WithdrawSubmission `protobuf:"bytes,1004,opt,name=withdraw_submission,json=withdrawSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_ProposalSubmission struct {
	ProposalSubmission *v1.ProposalSubmission `protobuf:"bytes,1005,opt,name=proposal_submission,json=proposalSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_VoteSubmission struct {
	VoteSubmission *v1.VoteSubmission `protobuf:"bytes,1006,opt,name=vote_submission,json=voteSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_LiquidityProvisionSubmission struct {
	LiquidityProvisionSubmission *v1.LiquidityProvisionSubmission `protobuf:"bytes,1007,opt,name=liquidity_provision_submission,json=liquidityProvisionSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_NodeRegistration struct {
	NodeRegistration *v1.NodeRegistration `protobuf:"bytes,2001,opt,name=node_registration,json=nodeRegistration,proto3,oneof"`
}

type SubmitTransactionRequest_NodeVote struct {
	NodeVote *v1.NodeVote `protobuf:"bytes,2002,opt,name=node_vote,json=nodeVote,proto3,oneof"`
}

type SubmitTransactionRequest_NodeSignature struct {
	NodeSignature *v1.NodeSignature `protobuf:"bytes,2003,opt,name=node_signature,json=nodeSignature,proto3,oneof"`
}

type SubmitTransactionRequest_ChainEvent struct {
	ChainEvent *v1.ChainEvent `protobuf:"bytes,2004,opt,name=chain_event,json=chainEvent,proto3,oneof"`
}

type SubmitTransactionRequest_OracleDataSubmission struct {
	OracleDataSubmission *v1.OracleDataSubmission `protobuf:"bytes,3001,opt,name=oracle_data_submission,json=oracleDataSubmission,proto3,oneof"`
}

func (*SubmitTransactionRequest_OrderSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_OrderCancellation) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_OrderAmendment) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_WithdrawSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_ProposalSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_VoteSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_LiquidityProvisionSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_NodeRegistration) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_NodeVote) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_NodeSignature) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_ChainEvent) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_OracleDataSubmission) isSubmitTransactionRequest_Command() {}

func (m *SubmitTransactionRequest) GetCommand() isSubmitTransactionRequest_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *SubmitTransactionRequest) GetOrderSubmission() *v1.OrderSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_OrderSubmission); ok {
		return x.OrderSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetOrderCancellation() *v1.OrderCancellation {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_OrderCancellation); ok {
		return x.OrderCancellation
	}
	return nil
}

func (m *SubmitTransactionRequest) GetOrderAmendment() *v1.OrderAmendment {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_OrderAmendment); ok {
		return x.OrderAmendment
	}
	return nil
}

func (m *SubmitTransactionRequest) GetWithdrawSubmission() *v1.WithdrawSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_WithdrawSubmission); ok {
		return x.WithdrawSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetProposalSubmission() *v1.ProposalSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_ProposalSubmission); ok {
		return x.ProposalSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetVoteSubmission() *v1.VoteSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_VoteSubmission); ok {
		return x.VoteSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetLiquidityProvisionSubmission() *v1.LiquidityProvisionSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_LiquidityProvisionSubmission); ok {
		return x.LiquidityProvisionSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetNodeRegistration() *v1.NodeRegistration {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_NodeRegistration); ok {
		return x.NodeRegistration
	}
	return nil
}

func (m *SubmitTransactionRequest) GetNodeVote() *v1.NodeVote {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_NodeVote); ok {
		return x.NodeVote
	}
	return nil
}

func (m *SubmitTransactionRequest) GetNodeSignature() *v1.NodeSignature {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_NodeSignature); ok {
		return x.NodeSignature
	}
	return nil
}

func (m *SubmitTransactionRequest) GetChainEvent() *v1.ChainEvent {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_ChainEvent); ok {
		return x.ChainEvent
	}
	return nil
}

func (m *SubmitTransactionRequest) GetOracleDataSubmission() *v1.OracleDataSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_OracleDataSubmission); ok {
		return x.OracleDataSubmission
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubmitTransactionRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubmitTransactionRequest_OrderSubmission)(nil),
		(*SubmitTransactionRequest_OrderCancellation)(nil),
		(*SubmitTransactionRequest_OrderAmendment)(nil),
		(*SubmitTransactionRequest_WithdrawSubmission)(nil),
		(*SubmitTransactionRequest_ProposalSubmission)(nil),
		(*SubmitTransactionRequest_VoteSubmission)(nil),
		(*SubmitTransactionRequest_LiquidityProvisionSubmission)(nil),
		(*SubmitTransactionRequest_NodeRegistration)(nil),
		(*SubmitTransactionRequest_NodeVote)(nil),
		(*SubmitTransactionRequest_NodeSignature)(nil),
		(*SubmitTransactionRequest_ChainEvent)(nil),
		(*SubmitTransactionRequest_OracleDataSubmission)(nil),
	}
}

func init() {
	proto.RegisterType((*SubmitTransactionRequest)(nil), "vega.wallet.v1.SubmitTransactionRequest")
}

func init() { proto.RegisterFile("wallet/v1/wallet.proto", fileDescriptor_73a45fb539ff6b1e) }

var fileDescriptor_73a45fb539ff6b1e = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0xc7, 0xd3, 0xe7, 0x62, 0x6d, 0x3d, 0x3d, 0xdd, 0x66, 0xd0, 0x16, 0xaa, 0x0a, 0xca, 0x98,
	0x50, 0x85, 0x50, 0xaa, 0xc2, 0x1d, 0x57, 0xb0, 0x81, 0x14, 0xc4, 0xc4, 0x26, 0x97, 0x37, 0x71,
	0x13, 0xb9, 0x89, 0xd5, 0x59, 0xa4, 0x3e, 0x99, 0xe3, 0xa4, 0xea, 0x47, 0xe3, 0x33, 0x70, 0xc5,
	0xcb, 0x97, 0xe0, 0xfd, 0x2b, 0x20, 0xc7, 0x49, 0x49, 0x9a, 0x96, 0xbb, 0xd3, 0xbf, 0xff, 0xe7,
	0x77, 0x72, 0x8e, 0x4f, 0x8d, 0xf6, 0xe7, 0x34, 0x0c, 0x99, 0x1a, 0xa6, 0xa3, 0xa1, 0x89, 0x9c,
	0x48, 0x82, 0x02, 0xdc, 0x49, 0xd9, 0x94, 0x3a, 0xb9, 0x94, 0x8e, 0xba, 0x5d, 0x1f, 0x66, 0x33,
	0x2a, 0x82, 0x58, 0x3b, 0x8b, 0xd8, 0x78, 0xbb, 0x47, 0xe5, 0xb3, 0x94, 0x86, 0x3c, 0xa0, 0x0a,
	0xa4, 0xb7, 0xe2, 0xba, 0x56, 0x76, 0x81, 0xa4, 0x7e, 0xc8, 0xf2, 0xa3, 0xc3, 0x0f, 0x2d, 0x64,
	0x8f, 0x93, 0xc9, 0x8c, 0xab, 0x17, 0x92, 0x8a, 0x98, 0xfa, 0x8a, 0x83, 0x20, 0xec, 0x32, 0x61,
	0xb1, 0xc2, 0x07, 0xa8, 0x19, 0x25, 0x13, 0xef, 0x1d, 0x5b, 0xd8, 0x8d, 0x7e, 0x63, 0xd0, 0x26,
	0x5b, 0x51, 0x32, 0x79, 0xc6, 0x16, 0xb8, 0x87, 0xda, 0x91, 0x84, 0x88, 0x4e, 0xa9, 0x62, 0xf6,
	0x7f, 0xfd, 0xc6, 0xa0, 0x45, 0xfe, 0x0a, 0xf8, 0x0c, 0xed, 0x82, 0x0c, 0x98, 0xf4, 0x62, 0x0d,
	0x8e, 0x63, 0x0e, 0xc2, 0xfe, 0xda, 0xec, 0x37, 0x06, 0xdb, 0xf7, 0x6e, 0x3a, 0x59, 0x73, 0xcb,
	0xef, 0x4b, 0x47, 0xce, 0x99, 0xb6, 0x8e, 0x97, 0x4e, 0xd7, 0x22, 0x3b, 0x50, 0x95, 0xf0, 0x4b,
	0x84, 0x0d, 0xd0, 0xa7, 0xc2, 0x67, 0x61, 0x48, 0xf5, 0x47, 0xda, 0xdf, 0x0c, 0xf2, 0xd6, 0x06,
	0xe4, 0x49, 0xc9, 0xeb, 0x5a, 0x64, 0x0f, 0x56, 0x45, 0x7c, 0x8a, 0x4c, 0x25, 0x8f, 0xce, 0x98,
	0x08, 0x66, 0x4c, 0x28, 0xfb, 0xbb, 0x61, 0xf6, 0x37, 0x30, 0x1f, 0x15, 0x46, 0xd7, 0x22, 0x1d,
	0xa8, 0x28, 0xf8, 0x0d, 0xba, 0x32, 0xe7, 0xea, 0x22, 0x90, 0x74, 0x5e, 0x6e, 0xfc, 0x87, 0x21,
	0x1e, 0xd5, 0x89, 0xaf, 0x73, 0x77, 0xa5, 0x77, 0x3c, 0xaf, 0xa9, 0x9a, 0xac, 0x87, 0x0b, 0x31,
	0x0d, 0xcb, 0xe4, 0x9f, 0x1b, 0xc9, 0xe7, 0xb9, 0xbb, 0x4a, 0x8e, 0x6a, 0xaa, 0x9e, 0x40, 0x0a,
	0x8a, 0x95, 0xa9, 0xbf, 0x36, 0x4e, 0xe0, 0x15, 0x28, 0x56, 0x21, 0x76, 0xd2, 0x8a, 0x82, 0xe7,
	0xe8, 0x7a, 0xc8, 0x2f, 0x13, 0x1e, 0x70, 0xb5, 0xf0, 0x22, 0x09, 0x29, 0xd7, 0x72, 0x19, 0xfe,
	0xdb, 0xc0, 0x9d, 0x3a, 0xfc, 0xb4, 0x48, 0x3c, 0x2f, 0xf2, 0x2a, 0xa5, 0x7a, 0xe1, 0x3f, 0xce,
	0x31, 0x41, 0x7b, 0x02, 0x02, 0xe6, 0x49, 0x36, 0xe5, 0xb1, 0x92, 0x66, 0x3d, 0x3e, 0xee, 0x64,
	0xb5, 0x0e, 0xeb, 0xb5, 0x9e, 0x43, 0xc0, 0x48, 0xc9, 0xea, 0x5a, 0x64, 0x57, 0xac, 0x68, 0xf8,
	0x01, 0x6a, 0x67, 0x4c, 0xdd, 0xa3, 0xfd, 0xc9, 0xb0, 0xba, 0xeb, 0x59, 0x7a, 0x30, 0xae, 0x45,
	0x5a, 0x22, 0x8f, 0xf1, 0x53, 0xd4, 0xc9, 0x72, 0x63, 0x3e, 0x15, 0x54, 0x25, 0x92, 0xd9, 0x9f,
	0x0d, 0xe0, 0xc6, 0x7a, 0xc0, 0xb8, 0xf0, 0xb9, 0x16, 0xf9, 0x5f, 0x94, 0x05, 0xfc, 0x10, 0x6d,
	0xfb, 0x17, 0x94, 0x0b, 0x8f, 0xa5, 0x7a, 0x3f, 0xbf, 0x18, 0x4e, 0xaf, 0xce, 0x39, 0xd1, 0xae,
	0x27, 0xa9, 0xd9, 0x4d, 0xe4, 0x2f, 0x7f, 0x61, 0x0f, 0xed, 0x9b, 0xbf, 0xbc, 0x17, 0x50, 0x45,
	0xcb, 0xb7, 0xf1, 0xfe, 0x20, 0x83, 0xdd, 0x5e, 0xb7, 0xec, 0x3a, 0xe1, 0x31, 0x55, 0xb4, 0x72,
	0x0b, 0x57, 0x61, 0x8d, 0x7e, 0xdc, 0x46, 0xcd, 0x3c, 0xf7, 0xf8, 0xee, 0xdb, 0x3b, 0x3e, 0x04,
	0x2c, 0x03, 0x66, 0xef, 0x8b, 0x0f, 0xa1, 0xc3, 0x61, 0x98, 0xc5, 0xf1, 0x50, 0xcb, 0xc3, 0xe5,
	0xab, 0x37, 0xd9, 0xca, 0xe4, 0xfb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x22, 0x63, 0xfd, 0xbf,
	0x09, 0x05, 0x00, 0x00,
}
