package generator

import (
	"log"
	"seeder/bootstrap"
	"sync"
	"fmt"
)

type IDBufferSegmentManager struct {


	bizTag string

	muCreate sync.Mutex
	muTagPool sync.Mutex
	tagPool   map[string]*IDBufferSegment

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
	seg, has := manager.GetSegmentFromPool(bizTag)
	if !has {
		manager.muCreate.Lock()
		defer manager.muCreate.Unlock()
		seg, has = manager.GetSegmentFromPool(bizTag)
		if !has {
			seg = manager.CreateBizTagSegment(bizTag)
			if seg == nil {
				log.Fatal("segment nil")
			}
		}

	}
	manager.application.GetLogger().Error(fmt.Sprintf("manager segment %p", seg))

	return seg

}
func (manager *IDBufferSegmentManager) GetSegmentFromPool(bizTag string) (seg *IDBufferSegment, has bool) {
	manager.muTagPool.Lock()
	defer manager.muTagPool.Unlock()
	seg, has = manager.tagPool[bizTag]
	return
}
func (manager *IDBufferSegmentManager) SegmentManager(bizTag string, seg chan *IDBufferSegment) {
	seg <- manager.CreateBizTagSegment(bizTag)
}
func (manager *IDBufferSegmentManager) CreateBizTagSegment(bizTag string) *IDBufferSegment {

	segment := NewIDBufferSegment(bizTag, manager.application)

	manager.application.GetLogger().Debug("Manger  Segment  CreateMasterIDBuffer ")

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
