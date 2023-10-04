/*
Package model provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/model/model.proto).

All commands for services to model management are provided by this adapter.
*/
package model

import (
	api "github.com/Goboolean/command-server/api/grpc/model"
	"sync"
)

type Adapter struct {
	api.UnimplementedModelServiceServer
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
