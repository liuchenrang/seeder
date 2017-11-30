package generator
import "sync/atomic"
type IDBuffer struct{
	currentId uint64
	maxId uint64
}

func (buffer *IDBuffer) GetId() uint64 {
	pint := & buffer.currentId
	atomic.AddUint64(pint, 1)
	return buffer.currentId;
}
func (buffer *IDBuffer) IsUseOut() bool {
	return false
}
func NewIDBuffer() *IDBuffer {
	return &IDBuffer{currentId:0}
}


