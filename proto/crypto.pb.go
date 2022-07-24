// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.2
// source: proto/crypto.proto

package proto

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

type Crypto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Votes int64  `protobuf:"varint,3,opt,name=votes,proto3" json:"votes,omitempty"`
}

func (x *Crypto) Reset() {
	*x = Crypto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crypto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crypto) ProtoMessage() {}

func (x *Crypto) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crypto.ProtoReflect.Descriptor instead.
func (*Crypto) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{0}
}

func (x *Crypto) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Crypto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Crypto) GetVotes() int64 {
	if x != nil {
		return x.Votes
	}
	return 0
}

type CreateCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto *Crypto `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *CreateCryptoRequest) Reset() {
	*x = CreateCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCryptoRequest) ProtoMessage() {}

func (x *CreateCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCryptoRequest.ProtoReflect.Descriptor instead.
func (*CreateCryptoRequest) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCryptoRequest) GetCrypto() *Crypto {
	if x != nil {
		return x.Crypto
	}
	return nil
}

type CreateCryptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto *Crypto `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *CreateCryptoResponse) Reset() {
	*x = CreateCryptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCryptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCryptoResponse) ProtoMessage() {}

func (x *CreateCryptoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCryptoResponse.ProtoReflect.Descriptor instead.
func (*CreateCryptoResponse) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCryptoResponse) GetCrypto() *Crypto {
	if x != nil {
		return x.Crypto
	}
	return nil
}

type GetCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCryptoRequest) Reset() {
	*x = GetCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCryptoRequest) ProtoMessage() {}

func (x *GetCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCryptoRequest.ProtoReflect.Descriptor instead.
func (*GetCryptoRequest) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{3}
}

func (x *GetCryptoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCryptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto *Crypto `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *GetCryptoResponse) Reset() {
	*x = GetCryptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCryptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCryptoResponse) ProtoMessage() {}

func (x *GetCryptoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCryptoResponse.ProtoReflect.Descriptor instead.
func (*GetCryptoResponse) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{4}
}

func (x *GetCryptoResponse) GetCrypto() *Crypto {
	if x != nil {
		return x.Crypto
	}
	return nil
}

type UpdateCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateCryptoRequest) Reset() {
	*x = UpdateCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCryptoRequest) ProtoMessage() {}

func (x *UpdateCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCryptoRequest.ProtoReflect.Descriptor instead.
func (*UpdateCryptoRequest) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCryptoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateCryptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto *Crypto `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *UpdateCryptoResponse) Reset() {
	*x = UpdateCryptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCryptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCryptoResponse) ProtoMessage() {}

func (x *UpdateCryptoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCryptoResponse.ProtoReflect.Descriptor instead.
func (*UpdateCryptoResponse) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCryptoResponse) GetCrypto() *Crypto {
	if x != nil {
		return x.Crypto
	}
	return nil
}

type DeleteCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCryptoRequest) Reset() {
	*x = DeleteCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCryptoRequest) ProtoMessage() {}

func (x *DeleteCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCryptoRequest.ProtoReflect.Descriptor instead.
func (*DeleteCryptoRequest) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCryptoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCryptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto *Crypto `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *DeleteCryptoResponse) Reset() {
	*x = DeleteCryptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCryptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCryptoResponse) ProtoMessage() {}

func (x *DeleteCryptoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCryptoResponse.ProtoReflect.Descriptor instead.
func (*DeleteCryptoResponse) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCryptoResponse) GetCrypto() *Crypto {
	if x != nil {
		return x.Crypto
	}
	return nil
}

type UpVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpVoteRequest) Reset() {
	*x = UpVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpVoteRequest) ProtoMessage() {}

func (x *UpVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpVoteRequest.ProtoReflect.Descriptor instead.
func (*UpVoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{9}
}

func (x *UpVoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto *Crypto `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *UpVoteResponse) Reset() {
	*x = UpVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpVoteResponse) ProtoMessage() {}

func (x *UpVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpVoteResponse.ProtoReflect.Descriptor instead.
func (*UpVoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{10}
}

func (x *UpVoteResponse) GetCrypto() *Crypto {
	if x != nil {
		return x.Crypto
	}
	return nil
}

type DownVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DownVoteRequest) Reset() {
	*x = DownVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownVoteRequest) ProtoMessage() {}

func (x *DownVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownVoteRequest.ProtoReflect.Descriptor instead.
func (*DownVoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{11}
}

func (x *DownVoteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DownVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto *Crypto `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *DownVoteResponse) Reset() {
	*x = DownVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownVoteResponse) ProtoMessage() {}

func (x *DownVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownVoteResponse.ProtoReflect.Descriptor instead.
func (*DownVoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_crypto_proto_rawDescGZIP(), []int{12}
}

func (x *DownVoteResponse) GetCrypto() *Crypto {
	if x != nil {
		return x.Crypto
	}
	return nil
}

var File_proto_crypto_proto protoreflect.FileDescriptor

var file_proto_crypto_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x06, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x22,
	0x3c, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x22, 0x3d, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x3a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x52, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0d, 0x55, 0x70, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x0e, 0x55, 0x70, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x32, 0x9e, 0x03, 0x0a, 0x0d, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x06, 0x55, 0x70, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x70, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x56, 0x6f, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_crypto_proto_rawDescOnce sync.Once
	file_proto_crypto_proto_rawDescData = file_proto_crypto_proto_rawDesc
)

func file_proto_crypto_proto_rawDescGZIP() []byte {
	file_proto_crypto_proto_rawDescOnce.Do(func() {
		file_proto_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_crypto_proto_rawDescData)
	})
	return file_proto_crypto_proto_rawDescData
}

var file_proto_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_crypto_proto_goTypes = []interface{}{
	(*Crypto)(nil),               // 0: proto.Crypto
	(*CreateCryptoRequest)(nil),  // 1: proto.CreateCryptoRequest
	(*CreateCryptoResponse)(nil), // 2: proto.CreateCryptoResponse
	(*GetCryptoRequest)(nil),     // 3: proto.GetCryptoRequest
	(*GetCryptoResponse)(nil),    // 4: proto.GetCryptoResponse
	(*UpdateCryptoRequest)(nil),  // 5: proto.UpdateCryptoRequest
	(*UpdateCryptoResponse)(nil), // 6: proto.UpdateCryptoResponse
	(*DeleteCryptoRequest)(nil),  // 7: proto.DeleteCryptoRequest
	(*DeleteCryptoResponse)(nil), // 8: proto.DeleteCryptoResponse
	(*UpVoteRequest)(nil),        // 9: proto.UpVoteRequest
	(*UpVoteResponse)(nil),       // 10: proto.UpVoteResponse
	(*DownVoteRequest)(nil),      // 11: proto.DownVoteRequest
	(*DownVoteResponse)(nil),     // 12: proto.DownVoteResponse
}
var file_proto_crypto_proto_depIdxs = []int32{
	0,  // 0: proto.CreateCryptoRequest.crypto:type_name -> proto.Crypto
	0,  // 1: proto.CreateCryptoResponse.crypto:type_name -> proto.Crypto
	0,  // 2: proto.GetCryptoResponse.crypto:type_name -> proto.Crypto
	0,  // 3: proto.UpdateCryptoResponse.crypto:type_name -> proto.Crypto
	0,  // 4: proto.DeleteCryptoResponse.crypto:type_name -> proto.Crypto
	0,  // 5: proto.UpVoteResponse.crypto:type_name -> proto.Crypto
	0,  // 6: proto.DownVoteResponse.crypto:type_name -> proto.Crypto
	1,  // 7: proto.CryptoService.CreateCrypto:input_type -> proto.CreateCryptoRequest
	3,  // 8: proto.CryptoService.GetCrypto:input_type -> proto.GetCryptoRequest
	5,  // 9: proto.CryptoService.UpdateCrypto:input_type -> proto.UpdateCryptoRequest
	7,  // 10: proto.CryptoService.DeleteCrypto:input_type -> proto.DeleteCryptoRequest
	9,  // 11: proto.CryptoService.UpVote:input_type -> proto.UpVoteRequest
	11, // 12: proto.CryptoService.DownVote:input_type -> proto.DownVoteRequest
	2,  // 13: proto.CryptoService.CreateCrypto:output_type -> proto.CreateCryptoResponse
	4,  // 14: proto.CryptoService.GetCrypto:output_type -> proto.GetCryptoResponse
	6,  // 15: proto.CryptoService.UpdateCrypto:output_type -> proto.UpdateCryptoResponse
	8,  // 16: proto.CryptoService.DeleteCrypto:output_type -> proto.DeleteCryptoResponse
	10, // 17: proto.CryptoService.UpVote:output_type -> proto.UpVoteResponse
	12, // 18: proto.CryptoService.DownVote:output_type -> proto.DownVoteResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_crypto_proto_init() }
func file_proto_crypto_proto_init() {
	if File_proto_crypto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_crypto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crypto); i {
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
		file_proto_crypto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCryptoRequest); i {
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
		file_proto_crypto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCryptoResponse); i {
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
		file_proto_crypto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCryptoRequest); i {
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
		file_proto_crypto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCryptoResponse); i {
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
		file_proto_crypto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCryptoRequest); i {
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
		file_proto_crypto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCryptoResponse); i {
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
		file_proto_crypto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCryptoRequest); i {
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
		file_proto_crypto_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCryptoResponse); i {
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
		file_proto_crypto_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpVoteRequest); i {
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
		file_proto_crypto_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpVoteResponse); i {
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
		file_proto_crypto_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownVoteRequest); i {
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
		file_proto_crypto_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownVoteResponse); i {
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
			RawDescriptor: file_proto_crypto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_crypto_proto_goTypes,
		DependencyIndexes: file_proto_crypto_proto_depIdxs,
		MessageInfos:      file_proto_crypto_proto_msgTypes,
	}.Build()
	File_proto_crypto_proto = out.File
	file_proto_crypto_proto_rawDesc = nil
	file_proto_crypto_proto_goTypes = nil
	file_proto_crypto_proto_depIdxs = nil
}
