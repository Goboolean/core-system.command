package persistence

import (
	"context"
	"fmt"
)

type MockAdapter struct {
	stock      []string
	model      []string
	simulation []string

	stockDBError      bool
	modelDBError      bool
	simulationDBError bool
}

func NewMockAdapter() *MockAdapter {
	return &MockAdapter{
		stock:      make([]string, 0),
		model:      make([]string, 0),
		simulation: make([]string, 0),

		stockDBError:      false,
		modelDBError:      false,
		simulationDBError: false,
	}
}

func (a *MockAdapter) SetStockDBError(err bool) {
	a.stockDBError = err
	return
}

func (a *MockAdapter) SetModelDBError(err bool) {
	a.modelDBError = err
	return
}

func (a *MockAdapter) SetSimulationDBError(err bool) {
	a.simulationDBError = err
	return
}

func (a *MockAdapter) IsExistStockById(ctx context.Context, id string) (bool, error) {
	if a.stockDBError {
		return false, fmt.Errorf("stock db error")
	}

	for _, stock := range a.stock {
		if stock == id {
			return true, nil
		}
	}
	return false, nil
}

func (a *MockAdapter) IsExistModelById(ctx context.Context, id string) (bool, error) {
	if a.modelDBError {
		return false, fmt.Errorf("model db error")
	}

	for _, model := range a.model {
		if model == id {
			return true, nil
		}
	}
	return false, nil
}

func (a *MockAdapter) IsExistSimulationByUUID(ctx context.Context, uuid string) (bool, error) {
	if a.simulationDBError {
		return false, fmt.Errorf("simulation db error")
	}

	for _, simulation := range a.simulation {
		if simulation == uuid {
			return true, nil
		}
	}
	return false, nil
}

func (a *MockAdapter) AddStock(id string) error {
	a.stock = append(a.stock, id)
	return nil
}

func (a *MockAdapter) AddModel(id string) error {
	a.model = append(a.model, id)
	return nil
}

func (a *MockAdapter) AddSimulation(uuid string) error {
	a.simulation = append(a.simulation, uuid)
	return nil
}

func (a *MockAdapter) ClearStock() error {
	a.stock = make([]string, 0)
	return nil
}

func (a *MockAdapter) ClearModel() error {
	a.model = make([]string, 0)
	return nil
}

func (a *MockAdapter) ClearSimulation() error {
	a.simulation = make([]string, 0)
	return nil
}
