/*
Package out

out 패키지는 서비스 로직의 요청 인터페이스를 제공합니다.
*/
package out

import "context"

type TaskPort interface {
	// uuid, model_id, stock_id
	AllocateTask(ctx context.Context, uuid string, modelId string, stockId string) error
}
