package generator

type IDBufferSegment struct {
	currentIDBuffer IDBuffer
	idList []IDBuffer
}

func (segment IDBufferSegment) GetId() int  {

	return 0;
}

func NewIDBufferSegment() (IDBufferSegment) {
	return  IDBufferSegment{}
}
