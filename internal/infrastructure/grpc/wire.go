//go:build wireinject
// +build wireinject

package grpcserver

import (
	"context"
	metadataadapter "github.com/Goboolean/command-server/internal/adapter/grpc/metadata"
	modeladapter "github.com/Goboolean/command-server/internal/adapter/grpc/model"
	useradapter "github.com/Goboolean/command-server/internal/adapter/grpc/user"
	"github.com/google/wire"
)

func initializeHost(ctx context.Context) (*Host, error) {
	wire.Build(New, metadataadapter.New, modeladapter.New, useradapter.New)
	return &Host{}, nil
}
