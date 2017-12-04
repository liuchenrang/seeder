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
	segment *IDBufferSegment
	tagPool map[string] *IDBufferSegment
}

func (manager *IDBufferSegmentManager) GetId(bizTag string) uint64 {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.CreateBizTagSegment()

	manager.segment = manager.tagPool[bizTag]
	if  manager.segment == nil{
		log.Fatal("bizTag " , bizTag, " not create")

	}
	return manager.segment.GetId()
}
func (manager *IDBufferSegmentManager) CreateBizTagSegment() *IDBufferSegment {


	_, has := manager.tagPool[manager.bizTag]
	logger.Debug("init ", manager.bizTag, has)

	if  has == false {
		manager.segment = NewIDBufferSegment(manager.bizTag, manager.config)
		manager.segment.CreateMasterIDBuffer(manager.bizTag)
		logger.Debug(" Segment Out CreateMasterIDBuffer ",manager.segment.masterIDBuffer.GetId())
		manager.tagPool[manager.bizTag] = manager.segment
		go func() {
			for {
				monitor := NewMonitor(manager.segment)
				monitor.SetVigilantValue(5)
				vigilant := monitor.IsOutVigilantValue()
				if vigilant {
					fmt.Println(" Over call CreateSlaveIDBuffer ", manager.bizTag)
					manager.segment.CreateSlaveIDBuffer(manager.bizTag)
					manager.segment.GetMasterIdBuffer().GetStats().Clear()
				}

			}
		}()
	}else{
		logger.Debug("load  ",has)
		manager.segment = manager.tagPool[manager.bizTag]
		manager.segment.CreateMasterIDBuffer(manager.bizTag)

	}

	return manager.segment
}

func NewIDBufferSegmentManager(bizTag string, config config.SeederConfig) *IDBufferSegmentManager {
	manager := &IDBufferSegmentManager{config: config, bizTag: bizTag, tagPool: make(map[string] *IDBufferSegment), lock: &sync.Mutex{}}

	//manager.CreateBizTagSegment()

	return manager
}
