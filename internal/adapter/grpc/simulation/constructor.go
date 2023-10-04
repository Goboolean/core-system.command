/*
Package simulation provides the implementation of the gRPC methods.

The interfaces are defined in the proto file(api/grpc/simulation/simulation.proto).

All commands for services to model management are provided by this adapter.
*/
package simulation

import (
	grpcapi "github.com/Goboolean/core-system.command/api/grpc/simulation"
	"sync"
)

type Adapter struct {
	grpcapi.UnimplementedSimulationServiceServer
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
