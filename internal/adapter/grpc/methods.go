/**
 * Implementation of the gRPC methods.
 * The interfaces are defined in the proto file.
 * api/grpc/metadata.proto
 *
 * All command server's services are provided by this adapter.
 */

package adapter

import (
	"context"
	"fmt"
	grpcapi "github.com/Goboolean/command-server/api/grpc"
)

func (a *MetadataCommandAdapter) Login(ctx context.Context, in *grpcapi.LoginRequest) (nil *grpcapi.LoginResponse, err error) {
	id, password := in.Id, in.Password
	fmt.Println(id, password)
	return nil, fmt.Errorf("not implemented")
}

func (a *MetadataCommandAdapter) Logout(ctx context.Context, in *grpcapi.LogoutRequest) (nil *grpcapi.LogoutResponse, err error) {
	userToken := in.UserToken
	fmt.Println(userToken)
	return nil, fmt.Errorf("not implemented")
}

func (a *MetadataCommandAdapter) GetUserInfo(ctx context.Context, in *grpcapi.GetUserInfoRequest) (nil *grpcapi.GetUserInfoResponse, err error) {
	userToken := in.UserToken
	fmt.Println(userToken)
	return nil, fmt.Errorf("not implemented")
}

func (a *MetadataCommandAdapter) GetStockInfo(ctx context.Context, in *grpcapi.GetStockInfoRequest) (nil *grpcapi.GetStockInfoResponse, err error) {
	stockName := in.StockId
	fmt.Println(stockName)
	return nil, fmt.Errorf("not implemented")
}

func (a *MetadataCommandAdapter) GetModelInfo(ctx context.Context, in *grpcapi.GetModelInfoRequest) (nil *grpcapi.GetModelInfoResponse, err error) {
	userToken := in.UserToken
	fmt.Println(userToken)
	return nil, fmt.Errorf("not implemented")
}

func (a *MetadataCommandAdapter) GetStockList(ctx context.Context, in *grpcapi.GetStockListRequest) (nil *grpcapi.GetStockListResponse, err error) {
	fmt.Println("GetStockList")
	return nil, fmt.Errorf("not implemented")
}

func (a *MetadataCommandAdapter) GetModelList(ctx context.Context, in *grpcapi.GetModelListRequest) (nil *grpcapi.GetModelListResponse, err error) {
	userToken := in.UserToken
	fmt.Println(userToken)
	return nil, fmt.Errorf("not implemented")
}

func (a *MetadataCommandAdapter) GetExperimentResult(ctx context.Context, in *grpcapi.GetExperimentResultRequest) (nil *grpcapi.GetExperimentResultResponse, err error) {
	userToken := in.UserToken
	modelId := in.ModelId
	stockId := in.StockId
	parameter := in.Parameter
	startDate := in.StartDate
	endDate := in.EndDate
	fmt.Println(userToken, modelId, stockId, parameter, startDate, endDate)
	return nil, fmt.Errorf("not implemented")
}
