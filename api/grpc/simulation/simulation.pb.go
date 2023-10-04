//
//Package simulation provides the interface of the gRPC methods.
//
//All commands for simulation are provided by this interfaces.
//
//Generate gRPC code:
//protoc --go_out=. --go_opt=paths=source_relative \
//--go-grpc_out=. --go-grpc_opt=paths=source_relative \
//api/grpc/simulation/simulation.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: api/grpc/simulation/simulation.proto

package simulation

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

// need implement about parameter
type SimulationModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserToken string  `protobuf:"bytes,1,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"`
	ModelId   string  `protobuf:"bytes,2,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	StockId   string  `protobuf:"bytes,3,opt,name=stock_id,json=stockId,proto3" json:"stock_id,omitempty"`
	Parameter []int64 `protobuf:"varint,4,rep,packed,name=parameter,proto3" json:"parameter,omitempty"`
}

func (x *SimulationModelRequest) Reset() {
	*x = SimulationModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_simulation_simulation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulationModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulationModelRequest) ProtoMessage() {}

func (x *SimulationModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_simulation_simulation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulationModelRequest.ProtoReflect.Descriptor instead.
func (*SimulationModelRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_simulation_simulation_proto_rawDescGZIP(), []int{0}
}

func (x *SimulationModelRequest) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

func (x *SimulationModelRequest) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

func (x *SimulationModelRequest) GetStockId() string {
	if x != nil {
		return x.StockId
	}
	return ""
}

func (x *SimulationModelRequest) GetParameter() []int64 {
	if x != nil {
		return x.Parameter
	}
	return nil
}

type SimulationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`                             // simulation status
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"` // token for accessing join server
}

func (x *SimulationStatus) Reset() {
	*x = SimulationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_simulation_simulation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulationStatus) ProtoMessage() {}

func (x *SimulationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_simulation_simulation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulationStatus.ProtoReflect.Descriptor instead.
func (*SimulationStatus) Descriptor() ([]byte, []int) {
	return file_api_grpc_simulation_simulation_proto_rawDescGZIP(), []int{1}
}

func (x *SimulationStatus) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SimulationStatus) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

// Simulation result information query command
type GetSimulationResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserToken    string `protobuf:"bytes,1,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"` // To verify access to information
	SimulationId string `protobuf:"bytes,2,opt,name=simulation_id,json=simulationId,proto3" json:"simulation_id,omitempty"`
}

func (x *GetSimulationResultRequest) Reset() {
	*x = GetSimulationResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_simulation_simulation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSimulationResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSimulationResultRequest) ProtoMessage() {}

func (x *GetSimulationResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_simulation_simulation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSimulationResultRequest.ProtoReflect.Descriptor instead.
func (*GetSimulationResultRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_simulation_simulation_proto_rawDescGZIP(), []int{2}
}

func (x *GetSimulationResultRequest) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

func (x *GetSimulationResultRequest) GetSimulationId() string {
	if x != nil {
		return x.SimulationId
	}
	return ""
}

type GetSimulationResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *GetSimulationResultResponse) Reset() {
	*x = GetSimulationResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_simulation_simulation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSimulationResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSimulationResultResponse) ProtoMessage() {}

func (x *GetSimulationResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_simulation_simulation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSimulationResultResponse.ProtoReflect.Descriptor instead.
func (*GetSimulationResultResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_simulation_simulation_proto_rawDescGZIP(), []int{3}
}

func (x *GetSimulationResultResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetSimulationResultResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

// Simulation list information query command
type GetSimulationListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserToken string `protobuf:"bytes,1,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"` // To verify access to information
}

func (x *GetSimulationListRequest) Reset() {
	*x = GetSimulationListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_simulation_simulation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSimulationListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSimulationListRequest) ProtoMessage() {}

func (x *GetSimulationListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_simulation_simulation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSimulationListRequest.ProtoReflect.Descriptor instead.
func (*GetSimulationListRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_simulation_simulation_proto_rawDescGZIP(), []int{4}
}

func (x *GetSimulationListRequest) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

type GetSimulationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SimulationId []string `protobuf:"bytes,1,rep,name=simulation_id,json=simulationId,proto3" json:"simulation_id,omitempty"`
}

func (x *GetSimulationListResponse) Reset() {
	*x = GetSimulationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_simulation_simulation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSimulationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSimulationListResponse) ProtoMessage() {}

func (x *GetSimulationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_simulation_simulation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSimulationListResponse.ProtoReflect.Descriptor instead.
func (*GetSimulationListResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_simulation_simulation_proto_rawDescGZIP(), []int{5}
}

func (x *GetSimulationListResponse) GetSimulationId() []string {
	if x != nil {
		return x.SimulationId
	}
	return nil
}

var File_api_grpc_simulation_simulation_proto protoreflect.FileDescriptor

var file_api_grpc_simulation_simulation_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x8b, 0x01, 0x0a, 0x16, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x22, 0x4d, 0x0a, 0x10, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x60, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x58, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x39, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x40, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xbc, 0x02, 0x0a, 0x11, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57,
	0x0a, 0x0f, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x22, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x6a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x26,
	0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x62, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x73, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_simulation_simulation_proto_rawDescOnce sync.Once
	file_api_grpc_simulation_simulation_proto_rawDescData = file_api_grpc_simulation_simulation_proto_rawDesc
)

func file_api_grpc_simulation_simulation_proto_rawDescGZIP() []byte {
	file_api_grpc_simulation_simulation_proto_rawDescOnce.Do(func() {
		file_api_grpc_simulation_simulation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_simulation_simulation_proto_rawDescData)
	})
	return file_api_grpc_simulation_simulation_proto_rawDescData
}

var file_api_grpc_simulation_simulation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_grpc_simulation_simulation_proto_goTypes = []interface{}{
	(*SimulationModelRequest)(nil),      // 0: simulation.SimulationModelRequest
	(*SimulationStatus)(nil),            // 1: simulation.SimulationStatus
	(*GetSimulationResultRequest)(nil),  // 2: simulation.GetSimulationResultRequest
	(*GetSimulationResultResponse)(nil), // 3: simulation.GetSimulationResultResponse
	(*GetSimulationListRequest)(nil),    // 4: simulation.GetSimulationListRequest
	(*GetSimulationListResponse)(nil),   // 5: simulation.GetSimulationListResponse
}
var file_api_grpc_simulation_simulation_proto_depIdxs = []int32{
	0, // 0: simulation.SimulationService.SimulationModel:input_type -> simulation.SimulationModelRequest
	2, // 1: simulation.SimulationService.GetSimulationResult:input_type -> simulation.GetSimulationResultRequest
	4, // 2: simulation.SimulationService.GetSimulationList:input_type -> simulation.GetSimulationListRequest
	1, // 3: simulation.SimulationService.SimulationModel:output_type -> simulation.SimulationStatus
	3, // 4: simulation.SimulationService.GetSimulationResult:output_type -> simulation.GetSimulationResultResponse
	5, // 5: simulation.SimulationService.GetSimulationList:output_type -> simulation.GetSimulationListResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_grpc_simulation_simulation_proto_init() }
func file_api_grpc_simulation_simulation_proto_init() {
	if File_api_grpc_simulation_simulation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_simulation_simulation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulationModelRequest); i {
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
		file_api_grpc_simulation_simulation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulationStatus); i {
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
		file_api_grpc_simulation_simulation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSimulationResultRequest); i {
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
		file_api_grpc_simulation_simulation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSimulationResultResponse); i {
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
		file_api_grpc_simulation_simulation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSimulationListRequest); i {
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
		file_api_grpc_simulation_simulation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSimulationListResponse); i {
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
			RawDescriptor: file_api_grpc_simulation_simulation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_simulation_simulation_proto_goTypes,
		DependencyIndexes: file_api_grpc_simulation_simulation_proto_depIdxs,
		MessageInfos:      file_api_grpc_simulation_simulation_proto_msgTypes,
	}.Build()
	File_api_grpc_simulation_simulation_proto = out.File
	file_api_grpc_simulation_simulation_proto_rawDesc = nil
	file_api_grpc_simulation_simulation_proto_goTypes = nil
	file_api_grpc_simulation_simulation_proto_depIdxs = nil
}
