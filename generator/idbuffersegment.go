package generator

import (
	"sync"
	"seeder/config"
	"seeder/logger"
	"seeder/bootstrap"
)

type IDBufferSegment struct {
	changeLock     sync.Mutex
	masterIDBuffer *IDBuffer
	slaveIdBuffer  *IDBuffer
	bizTag         string
	config config.SeederConfig
	
	SeederLogger.Logger
	application *bootstrap.Application
}

func (segment *IDBufferSegment) GetId() uint64 {
	segment.application.GetLogger().Debug( " segment nil ", segment == nil)
	idBuf := segment.masterIDBuffer

	id := idBuf.GetId();

	return id
}

func (segment *IDBufferSegment) CreateMasterIDBuffer(bizTag string) *IDBuffer {
	segment.changeLock.Lock()
	defer segment.changeLock.Unlock()

	segment.masterIDBuffer = NewIDBuffer(bizTag, segment.application)
	flushDB := make(chan string)
	go func() {
		segment.masterIDBuffer.Flush(flushDB)
	}()
	segment.application.GetLogger().Debug(" Segment CreateMasterIDBuffer ",segment.masterIDBuffer)
	return segment.masterIDBuffer
}
func (segment *IDBufferSegment) CreateSlaveIDBuffer(bizTag string) *IDBuffer {
	segment.slaveIdBuffer = NewIDBuffer(bizTag, segment.application)
	return segment.slaveIdBuffer
}
func (segment *IDBufferSegment) SetBizTag(bizTag string) {
	segment.bizTag = bizTag
}
func (segment *IDBufferSegment) GetBizTag() string {
	return segment.bizTag
}
func (segment *IDBufferSegment) GetMasterIdBuffer() *IDBuffer {
	return segment.masterIDBuffer
}
func (segment *IDBufferSegment) GetSlaveIdBuffer() *IDBuffer {
	return segment.slaveIdBuffer
}


func (segment *IDBufferSegment) ChangeSlaveToMaster() {
	segment.application.GetLogger().Debug(segment.bizTag + " changeSlaveToMaster")
	segment.changeLock.Lock()
	if segment.slaveIdBuffer == nil {
		segment.CreateSlaveIDBuffer(segment.bizTag)

	}
	flushDB := make(chan string)
	go func() {
		segment.masterIDBuffer.Flush(flushDB)
	}()
	<-flushDB
	segment.masterIDBuffer = segment.slaveIdBuffer

	segment.slaveIdBuffer = NewIDBuffer(segment.bizTag, segment.application)
	segment.changeLock.Unlock()
}


func NewIDBufferSegment(bizTag string,   application *bootstrap.Application) (*IDBufferSegment) {

	segment := &IDBufferSegment{application: application}

	segment.SetBizTag(bizTag)
	return segment
}