/*
Package model

ModelManager 객체는 Model에 대한 비지니스 로직을 담당합니다.
특정 모델에 대한 Simulation 요청이 들어온 경우, 서비스 로직을 실행한 이후
simulation manager에게 작업 할당 책임을 전달합니다.
*/
package model

import (
	"github.com/Goboolean/command-server/internal/domain/port/out"
	"sync"
)

type Manager struct {
	db out.ModelDBPort
}

var (
	instance *Manager
	once     sync.Once
)

func New(db out.ModelDBPort) *Manager {

	once.Do(func() {
		instance = &Manager{
			db: db,
		}
	})

	return instance
}
