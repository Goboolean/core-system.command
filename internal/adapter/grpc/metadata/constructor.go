/*
Package metadata provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/metadata/metadata.proto).

All commands for data requests are provided by this adapter.
*/
package metadata

import (
	api "github.com/Goboolean/core-system.command/api/grpc/metadata"
	"sync"
)

type Adapter struct {
	api.UnimplementedMetadataServiceServer
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
