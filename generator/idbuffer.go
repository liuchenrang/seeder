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
}

func (buffer *IDBuffer) GetId() uint64 {
	buffer.monitor.GetStats().Dig()
	pint := & buffer.currentId
	atomic.AddUint64(pint, 1)
	fmt.Println("dig ", buffer.monitor.GetStats().GetTotal())
	return buffer.currentId;
}
func (buffer *IDBuffer) IsUseOut() bool {
	return false
}
func (buffer *IDBuffer) flush(tagChan <-chan string, stepChan chan<- int) bool {
	return false
}
func (buffer *IDBuffer) Init()  {
	buffer.monitor = monitor.NewMonitor()

}
func NewIDBuffer(bizTag string) *IDBuffer {
	buffer := &IDBuffer{currentId: 0}
	buffer.Init()
	return buffer
}


