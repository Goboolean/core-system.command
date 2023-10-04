/*
Package main is the entry point of the command server.
*/
package main

import (
	"context"
	"github.com/Goboolean/core-system.command/cmd/inject"
	"github.com/Goboolean/core-system.command/internal/util/log"
	"os"
)

func main() {
	log.Init()
	defer log.Logger.Sync()

	// Set up environment variables.
	err := os.Setenv("COMMAND_SERVER_PORT", "50051")
	if err != nil {
		log.Logger.Fatal(err.Error())
	}

	// Set up server.
	ctx, cancel := context.WithCancel(context.Background())
	host, err := inject.InitializeHost(ctx)
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	go host.Run(ctx)

	select {
	case <-ctx.Done():
		// Every fatal error will be caught here.
		if err := ctx.Err(); err != nil {
			log.Logger.Fatal(err.Error())
		}
		cancel()
	}
}
