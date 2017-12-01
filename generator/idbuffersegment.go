package generator

import (
	"sync"
	"seeder/logger"
)
var (
	logger logger.Logger{}
)
type IDBufferSegment struct {
	changeLock sync.Mutex
	masterIDBuffer *IDBuffer
	slaveIdBuffer *IDBuffer
	bizTag string
}

func (segment *IDBufferSegment) GetId() uint64  {
	idBuf := segment.masterIDBuffer
	return idBuf.GetId();
}


func (segment *IDBufferSegment) CreateMasterIDBuffer(bizTag string)  *IDBuffer {
	segment.masterIDBuffer = NewIDBuffer(bizTag)
	return segment.masterIDBuffer
}
func (segment *IDBufferSegment) CreateSlaveIDBuffer(bizTag string)  *IDBuffer {
	segment.slaveIdBuffer = NewIDBuffer(bizTag)
	return segment.slaveIdBuffer
}
func (segment *IDBufferSegment) SetBizTag(bizTag string)   {
	segment.bizTag = bizTag
}
func (segment *IDBufferSegment) GetMasterIdBuffer() *IDBuffer   {
	return segment.masterIDBuffer
}
func NewIDBufferSegment(bizTag string) (*IDBufferSegment) {
	segment :=  &IDBufferSegment{}
	segment.SetBizTag(bizTag)
	return segment
}
func NewSegmentBizTag(bizTag string) (*IDBufferSegment) {
	segment :=  NewIDBufferSegment(bizTag)
	return segment
}
func (segment *IDBufferSegment) ChangeSlaveToMaster()  {
	logger.Debug(segment.bizTag + " changeSlaveToMaster")
	segment.changeLock.Lock()
	segment.masterIDBuffer = segment.slaveIdBuffer
	segment.slaveIdBuffer = NewIDBuffer(segment.bizTag)
	defer segment.changeLock.Unlock()
}


