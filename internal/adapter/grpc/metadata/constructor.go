/*
Package grpcadapter provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/metadata/metadata.proto).

All commands for data requests are provided by this adapter.
*/
package grpcadapter

import (
	grpcapi "github.com/Goboolean/command-server/api/grpc/metadata"
	"sync"
)

type MetadataAdapter struct {
	grpcapi.UnimplementedMetadataServiceServer
}

var (
	instance *MetadataAdapter
	once     sync.Once
)

func New() *MetadataAdapter {
	once.Do(func() {
		instance = &MetadataAdapter{}
	})

	return instance
}
