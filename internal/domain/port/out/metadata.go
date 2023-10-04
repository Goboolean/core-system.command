package out

import "context"

type MetadataDBPort interface {
	IsExistStockById(ctx context.Context, id string) (bool, error)
}
