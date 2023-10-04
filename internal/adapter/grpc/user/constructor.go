/*
Package user provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/user/user.proto).

All commands about authorization services are provided by this adapter.
*/
package user

import (
	api "github.com/Goboolean/core-system.command/api/grpc/user"
	"sync"
)

type Adapter struct {
	api.UnimplementedUserServiceServer
}

var (
	instance *Adapter
	once     sync.Once
)

func New() *Adapter {
	once.Do(func() {
		instance = &Adapter{}
	})

	return instance
}
