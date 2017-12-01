package stats

import "sync/atomic"

type Stats struct {
	total uint64
}

//已分配数目
func (stats *Stats) GetTotal() uint64 {
	return stats.total
}

//分配计数
func (stats *Stats) Dig() {
	pint := &stats.total
	atomic.AddUint64(pint, 1)
}

//清空计数
func (stats *Stats) Clear() {
	stats.total = 0
}
