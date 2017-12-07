package generator

import (
	"seeder/bootstrap"
	"sync"
)

type IDBufferSegment struct {
	mu     sync.Mutex
	masterIDBuffer *IDBuffer
	slaveIdBuffer  *IDBuffer

	bizTag         string
	application *bootstrap.Application
}

func (segment *IDBufferSegment) GetId() (id uint64) {
	var idBuffer *IDBuffer
	for {

		idBuffer = segment.GetMasterIdBuffer()
		id, _ = idBuffer.GetId()
		if id <= 0  {

			//debug := fmt.Sprint(
			//	" BeforeMasterIDBufferInfo",
			//	" currentid", segment.masterIDBuffer.GetCurrentId(),
			//	" isUserOut", segment.masterIDBuffer.IsUseOut(),
			//	" max", segment.masterIDBuffer.GetMaxId())
			//debug += fmt.Sprint(
			//	" slaveIdBufferInfo",
			//	" currentid", segment.slaveIdBuffer.GetCurrentId(),
			//	" isUserOut", segment.slaveIdBuffer.IsUseOut(),
			//	" max", segment.slaveIdBuffer.GetMaxId())
			//
			////segment.masterIDBuffer = segment.slaveIdBuffer
			//
			//debug += fmt.Sprint(
			//	" AfterMasterIDBufferInfo",
			//	" currentid", segment.masterIDBuffer.GetCurrentId(),
			//	" isUserOut", segment.masterIDBuffer.IsUseOut(),
			//	" max", segment.masterIDBuffer.GetMaxId())
			//debug += fmt.Sprint(
			//	" slaveIdBufferInfo",
			//	" currentid", segment.slaveIdBuffer.GetCurrentId(),
			//	" isUserOut", segment.slaveIdBuffer.IsUseOut(),
			//	" max", segment.slaveIdBuffer.GetMaxId())
			//segment.application.GetLogger().Debug(debug)
			//
			//continue

			//if segment.IsMasterUserOut() {
			//	segment.application.GetLogger().Debug("ChangeSlaveToMaster", segment.IsMasterUserOut(), segment.masterIDBuffer.IsUseOut())
				segment.ChangeSlaveToMaster()
			//
			//} else {
				segment.application.GetLogger().Debug("IsMasterUserOut 0 ")
			//}

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
	segment.masterIDBuffer = NewIDBuffer(bizTag, segment.application)
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
	segment.mu.Lock()
	defer segment.mu.Unlock()
	segment.application.GetLogger().Debug(segment.bizTag + " changeSlaveToMaster")
	if segment.IsMasterUserOut() {
		if segment.slaveIdBuffer == nil {
			segment.CreateSlaveIDBuffer(segment.bizTag)
		} else {
			segment.slaveIdBuffer = NewIDBuffer(segment.bizTag, segment.application)
		}
		segment.masterIDBuffer = segment.slaveIdBuffer
	}
}

func NewIDBufferSegment(bizTag string, application *bootstrap.Application) *IDBufferSegment {
	segment := &IDBufferSegment{application: application}
	segment.SetBizTag(bizTag)
	segment.CreateMasterIDBuffer(segment.bizTag)
	return segment
}
