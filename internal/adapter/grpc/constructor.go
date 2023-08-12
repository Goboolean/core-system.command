/*
Package grpcadapter provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/metadata.proto).

All command server's services are provided by this adapter.
*/
package grpcadapter

import (
	grpcapi "github.com/Goboolean/command-server/api/grpc"
	"sync"
)

type Adapter struct {
	grpcapi.UnimplementedMetadataServiceServer
	grpcapi.UnimplementedModelServiceServer
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
