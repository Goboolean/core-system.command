/*
Package grpcadapter provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/user/user.proto).

All commands about authorization services are provided by this adapter.
*/
package grpcadapter

import (
	"context"
	"fmt"
	grpcapi "github.com/Goboolean/command-server/api/grpc/user"
)

func (a *UserAdapter) Login(ctx context.Context, in *grpcapi.LoginRequest) (nil *grpcapi.LoginResponse, err error) {
	id, password := in.Id, in.Password
	fmt.Println(id, password)
	return nil, fmt.Errorf("not implemented")
}
