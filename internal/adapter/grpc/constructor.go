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

type MetadataCommandAdapter struct {
	grpcapi.UnimplementedMetadataServiceServer
}

var (
	instance *MetadataCommandAdapter
	once     sync.Once
)

func New() *MetadataCommandAdapter {
	once.Do(func() {
		instance = &MetadataCommandAdapter{}
	})

	return instance
}
