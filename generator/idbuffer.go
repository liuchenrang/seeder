package generator

type IDBuffer struct{
	currentId int
	maxId int
}

func (buffer IDBuffer) GetId() int {
	return 0
}
func NewIDBuffer() IDBuffer {
	return IDBuffer{}
}
