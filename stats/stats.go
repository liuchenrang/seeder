package stats

import (
	"sync"
	"sync/atomic"
)

type Stats struct {
	muTotal sync.RWMutex
	total   uint64
	Stop    bool
}

//已分配数目
func (stats *Stats) GetTotal() uint64 {
	stats.muTotal.RLock()
	defer stats.muTotal.RUnlock()
	return stats.total
}

//分配计数
func (stats *Stats) Dig() {
	stats.muTotal.Lock()
	defer stats.muTotal.Unlock()
	pint := &stats.total
	atomic.AddUint64(pint, 1)
}

func (stats *Stats) DoStop() {
	stats.Stop = true
}

//清空计数
func (stats *Stats) Clear() {
	stats.total = 0
}
