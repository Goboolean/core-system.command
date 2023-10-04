/*
Package task

Manager 객체는 현재 진행되는 작업을 관리하는 객체입니다.
*/
package task

import (
	"context"
	"fmt"
	"github.com/Goboolean/core-system.command/internal/domain/entity"
	"github.com/Goboolean/core-system.command/internal/util/hash"
	"github.com/Goboolean/core-system.command/internal/util/log"
	"github.com/Goboolean/core-system.command/internal/util/msg"
	"go.uber.org/zap"
)

// IsExistTask가 먼저 호출된 이후, 호출되어야 함!
// 할당 후에는, 클라이언트 고유 키값이 제공됨. 이는 클라이언트 구독 취소에 사용됨.
func (m *Manager) Allocate(ctx context.Context, modelId string, stockId string, client *entity.Client) (string, error) {
	uuid := hash.Sha256(modelId + stockId)

	// Allocate to kafka
	err := m.kafka.AllocateTask(ctx, uuid, modelId, stockId)
	if err != nil {
		log.Logger.Error(err.Error())
		go client.Send(msg.ServerError(), false)
		return "", err
	}

	// Insert to map
	if !m.insertTask(uuid) {
		log.Logger.Error("task already exist, but try to allocate task. before calling the Allocate function, make sure that the IsExistTask function is called.")
		return "", fmt.Errorf("task already exist")
	}

	// Add subscriber
	clientKey := m.AddSubscriber(uuid, client)
	log.Logger.Info("task allocated", zap.String("uuid", uuid), zap.String("clientKey", clientKey))

	return clientKey, nil
}

func (m *Manager) IsExistTask(uuid string) bool {
	_, ok := m.tasks[uuid]
	return ok
}

func (m *Manager) insertTask(uuid string) bool {
	if m.IsExistTask(uuid) {
		log.Logger.Debug("fail to insert task. task already exist", zap.String("uuid", uuid))
		return false
	}
	m.tasks[uuid] = NewTask()
	log.Logger.Debug("task inserted", zap.String("uuid", uuid))
	return true
}

func (m *Manager) RemoveTask(uuid string) bool {
	if m.IsExistTask(uuid) == false {
		log.Logger.Debug("fail to remove task. task not exist", zap.String("uuid", uuid))
		return false
	}
	delete(m.tasks, uuid)
	log.Logger.Debug("task removed", zap.String("uuid", uuid))
	return true
}

func (m *Manager) AddSubscriber(uuid string, client *entity.Client) string {
	clientKey := m.tasks[uuid].AddSubscriber(client)
	log.Logger.Debug("subscriber added", zap.String("uuid", uuid), zap.String("clientKey", clientKey))
	return clientKey
}

func (m *Manager) RemoveSubscriber(uuid string, key string) bool {
	ok := m.tasks[uuid].RemoveSubscriber(key)
	if !ok {
		log.Logger.Debug("fail to remove subscriber. subscriber not exist", zap.String("uuid", uuid), zap.String("clientKey", key))
	} else {
		log.Logger.Debug("subscriber removed", zap.String("uuid", uuid), zap.String("clientKey", key))
	}
	return ok
}

func (m *Manager) Notify(uuid string, stat int32, accTok string, buy bool) {
	log.Logger.Debug("notify message to subscribers", zap.String("uuid", uuid), zap.String("msg", msg.String(msg.NewMsg(stat, accTok))))
	m.tasks[uuid].Notify(stat, accTok, buy)
}

// for debug
func (m *Manager) GetSizeOfTask() int {
	return len(m.tasks)
}

func (m *Manager) RemoveTaskAll() {
	m.tasks = make(map[string]*Task)
}
