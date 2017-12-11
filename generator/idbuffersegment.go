package generator

import (
	"seeder/bootstrap"
	"sync"
)

var (
	SegmentBizTag        = make(chan string)
	SegmentSlaveChangeBizTag        = make(chan string)
	SegmentCreateMasterBizTag        = make(chan string)
	SegmentCreateSlaveBizTag        = make(chan string)
	SegmentBizTagIDBuffer = make(chan *IDBuffer)
)

type IDBufferSegment struct {
	mu             sync.Mutex
	muM           sync.Mutex
	masterIDBuffer *IDBuffer
	slaveIdBuffer  *IDBuffer

	bizTag      string
	application *bootstrap.Application
}

func (segment *IDBufferSegment) GetId() (id uint64) {

	var idBuffer *IDBuffer
	for {
		idBuffer = segment.GetMasterIdBuffer()
		id, _ = idBuffer.GetId()
		if id <= 0 {
			segment.ChangeSlaveToMaster()
			segment.application.GetLogger().Debug("IsMasterUserOut 0 ")
		} else {
			break
		}
	}
	return id
}
func (segment *IDBufferSegment) IsMasterUserOut() bool {
	segment.mu.Lock()
	defer segment.mu.Unlock()
	return segment.masterIDBuffer.IsUseOut()
}
func (segment *IDBufferSegment) CreateMasterIDBuffer(bizTag string) *IDBuffer {
	SegmentBizTag<-bizTag
	segment.masterIDBuffer = <-SegmentBizTagIDBuffer
	return segment.masterIDBuffer
}
func (segment *IDBufferSegment) CreateSlaveIDBuffer(bizTag string) *IDBuffer {
	SegmentBizTag<-segment.bizTag
	segment.slaveIdBuffer = <-SegmentBizTagIDBuffer
	return segment.slaveIdBuffer
}
func (segment *IDBufferSegment) SetBizTag(bizTag string) {
	segment.bizTag = bizTag
}
func (segment *IDBufferSegment) GetBizTag() string {
	return segment.bizTag
}
func (segment *IDBufferSegment) GetMasterIdBuffer() *IDBuffer {
	segment.mu.Lock()
	defer segment.mu.Unlock()
	if segment.masterIDBuffer != nil && !segment.masterIDBuffer.IsUseOut(){
		return segment.masterIDBuffer
	}
	SegmentBizTag<-segment.bizTag
	segment.masterIDBuffer = <-SegmentBizTagIDBuffer
	return segment.masterIDBuffer
}
func (segment *IDBufferSegment) GetSlaveIdBuffer() *IDBuffer {
	return segment.slaveIdBuffer
}

func (segment *IDBufferSegment) ChangeSlaveToMaster() {
	segment.application.GetLogger().Debug(segment.bizTag + " changeSlaveToMaster")
	if segment.IsMasterUserOut() {
		if segment.slaveIdBuffer == nil {
			segment.CreateSlaveIDBuffer(segment.bizTag)
		} else {
			if segment.slaveIdBuffer.IsUseOut() {
				segment.CreateSlaveIDBuffer(segment.bizTag)
			}
		}
		SegmentSlaveChangeBizTag<-segment.bizTag
	}
}
func (segment *IDBufferSegment) Close() {

	if segment.masterIDBuffer != nil {
		//segment.masterIDBuffer.Wg.Wait()
	}
	if segment.slaveIdBuffer != nil {
		//segment.slaveIdBuffer.Wg.Wait()
	}
}
func (segment *IDBufferSegment) BufferManager() {
	for{
		<-SegmentBizTag
		SegmentBizTagIDBuffer <- NewIDBuffer(segment.bizTag, segment.application)
	}
}
func (segment *IDBufferSegment) ReceiveChangeSlave() {
	for{
		<-SegmentSlaveChangeBizTag
		if segment.slaveIdBuffer == nil {
			SegmentBizTag<-segment.bizTag
			segment.slaveIdBuffer = <-SegmentBizTagIDBuffer
		}
		segment.masterIDBuffer = segment.slaveIdBuffer
		SegmentCreateSlaveBizTag<-segment.bizTag
	}
}

func (segment *IDBufferSegment) ReceiveCreateSlave() {
	for{
		<-SegmentSlaveChangeBizTag
		segment.slaveIdBuffer = NewIDBuffer(segment.bizTag, segment.application)
	}
}
func (segment *IDBufferSegment) ReceiveCreateMaster() {
	for{
		<-SegmentCreateSlaveBizTag
		segment.slaveIdBuffer = NewIDBuffer(segment.bizTag, segment.application)
	}
}

func NewIDBufferSegment(bizTag string, application *bootstrap.Application) *IDBufferSegment {
	segment := &IDBufferSegment{application: application, bizTag:bizTag}
	return segment
}
