package stats

import "sync/atomic"

type Stats struct{
	total  int32
}
func (stats *Stats) GetTotal() int32{
	return stats.total
}
func (stats *Stats) Dig(){
	pint := &stats.total;
	atomic.AddInt32(pint,1)
}
func (stats *Stats) Clear(){
	stats.total = 0
}

