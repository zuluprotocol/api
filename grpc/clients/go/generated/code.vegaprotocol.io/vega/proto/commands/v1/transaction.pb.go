// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: commands/v1/transaction.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type InputData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A random number used to provided uniqueness and prevents
	// against replay attack.
	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The block height associated to the transaction.
	// This should always be current height of the node at the time of sending the Tx.
	// BlockHeight is used as a mechanism for replay protection.
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Types that are assignable to Command:
	//	*InputData_OrderSubmission
	//	*InputData_OrderCancellation
	//	*InputData_OrderAmendment
	//	*InputData_WithdrawSubmission
	//	*InputData_ProposalSubmission
	//	*InputData_VoteSubmission
	//	*InputData_LiquidityProvisionSubmission
	//	*InputData_NodeRegistration
	//	*InputData_NodeVote
	//	*InputData_NodeSignature
	//	*InputData_ChainEvent
	//	*InputData_OracleDataSubmission
	Command isInputData_Command `protobuf_oneof:"command"`
}

func (x *InputData) Reset() {
	*x = InputData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_v1_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputData) ProtoMessage() {}

func (x *InputData) ProtoReflect() protoreflect.Message {
	mi := &file_commands_v1_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputData.ProtoReflect.Descriptor instead.
func (*InputData) Descriptor() ([]byte, []int) {
	return file_commands_v1_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *InputData) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *InputData) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (m *InputData) GetCommand() isInputData_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *InputData) GetOrderSubmission() *OrderSubmission {
	if x, ok := x.GetCommand().(*InputData_OrderSubmission); ok {
		return x.OrderSubmission
	}
	return nil
}

func (x *InputData) GetOrderCancellation() *OrderCancellation {
	if x, ok := x.GetCommand().(*InputData_OrderCancellation); ok {
		return x.OrderCancellation
	}
	return nil
}

func (x *InputData) GetOrderAmendment() *OrderAmendment {
	if x, ok := x.GetCommand().(*InputData_OrderAmendment); ok {
		return x.OrderAmendment
	}
	return nil
}

func (x *InputData) GetWithdrawSubmission() *WithdrawSubmission {
	if x, ok := x.GetCommand().(*InputData_WithdrawSubmission); ok {
		return x.WithdrawSubmission
	}
	return nil
}

func (x *InputData) GetProposalSubmission() *ProposalSubmission {
	if x, ok := x.GetCommand().(*InputData_ProposalSubmission); ok {
		return x.ProposalSubmission
	}
	return nil
}

func (x *InputData) GetVoteSubmission() *VoteSubmission {
	if x, ok := x.GetCommand().(*InputData_VoteSubmission); ok {
		return x.VoteSubmission
	}
	return nil
}

func (x *InputData) GetLiquidityProvisionSubmission() *LiquidityProvisionSubmission {
	if x, ok := x.GetCommand().(*InputData_LiquidityProvisionSubmission); ok {
		return x.LiquidityProvisionSubmission
	}
	return nil
}

func (x *InputData) GetNodeRegistration() *NodeRegistration {
	if x, ok := x.GetCommand().(*InputData_NodeRegistration); ok {
		return x.NodeRegistration
	}
	return nil
}

func (x *InputData) GetNodeVote() *NodeVote {
	if x, ok := x.GetCommand().(*InputData_NodeVote); ok {
		return x.NodeVote
	}
	return nil
}

func (x *InputData) GetNodeSignature() *NodeSignature {
	if x, ok := x.GetCommand().(*InputData_NodeSignature); ok {
		return x.NodeSignature
	}
	return nil
}

func (x *InputData) GetChainEvent() *ChainEvent {
	if x, ok := x.GetCommand().(*InputData_ChainEvent); ok {
		return x.ChainEvent
	}
	return nil
}

func (x *InputData) GetOracleDataSubmission() *OracleDataSubmission {
	if x, ok := x.GetCommand().(*InputData_OracleDataSubmission); ok {
		return x.OracleDataSubmission
	}
	return nil
}

type isInputData_Command interface {
	isInputData_Command()
}

type InputData_OrderSubmission struct {
	// user commands
	OrderSubmission *OrderSubmission `protobuf:"bytes,1001,opt,name=order_submission,json=orderSubmission,proto3,oneof"`
}

type InputData_OrderCancellation struct {
	OrderCancellation *OrderCancellation `protobuf:"bytes,1002,opt,name=order_cancellation,json=orderCancellation,proto3,oneof"`
}

type InputData_OrderAmendment struct {
	OrderAmendment *OrderAmendment `protobuf:"bytes,1003,opt,name=order_amendment,json=orderAmendment,proto3,oneof"`
}

type InputData_WithdrawSubmission struct {
	WithdrawSubmission *WithdrawSubmission `protobuf:"bytes,1004,opt,name=withdraw_submission,json=withdrawSubmission,proto3,oneof"`
}

type InputData_ProposalSubmission struct {
	ProposalSubmission *ProposalSubmission `protobuf:"bytes,1005,opt,name=proposal_submission,json=proposalSubmission,proto3,oneof"`
}

type InputData_VoteSubmission struct {
	VoteSubmission *VoteSubmission `protobuf:"bytes,1006,opt,name=vote_submission,json=voteSubmission,proto3,oneof"`
}

type InputData_LiquidityProvisionSubmission struct {
	LiquidityProvisionSubmission *LiquidityProvisionSubmission `protobuf:"bytes,1007,opt,name=liquidity_provision_submission,json=liquidityProvisionSubmission,proto3,oneof"`
}

type InputData_NodeRegistration struct {
	// validator commands
	NodeRegistration *NodeRegistration `protobuf:"bytes,2001,opt,name=node_registration,json=nodeRegistration,proto3,oneof"`
}

type InputData_NodeVote struct {
	NodeVote *NodeVote `protobuf:"bytes,2002,opt,name=node_vote,json=nodeVote,proto3,oneof"`
}

type InputData_NodeSignature struct {
	NodeSignature *NodeSignature `protobuf:"bytes,2003,opt,name=node_signature,json=nodeSignature,proto3,oneof"`
}

type InputData_ChainEvent struct {
	ChainEvent *ChainEvent `protobuf:"bytes,2004,opt,name=chain_event,json=chainEvent,proto3,oneof"`
}

type InputData_OracleDataSubmission struct {
	// Oracles
	OracleDataSubmission *OracleDataSubmission `protobuf:"bytes,3001,opt,name=oracle_data_submission,json=oracleDataSubmission,proto3,oneof"`
}

func (*InputData_OrderSubmission) isInputData_Command() {}

func (*InputData_OrderCancellation) isInputData_Command() {}

func (*InputData_OrderAmendment) isInputData_Command() {}

func (*InputData_WithdrawSubmission) isInputData_Command() {}

func (*InputData_ProposalSubmission) isInputData_Command() {}

func (*InputData_VoteSubmission) isInputData_Command() {}

func (*InputData_LiquidityProvisionSubmission) isInputData_Command() {}

func (*InputData_NodeRegistration) isInputData_Command() {}

func (*InputData_NodeVote) isInputData_Command() {}

func (*InputData_NodeSignature) isInputData_Command() {}

func (*InputData_ChainEvent) isInputData_Command() {}

func (*InputData_OracleDataSubmission) isInputData_Command() {}

// Represents a transaction to be sent to Vega.
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One of the set of Vega commands (proto marshalled).
	InputData []byte `protobuf:"bytes,1,opt,name=input_data,json=inputData,proto3" json:"input_data,omitempty"`
	// The signature of the inputData
	Signature *Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// The sender of the transaction.
	// Any of the following would be valid:
	//
	// Types that are assignable to From:
	//	*Transaction_Address
	//	*Transaction_PubKey
	From isTransaction_From `protobuf_oneof:"from"`
	// A version of the transaction, to be used
	// in the future in case we want to implement
	// changes to the Transaction format
	Version uint32 `protobuf:"varint,2000,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_v1_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_commands_v1_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_commands_v1_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetInputData() []byte {
	if x != nil {
		return x.InputData
	}
	return nil
}

func (x *Transaction) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (m *Transaction) GetFrom() isTransaction_From {
	if m != nil {
		return m.From
	}
	return nil
}

func (x *Transaction) GetAddress() []byte {
	if x, ok := x.GetFrom().(*Transaction_Address); ok {
		return x.Address
	}
	return nil
}

func (x *Transaction) GetPubKey() []byte {
	if x, ok := x.GetFrom().(*Transaction_PubKey); ok {
		return x.PubKey
	}
	return nil
}

func (x *Transaction) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type isTransaction_From interface {
	isTransaction_From()
}

type Transaction_Address struct {
	// The address of the sender.
	Address []byte `protobuf:"bytes,1001,opt,name=address,proto3,oneof"`
}

type Transaction_PubKey struct {
	// The public key of the sender.
	PubKey []byte `protobuf:"bytes,1002,opt,name=pub_key,json=pubKey,proto3,oneof"`
}

func (*Transaction_Address) isTransaction_From() {}

func (*Transaction_PubKey) isTransaction_From() {}

// A signature to be authenticate a transaction
// and to be verified by the vega network
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bytes of the signature
	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	// The algorithm used to create the signature
	Algo string `protobuf:"bytes,2,opt,name=algo,proto3" json:"algo,omitempty"`
	// The version of the signature used to create the signature
	Version uint32 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_v1_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_commands_v1_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_commands_v1_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *Signature) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Signature) GetAlgo() string {
	if x != nil {
		return x.Algo
	}
	return ""
}

func (x *Signature) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_commands_v1_transaction_proto protoreflect.FileDescriptor

var file_commands_v1_transaction_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe,
	0x08, 0x0a, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x4f, 0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xea, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x11, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a,
	0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0xeb, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x41, 0x6d, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x41, 0x6d, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x58, 0x0a, 0x13, 0x77,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0xec, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x65, 0x67, 0x61,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x12, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xed, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x12, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x4c, 0x0a, 0x0f, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0xee, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x65, 0x67, 0x61,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x76,
	0x6f, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x77, 0x0a,
	0x1e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0xef, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x1c, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x11, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xd1, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x18, 0xd2, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x6f,
	0x64, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0xd3, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x48, 0x00, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0xd4, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x5f, 0x0a, 0x16, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xb9, 0x17,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x14,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22,
	0xc3, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x12, 0x19, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xd0, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x4f, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x6c, 0x67, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x4f, 0x0a, 0x20, 0x69, 0x6f, 0x2e, 0x76, 0x65, 0x67,
	0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x5a, 0x2b, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x76, 0x65, 0x67, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x69, 0x6f,
	0x2f, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commands_v1_transaction_proto_rawDescOnce sync.Once
	file_commands_v1_transaction_proto_rawDescData = file_commands_v1_transaction_proto_rawDesc
)

func file_commands_v1_transaction_proto_rawDescGZIP() []byte {
	file_commands_v1_transaction_proto_rawDescOnce.Do(func() {
		file_commands_v1_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_commands_v1_transaction_proto_rawDescData)
	})
	return file_commands_v1_transaction_proto_rawDescData
}

var file_commands_v1_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_commands_v1_transaction_proto_goTypes = []interface{}{
	(*InputData)(nil),                    // 0: vega.commands.v1.InputData
	(*Transaction)(nil),                  // 1: vega.commands.v1.Transaction
	(*Signature)(nil),                    // 2: vega.commands.v1.Signature
	(*OrderSubmission)(nil),              // 3: vega.commands.v1.OrderSubmission
	(*OrderCancellation)(nil),            // 4: vega.commands.v1.OrderCancellation
	(*OrderAmendment)(nil),               // 5: vega.commands.v1.OrderAmendment
	(*WithdrawSubmission)(nil),           // 6: vega.commands.v1.WithdrawSubmission
	(*ProposalSubmission)(nil),           // 7: vega.commands.v1.ProposalSubmission
	(*VoteSubmission)(nil),               // 8: vega.commands.v1.VoteSubmission
	(*LiquidityProvisionSubmission)(nil), // 9: vega.commands.v1.LiquidityProvisionSubmission
	(*NodeRegistration)(nil),             // 10: vega.commands.v1.NodeRegistration
	(*NodeVote)(nil),                     // 11: vega.commands.v1.NodeVote
	(*NodeSignature)(nil),                // 12: vega.commands.v1.NodeSignature
	(*ChainEvent)(nil),                   // 13: vega.commands.v1.ChainEvent
	(*OracleDataSubmission)(nil),         // 14: vega.commands.v1.OracleDataSubmission
}
var file_commands_v1_transaction_proto_depIdxs = []int32{
	3,  // 0: vega.commands.v1.InputData.order_submission:type_name -> vega.commands.v1.OrderSubmission
	4,  // 1: vega.commands.v1.InputData.order_cancellation:type_name -> vega.commands.v1.OrderCancellation
	5,  // 2: vega.commands.v1.InputData.order_amendment:type_name -> vega.commands.v1.OrderAmendment
	6,  // 3: vega.commands.v1.InputData.withdraw_submission:type_name -> vega.commands.v1.WithdrawSubmission
	7,  // 4: vega.commands.v1.InputData.proposal_submission:type_name -> vega.commands.v1.ProposalSubmission
	8,  // 5: vega.commands.v1.InputData.vote_submission:type_name -> vega.commands.v1.VoteSubmission
	9,  // 6: vega.commands.v1.InputData.liquidity_provision_submission:type_name -> vega.commands.v1.LiquidityProvisionSubmission
	10, // 7: vega.commands.v1.InputData.node_registration:type_name -> vega.commands.v1.NodeRegistration
	11, // 8: vega.commands.v1.InputData.node_vote:type_name -> vega.commands.v1.NodeVote
	12, // 9: vega.commands.v1.InputData.node_signature:type_name -> vega.commands.v1.NodeSignature
	13, // 10: vega.commands.v1.InputData.chain_event:type_name -> vega.commands.v1.ChainEvent
	14, // 11: vega.commands.v1.InputData.oracle_data_submission:type_name -> vega.commands.v1.OracleDataSubmission
	2,  // 12: vega.commands.v1.Transaction.signature:type_name -> vega.commands.v1.Signature
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_commands_v1_transaction_proto_init() }
func file_commands_v1_transaction_proto_init() {
	if File_commands_v1_transaction_proto != nil {
		return
	}
	file_commands_v1_commands_proto_init()
	file_commands_v1_validator_commands_proto_init()
	file_commands_v1_oracles_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commands_v1_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputData); i {
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
		file_commands_v1_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_commands_v1_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
	file_commands_v1_transaction_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*InputData_OrderSubmission)(nil),
		(*InputData_OrderCancellation)(nil),
		(*InputData_OrderAmendment)(nil),
		(*InputData_WithdrawSubmission)(nil),
		(*InputData_ProposalSubmission)(nil),
		(*InputData_VoteSubmission)(nil),
		(*InputData_LiquidityProvisionSubmission)(nil),
		(*InputData_NodeRegistration)(nil),
		(*InputData_NodeVote)(nil),
		(*InputData_NodeSignature)(nil),
		(*InputData_ChainEvent)(nil),
		(*InputData_OracleDataSubmission)(nil),
	}
	file_commands_v1_transaction_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Transaction_Address)(nil),
		(*Transaction_PubKey)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commands_v1_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commands_v1_transaction_proto_goTypes,
		DependencyIndexes: file_commands_v1_transaction_proto_depIdxs,
		MessageInfos:      file_commands_v1_transaction_proto_msgTypes,
	}.Build()
	File_commands_v1_transaction_proto = out.File
	file_commands_v1_transaction_proto_rawDesc = nil
	file_commands_v1_transaction_proto_goTypes = nil
	file_commands_v1_transaction_proto_depIdxs = nil
}