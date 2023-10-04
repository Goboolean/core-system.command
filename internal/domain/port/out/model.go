package out

import "context"

type ModelDBPort interface {
	IsExistModelById(ctx context.Context, id string) (bool, error)
}
