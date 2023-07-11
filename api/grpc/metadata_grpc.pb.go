// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: api/grpc/metadata.proto

package grpcapi

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

// MetadataServiceClient is the client API for MetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	GetStockInfo(ctx context.Context, in *GetStockInfoRequest, opts ...grpc.CallOption) (*GetStockInfoResponse, error)
	GetModelInfo(ctx context.Context, in *GetModelInfoRequest, opts ...grpc.CallOption) (*GetModelInfoResponse, error)
	GetStockList(ctx context.Context, in *GetStockListRequest, opts ...grpc.CallOption) (*GetStockListResponse, error)
	GetModelList(ctx context.Context, in *GetModelListRequest, opts ...grpc.CallOption) (*GetModelListResponse, error)
	GetExperimentResult(ctx context.Context, in *GetExperimentResultRequest, opts ...grpc.CallOption) (*GetExperimentResultResponse, error)
}

type metadataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataServiceClient(cc grpc.ClientConnInterface) MetadataServiceClient {
	return &metadataServiceClient{cc}
}

func (c *metadataServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/grpcapi.MetadataService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/grpcapi.MetadataService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, "/grpcapi.MetadataService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) GetStockInfo(ctx context.Context, in *GetStockInfoRequest, opts ...grpc.CallOption) (*GetStockInfoResponse, error) {
	out := new(GetStockInfoResponse)
	err := c.cc.Invoke(ctx, "/grpcapi.MetadataService/GetStockInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) GetModelInfo(ctx context.Context, in *GetModelInfoRequest, opts ...grpc.CallOption) (*GetModelInfoResponse, error) {
	out := new(GetModelInfoResponse)
	err := c.cc.Invoke(ctx, "/grpcapi.MetadataService/GetModelInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) GetStockList(ctx context.Context, in *GetStockListRequest, opts ...grpc.CallOption) (*GetStockListResponse, error) {
	out := new(GetStockListResponse)
	err := c.cc.Invoke(ctx, "/grpcapi.MetadataService/GetStockList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) GetModelList(ctx context.Context, in *GetModelListRequest, opts ...grpc.CallOption) (*GetModelListResponse, error) {
	out := new(GetModelListResponse)
	err := c.cc.Invoke(ctx, "/grpcapi.MetadataService/GetModelList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) GetExperimentResult(ctx context.Context, in *GetExperimentResultRequest, opts ...grpc.CallOption) (*GetExperimentResultResponse, error) {
	out := new(GetExperimentResultResponse)
	err := c.cc.Invoke(ctx, "/grpcapi.MetadataService/GetExperimentResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServiceServer is the server API for MetadataService service.
// All implementations must embed UnimplementedMetadataServiceServer
// for forward compatibility
type MetadataServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	GetStockInfo(context.Context, *GetStockInfoRequest) (*GetStockInfoResponse, error)
	GetModelInfo(context.Context, *GetModelInfoRequest) (*GetModelInfoResponse, error)
	GetStockList(context.Context, *GetStockListRequest) (*GetStockListResponse, error)
	GetModelList(context.Context, *GetModelListRequest) (*GetModelListResponse, error)
	GetExperimentResult(context.Context, *GetExperimentResultRequest) (*GetExperimentResultResponse, error)
	mustEmbedUnimplementedMetadataServiceServer()
}

// UnimplementedMetadataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetadataServiceServer struct {
}

func (UnimplementedMetadataServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedMetadataServiceServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedMetadataServiceServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedMetadataServiceServer) GetStockInfo(context.Context, *GetStockInfoRequest) (*GetStockInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockInfo not implemented")
}
func (UnimplementedMetadataServiceServer) GetModelInfo(context.Context, *GetModelInfoRequest) (*GetModelInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelInfo not implemented")
}
func (UnimplementedMetadataServiceServer) GetStockList(context.Context, *GetStockListRequest) (*GetStockListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockList not implemented")
}
func (UnimplementedMetadataServiceServer) GetModelList(context.Context, *GetModelListRequest) (*GetModelListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelList not implemented")
}
func (UnimplementedMetadataServiceServer) GetExperimentResult(context.Context, *GetExperimentResultRequest) (*GetExperimentResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExperimentResult not implemented")
}
func (UnimplementedMetadataServiceServer) mustEmbedUnimplementedMetadataServiceServer() {}

// UnsafeMetadataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServiceServer will
// result in compilation errors.
type UnsafeMetadataServiceServer interface {
	mustEmbedUnimplementedMetadataServiceServer()
}

func RegisterMetadataServiceServer(s grpc.ServiceRegistrar, srv MetadataServiceServer) {
	s.RegisterService(&MetadataService_ServiceDesc, srv)
}

func _MetadataService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.MetadataService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.MetadataService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.MetadataService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_GetStockInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStockInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetStockInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.MetadataService/GetStockInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetStockInfo(ctx, req.(*GetStockInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_GetModelInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetModelInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.MetadataService/GetModelInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetModelInfo(ctx, req.(*GetModelInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_GetStockList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStockListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetStockList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.MetadataService/GetStockList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetStockList(ctx, req.(*GetStockListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_GetModelList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetModelList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.MetadataService/GetModelList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetModelList(ctx, req.(*GetModelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_GetExperimentResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExperimentResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetExperimentResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.MetadataService/GetExperimentResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetExperimentResult(ctx, req.(*GetExperimentResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetadataService_ServiceDesc is the grpc.ServiceDesc for MetadataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetadataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.MetadataService",
	HandlerType: (*MetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _MetadataService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _MetadataService_Logout_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _MetadataService_GetUserInfo_Handler,
		},
		{
			MethodName: "GetStockInfo",
			Handler:    _MetadataService_GetStockInfo_Handler,
		},
		{
			MethodName: "GetModelInfo",
			Handler:    _MetadataService_GetModelInfo_Handler,
		},
		{
			MethodName: "GetStockList",
			Handler:    _MetadataService_GetStockList_Handler,
		},
		{
			MethodName: "GetModelList",
			Handler:    _MetadataService_GetModelList_Handler,
		},
		{
			MethodName: "GetExperimentResult",
			Handler:    _MetadataService_GetExperimentResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/metadata.proto",
}
