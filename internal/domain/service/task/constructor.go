/*
Package task

Manager 객체는 현재 진행되는 작업을 관리하는 객체입니다.
*/
package task

import (
	"github.com/Goboolean/core-system.command/internal/domain/port/out"
	"go.uber.org/zap"
	"sync"
)

type Manager struct {
	tasks map[string]*Task
	kafka out.TaskPort
}

var (
	instance *Manager
	once     sync.Once
	logger   *zap.Logger
)

func New(kafka out.TaskPort) *Manager {

	once.Do(func() {
		instance = &Manager{
			tasks: make(map[string]*Task),
			kafka: kafka,
		}
	})

	return instance
}
