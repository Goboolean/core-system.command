package broker

import (
	"context"
	"errors"
)

type MockAdapter struct {
	ERR bool
}

func NewMockAdapter(err bool) *MockAdapter {
	return &MockAdapter{
		ERR: err,
	}
}

func (m *MockAdapter) SetError(err bool) {
	m.ERR = err
	return
}

func (m *MockAdapter) AllocateTask(ctx context.Context, uuid string, modelId string, stockId string) error {
	if m.ERR {
		return errors.New("kafka infra error")
	}
	return nil
}
