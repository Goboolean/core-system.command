/*
Package grpcadapter provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/model/model.proto).

All commands for services to model management are provided by this adapter.
*/
package grpcadapter

import (
	grpcapi "github.com/Goboolean/command-server/api/grpc/model"
	"sync"
)

type ModelAdapter struct {
	grpcapi.UnimplementedModelServiceServer
}

var (
	instance *ModelAdapter
	once     sync.Once
)

func New() *ModelAdapter {
	once.Do(func() {
		instance = &ModelAdapter{}
	})

	return instance
}
