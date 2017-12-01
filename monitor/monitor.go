package monitor

import (
	"../stats"
	"../generator"
)

type Monitor struct {
	threshold uint64
	segment  *segment
}

func (m *Monitor) SetVigilantValue(threshold: uint64) {
	m.threshold = threshold
}
func (m *Monitor) IsOutVigilantValue() {
	return m.segment.masterIDBuffer.total >= m.threshold
}
func (m *Monitor) Event(tag <-chan string) {

}
func (m *Monitor) GetStats() *Stats {
	return m.segment.masterIDBuffer.stats
}

func NewMonitor(seg *IDBufferSegment) *Monitor {
	return &Monitor{segment: seg}
}
