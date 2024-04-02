// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: services.proto

package services

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

type TextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TextRequest) Reset() {
	*x = TextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextRequest) ProtoMessage() {}

func (x *TextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextRequest.ProtoReflect.Descriptor instead.
func (*TextRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{0}
}

func (x *TextRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CapitalizedText string `protobuf:"bytes,1,opt,name=capitalized_text,json=capitalizedText,proto3" json:"capitalized_text,omitempty"`
}

func (x *TextResponse) Reset() {
	*x = TextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextResponse) ProtoMessage() {}

func (x *TextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextResponse.ProtoReflect.Descriptor instead.
func (*TextResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{1}
}

func (x *TextResponse) GetCapitalizedText() string {
	if x != nil {
		return x.CapitalizedText
	}
	return ""
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{2}
}

func (x *HeartbeatRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{3}
}

type ClientToMasterUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientToMasterUploadRequest) Reset() {
	*x = ClientToMasterUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientToMasterUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientToMasterUploadRequest) ProtoMessage() {}

func (x *ClientToMasterUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientToMasterUploadRequest.ProtoReflect.Descriptor instead.
func (*ClientToMasterUploadRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{4}
}

type ClientToMasterUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Port      string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ClientToMasterUploadResponse) Reset() {
	*x = ClientToMasterUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientToMasterUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientToMasterUploadResponse) ProtoMessage() {}

func (x *ClientToMasterUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientToMasterUploadResponse.ProtoReflect.Descriptor instead.
func (*ClientToMasterUploadResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{5}
}

func (x *ClientToMasterUploadResponse) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *ClientToMasterUploadResponse) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type ClientToDataKeeperUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *ClientToDataKeeperUploadRequest) Reset() {
	*x = ClientToDataKeeperUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientToDataKeeperUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientToDataKeeperUploadRequest) ProtoMessage() {}

func (x *ClientToDataKeeperUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientToDataKeeperUploadRequest.ProtoReflect.Descriptor instead.
func (*ClientToDataKeeperUploadRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{6}
}

func (x *ClientToDataKeeperUploadRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type ClientToDataKeeperUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientToDataKeeperUploadResponse) Reset() {
	*x = ClientToDataKeeperUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientToDataKeeperUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientToDataKeeperUploadResponse) ProtoMessage() {}

func (x *ClientToDataKeeperUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientToDataKeeperUploadResponse.ProtoReflect.Descriptor instead.
func (*ClientToDataKeeperUploadResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{7}
}

var File_services_proto protoreflect.FileDescriptor

var file_services_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x0b, 0x54, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x39, 0x0a,
	0x0c, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x22, 0x2b, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x1c, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x3d, 0x0a, 0x1f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x20, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xed, 0x02, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0e,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x36, 0x0a,
	0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54,
	0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x25, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54,
	0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x18,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x61,
	0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1e, 0x5a, 0x1c, 0x77, 0x69, 0x72, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x61, 0x62, 0x5f,
	0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_proto_rawDescOnce sync.Once
	file_services_proto_rawDescData = file_services_proto_rawDesc
)

func file_services_proto_rawDescGZIP() []byte {
	file_services_proto_rawDescOnce.Do(func() {
		file_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_proto_rawDescData)
	})
	return file_services_proto_rawDescData
}

var file_services_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_proto_goTypes = []interface{}{
	(*TextRequest)(nil),                      // 0: services.TextRequest
	(*TextResponse)(nil),                     // 1: services.TextResponse
	(*HeartbeatRequest)(nil),                 // 2: services.HeartbeatRequest
	(*HeartbeatResponse)(nil),                // 3: services.HeartbeatResponse
	(*ClientToMasterUploadRequest)(nil),      // 4: services.clientToMasterUploadRequest
	(*ClientToMasterUploadResponse)(nil),     // 5: services.clientToMasterUploadResponse
	(*ClientToDataKeeperUploadRequest)(nil),  // 6: services.clientToDataKeeperUploadRequest
	(*ClientToDataKeeperUploadResponse)(nil), // 7: services.clientToDataKeeperUploadResponse
}
var file_services_proto_depIdxs = []int32{
	2, // 0: services.Services.TrackHeartbeat:input_type -> services.HeartbeatRequest
	0, // 1: services.Services.Hello:input_type -> services.TextRequest
	4, // 2: services.Services.clientToMasterUpload:input_type -> services.clientToMasterUploadRequest
	6, // 3: services.Services.clientToDataKeeperUpload:input_type -> services.clientToDataKeeperUploadRequest
	3, // 4: services.Services.TrackHeartbeat:output_type -> services.HeartbeatResponse
	1, // 5: services.Services.Hello:output_type -> services.TextResponse
	5, // 6: services.Services.clientToMasterUpload:output_type -> services.clientToMasterUploadResponse
	7, // 7: services.Services.clientToDataKeeperUpload:output_type -> services.clientToDataKeeperUploadResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_proto_init() }
func file_services_proto_init() {
	if File_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextRequest); i {
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
		file_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextResponse); i {
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
		file_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
		file_services_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientToMasterUploadRequest); i {
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
		file_services_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientToMasterUploadResponse); i {
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
		file_services_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientToDataKeeperUploadRequest); i {
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
		file_services_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientToDataKeeperUploadResponse); i {
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
			RawDescriptor: file_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_proto_goTypes,
		DependencyIndexes: file_services_proto_depIdxs,
		MessageInfos:      file_services_proto_msgTypes,
	}.Build()
	File_services_proto = out.File
	file_services_proto_rawDesc = nil
	file_services_proto_goTypes = nil
	file_services_proto_depIdxs = nil
}
