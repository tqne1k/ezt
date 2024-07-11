// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: grpc/model/ezt.proto

package model

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

type GetTunnelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *GetTunnelRequest) Reset() {
	*x = GetTunnelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_ezt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTunnelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTunnelRequest) ProtoMessage() {}

func (x *GetTunnelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_ezt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTunnelRequest.ProtoReflect.Descriptor instead.
func (*GetTunnelRequest) Descriptor() ([]byte, []int) {
	return file_grpc_model_ezt_proto_rawDescGZIP(), []int{0}
}

func (x *GetTunnelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetTunnelRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type GetTunnelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PublicKey     string      `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	ListeningPort string      `protobuf:"bytes,3,opt,name=listening_port,json=listeningPort,proto3" json:"listening_port,omitempty"`
	PeerInfo      []*PeerInfo `protobuf:"bytes,4,rep,name=peer_info,json=peerInfo,proto3" json:"peer_info,omitempty"`
}

func (x *GetTunnelResponse) Reset() {
	*x = GetTunnelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_ezt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTunnelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTunnelResponse) ProtoMessage() {}

func (x *GetTunnelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_ezt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTunnelResponse.ProtoReflect.Descriptor instead.
func (*GetTunnelResponse) Descriptor() ([]byte, []int) {
	return file_grpc_model_ezt_proto_rawDescGZIP(), []int{1}
}

func (x *GetTunnelResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetTunnelResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *GetTunnelResponse) GetListeningPort() string {
	if x != nil {
		return x.ListeningPort
	}
	return ""
}

func (x *GetTunnelResponse) GetPeerInfo() []*PeerInfo {
	if x != nil {
		return x.PeerInfo
	}
	return nil
}

type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey           string `protobuf:"bytes,6,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Endpoint            string `protobuf:"bytes,7,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	AllowedIps          string `protobuf:"bytes,8,opt,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
	LatestHandshake     string `protobuf:"bytes,9,opt,name=latest_handshake,json=latestHandshake,proto3" json:"latest_handshake,omitempty"`
	Transfer            string `protobuf:"bytes,10,opt,name=transfer,proto3" json:"transfer,omitempty"`
	PersistentKeepalive string `protobuf:"bytes,11,opt,name=persistent_keepalive,json=persistentKeepalive,proto3" json:"persistent_keepalive,omitempty"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_ezt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_ezt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_grpc_model_ezt_proto_rawDescGZIP(), []int{2}
}

func (x *PeerInfo) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *PeerInfo) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *PeerInfo) GetAllowedIps() string {
	if x != nil {
		return x.AllowedIps
	}
	return ""
}

func (x *PeerInfo) GetLatestHandshake() string {
	if x != nil {
		return x.LatestHandshake
	}
	return ""
}

func (x *PeerInfo) GetTransfer() string {
	if x != nil {
		return x.Transfer
	}
	return ""
}

func (x *PeerInfo) GetPersistentKeepalive() string {
	if x != nil {
		return x.PersistentKeepalive
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_ezt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_ezt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_grpc_model_ezt_proto_rawDescGZIP(), []int{3}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_ezt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_ezt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_grpc_model_ezt_proto_rawDescGZIP(), []int{4}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_model_ezt_proto protoreflect.FileDescriptor

var file_grpc_model_ezt_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x7a, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x45, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x22, 0x9a, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0xe0, 0x01, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x70, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x73,
	0x68, 0x61, 0x6b, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x12, 0x31, 0x0a, 0x14, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6b,
	0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x3e, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x2f,
	0x0a, 0x05, 0x53, 0x61, 0x79, 0x48, 0x69, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32,
	0x50, 0x0a, 0x0a, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x14, 0x5a, 0x12, 0x65, 0x7a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_model_ezt_proto_rawDescOnce sync.Once
	file_grpc_model_ezt_proto_rawDescData = file_grpc_model_ezt_proto_rawDesc
)

func file_grpc_model_ezt_proto_rawDescGZIP() []byte {
	file_grpc_model_ezt_proto_rawDescOnce.Do(func() {
		file_grpc_model_ezt_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_model_ezt_proto_rawDescData)
	})
	return file_grpc_model_ezt_proto_rawDescData
}

var file_grpc_model_ezt_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_model_ezt_proto_goTypes = []interface{}{
	(*GetTunnelRequest)(nil),  // 0: grpc.GetTunnelRequest
	(*GetTunnelResponse)(nil), // 1: grpc.GetTunnelResponse
	(*PeerInfo)(nil),          // 2: grpc.PeerInfo
	(*HelloRequest)(nil),      // 3: grpc.HelloRequest
	(*HelloReply)(nil),        // 4: grpc.HelloReply
}
var file_grpc_model_ezt_proto_depIdxs = []int32{
	2, // 0: grpc.GetTunnelResponse.peer_info:type_name -> grpc.PeerInfo
	3, // 1: grpc.HealthCheck.SayHi:input_type -> grpc.HelloRequest
	0, // 2: grpc.TunnelInfo.GetTunnelInfo:input_type -> grpc.GetTunnelRequest
	4, // 3: grpc.HealthCheck.SayHi:output_type -> grpc.HelloReply
	1, // 4: grpc.TunnelInfo.GetTunnelInfo:output_type -> grpc.GetTunnelResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_model_ezt_proto_init() }
func file_grpc_model_ezt_proto_init() {
	if File_grpc_model_ezt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_model_ezt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTunnelRequest); i {
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
		file_grpc_model_ezt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTunnelResponse); i {
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
		file_grpc_model_ezt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
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
		file_grpc_model_ezt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_grpc_model_ezt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_grpc_model_ezt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_grpc_model_ezt_proto_goTypes,
		DependencyIndexes: file_grpc_model_ezt_proto_depIdxs,
		MessageInfos:      file_grpc_model_ezt_proto_msgTypes,
	}.Build()
	File_grpc_model_ezt_proto = out.File
	file_grpc_model_ezt_proto_rawDesc = nil
	file_grpc_model_ezt_proto_goTypes = nil
	file_grpc_model_ezt_proto_depIdxs = nil
}