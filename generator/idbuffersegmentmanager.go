package generator

import (
	"log"
	"seeder/bootstrap"
	"sync"
)

type IDBufferSegmentManager struct {
	bizTag string

	lock    sync.Mutex
	tagPool map[string]*IDBufferSegment

	application *bootstrap.Application
}

func (manager *IDBufferSegmentManager) GetId(bizTag string) (id uint64, e error) {
	segment := manager.GetSegmentByBizTag(bizTag)

	id = segment.GetId()
	return id, nil
}
func (manager *IDBufferSegmentManager) AddSegmentToPool(bizTag string, segment *IDBufferSegment) {
	manager.tagPool[bizTag] = segment
}

func (manager *IDBufferSegmentManager) GetSegmentByBizTag(bizTag string) *IDBufferSegment {
	_, has := manager.tagPool[bizTag]

	if !has {
		// 仅有一个人可以申请创建BizTagSegment
		manager.lock.Lock()
		defer manager.lock.Unlock()

		_, has = manager.tagPool[bizTag]
		if !has {
			segment := manager.CreateBizTagSegment(bizTag)
			if segment == nil {
				log.Fatal("segment nil")
			}
			manager.AddSegmentToPool(bizTag, segment)
		}

	}
	return manager.tagPool[bizTag]

}

func (manager *IDBufferSegmentManager) CreateBizTagSegment(bizTag string) *IDBufferSegment {

	segment := NewIDBufferSegment(bizTag, manager.application)

	manager.application.GetLogger().Debug("Manger  Segment  CreateMasterIDBuffer ")

	//go func() {
	//	monitor := NewMonitor(segment, manager.application)
	//	for {
	//		time.Sleep(time.Millisecond * 100)
	//		vigilanValue := manager.application.GetConfig().Monitior.VigilantValue
	//		manager.application.GetLogger().Debug("NewMonitor timer ", bizTag, "Vigilant", vigilanValue)
	//		if vigilanValue <= 100 {
	//			monitor.SetVigilantValue(vigilanValue)
	//			vigilant := monitor.IsOutVigilantValue()
	//			if vigilant && !segment.GetMasterIdBuffer().GetStats().Stop{
	//				manager.application.GetLogger().Debug(" Over call CreateSlaveIDBuffer ", bizTag)
	//				segment.CreateSlaveIDBuffer(bizTag)
	//				segment.GetMasterIdBuffer().GetStats().Stop = true
	//			}
	//		}
	//
	//	}
	//}()
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
