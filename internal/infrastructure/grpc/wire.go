//go:build wireinject
// +build wireinject

package grpcserver

import (
	"context"
	grpcadapter "github.com/Goboolean/command-server/internal/adapter/grpc"
	"github.com/google/wire"
)

func initializeHost(ctx context.Context) (*Host, error) {
	wire.Build(grpcadapter.New, New)
	return &Host{}, nil
}
