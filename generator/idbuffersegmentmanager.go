package generator

import (
	"log"
	"seeder/bootstrap"
	"sync"
)

type IDBufferSegmentManager struct {
	bizTag string

	lock    sync.RWMutex
	tagPool map[string]*IDBufferSegment

	application *bootstrap.Application
}

func (manager *IDBufferSegmentManager) GetId(bizTag string) (id uint64, e error) {
	manager.lock.RLock()
	segment := manager.GetSegmentByBizTag(bizTag)
	manager.lock.RUnlock()

	if segment == nil {
		manager.lock.Lock()
		segment = manager.GetSegmentByBizTag(bizTag)
		if segment == nil {
			segment = manager.CreateBizTagSegment(bizTag)
			if segment == nil {
				log.Fatal("segment nil")
			}
			manager.AddSegmentToPool(bizTag, segment)
		}
		manager.lock.Unlock()
	}
	id = segment.GetId()
	return id, nil
}
func (manager *IDBufferSegmentManager) AddSegmentToPool(bizTag string, segment *IDBufferSegment) {
	manager.tagPool[bizTag] = segment
}

func (manager *IDBufferSegmentManager) GetSegmentByBizTag(bizTag string) *IDBufferSegment {
	return manager.tagPool[bizTag]
}

func (manager *IDBufferSegmentManager) CreateBizTagSegment(bizTag string) *IDBufferSegment {

	segment := NewIDBufferSegment(bizTag, manager.application)

	manager.application.GetLogger().Debug("Manger  Segment  CreateMasterIDBuffer ")

	segment.CreateMasterIDBuffer(bizTag)

	go func() {
		// monitor := NewMonitor(segment, manager.application)
		// for {
		// 	time.Sleep(time.Millisecond * 100)
		// 	manager.application.GetLogger().Debug("NewMonitor timer ", bizTag, "Vigilant", manager.application.GetConfig().Monitior.VigilantValue)
		// 	monitor.SetVigilantValue(manager.application.GetConfig().Monitior.VigilantValue)
		// 	vigilant := monitor.IsOutVigilantValue()
		// 	if vigilant {
		// 		manager.application.GetLogger().Debug(" Over call CreateSlaveIDBuffer ", bizTag)
		// 		segment.CreateSlaveIDBuffer(bizTag)
		// 		segment.GetMasterIdBuffer().GetStats().Clear()
		// 	}

		// }
	}()
	return segment

}

func NewIDBufferSegmentManager(application *bootstrap.Application) *IDBufferSegmentManager {

	manager := &IDBufferSegmentManager{application: application, tagPool: make(map[string]*IDBufferSegment)}
	return manager
}
