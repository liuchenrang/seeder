package generator


type IDBufferSegment struct {
	currentIdBuffer *IDBuffer
	ids [] *IDBuffer
	bizTag string
}

func (segment *IDBufferSegment) GetId() uint64  {
	idBuf := segment.currentIdBuffer
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
			segment.currentIdBuffer = idBuf
		}else{
			hasOneUse++
		}
	}
	if hasOneUse == 1 {
		go segment.currentIdBuffer.flush(tagChan,tagStep)
	} else if hasOneUse == 2 {
		go segment.currentIdBuffer.flush(tagChan,tagStep)
		segment.currentIdBuffer.maxId = <-tagStep
	}
	return segment.currentIdBuffer;

}

func (segment *IDBufferSegment) Init(bizTag string) bool  {
	segment.bizTag = bizTag
	idBuffer := NewIDBuffer(bizTag)
	segment.ids = append(segment.ids, idBuffer)
	segment.currentIdBuffer = segment.selectIdBuffer()
	return true;
}
func (segment *IDBufferSegment) SetBizTag(bizTag string)   {
	segment.bizTag = bizTag
}

func NewIDBufferSegment(bizTag string) (*IDBufferSegment) {
	segment :=  &IDBufferSegment{}
	segment.Init(bizTag)
	return segment
}


