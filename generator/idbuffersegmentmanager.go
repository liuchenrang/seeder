package generator

import (
	"sync"
	"log"
	"seeder/bootstrap"
	"time"
)


type IDBufferSegmentManager struct {
	bizTag  string

	lock *sync.Mutex
	tagPool map[string] *IDBufferSegment

	application *bootstrap.Application
}

func (manager *IDBufferSegmentManager) GetId(bizTag string) uint64 {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	segment := manager.getSegmentByBizTag(bizTag)
	if  segment == nil{
		log.Fatal("bizTag " , bizTag, " not create")
	}
	var id uint64;


	for  {
		id = segment.GetId()
		if id <= 0 {
			segment.ChangeSlaveToMaster()
			manager.application.GetLogger().Debug("ChangeSlaveToMaster ", id)

			id = segment.GetId()
			manager.application.GetLogger().Debug("ChangeSlaveToMasterId ", id)
		}else{
			break
		}
	}

	return id
}
func (manager *IDBufferSegmentManager) getSegmentByBizTag(bizTag string)  *IDBufferSegment {
	_, has := manager.tagPool[bizTag]
	if !has  {
		return manager.CreateBizTagSegment(bizTag)
	}
	return manager.tagPool[bizTag]
}

func (manager *IDBufferSegmentManager) CreateBizTagSegment(bizTag string) *IDBufferSegment {

	_, has := manager.tagPool[bizTag]
	manager.application.GetLogger().Debug("init ", bizTag, has)

	if  has == false {

		segment := NewIDBufferSegment(bizTag, manager.application)
		manager.application.GetLogger().Debug("Manger  Segment  CreateMasterIDBuffer ")

		segment.CreateMasterIDBuffer(bizTag)

		manager.tagPool[bizTag] = segment
		go func() {
			for {
				time.Sleep(time.Millisecond*100)
				manager.application.GetLogger().Debug("NewMonitor timer ", bizTag)

				monitor := NewMonitor(segment,  manager.application)
				monitor.SetVigilantValue(5)
				vigilant := monitor.IsOutVigilantValue()
				if vigilant {
					manager.application.GetLogger().Debug(" Over call CreateSlaveIDBuffer ", bizTag)
					segment.CreateSlaveIDBuffer(bizTag)
					segment.GetMasterIdBuffer().GetStats().Clear()
				}

			}
		}()
	}
	return  manager.tagPool[bizTag]

}

func NewIDBufferSegmentManager(application *bootstrap.Application) *IDBufferSegmentManager {

	manager := &IDBufferSegmentManager{application: application, tagPool: make(map[string] *IDBufferSegment), lock: &sync.Mutex{}}
	return manager
}
