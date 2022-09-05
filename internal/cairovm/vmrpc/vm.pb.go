// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: vm.proto

package vmrpc

import (
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

type NodeType int32

const (
	NodeType_EDGE_NODE   NodeType = 0
	NodeType_BINARY_NODE NodeType = 1
)

// Enum value maps for NodeType.
var (
	NodeType_name = map[int32]string{
		0: "EDGE_NODE",
		1: "BINARY_NODE",
	}
	NodeType_value = map[string]int32{
		"EDGE_NODE":   0,
		"BINARY_NODE": 1,
	}
)

func (x NodeType) Enum() *NodeType {
	p := new(NodeType)
	*p = x
	return p
}

func (x NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_vm_proto_enumTypes[0].Descriptor()
}

func (NodeType) Type() protoreflect.EnumType {
	return &file_vm_proto_enumTypes[0]
}

func (x NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeType.Descriptor instead.
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return file_vm_proto_rawDescGZIP(), []int{0}
}

type GetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetValueRequest) Reset() {
	*x = GetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueRequest) ProtoMessage() {}

func (x *GetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueRequest.ProtoReflect.Descriptor instead.
func (*GetValueRequest) Descriptor() ([]byte, []int) {
	return file_vm_proto_rawDescGZIP(), []int{0}
}

func (x *GetValueRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type VMTrieNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Len    uint32   `protobuf:"varint,1,opt,name=len,proto3" json:"len,omitempty"`
	Bottom []byte   `protobuf:"bytes,2,opt,name=bottom,proto3" json:"bottom,omitempty"`
	Path   []byte   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Left   []byte   `protobuf:"bytes,4,opt,name=left,proto3" json:"left,omitempty"`
	Right  []byte   `protobuf:"bytes,5,opt,name=right,proto3" json:"right,omitempty"`
	Type   NodeType `protobuf:"varint,6,opt,name=type,proto3,enum=NodeType" json:"type,omitempty"`
}

func (x *VMTrieNode) Reset() {
	*x = VMTrieNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMTrieNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMTrieNode) ProtoMessage() {}

func (x *VMTrieNode) ProtoReflect() protoreflect.Message {
	mi := &file_vm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMTrieNode.ProtoReflect.Descriptor instead.
func (*VMTrieNode) Descriptor() ([]byte, []int) {
	return file_vm_proto_rawDescGZIP(), []int{1}
}

func (x *VMTrieNode) GetLen() uint32 {
	if x != nil {
		return x.Len
	}
	return 0
}

func (x *VMTrieNode) GetBottom() []byte {
	if x != nil {
		return x.Bottom
	}
	return nil
}

func (x *VMTrieNode) GetPath() []byte {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *VMTrieNode) GetLeft() []byte {
	if x != nil {
		return x.Left
	}
	return nil
}

func (x *VMTrieNode) GetRight() []byte {
	if x != nil {
		return x.Right
	}
	return nil
}

func (x *VMTrieNode) GetType() NodeType {
	if x != nil {
		return x.Type
	}
	return NodeType_EDGE_NODE
}

type VMContractState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractHash []byte `protobuf:"bytes,1,opt,name=contractHash,proto3" json:"contractHash,omitempty"`
	StorageRoot  []byte `protobuf:"bytes,2,opt,name=storageRoot,proto3" json:"storageRoot,omitempty"`
	Height       uint32 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *VMContractState) Reset() {
	*x = VMContractState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMContractState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMContractState) ProtoMessage() {}

func (x *VMContractState) ProtoReflect() protoreflect.Message {
	mi := &file_vm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMContractState.ProtoReflect.Descriptor instead.
func (*VMContractState) Descriptor() ([]byte, []int) {
	return file_vm_proto_rawDescGZIP(), []int{2}
}

func (x *VMContractState) GetContractHash() []byte {
	if x != nil {
		return x.ContractHash
	}
	return nil
}

func (x *VMContractState) GetStorageRoot() []byte {
	if x != nil {
		return x.StorageRoot
	}
	return nil
}

func (x *VMContractState) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type VMContractDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *VMContractDefinition) Reset() {
	*x = VMContractDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMContractDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMContractDefinition) ProtoMessage() {}

func (x *VMContractDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_vm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMContractDefinition.ProtoReflect.Descriptor instead.
func (*VMContractDefinition) Descriptor() ([]byte, []int) {
	return file_vm_proto_rawDescGZIP(), []int{3}
}

func (x *VMContractDefinition) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type VMCallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calldata        [][]byte `protobuf:"bytes,1,rep,name=calldata,proto3" json:"calldata,omitempty"`
	CallerAddress   []byte   `protobuf:"bytes,2,opt,name=caller_address,json=callerAddress,proto3" json:"caller_address,omitempty"`
	ContractAddress []byte   `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	ClassHash       []byte   `protobuf:"bytes,4,opt,name=class_hash,json=classHash,proto3" json:"class_hash,omitempty"`
	Root            []byte   `protobuf:"bytes,5,opt,name=root,proto3" json:"root,omitempty"`
	Selector        []byte   `protobuf:"bytes,6,opt,name=selector,proto3" json:"selector,omitempty"`
	Sequencer       []byte   `protobuf:"bytes,7,opt,name=sequencer,proto3" json:"sequencer,omitempty"`
}

func (x *VMCallRequest) Reset() {
	*x = VMCallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMCallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMCallRequest) ProtoMessage() {}

func (x *VMCallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMCallRequest.ProtoReflect.Descriptor instead.
func (*VMCallRequest) Descriptor() ([]byte, []int) {
	return file_vm_proto_rawDescGZIP(), []int{4}
}

func (x *VMCallRequest) GetCalldata() [][]byte {
	if x != nil {
		return x.Calldata
	}
	return nil
}

func (x *VMCallRequest) GetCallerAddress() []byte {
	if x != nil {
		return x.CallerAddress
	}
	return nil
}

func (x *VMCallRequest) GetContractAddress() []byte {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

func (x *VMCallRequest) GetClassHash() []byte {
	if x != nil {
		return x.ClassHash
	}
	return nil
}

func (x *VMCallRequest) GetRoot() []byte {
	if x != nil {
		return x.Root
	}
	return nil
}

func (x *VMCallRequest) GetSelector() []byte {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *VMCallRequest) GetSequencer() []byte {
	if x != nil {
		return x.Sequencer
	}
	return nil
}

type VMCallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retdata [][]byte `protobuf:"bytes,1,rep,name=retdata,proto3" json:"retdata,omitempty"`
}

func (x *VMCallResponse) Reset() {
	*x = VMCallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMCallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMCallResponse) ProtoMessage() {}

func (x *VMCallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMCallResponse.ProtoReflect.Descriptor instead.
func (*VMCallResponse) Descriptor() ([]byte, []int) {
	return file_vm_proto_rawDescGZIP(), []int{5}
}

func (x *VMCallResponse) GetRetdata() [][]byte {
	if x != nil {
		return x.Retdata
	}
	return nil
}

var File_vm_proto protoreflect.FileDescriptor

var file_vm_proto_rawDesc = []byte{
	0x0a, 0x08, 0x76, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x93, 0x01, 0x0a, 0x0a, 0x56, 0x4d, 0x54, 0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6c, 0x65, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x65, 0x66, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6f, 0x0a, 0x0f, 0x56, 0x4d, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x2c, 0x0a, 0x14, 0x56, 0x4d, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xea, 0x01, 0x0a, 0x0d, 0x56, 0x4d, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x63, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x72, 0x22, 0x2a, 0x0a, 0x0e, 0x56, 0x4d, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x2a, 0x0a,
	0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x44, 0x47,
	0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x49, 0x4e, 0x41,
	0x52, 0x59, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x01, 0x32, 0xc2, 0x01, 0x0a, 0x0e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x72, 0x69, 0x63, 0x69, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x56, 0x4d, 0x54, 0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x56, 0x4d, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x56, 0x4d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x32, 0x2f,
	0x0a, 0x02, 0x56, 0x4d, 0x12, 0x29, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x56,
	0x4d, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x56,
	0x4d, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x45, 0x74, 0x68, 0x2f, 0x6a, 0x75, 0x6e, 0x6f,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x6d, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vm_proto_rawDescOnce sync.Once
	file_vm_proto_rawDescData = file_vm_proto_rawDesc
)

func file_vm_proto_rawDescGZIP() []byte {
	file_vm_proto_rawDescOnce.Do(func() {
		file_vm_proto_rawDescData = protoimpl.X.CompressGZIP(file_vm_proto_rawDescData)
	})
	return file_vm_proto_rawDescData
}

var file_vm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vm_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vm_proto_goTypes = []interface{}{
	(NodeType)(0),                // 0: NodeType
	(*GetValueRequest)(nil),      // 1: GetValueRequest
	(*VMTrieNode)(nil),           // 2: VMTrieNode
	(*VMContractState)(nil),      // 3: VMContractState
	(*VMContractDefinition)(nil), // 4: VMContractDefinition
	(*VMCallRequest)(nil),        // 5: VMCallRequest
	(*VMCallResponse)(nil),       // 6: VMCallResponse
}
var file_vm_proto_depIdxs = []int32{
	0, // 0: VMTrieNode.type:type_name -> NodeType
	1, // 1: StorageAdapter.GetPatriciaNode:input_type -> GetValueRequest
	1, // 2: StorageAdapter.GetContractState:input_type -> GetValueRequest
	1, // 3: StorageAdapter.GetContractDefinition:input_type -> GetValueRequest
	5, // 4: VM.Call:input_type -> VMCallRequest
	2, // 5: StorageAdapter.GetPatriciaNode:output_type -> VMTrieNode
	3, // 6: StorageAdapter.GetContractState:output_type -> VMContractState
	4, // 7: StorageAdapter.GetContractDefinition:output_type -> VMContractDefinition
	6, // 8: VM.Call:output_type -> VMCallResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vm_proto_init() }
func file_vm_proto_init() {
	if File_vm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueRequest); i {
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
		file_vm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMTrieNode); i {
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
		file_vm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMContractState); i {
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
		file_vm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMContractDefinition); i {
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
		file_vm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMCallRequest); i {
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
		file_vm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMCallResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_vm_proto_goTypes,
		DependencyIndexes: file_vm_proto_depIdxs,
		EnumInfos:         file_vm_proto_enumTypes,
		MessageInfos:      file_vm_proto_msgTypes,
	}.Build()
	File_vm_proto = out.File
	file_vm_proto_rawDesc = nil
	file_vm_proto_goTypes = nil
	file_vm_proto_depIdxs = nil
}