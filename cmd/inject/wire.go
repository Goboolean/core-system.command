//go:build wireinject
// +build wireinject

package inject

import (
	"context"
	metadataAdapter "github.com/Goboolean/core-system.command/internal/adapter/grpc/metadata"
	modelAdapter "github.com/Goboolean/core-system.command/internal/adapter/grpc/model"
	simulationAdapter "github.com/Goboolean/core-system.command/internal/adapter/grpc/simulation"
	userAdapter "github.com/Goboolean/core-system.command/internal/adapter/grpc/user"
	"github.com/Goboolean/core-system.command/internal/adapter/persistence"
	"github.com/Goboolean/core-system.command/internal/domain/port/out"
	"github.com/Goboolean/core-system.command/internal/domain/service/metadata"
	"github.com/Goboolean/core-system.command/internal/domain/service/model"
	"github.com/Goboolean/core-system.command/internal/domain/service/simulation"
	"github.com/Goboolean/core-system.command/internal/domain/service/task"
	grpcServer "github.com/Goboolean/core-system.command/internal/infrastructure/grpc"
	"github.com/google/wire"
)

// Inject for grpc server Host
func InitializeHost(ctx context.Context) (*grpcServer.Host, error) {
	wire.Build(grpcServer.New, metadataAdapter.New, modelAdapter.New, simulationAdapter.New, userAdapter.New)
	return &grpcServer.Host{}, nil
}

// Inject for task manager(Mock)
func InitializeMockTaskManager(ctx context.Context, kafka out.TaskPort) (*task.Manager, error) {
	wire.Build(task.New)
	return &task.Manager{}, nil
}

// Inject for simulation manager(Mock)

func ProviderOutModelDBPort(db *persistence.MockAdapter) out.ModelDBPort {
	return db
}

func ProviderOutMetadataDBPort(db *persistence.MockAdapter) out.MetadataDBPort {
	return db
}

func ProviderOutSimulationDBPort(db *persistence.MockAdapter) out.SimulationDBPort {
	return db
}

func InitializeMockSimulationManager(ctx context.Context, tm *task.Manager, db *persistence.MockAdapter) (*simulation.Manager, error) {
	wire.Build(model.New, metadata.New, simulation.New, ProviderOutModelDBPort, ProviderOutMetadataDBPort, ProviderOutSimulationDBPort)
	return &simulation.Manager{}, nil
}
