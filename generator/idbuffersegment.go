package generator

import (
	"sync"
	"seeder/config"
	"seeder/logger"
)

type IDBufferSegment struct {
	changeLock     sync.Mutex
	masterIDBuffer *IDBuffer
	slaveIdBuffer  *IDBuffer
	bizTag         string
	config config.SeederConfig
	
	SeederLogger.Logger
}

func (segment *IDBufferSegment) GetId() uint64 {
	segment.Debug( " segment nil ", segment == nil)
	idBuf := segment.masterIDBuffer

	return idBuf.GetId();
}

func (segment *IDBufferSegment) CreateMasterIDBuffer(bizTag string) *IDBuffer {
	segment.changeLock.Lock()
	defer segment.changeLock.Unlock()

	segment.masterIDBuffer = NewIDBuffer(bizTag, segment.config)
	flushDB := make(chan string)
	go func() {
		segment.masterIDBuffer.Flush(flushDB)
	}()
	segment.Debug(" Segment CreateMasterIDBuffer ",segment.masterIDBuffer)
	return segment.masterIDBuffer
}
func (segment *IDBufferSegment) CreateSlaveIDBuffer(bizTag string) *IDBuffer {
	segment.slaveIdBuffer = NewIDBuffer(bizTag, segment.config)
	return segment.slaveIdBuffer
}
func (segment *IDBufferSegment) SetBizTag(bizTag string) {
	segment.bizTag = bizTag
}
func (segment *IDBufferSegment) GetMasterIdBuffer() *IDBuffer {
	return segment.masterIDBuffer
}
func (segment *IDBufferSegment) GetSlaveIdBuffer() *IDBuffer {
	return segment.slaveIdBuffer
}
func NewIDBufferSegment(bizTag string,  config config.SeederConfig) (*IDBufferSegment) {

	segment := &IDBufferSegment{config: config}

	segment.SetBizTag(bizTag)
	return segment
}

func (segment *IDBufferSegment) ChangeSlaveToMaster() {
	segment.Debug(segment.bizTag + " changeSlaveToMaster")
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
	segment.slaveIdBuffer = NewIDBuffer(segment.bizTag, segment.config)
	segment.changeLock.Unlock()
}
