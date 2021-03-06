// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.5.0
// source: user.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	common_pb "go_project_template/pb/gen-go/common_pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetUserNameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseReq *common_pb.BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq,proto3" json:"base_req,omitempty"`
}

func (x *GetUserNameReq) Reset() {
	*x = GetUserNameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserNameReq) ProtoMessage() {}

func (x *GetUserNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserNameReq.ProtoReflect.Descriptor instead.
func (*GetUserNameReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserNameReq) GetBaseReq() *common_pb.BaseReq {
	if x != nil {
		return x.BaseReq
	}
	return nil
}

type GetUserNameRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetUserNameRsp) Reset() {
	*x = GetUserNameRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserNameRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserNameRsp) ProtoMessage() {}

func (x *GetUserNameRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserNameRsp.ProtoReflect.Descriptor instead.
func (*GetUserNameRsp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserNameRsp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateUserNameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseReq *common_pb.BaseReq `protobuf:"bytes,1,opt,name=baseReq,proto3" json:"baseReq,omitempty"`
	NewName string             `protobuf:"bytes,2,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
}

func (x *UpdateUserNameReq) Reset() {
	*x = UpdateUserNameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserNameReq) ProtoMessage() {}

func (x *UpdateUserNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserNameReq.ProtoReflect.Descriptor instead.
func (*UpdateUserNameReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserNameReq) GetBaseReq() *common_pb.BaseReq {
	if x != nil {
		return x.BaseReq
	}
	return nil
}

func (x *UpdateUserNameReq) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type UpdateUserNameRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUserNameRsp) Reset() {
	*x = UpdateUserNameRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserNameRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserNameRsp) ProtoMessage() {}

func (x *UpdateUserNameRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserNameRsp.ProtoReflect.Descriptor instead.
func (*UpdateUserNameRsp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x08, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a,
	0x07, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x52, 0x07, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x73, 0x70, 0x32, 0x84, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x53, 0x76, 0x63,
	0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_proto_goTypes = []interface{}{
	(*GetUserNameReq)(nil),    // 0: pb.GetUserNameReq
	(*GetUserNameRsp)(nil),    // 1: pb.GetUserNameRsp
	(*UpdateUserNameReq)(nil), // 2: pb.UpdateUserNameReq
	(*UpdateUserNameRsp)(nil), // 3: pb.UpdateUserNameRsp
	(*common_pb.BaseReq)(nil), // 4: pb.BaseReq
}
var file_user_proto_depIdxs = []int32{
	4, // 0: pb.GetUserNameReq.base_req:type_name -> pb.BaseReq
	4, // 1: pb.UpdateUserNameReq.baseReq:type_name -> pb.BaseReq
	0, // 2: pb.UserSvc.GetUserName:input_type -> pb.GetUserNameReq
	2, // 3: pb.UserSvc.UpdateUserName:input_type -> pb.UpdateUserNameReq
	1, // 4: pb.UserSvc.GetUserName:output_type -> pb.GetUserNameRsp
	3, // 5: pb.UserSvc.UpdateUserName:output_type -> pb.UpdateUserNameRsp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserNameReq); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserNameRsp); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserNameReq); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserNameRsp); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserSvcClient is the client API for UserSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSvcClient interface {
	GetUserName(ctx context.Context, in *GetUserNameReq, opts ...grpc.CallOption) (*GetUserNameRsp, error)
	UpdateUserName(ctx context.Context, in *UpdateUserNameReq, opts ...grpc.CallOption) (*UpdateUserNameRsp, error)
}

type userSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSvcClient(cc grpc.ClientConnInterface) UserSvcClient {
	return &userSvcClient{cc}
}

func (c *userSvcClient) GetUserName(ctx context.Context, in *GetUserNameReq, opts ...grpc.CallOption) (*GetUserNameRsp, error) {
	out := new(GetUserNameRsp)
	err := c.cc.Invoke(ctx, "/pb.UserSvc/GetUserName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) UpdateUserName(ctx context.Context, in *UpdateUserNameReq, opts ...grpc.CallOption) (*UpdateUserNameRsp, error) {
	out := new(UpdateUserNameRsp)
	err := c.cc.Invoke(ctx, "/pb.UserSvc/UpdateUserName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSvcServer is the server API for UserSvc service.
type UserSvcServer interface {
	GetUserName(context.Context, *GetUserNameReq) (*GetUserNameRsp, error)
	UpdateUserName(context.Context, *UpdateUserNameReq) (*UpdateUserNameRsp, error)
}

// UnimplementedUserSvcServer can be embedded to have forward compatible implementations.
type UnimplementedUserSvcServer struct {
}

func (*UnimplementedUserSvcServer) GetUserName(context.Context, *GetUserNameReq) (*GetUserNameRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserName not implemented")
}
func (*UnimplementedUserSvcServer) UpdateUserName(context.Context, *UpdateUserNameReq) (*UpdateUserNameRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserName not implemented")
}

func RegisterUserSvcServer(s *grpc.Server, srv UserSvcServer) {
	s.RegisterService(&_UserSvc_serviceDesc, srv)
}

func _UserSvc_GetUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).GetUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserSvc/GetUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).GetUserName(ctx, req.(*GetUserNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_UpdateUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).UpdateUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserSvc/UpdateUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).UpdateUserName(ctx, req.(*UpdateUserNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserSvc",
	HandlerType: (*UserSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserName",
			Handler:    _UserSvc_GetUserName_Handler,
		},
		{
			MethodName: "UpdateUserName",
			Handler:    _UserSvc_UpdateUserName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
