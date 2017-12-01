package stats

import "sync/atomic"

type Stats struct {
	total int32
}

//已分配数目
func (stats *Stats) GetTotal() int32 {
	return stats.total
}

//分配计数
func (stats *Stats) Dig() {
	pint := &stats.total
	atomic.AddInt32(pint, 1)
}

//清空计数
func (stats *Stats) Clear() {
	stats.total = 0
}
