package generator

import (
	"seeder/stats"
	"seeder/bootstrap"
)

type Monitor struct {
	threshold uint64
	segment  *IDBufferSegment
	application *bootstrap.Application
}

func (m *Monitor) SetVigilantValue(threshold uint64) {
	m.threshold = threshold
}
func (m *Monitor) IsOutVigilantValue() bool {
	i := m.segment.GetMasterIdBuffer().GetStats().GetTotal()
	m.application.GetLogger().Debug(m.segment.GetBizTag() , "total", i)

	return i >= m.threshold
}
func (m *Monitor) Event(tag <-chan string) {

}
func (m *Monitor) GetStats() *stats.Stats {
	return m.segment.GetMasterIdBuffer().GetStats()
}

func NewMonitor(seg *IDBufferSegment, application *bootstrap.Application) *Monitor {
	return &Monitor{segment: seg, application:application}
}
