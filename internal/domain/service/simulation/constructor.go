/*
Package simulation

Manager 객체는 simulation을 관리하는 객체입니다.

simulation 할당을 명령하며, 할당된 simulation을 관리합니다.
또한 진행 중인 simulation 현황을 관리하며,
simulation의 결과를 관리합니다.
*/
package simulation

import (
	"github.com/Goboolean/core-system.command/internal/domain/port/out"
	"github.com/Goboolean/core-system.command/internal/domain/service/metadata"
	"github.com/Goboolean/core-system.command/internal/domain/service/model"
	"github.com/Goboolean/core-system.command/internal/domain/service/task"
	"sync"
)

type Manager struct {
	model    *model.Manager
	metadata *metadata.Manager
	task     *task.Manager
	db       out.SimulationDBPort
}

var (
	instance *Manager
	once     sync.Once
)

type checkInfo struct {
	NAME  string
	EXIST bool
	ERR   error
}

func New(model *model.Manager, metadata *metadata.Manager, task *task.Manager, db out.SimulationDBPort) *Manager {

	once.Do(func() {
		instance = &Manager{
			model:    model,
			metadata: metadata,
			task:     task,
			db:       db,
		}
	})

	return instance
}
