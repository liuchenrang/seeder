package generator

import (
	"sync/atomic"
	"fmt"
	"seeder/stats"
)
type IDBuffer struct{
	currentId uint64
	maxId uint64
	stats *stats.Stats
	bizTag string
}

func (buffer *IDBuffer) GetId() uint64 {

	buffer.stats.Dig()
	pint := & buffer.currentId
	atomic.AddUint64(pint, 1)
	fmt.Println("dig ", buffer.stats.GetTotal())
	return buffer.currentId;
}
func (buffer *IDBuffer) IsUseOut() bool {
	fmt.Println("currentId", buffer.currentId)
	fmt.Println("maxId", buffer.maxId)
	fmt.Println("compare", buffer.maxId < buffer.currentId)
	isUseOut := buffer.currentId > buffer.maxId
	return isUseOut
}
func (buffer *IDBuffer) flush(tagChan <-chan string, tagStepChan chan<- uint64) bool {
	buffer.stats.Clear()
	tagStepChan <- 2000
	return false
}
func (buffer *IDBuffer) Init(bizTag string)  {
	buffer.stats = &stats.Stats{}
}
func NewIDBuffer(bizTag string) *IDBuffer {
	make := TypeIDMake{}
	make.Make().GetId(bizTag, 1)
	buffer := &IDBuffer{currentId: 0, maxId: 1000}  //
	buffer.Init(bizTag)
	return buffer
}


