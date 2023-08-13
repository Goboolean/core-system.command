/*
Package grpcadapter provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/model/model.proto).

All commands for services to model management are provided by this adapter.
*/
package grpcadapter

import (
	"context"
	"fmt"
	"github.com/Goboolean/command-server/api/grpc/model"
)

func (a *ModelAdapter) RegisterModel(ctx context.Context, in *model.RegisterModelRequest) (nil *model.RegisterModelResponse, err error) {
	userToken := in.UserToken
	fmt.Println(userToken)
	return nil, fmt.Errorf("not implemented")
}

// SimulationModel
//
// This method requests a model simulation based on the model information received in the request.
// The return type is stream, which provides a simulation state.
// When the simulation is finished, return the access token to the join server at the end of the stream and close the stream.
func (a *ModelAdapter) SimulationModel(in *model.SimulationModelRequest, stream model.ModelService_SimulationModelServer) error {

	// TODO: Implement the simulation logic.
	// - simulation 상태 정보를 반환하는 channel을 service 도메인으로부터 받아온다.
	// - channel에 결과가 추가될 때마다 stream에 전송한다.
	// - channel의 마지막 데이터는 access_token을 가짐을 보장한다.
	// - channel이 닫히면 stream도 닫는다.

	ch := make(chan *model.SimulationStatus)

	// ==============================================================================================================
	go func() {
		ch <- &model.SimulationStatus{
			Status:      1,
			AccessToken: "access_token",
		}
		close(ch)
	}()

	for s := range ch {
		if err := stream.Send(s); err != nil {
			return err
		}
	}
	// ==============================================================================================================

	return nil
}

func (a *ModelAdapter) EditModel(ctx context.Context, in *model.EditModelRequest) (nil *model.EditModelResponse, err error) {
	userToken := in.UserToken
	modelId := in.ModelId
	fmt.Println(userToken, modelId)
	return nil, fmt.Errorf("not implemented")
}

func (a *ModelAdapter) DeleteModel(ctx context.Context, in *model.DeleteModelRequest) (nil *model.DeleteModelResponse, err error) {
	userToken := in.UserToken
	modelId := in.ModelId
	fmt.Println(userToken, modelId)
	return nil, fmt.Errorf("not implemented")
}
