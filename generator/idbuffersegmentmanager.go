package generator

import (
	"seeder/logger"
	"fmt"
	"seeder/config"
)
var logger SeederLogger.Logger

type IDBufferSegmentManager struct{
		segment *IDBufferSegment
}
func (manager *IDBufferSegmentManager) GetId(bizTag string) uint64{
	return manager.segment.GetId()
}
func NewIDBufferSegmentManager(bizTag string, config config.SeederConfig) *IDBufferSegmentManager{
	segment := NewIDBufferSegment(bizTag, config)
	segment.CreateMasterIDBuffer(bizTag)
	go func(){
		for{
			monitor := NewMonitor(segment)
			monitor.SetVigilantValue(5)
			vigilant := monitor.IsOutVigilantValue()
			if vigilant {
				fmt.Println(" Over call CreateSlaveIDBuffer ",bizTag)
				segment.CreateSlaveIDBuffer(bizTag)
				segment.GetMasterIdBuffer().GetStats().Clear()
			}

		}
	}()
	return &IDBufferSegmentManager{segment:segment}
}
