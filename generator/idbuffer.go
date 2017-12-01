package generator

import (
	"sync/atomic"
	"seeder/stats"
	"sync"
)
type IDBuffer struct{
	currentId uint64
	maxId uint64
	stats *stats.Stats
	bizTag string
}
var m = sync.Mutex{}
func (buffer *IDBuffer) GetId() uint64 {

	buffer.stats.Dig()
	pint := & buffer.currentId
	atomic.AddUint64(pint, 1)
	return buffer.currentId;
}
func (buffer *IDBuffer) IsUseOut() bool {
	isUseOut := buffer.currentId > buffer.maxId
	return isUseOut
}
func (buffer *IDBuffer) flush(tagChan <-chan string, tagStepChan chan<- uint64) bool {
	buffer.stats.Clear()
	tagStepChan <- 2000
	return false
}

func NewIDBuffer(bizTag string) *IDBuffer {
	IdChan := make(chan map[string]uint64)
	typeIdMake := TypeIDMake{}
	go func(){
		maxId, cacheStep, _ := typeIdMake.Make(bizTag).GenerateSegment(bizTag)
		find := make(map[string]uint64)
		find["maxId"] = maxId
		find["cacheStep"] = cacheStep
		IdChan <- find
	}()
	row := make(map[string]uint64)
	row = <-IdChan
	buffer := &IDBuffer{currentId: row["maxId"], maxId: row["maxId"]  +  row["cacheStep"] , stats: &stats.Stats{}}  //
	return buffer
}


