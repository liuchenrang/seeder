package generator

import (
	"seeder/monitor"
	"seeder/pool"
)

type IDBufferSegmentManager struct{
		pool pool.Pool
		segment *IDBufferSegment
}
func (manager *IDBufferSegmentManager) GetId(bizTag string) uint64{
	return manager.segment.GetId()
}
func NewIDBufferSegmentManager(bizTag string) *IDBufferSegmentManager{
	segment := NewIDBufferSegment(bizTag)
	segment.CreateMasterIDBuffer(bizTag)
	go func(){
		for{
			monitor := monitor.NewMonitor(segment)
			monitor.SetVigilantValue(200)
			vigilant := monitor.IsOutVigilantValue()
			if vigilant {
				segment.CreateSlaveIDBuffer(bizTag)
			}
		}
	}()
	return &IDBufferSegmentManager{segment:segment}
}
