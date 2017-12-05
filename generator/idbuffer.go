package generator

import (
	"sync/atomic"
	"seeder/stats"
	"sync"
	"seeder/generator/idgen"
	"seeder/bootstrap"
)
type IDBuffer struct{
	currentId uint64
	maxId uint64
	stats *stats.Stats
	bizTag string
	lck *sync.Mutex
	isUseOut bool
	db idgen.IDGen
	 

	application *bootstrap.Application

}

func (buffer *IDBuffer) GetId() uint64 {
	out := buffer.IsUseOut()
	if out {
		return 0
	}
	buffer.stats.Dig()
	pint := & buffer.currentId
	atomic.AddUint64(pint, 1)
	return buffer.currentId;
}
func (buffer *IDBuffer) GetStats()  *stats.Stats {
	//buffer.application.Get("globalLogger").(log4go.Logger).Debug("buffer  nil ", buffer == nil)
	return buffer.stats
}
func (buffer *IDBuffer) IsUseOut() bool {
	if buffer.isUseOut {
		return buffer.isUseOut
	}
	buffer.lck.Lock()
	buffer.isUseOut = buffer.currentId > buffer.maxId
	buffer.lck.Unlock()
	return buffer.isUseOut
}
 func (buffer *IDBuffer) Flush(tagChan chan string)  {
	 buffer.db.UpdateStep(buffer.bizTag)
	 tagChan <- "finish"
	 buffer.application.GetLogger().Debug("Do IDBuffer Write"  , <-tagChan)
 }

func NewIDBuffer(bizTag string, application *bootstrap.Application) *IDBuffer {

	IdChan := make(chan map[string]uint64)
	typeIdMake := TypeIDMake{}
	dbGen := typeIdMake.Make(bizTag, application)
	go func(){
		maxId, cacheStep, _ := dbGen.GenerateSegment(bizTag)
		find := make(map[string]uint64)
		find["maxId"] = maxId
		find["cacheStep"] = cacheStep
		IdChan <- find
	}()
	row := make(map[string]uint64)
	row = <-IdChan
	buffer := &IDBuffer{
		bizTag:bizTag, currentId: row["maxId"], maxId: row["maxId"]  +  row["cacheStep"] , stats: &stats.Stats{}, lck:&sync.Mutex{}, db: dbGen, isUseOut:false,
		application:application,
	}  //

	return buffer
}


