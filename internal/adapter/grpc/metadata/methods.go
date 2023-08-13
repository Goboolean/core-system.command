/*
Package grpcadapter provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/metadata/metadata.proto).

All commands for data requests are provided by this adapter.
*/
package grpcadapter

import (
	"context"
	"fmt"
	grpcapi "github.com/Goboolean/command-server/api/grpc/metadata"
)

func (a *MetadataAdapter) GetUserInfo(ctx context.Context, in *grpcapi.GetUserInfoRequest) (nil *grpcapi.GetUserInfoResponse, err error) {
	userToken := in.GetUserToken()
	fmt.Println(userToken)
	return nil, fmt.Errorf("not implemented")
}
