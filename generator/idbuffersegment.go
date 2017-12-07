package generator

import (
	"seeder/bootstrap"
	"seeder/config"
	"seeder/logger"
	"sync"
)

type IDBufferSegment struct {
	changeLock     sync.Mutex
	masterIDBuffer *IDBuffer
	slaveIdBuffer  *IDBuffer
	bizTag         string
	config         config.SeederConfig

	SeederLogger.Logger
	application *bootstrap.Application
}

func (segment *IDBufferSegment) GetId() (id uint64) {
	idBuf := segment.masterIDBuffer
	segment.application.GetLogger().Debug("SegmentAddress ", segment)
	for {
		id, _ = idBuf.GetId()
		if id <= 0 {
			segment.application.GetLogger().Debug(
				"UserOut",
				"currentid", segment.masterIDBuffer.GetCurrentId(),
				"isUserOut", segment.masterIDBuffer.IsUseOut(),
				"max", segment.masterIDBuffer.GetMaxId())

			if segment.IsMasterUserOut() {
				segment.application.GetLogger().Debug("ChangeSlaveToMaster", segment.IsMasterUserOut(), segment.masterIDBuffer.IsUseOut())
				segment.ChangeSlaveToMaster()
			} else {
				idBuf = segment.masterIDBuffer
				segment.application.GetLogger().Debug("IsMasterUserOut 0 ")
			}
		} else {
			break
		}
	}
	return id
}
func (segment *IDBufferSegment) IsMasterUserOut() bool {
	return segment.masterIDBuffer.IsUseOut()
}
func (segment *IDBufferSegment) CreateMasterIDBuffer(bizTag string) *IDBuffer {
	segment.changeLock.Lock()
	defer segment.changeLock.Unlock()
	segment.masterIDBuffer = NewIDBuffer(bizTag, segment.application)
	segment.masterIDBuffer.Flush()
	return segment.masterIDBuffer
}
func (segment *IDBufferSegment) CreateSlaveIDBuffer(bizTag string) *IDBuffer {
	segment.slaveIdBuffer = NewIDBuffer(bizTag, segment.application)
	segment.slaveIdBuffer.Flush()
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
	segment.changeLock.Lock()
	defer segment.changeLock.Unlock()
	segment.application.GetLogger().Debug(segment.bizTag + " changeSlaveToMaster")
	if segment.slaveIdBuffer == nil {
		segment.CreateSlaveIDBuffer(segment.bizTag)
	}
	segment.masterIDBuffer = segment.slaveIdBuffer
	segment.slaveIdBuffer = NewIDBuffer(segment.bizTag, segment.application)
}

func NewIDBufferSegment(bizTag string, application *bootstrap.Application) *IDBufferSegment {
	segment := &IDBufferSegment{application: application}
	segment.SetBizTag(bizTag)
	return segment
}
