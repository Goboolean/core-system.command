/*
Package client is the client for the gRPC server.

The interfaces are defined in the proto file(api/grpc/metadata.proto).

You can use it when writing the client code.

This code is not used when the command server is operating.
*/
package main

import (
	"context"
	"flag"
	"github.com/Goboolean/core-system.command/api/grpc/user"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Login(ctx, &user.LoginRequest{Id: "test", Password: "test"})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Printf("Login Token: %s", r.GetUserToken())
}
