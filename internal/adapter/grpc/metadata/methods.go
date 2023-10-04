/*
Package metadata provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/metadata/metadata.proto).

All commands for data requests are provided by this adapter.
*/
package metadata

import (
	"context"
	"fmt"
	api "github.com/Goboolean/core-system.command/api/grpc/metadata"
)

func (a *Adapter) GetUserInfo(ctx context.Context, in *api.GetUserInfoRequest) (nil *api.GetUserInfoResponse, err error) {
	userToken := in.GetUserToken()
	fmt.Println(userToken)
	return nil, fmt.Errorf("not implemented")
}
