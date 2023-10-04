package metadata

import "context"

func (m *Manager) IsExistStockById(ctx context.Context, id string) (bool, error) {
	return m.db.IsExistStockById(ctx, id)
}
