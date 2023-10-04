/*
Package grpcserver is a server that accepts grpc requests and can be used by replacing the implementation adapter as needed.

The interfaces are defined in the proto file(api/grpc/metadata.proto).

All command server's services are provided by grpcadapter.
*/
package grpcserver

import (
	"context"
	"fmt"
	metadataapi "github.com/Goboolean/command-server/api/grpc/metadata"
	modelapi "github.com/Goboolean/command-server/api/grpc/model"
	simulationapi "github.com/Goboolean/command-server/api/grpc/simulation"
	userapi "github.com/Goboolean/command-server/api/grpc/user"
	metadata "github.com/Goboolean/command-server/internal/adapter/grpc/metadata"
	model "github.com/Goboolean/command-server/internal/adapter/grpc/model"
	simulation "github.com/Goboolean/command-server/internal/adapter/grpc/simulation"
	user "github.com/Goboolean/command-server/internal/adapter/grpc/user"
	"google.golang.org/grpc"
	"net"
	"os"
	"sync"
)

type Host struct {
	server            *grpc.Server
	metadataAdapter   *metadata.Adapter
	modelAdapter      *model.Adapter
	simulationAdapter *simulation.Adapter
	userAdapter       *user.Adapter
}

var (
	instance *Host
	once     sync.Once
)

func New(metaA *metadata.Adapter, modelA *model.Adapter, simulationA *simulation.Adapter, userA *user.Adapter) *Host {
	once.Do(func() {
		instance = &Host{
			metadataAdapter:   metaA,
			modelAdapter:      modelA,
			simulationAdapter: simulationA,
			userAdapter:       userA,
		}
	})

	return instance
}

// Run starts the grpc server
//
// The Run method is designed to run as goroutine.
// ctx is used to control the flow of work.
// The adapter is an implementation of the grpc protocol (interface). You can change the implementation as needed.
func (h *Host) Run(ctx context.Context) {

	port, flag := os.LookupEnv("COMMAND_SERVER_PORT")
	if !flag {
		fmt.Println("Command server port required in environment variable.")
		panic("Command server port required. Please set COMMAND_SERVER_PORT environment variable.")
	}

	// listen on port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Println("Error listening gRPC server port at: "+port, err)
		panic(err)
	}

	// make new grpc server and register adapter
	h.server = grpc.NewServer()
	metadataapi.RegisterMetadataServiceServer(h.server, h.metadataAdapter)
	modelapi.RegisterModelServiceServer(h.server, h.modelAdapter)
	simulationapi.RegisterSimulationServiceServer(h.server, h.simulationAdapter)
	userapi.RegisterUserServiceServer(h.server, h.userAdapter)

	// start server
	go func() {
		if err := h.server.Serve(lis); err != nil {
			fmt.Println("Error starting gRPC server", err)
			panic(err)
		}
	}()

	// wait for context to be done
	select {
	case <-ctx.Done():
		h.Close()
	}
}

func (h *Host) Close() {
	h.server.GracefulStop()
}
