/*
Package main is the entry point of the command server.
*/
package main

import (
	"context"
	grpcserver "github.com/Goboolean/command-server/internal/infrastructure/grpc"
	"log"
	"os"
)

func main() {

	// Set up environment variables.
	err := os.Setenv("COMMAND_SERVER_PORT", "50051")
	if err != nil {
		log.Fatal(err)
	}

	// Set up server.
	ctx, cancel := context.WithCancel(context.Background())
	host, err := grpcserver.InitializeHost(ctx)
	if err != nil {
		log.Fatal(err)
	}
	go host.Run(ctx)

	select {
	case <-ctx.Done():
		// Every fatal error will be caught here.
		if err := ctx.Err(); err != nil {
			log.Fatal(err)
		}
		cancel()
	}
}
