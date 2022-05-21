// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/tailnets.proto

package api

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

type Tailnet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tailnet) Reset() {
	*x = Tailnet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tailnets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tailnet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tailnet) ProtoMessage() {}

func (x *Tailnet) ProtoReflect() protoreflect.Message {
	mi := &file_api_tailnets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tailnet.ProtoReflect.Descriptor instead.
func (*Tailnet) Descriptor() ([]byte, []int) {
	return file_api_tailnets_proto_rawDescGZIP(), []int{0}
}

func (x *Tailnet) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tailnet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateTailnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateTailnetRequest) Reset() {
	*x = CreateTailnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tailnets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTailnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTailnetRequest) ProtoMessage() {}

func (x *CreateTailnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tailnets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTailnetRequest.ProtoReflect.Descriptor instead.
func (*CreateTailnetRequest) Descriptor() ([]byte, []int) {
	return file_api_tailnets_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTailnetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateTailnetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tailnet *Tailnet `protobuf:"bytes,1,opt,name=tailnet,proto3" json:"tailnet,omitempty"`
}

func (x *CreateTailnetResponse) Reset() {
	*x = CreateTailnetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tailnets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTailnetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTailnetResponse) ProtoMessage() {}

func (x *CreateTailnetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tailnets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTailnetResponse.ProtoReflect.Descriptor instead.
func (*CreateTailnetResponse) Descriptor() ([]byte, []int) {
	return file_api_tailnets_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTailnetResponse) GetTailnet() *Tailnet {
	if x != nil {
		return x.Tailnet
	}
	return nil
}

type GetTailnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTailnetRequest) Reset() {
	*x = GetTailnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tailnets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTailnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTailnetRequest) ProtoMessage() {}

func (x *GetTailnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tailnets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTailnetRequest.ProtoReflect.Descriptor instead.
func (*GetTailnetRequest) Descriptor() ([]byte, []int) {
	return file_api_tailnets_proto_rawDescGZIP(), []int{3}
}

func (x *GetTailnetRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTailnetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tailnet *Tailnet `protobuf:"bytes,1,opt,name=tailnet,proto3" json:"tailnet,omitempty"`
}

func (x *GetTailnetResponse) Reset() {
	*x = GetTailnetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tailnets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTailnetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTailnetResponse) ProtoMessage() {}

func (x *GetTailnetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tailnets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTailnetResponse.ProtoReflect.Descriptor instead.
func (*GetTailnetResponse) Descriptor() ([]byte, []int) {
	return file_api_tailnets_proto_rawDescGZIP(), []int{4}
}

func (x *GetTailnetResponse) GetTailnet() *Tailnet {
	if x != nil {
		return x.Tailnet
	}
	return nil
}

type ListTailnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTailnetRequest) Reset() {
	*x = ListTailnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tailnets_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTailnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTailnetRequest) ProtoMessage() {}

func (x *ListTailnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tailnets_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTailnetRequest.ProtoReflect.Descriptor instead.
func (*ListTailnetRequest) Descriptor() ([]byte, []int) {
	return file_api_tailnets_proto_rawDescGZIP(), []int{5}
}

type ListTailnetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tailnet []*Tailnet `protobuf:"bytes,1,rep,name=tailnet,proto3" json:"tailnet,omitempty"`
}

func (x *ListTailnetResponse) Reset() {
	*x = ListTailnetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tailnets_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTailnetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTailnetResponse) ProtoMessage() {}

func (x *ListTailnetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tailnets_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTailnetResponse.ProtoReflect.Descriptor instead.
func (*ListTailnetResponse) Descriptor() ([]byte, []int) {
	return file_api_tailnets_proto_rawDescGZIP(), []int{6}
}

func (x *ListTailnetResponse) GetTailnet() []*Tailnet {
	if x != nil {
		return x.Tailnet
	}
	return nil
}

type DeleteTailnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TailnetId uint64 `protobuf:"varint,1,opt,name=tailnet_id,json=tailnetId,proto3" json:"tailnet_id,omitempty"`
	Force     bool   `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *DeleteTailnetRequest) Reset() {
	*x = DeleteTailnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tailnets_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTailnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTailnetRequest) ProtoMessage() {}

func (x *DeleteTailnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tailnets_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTailnetRequest.ProtoReflect.Descriptor instead.
func (*DeleteTailnetRequest) Descriptor() ([]byte, []int) {
	return file_api_tailnets_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTailnetRequest) GetTailnetId() uint64 {
	if x != nil {
		return x.TailnetId
	}
	return 0
}

func (x *DeleteTailnetRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type DeleteTailnetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTailnetResponse) Reset() {
	*x = DeleteTailnetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tailnets_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTailnetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTailnetResponse) ProtoMessage() {}

func (x *DeleteTailnetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tailnets_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTailnetResponse.ProtoReflect.Descriptor instead.
func (*DeleteTailnetResponse) Descriptor() ([]byte, []int) {
	return file_api_tailnets_proto_rawDescGZIP(), []int{8}
}

var File_api_tailnets_proto protoreflect.FileDescriptor

var file_api_tailnets_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x2d, 0x0a, 0x07, 0x54, 0x61, 0x69,
	0x6c, 0x6e, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x07, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61,
	0x69, 0x6c, 0x6e, 0x65, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x61, 0x69, 0x6c,
	0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x07, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52,
	0x07, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x69,
	0x6c, 0x6e, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x22, 0x4b, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x61, 0x69, 0x6c, 0x6e,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6a, 0x73, 0x69, 0x65, 0x62, 0x65, 0x6e, 0x73, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_tailnets_proto_rawDescOnce sync.Once
	file_api_tailnets_proto_rawDescData = file_api_tailnets_proto_rawDesc
)

func file_api_tailnets_proto_rawDescGZIP() []byte {
	file_api_tailnets_proto_rawDescOnce.Do(func() {
		file_api_tailnets_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_tailnets_proto_rawDescData)
	})
	return file_api_tailnets_proto_rawDescData
}

var file_api_tailnets_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_tailnets_proto_goTypes = []interface{}{
	(*Tailnet)(nil),               // 0: api.Tailnet
	(*CreateTailnetRequest)(nil),  // 1: api.CreateTailnetRequest
	(*CreateTailnetResponse)(nil), // 2: api.CreateTailnetResponse
	(*GetTailnetRequest)(nil),     // 3: api.GetTailnetRequest
	(*GetTailnetResponse)(nil),    // 4: api.GetTailnetResponse
	(*ListTailnetRequest)(nil),    // 5: api.ListTailnetRequest
	(*ListTailnetResponse)(nil),   // 6: api.ListTailnetResponse
	(*DeleteTailnetRequest)(nil),  // 7: api.DeleteTailnetRequest
	(*DeleteTailnetResponse)(nil), // 8: api.DeleteTailnetResponse
}
var file_api_tailnets_proto_depIdxs = []int32{
	0, // 0: api.CreateTailnetResponse.tailnet:type_name -> api.Tailnet
	0, // 1: api.GetTailnetResponse.tailnet:type_name -> api.Tailnet
	0, // 2: api.ListTailnetResponse.tailnet:type_name -> api.Tailnet
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_tailnets_proto_init() }
func file_api_tailnets_proto_init() {
	if File_api_tailnets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_tailnets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tailnet); i {
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
		file_api_tailnets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTailnetRequest); i {
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
		file_api_tailnets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTailnetResponse); i {
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
		file_api_tailnets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTailnetRequest); i {
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
		file_api_tailnets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTailnetResponse); i {
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
		file_api_tailnets_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTailnetRequest); i {
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
		file_api_tailnets_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTailnetResponse); i {
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
		file_api_tailnets_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTailnetRequest); i {
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
		file_api_tailnets_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTailnetResponse); i {
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
			RawDescriptor: file_api_tailnets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_tailnets_proto_goTypes,
		DependencyIndexes: file_api_tailnets_proto_depIdxs,
		MessageInfos:      file_api_tailnets_proto_msgTypes,
	}.Build()
	File_api_tailnets_proto = out.File
	file_api_tailnets_proto_rawDesc = nil
	file_api_tailnets_proto_goTypes = nil
	file_api_tailnets_proto_depIdxs = nil
}
