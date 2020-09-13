// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: grpcinventoryserver.proto

package grpcinventoryclient

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QuantityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string `protobuf:"bytes,1,opt,name=productID,proto3" json:"productID,omitempty"`
}

func (x *QuantityRequest) Reset() {
	*x = QuantityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcinventoryserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuantityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuantityRequest) ProtoMessage() {}

func (x *QuantityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcinventoryserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuantityRequest.ProtoReflect.Descriptor instead.
func (*QuantityRequest) Descriptor() ([]byte, []int) {
	return file_grpcinventoryserver_proto_rawDescGZIP(), []int{0}
}

func (x *QuantityRequest) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

type QuantityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string `protobuf:"bytes,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *QuantityResponse) Reset() {
	*x = QuantityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcinventoryserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuantityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuantityResponse) ProtoMessage() {}

func (x *QuantityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcinventoryserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuantityResponse.ProtoReflect.Descriptor instead.
func (*QuantityResponse) Descriptor() ([]byte, []int) {
	return file_grpcinventoryserver_proto_rawDescGZIP(), []int{1}
}

func (x *QuantityResponse) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *QuantityResponse) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string `protobuf:"bytes,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ReservationRequest) Reset() {
	*x = ReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcinventoryserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationRequest) ProtoMessage() {}

func (x *ReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcinventoryserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationRequest.ProtoReflect.Descriptor instead.
func (*ReservationRequest) Descriptor() ([]byte, []int) {
	return file_grpcinventoryserver_proto_rawDescGZIP(), []int{2}
}

func (x *ReservationRequest) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *ReservationRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationID string `protobuf:"bytes,1,opt,name=reservationID,proto3" json:"reservationID,omitempty"`
}

func (x *ReservationResponse) Reset() {
	*x = ReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcinventoryserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationResponse) ProtoMessage() {}

func (x *ReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcinventoryserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationResponse.ProtoReflect.Descriptor instead.
func (*ReservationResponse) Descriptor() ([]byte, []int) {
	return file_grpcinventoryserver_proto_rawDescGZIP(), []int{3}
}

func (x *ReservationResponse) GetReservationID() string {
	if x != nil {
		return x.ReservationID
	}
	return ""
}

type RollBackReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationID string `protobuf:"bytes,1,opt,name=reservationID,proto3" json:"reservationID,omitempty"`
}

func (x *RollBackReservationRequest) Reset() {
	*x = RollBackReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcinventoryserver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollBackReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollBackReservationRequest) ProtoMessage() {}

func (x *RollBackReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcinventoryserver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollBackReservationRequest.ProtoReflect.Descriptor instead.
func (*RollBackReservationRequest) Descriptor() ([]byte, []int) {
	return file_grpcinventoryserver_proto_rawDescGZIP(), []int{4}
}

func (x *RollBackReservationRequest) GetReservationID() string {
	if x != nil {
		return x.ReservationID
	}
	return ""
}

type RollBackReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RollBackReservationResponse) Reset() {
	*x = RollBackReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcinventoryserver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollBackReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollBackReservationResponse) ProtoMessage() {}

func (x *RollBackReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcinventoryserver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollBackReservationResponse.ProtoReflect.Descriptor instead.
func (*RollBackReservationResponse) Descriptor() ([]byte, []int) {
	return file_grpcinventoryserver_proto_rawDescGZIP(), []int{5}
}

func (x *RollBackReservationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CommitReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationID string `protobuf:"bytes,1,opt,name=reservationID,proto3" json:"reservationID,omitempty"`
}

func (x *CommitReservationRequest) Reset() {
	*x = CommitReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcinventoryserver_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitReservationRequest) ProtoMessage() {}

func (x *CommitReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcinventoryserver_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitReservationRequest.ProtoReflect.Descriptor instead.
func (*CommitReservationRequest) Descriptor() ([]byte, []int) {
	return file_grpcinventoryserver_proto_rawDescGZIP(), []int{6}
}

func (x *CommitReservationRequest) GetReservationID() string {
	if x != nil {
		return x.ReservationID
	}
	return ""
}

type CommitReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CommitReservationResponse) Reset() {
	*x = CommitReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcinventoryserver_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitReservationResponse) ProtoMessage() {}

func (x *CommitReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcinventoryserver_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitReservationResponse.ProtoReflect.Descriptor instead.
func (*CommitReservationResponse) Descriptor() ([]byte, []int) {
	return file_grpcinventoryserver_proto_rawDescGZIP(), []int{7}
}

func (x *CommitReservationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_grpcinventoryserver_proto protoreflect.FileDescriptor

var file_grpcinventoryserver_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x72, 0x70,
	0x63, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x22, 0x2f, 0x0a, 0x0f, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x44, 0x22, 0x4c, 0x0a, 0x10, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x4e, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x3b, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x42, 0x0a, 0x1a,
	0x52, 0x6f, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x22, 0x37, 0x0a, 0x1b, 0x52, 0x6f, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x40, 0x0a, 0x18, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x35, 0x0a, 0x19, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x32, 0xd6, 0x03, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7a, 0x0a, 0x13, 0x52, 0x6f, 0x6c,
	0x6c, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x30, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_grpcinventoryserver_proto_rawDescOnce sync.Once
	file_grpcinventoryserver_proto_rawDescData = file_grpcinventoryserver_proto_rawDesc
)

func file_grpcinventoryserver_proto_rawDescGZIP() []byte {
	file_grpcinventoryserver_proto_rawDescOnce.Do(func() {
		file_grpcinventoryserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcinventoryserver_proto_rawDescData)
	})
	return file_grpcinventoryserver_proto_rawDescData
}

var file_grpcinventoryserver_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpcinventoryserver_proto_goTypes = []interface{}{
	(*QuantityRequest)(nil),             // 0: grpcinventoryserver.QuantityRequest
	(*QuantityResponse)(nil),            // 1: grpcinventoryserver.QuantityResponse
	(*ReservationRequest)(nil),          // 2: grpcinventoryserver.ReservationRequest
	(*ReservationResponse)(nil),         // 3: grpcinventoryserver.ReservationResponse
	(*RollBackReservationRequest)(nil),  // 4: grpcinventoryserver.RollBackReservationRequest
	(*RollBackReservationResponse)(nil), // 5: grpcinventoryserver.RollBackReservationResponse
	(*CommitReservationRequest)(nil),    // 6: grpcinventoryserver.CommitReservationRequest
	(*CommitReservationResponse)(nil),   // 7: grpcinventoryserver.CommitReservationResponse
}
var file_grpcinventoryserver_proto_depIdxs = []int32{
	0, // 0: grpcinventoryserver.InventoryService.GetAvailableQuantity:input_type -> grpcinventoryserver.QuantityRequest
	2, // 1: grpcinventoryserver.InventoryService.GetResevationToken:input_type -> grpcinventoryserver.ReservationRequest
	4, // 2: grpcinventoryserver.InventoryService.RollBackReservation:input_type -> grpcinventoryserver.RollBackReservationRequest
	6, // 3: grpcinventoryserver.InventoryService.CommitReservation:input_type -> grpcinventoryserver.CommitReservationRequest
	1, // 4: grpcinventoryserver.InventoryService.GetAvailableQuantity:output_type -> grpcinventoryserver.QuantityResponse
	3, // 5: grpcinventoryserver.InventoryService.GetResevationToken:output_type -> grpcinventoryserver.ReservationResponse
	5, // 6: grpcinventoryserver.InventoryService.RollBackReservation:output_type -> grpcinventoryserver.RollBackReservationResponse
	7, // 7: grpcinventoryserver.InventoryService.CommitReservation:output_type -> grpcinventoryserver.CommitReservationResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcinventoryserver_proto_init() }
func file_grpcinventoryserver_proto_init() {
	if File_grpcinventoryserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcinventoryserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuantityRequest); i {
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
		file_grpcinventoryserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuantityResponse); i {
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
		file_grpcinventoryserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationRequest); i {
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
		file_grpcinventoryserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationResponse); i {
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
		file_grpcinventoryserver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollBackReservationRequest); i {
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
		file_grpcinventoryserver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollBackReservationResponse); i {
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
		file_grpcinventoryserver_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitReservationRequest); i {
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
		file_grpcinventoryserver_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitReservationResponse); i {
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
			RawDescriptor: file_grpcinventoryserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcinventoryserver_proto_goTypes,
		DependencyIndexes: file_grpcinventoryserver_proto_depIdxs,
		MessageInfos:      file_grpcinventoryserver_proto_msgTypes,
	}.Build()
	File_grpcinventoryserver_proto = out.File
	file_grpcinventoryserver_proto_rawDesc = nil
	file_grpcinventoryserver_proto_goTypes = nil
	file_grpcinventoryserver_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InventoryServiceClient interface {
	GetAvailableQuantity(ctx context.Context, in *QuantityRequest, opts ...grpc.CallOption) (*QuantityResponse, error)
	GetResevationToken(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error)
	RollBackReservation(ctx context.Context, in *RollBackReservationRequest, opts ...grpc.CallOption) (*RollBackReservationResponse, error)
	CommitReservation(ctx context.Context, in *CommitReservationRequest, opts ...grpc.CallOption) (*CommitReservationResponse, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) GetAvailableQuantity(ctx context.Context, in *QuantityRequest, opts ...grpc.CallOption) (*QuantityResponse, error) {
	out := new(QuantityResponse)
	err := c.cc.Invoke(ctx, "/grpcinventoryserver.InventoryService/GetAvailableQuantity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetResevationToken(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, "/grpcinventoryserver.InventoryService/GetResevationToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) RollBackReservation(ctx context.Context, in *RollBackReservationRequest, opts ...grpc.CallOption) (*RollBackReservationResponse, error) {
	out := new(RollBackReservationResponse)
	err := c.cc.Invoke(ctx, "/grpcinventoryserver.InventoryService/RollBackReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) CommitReservation(ctx context.Context, in *CommitReservationRequest, opts ...grpc.CallOption) (*CommitReservationResponse, error) {
	out := new(CommitReservationResponse)
	err := c.cc.Invoke(ctx, "/grpcinventoryserver.InventoryService/CommitReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
type InventoryServiceServer interface {
	GetAvailableQuantity(context.Context, *QuantityRequest) (*QuantityResponse, error)
	GetResevationToken(context.Context, *ReservationRequest) (*ReservationResponse, error)
	RollBackReservation(context.Context, *RollBackReservationRequest) (*RollBackReservationResponse, error)
	CommitReservation(context.Context, *CommitReservationRequest) (*CommitReservationResponse, error)
}

// UnimplementedInventoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInventoryServiceServer struct {
}

func (*UnimplementedInventoryServiceServer) GetAvailableQuantity(context.Context, *QuantityRequest) (*QuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableQuantity not implemented")
}
func (*UnimplementedInventoryServiceServer) GetResevationToken(context.Context, *ReservationRequest) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResevationToken not implemented")
}
func (*UnimplementedInventoryServiceServer) RollBackReservation(context.Context, *RollBackReservationRequest) (*RollBackReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollBackReservation not implemented")
}
func (*UnimplementedInventoryServiceServer) CommitReservation(context.Context, *CommitReservationRequest) (*CommitReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitReservation not implemented")
}

func RegisterInventoryServiceServer(s *grpc.Server, srv InventoryServiceServer) {
	s.RegisterService(&_InventoryService_serviceDesc, srv)
}

func _InventoryService_GetAvailableQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuantityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetAvailableQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcinventoryserver.InventoryService/GetAvailableQuantity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetAvailableQuantity(ctx, req.(*QuantityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetResevationToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetResevationToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcinventoryserver.InventoryService/GetResevationToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetResevationToken(ctx, req.(*ReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_RollBackReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollBackReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).RollBackReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcinventoryserver.InventoryService/RollBackReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).RollBackReservation(ctx, req.(*RollBackReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_CommitReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).CommitReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcinventoryserver.InventoryService/CommitReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).CommitReservation(ctx, req.(*CommitReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InventoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcinventoryserver.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvailableQuantity",
			Handler:    _InventoryService_GetAvailableQuantity_Handler,
		},
		{
			MethodName: "GetResevationToken",
			Handler:    _InventoryService_GetResevationToken_Handler,
		},
		{
			MethodName: "RollBackReservation",
			Handler:    _InventoryService_RollBackReservation_Handler,
		},
		{
			MethodName: "CommitReservation",
			Handler:    _InventoryService_CommitReservation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcinventoryserver.proto",
}
