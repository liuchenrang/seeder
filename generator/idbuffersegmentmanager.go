package generator

import (
	"seeder/logger"
	"seeder/config"
	"sync"
	"fmt"
	"log"
)

var logger SeederLogger.Logger

type IDBufferSegmentManager struct {
	bizTag  string
	config  config.SeederConfig

	lock *sync.Mutex
	tagPool map[string] *IDBufferSegment
}

func (manager *IDBufferSegmentManager) GetId(bizTag string) uint64 {
	bizTag = bizTag
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.CreateBizTagSegment(bizTag)
	segment := manager.tagPool[bizTag]
	if  segment == nil{
		log.Fatal("bizTag " , bizTag, " not create")
	}
	return segment.GetId()
}
func (manager *IDBufferSegmentManager)getSegmentByBizTag(bizTag string)  *IDBufferSegment {
	_, has := manager.tagPool[bizTag]
	logger.Debug("init ", bizTag, has)
	if !has  {
		return manager.CreateBizTagSegment(bizTag)
	}
	return manager.tagPool[bizTag]
}
func (manager *IDBufferSegmentManager) CreateBizTagSegment(bizTag string) *IDBufferSegment {

	_, has := manager.tagPool[bizTag]
	logger.Debug("init ", bizTag, has)

	if  has == false {
		segment := NewIDBufferSegment(bizTag, manager.config)
		segment.CreateMasterIDBuffer(bizTag)
		logger.Debug(" Segment Out CreateMasterIDBuffer ",segment.masterIDBuffer.GetId())
		manager.tagPool[bizTag] = segment
		go func() {
			for {
				monitor := NewMonitor(segment)
				monitor.SetVigilantValue(5)
				vigilant := monitor.IsOutVigilantValue()
				if vigilant {
					fmt.Println(" Over call CreateSlaveIDBuffer ", bizTag)
					segment.CreateSlaveIDBuffer(bizTag)
					segment.GetMasterIdBuffer().GetStats().Clear()
				}

			}
		}()
	}
	return  manager.tagPool[bizTag]

}

func NewIDBufferSegmentManager(config config.SeederConfig) *IDBufferSegmentManager {
	manager := &IDBufferSegmentManager{config: config, tagPool: make(map[string] *IDBufferSegment), lock: &sync.Mutex{}}
	return manager
}
