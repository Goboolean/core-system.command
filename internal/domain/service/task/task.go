package task

import (
	"github.com/Goboolean/command-server/internal/domain/entity"
	"github.com/Goboolean/command-server/internal/util/msg"
	"github.com/google/uuid"
)

type Task struct {
	subscribers map[string]*entity.Client
}

func NewTask() *Task {
	return &Task{
		subscribers: make(map[string]*entity.Client),
	}
}

func (t *Task) AddSubscriber(client *entity.Client) string {
	key := uuid.New().String()
	t.subscribers[key] = client
	return key
}

func (t *Task) RemoveSubscriber(key string) bool {
	if _, ok := t.subscribers[key]; !ok {
		return false
	}
	delete(t.subscribers, key)
	return true
}

func (t *Task) Notify(stat int32, accTok string, buy bool) {
	for _, sub := range t.subscribers {
		// need check is channel closed
		go sub.Send(msg.NewMsg(stat, accTok), buy)
	}
}

// for debug
func (t *Task) GetSizeOfSubscribers() int {
	return len(t.subscribers)
}
