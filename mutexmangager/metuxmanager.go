package mutexmangager

import (
	"sync"
)

type MutexManager struct {
	muRw     sync.RWMutex
	muRwPool map[string]sync.RWMutex
	muPool   map[string]sync.Mutex
}

func (mm *MutexManager) Get(string) sync.Mutex {
	return sync.Mutex{}
}
func (mm *MutexManager) GetRw(string) {
	
}

func NewMutexManager() *MutexManager {
	return &MutexManager{}
}
