package simulation

import (
	"context"
	"github.com/Goboolean/command-server/internal/domain/entity"
	"github.com/Goboolean/command-server/internal/util/hash"
	"github.com/Goboolean/command-server/internal/util/log"
	"github.com/Goboolean/command-server/internal/util/msg"
	"go.uber.org/zap"
)

// client key 반환
func (m *Manager) Allocate(ctx context.Context, modelId string, stockId string, client *entity.Client) (string, error) {

	uuid := hash.Sha256(modelId + stockId)

	// task가 존재하는지 확인 -> 존재하면 단순히 구독
	if m.task.IsExistTask(uuid) {
		clientKey := m.task.AddSubscriber(uuid, client)
		log.Logger.Info("task already exist. subscribe client to task", zap.String("uuid", uuid), zap.String("clientKey", clientKey))
		return clientKey, nil
	}

	// simulation이 존재하는지 확인 -> 테스크 할당 필요하지 않음
	sim, err := m.db.IsExistSimulationByUUID(ctx, uuid)
	if err != nil {
		return "", err
	}
	if sim == true {
		go client.Send(msg.Success(uuid), false)
		log.Logger.Info("simulation already exist. return uuid", zap.String("uuid", uuid))
		return "", nil
	}

	// model, stock 정보 확인
	// channel을 잘 닫는 방법은?
	chkCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan checkInfo, 2)
	wgCh := make(chan struct{}, 2)

	go m.checkModel(chkCtx, modelId, ch, wgCh)
	go m.checkStock(chkCtx, stockId, ch, wgCh)

	// goroutine들이 종료되면, channel을 닫도록 함
	go func() {
		for i := 0; i < 2; i++ {
			<-wgCh
		}
		close(ch)
		close(wgCh)
	}()

	for i := 0; i < 2; i++ {
		select {
		case c := <-ch:
			if c.ERR != nil {
				return "", c.ERR
			}

			if c.EXIST == false {
				go client.Send(msg.NotFound(), false)
				log.Logger.Info(c.NAME+" not exist. return 404", zap.String("modelId", modelId), zap.String("stockId", stockId))
				return "", nil
			}
		}
	}

	// task 할당
	return m.task.Allocate(ctx, modelId, stockId, client)
}

func (m *Manager) checkModel(ctx context.Context, modelId string, ch chan<- checkInfo, wgCh chan<- struct{}) {
	model, err := m.model.IsExistModelById(ctx, modelId)
	select {
	case ch <- checkInfo{
		NAME:  "model",
		EXIST: model,
		ERR:   err,
	}:
	}
	wgCh <- struct{}{}
}

func (m *Manager) checkStock(ctx context.Context, stockId string, ch chan<- checkInfo, wgCh chan<- struct{}) {
	stock, err := m.metadata.IsExistStockById(ctx, stockId)
	select {
	case ch <- checkInfo{
		NAME:  "stock",
		EXIST: stock,
		ERR:   err,
	}:
	}
	wgCh <- struct{}{}
}
