// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: api/grpc/simulation/simulation.proto

package simulation

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

// SimulationServiceClient is the client API for SimulationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimulationServiceClient interface {
	// Returns to the stream what state the model being simulated is in.
	// When the simulation is finished, return access_token
	// and allow the result to be retrieved from the join server.
	SimulationModel(ctx context.Context, in *SimulationModelRequest, opts ...grpc.CallOption) (SimulationService_SimulationModelClient, error)
	GetSimulationResult(ctx context.Context, in *GetSimulationResultRequest, opts ...grpc.CallOption) (SimulationService_GetSimulationResultClient, error)
	GetSimulationList(ctx context.Context, in *GetSimulationListRequest, opts ...grpc.CallOption) (*GetSimulationListResponse, error)
}

type simulationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimulationServiceClient(cc grpc.ClientConnInterface) SimulationServiceClient {
	return &simulationServiceClient{cc}
}

func (c *simulationServiceClient) SimulationModel(ctx context.Context, in *SimulationModelRequest, opts ...grpc.CallOption) (SimulationService_SimulationModelClient, error) {
	stream, err := c.cc.NewStream(ctx, &SimulationService_ServiceDesc.Streams[0], "/simulation.SimulationService/SimulationModel", opts...)
	if err != nil {
		return nil, err
	}
	x := &simulationServiceSimulationModelClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SimulationService_SimulationModelClient interface {
	Recv() (*SimulationStatus, error)
	grpc.ClientStream
}

type simulationServiceSimulationModelClient struct {
	grpc.ClientStream
}

func (x *simulationServiceSimulationModelClient) Recv() (*SimulationStatus, error) {
	m := new(SimulationStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *simulationServiceClient) GetSimulationResult(ctx context.Context, in *GetSimulationResultRequest, opts ...grpc.CallOption) (SimulationService_GetSimulationResultClient, error) {
	stream, err := c.cc.NewStream(ctx, &SimulationService_ServiceDesc.Streams[1], "/simulation.SimulationService/GetSimulationResult", opts...)
	if err != nil {
		return nil, err
	}
	x := &simulationServiceGetSimulationResultClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SimulationService_GetSimulationResultClient interface {
	Recv() (*GetSimulationResultResponse, error)
	grpc.ClientStream
}

type simulationServiceGetSimulationResultClient struct {
	grpc.ClientStream
}

func (x *simulationServiceGetSimulationResultClient) Recv() (*GetSimulationResultResponse, error) {
	m := new(GetSimulationResultResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *simulationServiceClient) GetSimulationList(ctx context.Context, in *GetSimulationListRequest, opts ...grpc.CallOption) (*GetSimulationListResponse, error) {
	out := new(GetSimulationListResponse)
	err := c.cc.Invoke(ctx, "/simulation.SimulationService/GetSimulationList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimulationServiceServer is the server API for SimulationService service.
// All implementations must embed UnimplementedSimulationServiceServer
// for forward compatibility
type SimulationServiceServer interface {
	// Returns to the stream what state the model being simulated is in.
	// When the simulation is finished, return access_token
	// and allow the result to be retrieved from the join server.
	SimulationModel(*SimulationModelRequest, SimulationService_SimulationModelServer) error
	GetSimulationResult(*GetSimulationResultRequest, SimulationService_GetSimulationResultServer) error
	GetSimulationList(context.Context, *GetSimulationListRequest) (*GetSimulationListResponse, error)
	mustEmbedUnimplementedSimulationServiceServer()
}

// UnimplementedSimulationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSimulationServiceServer struct {
}

func (UnimplementedSimulationServiceServer) SimulationModel(*SimulationModelRequest, SimulationService_SimulationModelServer) error {
	return status.Errorf(codes.Unimplemented, "method SimulationModel not implemented")
}
func (UnimplementedSimulationServiceServer) GetSimulationResult(*GetSimulationResultRequest, SimulationService_GetSimulationResultServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSimulationResult not implemented")
}
func (UnimplementedSimulationServiceServer) GetSimulationList(context.Context, *GetSimulationListRequest) (*GetSimulationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimulationList not implemented")
}
func (UnimplementedSimulationServiceServer) mustEmbedUnimplementedSimulationServiceServer() {}

// UnsafeSimulationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimulationServiceServer will
// result in compilation errors.
type UnsafeSimulationServiceServer interface {
	mustEmbedUnimplementedSimulationServiceServer()
}

func RegisterSimulationServiceServer(s grpc.ServiceRegistrar, srv SimulationServiceServer) {
	s.RegisterService(&SimulationService_ServiceDesc, srv)
}

func _SimulationService_SimulationModel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimulationModelRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimulationServiceServer).SimulationModel(m, &simulationServiceSimulationModelServer{stream})
}

type SimulationService_SimulationModelServer interface {
	Send(*SimulationStatus) error
	grpc.ServerStream
}

type simulationServiceSimulationModelServer struct {
	grpc.ServerStream
}

func (x *simulationServiceSimulationModelServer) Send(m *SimulationStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _SimulationService_GetSimulationResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetSimulationResultRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimulationServiceServer).GetSimulationResult(m, &simulationServiceGetSimulationResultServer{stream})
}

type SimulationService_GetSimulationResultServer interface {
	Send(*GetSimulationResultResponse) error
	grpc.ServerStream
}

type simulationServiceGetSimulationResultServer struct {
	grpc.ServerStream
}

func (x *simulationServiceGetSimulationResultServer) Send(m *GetSimulationResultResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SimulationService_GetSimulationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSimulationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationServiceServer).GetSimulationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simulation.SimulationService/GetSimulationList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationServiceServer).GetSimulationList(ctx, req.(*GetSimulationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SimulationService_ServiceDesc is the grpc.ServiceDesc for SimulationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimulationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simulation.SimulationService",
	HandlerType: (*SimulationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSimulationList",
			Handler:    _SimulationService_GetSimulationList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SimulationModel",
			Handler:       _SimulationService_SimulationModel_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSimulationResult",
			Handler:       _SimulationService_GetSimulationResult_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/grpc/simulation/simulation.proto",
}
