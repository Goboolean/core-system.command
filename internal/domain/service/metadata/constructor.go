/*
Package metadata

Manager 객체는 metadata를 관리하는 객체입니다.
*/
package metadata

import (
	"github.com/Goboolean/command-server/internal/domain/port/out"
	"sync"
)

type Manager struct {
	db out.MetadataDBPort
}

var (
	instance *Manager
	once     sync.Once
)

func New(db out.MetadataDBPort) *Manager {

	once.Do(func() {
		instance = &Manager{
			db: db,
		}
	})

	return instance
}
