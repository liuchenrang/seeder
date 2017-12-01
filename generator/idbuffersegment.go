package generator

import "sync"

type IDBufferSegment struct {
	masterIDBuffer *IDBuffer
	slaveIdBuffer *IDBuffer
	ids [] *IDBuffer
	changeLock sync.Mutex
	bizTag string
}

func (segment *IDBufferSegment) GetId() uint64  {
	idBuf := segment.masterIDBuffer
	if idBuf.IsUseOut() {
		for idBuf = segment.selectIdBuffer(); !idBuf.IsUseOut(); {

		}
	}
	return idBuf.GetId();
}
func (segment *IDBufferSegment) selectIdBuffer() *IDBuffer {
	 tagChan := make(chan string);
	 tagStep := make(chan uint64);
	 hasOneUse := 0
	for _, idBuf := range segment.ids {
		if !idBuf.IsUseOut() {
			segment.masterIDBuffer = idBuf
		}else{
			hasOneUse++
		}
	}
	if hasOneUse == 1 {
		go segment.masterIDBuffer.flush(tagChan,tagStep)
	} else if hasOneUse == 2 {
		go segment.masterIDBuffer.flush(tagChan,tagStep)
		segment.masterIDBuffer.maxId = <-tagStep
	}
	return segment.masterIDBuffer;

}

func (segment *IDBufferSegment) Init(bizTag string) bool  {
	idBuffer := NewIDBuffer(bizTag)
	segment.ids = append(segment.ids, idBuffer)
	segment.masterIDBuffer = segment.selectIdBuffer()
	return true;
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
	segment.changeLock.Lock()
	segment.masterIDBuffer = segment.slaveIdBuffer
	segment.slaveIdBuffer = NewIDBuffer(segment.bizTag)
	defer segment.changeLock.Unlock()
}


