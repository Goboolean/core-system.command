/*
Package user provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/user/user.proto).

All commands about authorization services are provided by this adapter.
*/
package user

import (
	"context"
	"fmt"
	api "github.com/Goboolean/core-system.command/api/grpc/user"
)

func (a *Adapter) Login(ctx context.Context, in *api.LoginRequest) (nil *api.LoginResponse, err error) {
	id, password := in.Id, in.Password
	fmt.Println(id, password)
	return nil, fmt.Errorf("not implemented")
}
