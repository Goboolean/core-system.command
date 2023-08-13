/*
Package grpcadapter provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/user/user.proto).

All commands about authorization services are provided by this adapter.
*/
package grpcadapter

import (
	grpcapi "github.com/Goboolean/command-server/api/grpc/user"
	"sync"
)

type UserAdapter struct {
	grpcapi.UnimplementedUserServiceServer
}

var (
	instance *UserAdapter
	once     sync.Once
)

func New() *UserAdapter {
	once.Do(func() {
		instance = &UserAdapter{}
	})

	return instance
}
