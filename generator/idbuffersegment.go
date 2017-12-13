package generator

import (
	"fmt"
	"seeder/bootstrap"
	"sync"
	"time"
)

type IDBufferSegment struct {
	muGetId  sync.Mutex
	muChage        sync.Mutex
	muSlave        sync.Mutex
	masterIDBuffer *IDBuffer
	muMaster       sync.Mutex
	slaveIdBuffer  *IDBuffer

	muCreateBuffer sync.Mutex
	bizTag         string
	application    *bootstrap.Application
}

func (segment *IDBufferSegment) GetId() (id uint64) {
	var idBuffer *IDBuffer
	segment.muGetId.Lock()
	defer segment.muGetId.Unlock()
	for {
		idBuffer = segment.GetMasterIdBuffer()
		id, _ = idBuffer.GetId()
		segment.application.GetLogger().Error("Check current=", idBuffer.GetCurrentId(), "max=", idBuffer.GetMaxId(), fmt.Sprintf("this %p", idBuffer), fmt.Sprintf("segment %p", segment), fmt.Sprintf("out=%t", idBuffer.IsUseOut()))
		if idBuffer.IsUseOut() {
			segment.ChangeSlaveToMaster()
		} else {
			break
		}
	}
	segment.application.GetLogger().Error("Return ", "id", id, " current=", idBuffer.GetCurrentId(), "max=", idBuffer.GetMaxId(), fmt.Sprintf("this %p", idBuffer), fmt.Sprintf("segment %p", segment), fmt.Sprintf("out=%t", idBuffer.IsUseOut()))

	return id
}

func (segment *IDBufferSegment) IsMasterUserOut() bool {
	segment.muMaster.Lock()
	defer segment.muMaster.Unlock()

	return segment.masterIDBuffer.IsUseOut()
}
func (segment *IDBufferSegment) CreateMasterIDBuffer(bizTag string) *IDBuffer {
	segment.muMaster.Lock()
	defer segment.muMaster.Unlock()
	segment.masterIDBuffer = NewIDBuffer(bizTag, segment.application)
	return segment.masterIDBuffer
}
func (segment *IDBufferSegment) SetMasterIDBuffer(idBuf *IDBuffer) {
	segment.muMaster.Lock()
	defer segment.muMaster.Unlock()

	segment.masterIDBuffer = idBuf
}
func (segment *IDBufferSegment) CreateSlaveIDBuffer(bizTag string) *IDBuffer {
	segment.muSlave.Lock()
	defer segment.muSlave.Unlock()
	segment.slaveIdBuffer = NewIDBuffer(bizTag, segment.application)
	return segment.slaveIdBuffer
}
func (segment *IDBufferSegment) SetBizTag(bizTag string) {
	segment.bizTag = bizTag
}
func (segment *IDBufferSegment) CreateBuffer(bizTag string) *IDBuffer {
	segment.muCreateBuffer.Lock()
	defer segment.muCreateBuffer.Unlock()
	return NewIDBuffer(bizTag, segment.application)
}
func (segment *IDBufferSegment) GetBizTag() string {
	return segment.bizTag
}
func (segment *IDBufferSegment) GetMasterIdBuffer() *IDBuffer {
	segment.muMaster.Lock()
	defer segment.muMaster.Unlock()
	return segment.masterIDBuffer
}
func (segment *IDBufferSegment) GetSlaveIdBuffer() *IDBuffer {
	segment.muSlave.Lock()
	defer segment.muSlave.Unlock()
	return segment.slaveIdBuffer
}
func (segment *IDBufferSegment) SetSlaveIdBuffer(slave *IDBuffer) {
	segment.muSlave.Lock()
	defer segment.muSlave.Unlock()
	segment.slaveIdBuffer = slave
}

func (segment *IDBufferSegment) ChangeSlaveToMaster() {
	segment.muChage.Lock()
	defer segment.muChage.Unlock()

	if segment.IsMasterUserOut() {

		if segment.GetSlaveIdBuffer() == nil {
			segment.SetSlaveIdBuffer(segment.CreateBuffer(segment.bizTag))
		} else {
			if segment.GetSlaveIdBufferIsUseOut() {
				segment.SetSlaveIdBuffer(segment.CreateBuffer(segment.bizTag))
			}
		}
		segment.application.GetLogger().Error("ChangeSlaveToMaster ", fmt.Sprintf("master %p", segment.masterIDBuffer), fmt.Sprintf("slave %p", segment.slaveIdBuffer))
		segment.SetMasterIDBuffer(segment.slaveIdBuffer)
	}
}
func (segment *IDBufferSegment) GetSlaveIdBufferIsUseOut() bool {
	segment.muSlave.Lock()
	defer segment.muSlave.Unlock()
	return segment.slaveIdBuffer.IsUseOut()
}
func (segment *IDBufferSegment) Close() {

	if segment.masterIDBuffer != nil {
		//segment.masterIDBuffer.Wg.Wait()
	}
	if segment.slaveIdBuffer != nil {
		//segment.slaveIdBuffer.Wg.Wait()
	}
}
func (segment *IDBufferSegment) StartMonitor() {

	go func() {
		application := segment.application
		monitor := NewMonitor(segment, application)
		for {
			time.Sleep(time.Millisecond * 100)
			vigilanValue := application.GetConfig().Monitior.VigilantValue
			application.GetLogger().Fine("NewMonitor timer ", segment.bizTag, "Vigilant", vigilanValue)
			if vigilanValue <= 100 {
				monitor.SetVigilantValue(vigilanValue)
				vigilant := monitor.IsOutVigilantValue()
				if vigilant && !segment.GetMasterIdBuffer().GetStats().Stop {
					application.GetLogger().Info(" Over call CreateSlaveIDBuffer ", segment.bizTag)
					segment.CreateSlaveIDBuffer(segment.bizTag)
					segment.GetMasterIdBuffer().GetStats().DoStop()
				}
			}

		}
	}()
}
func NewIDBufferSegment(bizTag string, application *bootstrap.Application) *IDBufferSegment {
	segment := &IDBufferSegment{application: application}
	segment.SetBizTag(bizTag)
	segment.CreateMasterIDBuffer(segment.bizTag)
	//segment.StartMonitor()
	segment.application.GetLogger().Error("InitMaster ", fmt.Sprintf("master %p", segment.masterIDBuffer), fmt.Sprintf("slave ", segment.slaveIdBuffer))

	return segment
}
