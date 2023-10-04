package out

import "context"

type SimulationDBPort interface {
	IsExistSimulationByUUID(ctx context.Context, uuid string) (bool, error)
}
