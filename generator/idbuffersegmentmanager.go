package generator

import (
	"log"
	"seeder/bootstrap"
	"sync"
	"time"
)

type IDBufferSegmentManager struct {
	bizTag string

	muTagPool    sync.Mutex
	tagPool map[string]*IDBufferSegment

	application *bootstrap.Application
}

func (manager *IDBufferSegmentManager) GetId(bizTag string) (id uint64, e error) {
	segment := manager.GetSegmentByBizTag(bizTag)

	id = segment.GetId()
	return id, nil
}
func (manager *IDBufferSegmentManager) AddSegmentToPool(bizTag string, segment *IDBufferSegment) {
	manager.muTagPool.Lock()
	defer manager.muTagPool.Unlock()
	manager.tagPool[bizTag] = segment
}

func (manager *IDBufferSegmentManager) GetSegmentByBizTag(bizTag string) *IDBufferSegment {
	seg , has := manager.GetSegmentFromPool(bizTag)

	if !has {
			seg = manager.CreateBizTagSegment(bizTag)
			if seg == nil {
				log.Fatal("segment nil")
			}
	}
	return seg

}
func (manager *IDBufferSegmentManager) GetSegmentFromPool(bizTag string)(seg *IDBufferSegment ,has bool) {
	manager.muTagPool.Lock()
	defer manager.muTagPool.Unlock()
	seg , has = manager.tagPool[bizTag]
	return
}
func (manager *IDBufferSegmentManager) SegmentManager(bizTag  string, seg chan *IDBufferSegment){
	seg <- manager.CreateBizTagSegment(bizTag)
}
func (manager *IDBufferSegmentManager) CreateBizTagSegment(bizTag string) *IDBufferSegment {

	segment := NewIDBufferSegment(bizTag, manager.application)

	manager.application.GetLogger().Debug("Manger  Segment  CreateMasterIDBuffer ")

	go func() {
		monitor := NewMonitor(segment, manager.application)
		for {
			time.Sleep(time.Millisecond * 100)
			vigilanValue := manager.application.GetConfig().Monitior.VigilantValue
			manager.application.GetLogger().Debug("NewMonitor timer ", bizTag, "Vigilant", vigilanValue)
			if vigilanValue <= 100 {
				monitor.SetVigilantValue(vigilanValue)
				vigilant := monitor.IsOutVigilantValue()
				if vigilant && !segment.GetMasterIdBuffer().GetStats().Stop{
					manager.application.GetLogger().Debug(" Over call CreateSlaveIDBuffer ", bizTag)
					segment.CreateSlaveIDBuffer(bizTag)
					segment.GetMasterIdBuffer().GetStats().Stop = true
				}
			}

		}
	}()
	manager.AddSegmentToPool(bizTag, segment)

	return segment

}
func (manager *IDBufferSegmentManager) Stop() {
	for _, segment := range manager.tagPool {
		if segment.masterIDBuffer != nil {
			segment.Close()
		}
	}
}

func NewIDBufferSegmentManager(application *bootstrap.Application) *IDBufferSegmentManager {

	manager := &IDBufferSegmentManager{application: application, tagPool: make(map[string]*IDBufferSegment)}
	return manager
}
