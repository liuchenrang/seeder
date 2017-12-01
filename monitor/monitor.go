package monitor

import (
	"seeder/generator"
	"seeder/stats"
)

type Monitor struct {
	threshold uint64
	segment  *generator.IDBufferSegment
}

func (m *Monitor) SetVigilantValue(threshold uint64) {
	m.threshold = threshold
}
func (m *Monitor) IsOutVigilantValue() bool {
	i := m.segment.GetMasterIdBuffer().GetStats().GetTotal()
	return i >= m.threshold
}
func (m *Monitor) Event(tag <-chan string) {

}
func (m *Monitor) GetStats() *stats.Stats {
	return m.segment.GetMasterIdBuffer().GetStats()
}

func NewMonitor(seg *generator.IDBufferSegment) *Monitor {
	return &Monitor{segment: seg}
}
