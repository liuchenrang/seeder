package generator


type IDBufferSegment struct {
	currentIDBuffer *IDBuffer
	idList []*IDBuffer
}

func (segment *IDBufferSegment) GetId() uint64  {
	idBuf := segment.currentIDBuffer

	if idBuf.IsUseOut() {
		for idBuf = segment.selectIdBuffer(); !idBuf.IsUseOut(); {

		}
	}
	return idBuf.GetId();
}
func (segment *IDBufferSegment) selectIdBuffer() *IDBuffer {
	if segment.idList[0].IsUseOut() {

	} else {

	}
	return segment.idList[0]
}
func (segment *IDBufferSegment) Init() bool  {
	idbuffer := NewIDBuffer()
	segment.idList = append(segment.idList, idbuffer)
	segment.currentIDBuffer = segment.selectIdBuffer()
	return true;
}

func NewIDBufferSegment() (*IDBufferSegment) {
	segment :=  &IDBufferSegment{}
	segment.Init()
	return segment
}


