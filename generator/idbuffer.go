package generator

import (
	"sync/atomic"
	"seeder/monitor"
	"fmt"
)
type IDBuffer struct{
	currentId uint64
	maxId uint64
	monitor *monitor.Monitor
	bizTag string
}

func (buffer *IDBuffer) GetId() uint64 {

	buffer.monitor.GetStats().Dig()
	pint := & buffer.currentId
	atomic.AddUint64(pint, 1)
	fmt.Println("dig ", buffer.monitor.GetStats().GetTotal())
	return buffer.currentId;
}
func (buffer *IDBuffer) IsUseOut() bool {
	fmt.Println("currentId", buffer.currentId)
	fmt.Println("maxId", buffer.maxId)
	fmt.Println("compare", buffer.maxId < buffer.currentId)
	isUseOut := buffer.currentId > buffer.maxId
	return isUseOut
}
func (buffer *IDBuffer) flush(tagChan <-chan string, tagStepChan chan<- int) bool {
	tagStepChan <- 2000
	return false
}
func (buffer *IDBuffer) Init(bizTag string)  {
	buffer.monitor = monitor.NewMonitor()
	buffer.maxId = 1000;

}
func NewIDBuffer(bizTag string) *IDBuffer {
	buffer := &IDBuffer{currentId: 0}
	buffer.Init(bizTag)
	return buffer
}


