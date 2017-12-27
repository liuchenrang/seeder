package generator

import (
	"seeder/bootstrap"
	"seeder/stats"
)

type Monitor struct {
	threshold   uint8
	segment     *IDBufferSegment
	application *bootstrap.Application
}

func (m *Monitor) SetVigilantValue(threshold uint8) {
	m.threshold = threshold
}
func (m *Monitor) IsOutVigilantValue() bool {
	idBuffer := m.segment.GetMasterIdBuffer()
	total := idBuffer.GetCacheStep()
	useTotal := idBuffer.GetStats().GetTotal()
	var usePercent uint64
	if total > 0 {
		usePercent = (useTotal * 100 / total * 100) / 100
		m.application.GetLogger().Info(m.segment.GetBizTag(), " usePercent ", usePercent, "useTotal", useTotal, "total Step", total)
		m.application.GetLogger().Info("tag %s ,usePercent %d useTotal %d  total Step %d",m.segment.GetBizTag(), usePercent, useTotal, total)
	}
	return uint8(usePercent) >= m.threshold
}
func (m *Monitor) Event(tag <-chan string) {

}
func (m *Monitor) GetStats() *stats.Stats {
	return m.segment.GetMasterIdBuffer().GetStats()
}

func NewMonitor(seg *IDBufferSegment, application *bootstrap.Application) *Monitor {
	return &Monitor{segment: seg, application: application}
}
