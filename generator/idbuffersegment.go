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
	 tagStep := make(chan int);
	 go segment.currentIdBuffer.flush(tagChan,tagStep)
	 <-tagStep;
	 return segment.currentIdBuffer;

}
func (segment *IDBufferSegment) Init(bizTag string) bool  {
	segment.bizTag = bizTag
	idbuffer := NewIDBuffer(bizTag)
	segment.ids = append(segment.ids, idbuffer)
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


