package generator

import (
	"fmt"
	"seeder/bootstrap"
	"sync"
)

type IDBufferSegment struct {
	muGetId        sync.Mutex
	muChage        sync.Mutex

	muSlaveApply   sync.Mutex
	muSlave        sync.RWMutex
	masterIDBuffer *IDBuffer
	muMaster       sync.RWMutex
	slaveIdBuffer  *IDBuffer

	muCreateBuffer sync.Mutex
	bizTag         string
	application    *bootstrap.Application

	monitorCheck chan interface{}
}

func (segment *IDBufferSegment) GetId() (id uint64) {
	var idBuffer *IDBuffer
	for {
		idBuffer = segment.GetMasterIdBuffer()
		id, _ = idBuffer.GetId()
		segment.monitorCheck <- nil
		segment.application.GetLogger().Info("Check current=", idBuffer.GetCurrentId(), "max=", idBuffer.GetMaxId(), fmt.Sprintf("this %p", idBuffer), fmt.Sprintf("segment %p", segment), fmt.Sprintf("out=%t", idBuffer.IsUseOut()))
		if idBuffer.IsUseOut() {
			segment.ChangeSlaveToMaster()
		} else {
			break
		}
	}
	segment.application.GetLogger().Info("Return ", "id", id, " current=", idBuffer.GetCurrentId(), "max=", idBuffer.GetMaxId(), fmt.Sprintf("this %p", idBuffer), fmt.Sprintf("segment %p", segment), fmt.Sprintf("out=%t", idBuffer.IsUseOut()))

	return id
}

func (segment *IDBufferSegment) IsMasterUserOut() bool {
	segment.muMaster.RLock()
	defer segment.muMaster.RUnlock()

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
	segment.muMaster.RLock()
	defer segment.muMaster.RUnlock()
	return segment.masterIDBuffer
}
func (segment *IDBufferSegment) GetSlaveIdBuffer() *IDBuffer {
	segment.muSlave.RLock()
	defer segment.muSlave.RUnlock()
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
		segment.application.GetLogger().Info("ChangeSlaveToMaster ", fmt.Sprintf("master %p", segment.masterIDBuffer), fmt.Sprintf("slave %p", segment.slaveIdBuffer))
		segment.SetMasterIDBuffer(segment.ApplySlave())
	}
}
func (segment *IDBufferSegment) ApplySlave() *IDBuffer  {
	segment.muSlaveApply.Lock()
	defer segment.muSlaveApply.Unlock()
	if segment.GetSlaveIdBuffer() == nil {
		segment.application.GetLogger().Info(" ApplySlaveNilCreate ", segment.bizTag)
		segment.SetSlaveIdBuffer(segment.CreateBuffer(segment.bizTag))
	} else {
		if segment.GetSlaveIdBufferIsUseOut() {
			segment.application.GetLogger().Info(" ApplySlaveCreateSlave ", segment.bizTag)
			segment.SetSlaveIdBuffer(segment.CreateBuffer(segment.bizTag))
		} else {
			segment.application.GetLogger().Info(" UseMonitorSlave ", segment.bizTag)
		}
	}
	return segment.GetSlaveIdBuffer()
}
func (segment *IDBufferSegment) GetSlaveIdBufferIsUseOut() bool {
	segment.muSlave.RLock()
	defer segment.muSlave.RUnlock()
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
	segment.monitorCheck = make(chan interface{}, 10)
	go func(check chan interface{}) {
		application := segment.application
		monitor := NewMonitor(segment, application)
		for {
			<-segment.monitorCheck
			vigilanValue := application.GetConfig().Monitior.VigilantValue
			application.GetLogger().Info("NewMonitor timer ", segment.bizTag, "Vigilant", vigilanValue)
			if vigilanValue <= 100 {
				monitor.SetVigilantValue(vigilanValue)
				vigilant := monitor.IsOutVigilantValue()
				if vigilant && !segment.GetMasterIdBuffer().GetStats().Stop {
					application.GetLogger().Info(" OverCallCreateSlaveIDBuffer ", segment.bizTag)
					segment.ApplySlave()
					segment.GetMasterIdBuffer().GetStats().DoStop()
				}
			}

		}
	}(segment.monitorCheck)
}
func NewIDBufferSegment(bizTag string, application *bootstrap.Application) *IDBufferSegment {
	segment := &IDBufferSegment{application: application}
	segment.SetBizTag(bizTag)
	segment.CreateMasterIDBuffer(segment.bizTag)
	segment.StartMonitor()
	segment.application.GetLogger().Info("InitMaster ", fmt.Sprintf("master %p", segment.masterIDBuffer), fmt.Sprintf("slave ", segment.slaveIdBuffer))

	return segment
}
