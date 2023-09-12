// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetaServiceClient is the client API for MetaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetaServiceClient interface {
	List(ctx context.Context, in *MetaListRequest, opts ...grpc.CallOption) (*MetaListResponse, error)
	Get(ctx context.Context, in *MetaGetRequest, opts ...grpc.CallOption) (*MetaGetResponse, error)
	SetFavorite(ctx context.Context, in *SetFavoriteRequest, opts ...grpc.CallOption) (*SetFavoriteResponse, error)
	GetPage(ctx context.Context, in *GetPageRequest, opts ...grpc.CallOption) (*GetPageResponse, error)
}

type metaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaServiceClient(cc grpc.ClientConnInterface) MetaServiceClient {
	return &metaServiceClient{cc}
}

func (c *metaServiceClient) List(ctx context.Context, in *MetaListRequest, opts ...grpc.CallOption) (*MetaListResponse, error) {
	out := new(MetaListResponse)
	err := c.cc.Invoke(ctx, "/MetaService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) Get(ctx context.Context, in *MetaGetRequest, opts ...grpc.CallOption) (*MetaGetResponse, error) {
	out := new(MetaGetResponse)
	err := c.cc.Invoke(ctx, "/MetaService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) SetFavorite(ctx context.Context, in *SetFavoriteRequest, opts ...grpc.CallOption) (*SetFavoriteResponse, error) {
	out := new(SetFavoriteResponse)
	err := c.cc.Invoke(ctx, "/MetaService/SetFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) GetPage(ctx context.Context, in *GetPageRequest, opts ...grpc.CallOption) (*GetPageResponse, error) {
	out := new(GetPageResponse)
	err := c.cc.Invoke(ctx, "/MetaService/GetPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaServiceServer is the server API for MetaService service.
// All implementations must embed UnimplementedMetaServiceServer
// for forward compatibility
type MetaServiceServer interface {
	List(context.Context, *MetaListRequest) (*MetaListResponse, error)
	Get(context.Context, *MetaGetRequest) (*MetaGetResponse, error)
	SetFavorite(context.Context, *SetFavoriteRequest) (*SetFavoriteResponse, error)
	GetPage(context.Context, *GetPageRequest) (*GetPageResponse, error)
	mustEmbedUnimplementedMetaServiceServer()
}

// UnimplementedMetaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetaServiceServer struct {
}

func (UnimplementedMetaServiceServer) List(context.Context, *MetaListRequest) (*MetaListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMetaServiceServer) Get(context.Context, *MetaGetRequest) (*MetaGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMetaServiceServer) SetFavorite(context.Context, *SetFavoriteRequest) (*SetFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFavorite not implemented")
}
func (UnimplementedMetaServiceServer) GetPage(context.Context, *GetPageRequest) (*GetPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPage not implemented")
}
func (UnimplementedMetaServiceServer) mustEmbedUnimplementedMetaServiceServer() {}

// UnsafeMetaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetaServiceServer will
// result in compilation errors.
type UnsafeMetaServiceServer interface {
	mustEmbedUnimplementedMetaServiceServer()
}

func RegisterMetaServiceServer(s grpc.ServiceRegistrar, srv MetaServiceServer) {
	s.RegisterService(&MetaService_ServiceDesc, srv)
}

func _MetaService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MetaService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).List(ctx, req.(*MetaListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MetaService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).Get(ctx, req.(*MetaGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_SetFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).SetFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MetaService/SetFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).SetFavorite(ctx, req.(*SetFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_GetPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).GetPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MetaService/GetPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).GetPage(ctx, req.(*GetPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetaService_ServiceDesc is the grpc.ServiceDesc for MetaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MetaService",
	HandlerType: (*MetaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _MetaService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MetaService_Get_Handler,
		},
		{
			MethodName: "SetFavorite",
			Handler:    _MetaService_SetFavorite_Handler,
		},
		{
			MethodName: "GetPage",
			Handler:    _MetaService_GetPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagServiceClient interface {
	List(ctx context.Context, in *TagListRequest, opts ...grpc.CallOption) (*TagListResponse, error)
	SetFavorite(ctx context.Context, in *SetFavoriteRequest, opts ...grpc.CallOption) (*SetFavoriteResponse, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) List(ctx context.Context, in *TagListRequest, opts ...grpc.CallOption) (*TagListResponse, error) {
	out := new(TagListResponse)
	err := c.cc.Invoke(ctx, "/TagService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) SetFavorite(ctx context.Context, in *SetFavoriteRequest, opts ...grpc.CallOption) (*SetFavoriteResponse, error) {
	out := new(SetFavoriteResponse)
	err := c.cc.Invoke(ctx, "/TagService/SetFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
// All implementations must embed UnimplementedTagServiceServer
// for forward compatibility
type TagServiceServer interface {
	List(context.Context, *TagListRequest) (*TagListResponse, error)
	SetFavorite(context.Context, *SetFavoriteRequest) (*SetFavoriteResponse, error)
	mustEmbedUnimplementedTagServiceServer()
}

// UnimplementedTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (UnimplementedTagServiceServer) List(context.Context, *TagListRequest) (*TagListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTagServiceServer) SetFavorite(context.Context, *SetFavoriteRequest) (*SetFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFavorite not implemented")
}
func (UnimplementedTagServiceServer) mustEmbedUnimplementedTagServiceServer() {}

// UnsafeTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServiceServer will
// result in compilation errors.
type UnsafeTagServiceServer interface {
	mustEmbedUnimplementedTagServiceServer()
}

func RegisterTagServiceServer(s grpc.ServiceRegistrar, srv TagServiceServer) {
	s.RegisterService(&TagService_ServiceDesc, srv)
}

func _TagService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TagService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).List(ctx, req.(*TagListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_SetFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).SetFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TagService/SetFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).SetFavorite(ctx, req.(*SetFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagService_ServiceDesc is the grpc.ServiceDesc for TagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TagService_List_Handler,
		},
		{
			MethodName: "SetFavorite",
			Handler:    _TagService_SetFavorite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
