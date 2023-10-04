package model

import "context"

func (m *Manager) IsExistModelById(ctx context.Context, id string) (bool, error) {
	return m.db.IsExistModelById(ctx, id)
}
