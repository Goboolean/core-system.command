/*
Package model provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/model/model.proto).

All commands for services to model management are provided by this adapter.
*/
package model

import (
	"context"
	"fmt"
	api "github.com/Goboolean/core-system.command/api/grpc/model"
)

func (a *Adapter) RegisterModel(ctx context.Context, in *api.RegisterModelRequest) (nil *api.RegisterModelResponse, err error) {
	userToken := in.UserToken
	fmt.Println(userToken)
	return nil, fmt.Errorf("not implemented")
}
