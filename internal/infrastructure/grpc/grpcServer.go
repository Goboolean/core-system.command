/*
Package grpcserver is a server that accepts grpc requests and can be used by replacing the implementation adapter as needed.

The interfaces are defined in the proto file(api/grpc/metadata.proto).

All command server's services are provided by grpcadapter.
*/
package grpcserver

import (
	"context"
	"fmt"
	grpcapi "github.com/Goboolean/command-server/api/grpc"
	grpcadapter "github.com/Goboolean/command-server/internal/adapter/grpc"
	"google.golang.org/grpc"
	"net"
	"os"
	"sync"
)

type Host struct {
	server *grpc.Server
	impl   *grpcadapter.Adapter
}

var (
	instance *Host
	once     sync.Once
)

func New(adapter *grpcadapter.Adapter) *Host {
	once.Do(func() {
		instance = &Host{
			impl: adapter,
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
	grpcapi.RegisterMetadataServiceServer(h.server, h.impl)

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
