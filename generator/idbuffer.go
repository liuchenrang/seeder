package generator

import (
	"sync/atomic"
	"seeder/stats"
	"sync"
	"seeder/generator/idgen"
	"seeder/config"
	"seeder/logger"
)
type IDBuffer struct{
	currentId uint64
	maxId uint64
	stats *stats.Stats
	bizTag string
	lck *sync.Mutex
	isUseOut bool
	db idgen.IDGen
	
	SeederLogger.Logger

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
	//buffer.Debug("buffer  nil ", buffer == nil)
	return buffer.stats
}
func (buffer *IDBuffer) IsUseOut() bool {
	if buffer.isUseOut {
		return buffer.isUseOut
	}
	buffer.lck.Lock()
	buffer.isUseOut = buffer.currentId > buffer.maxId
	buffer.Debug("currentId ", buffer.currentId, "maxId", buffer.maxId,"isUseOut",  buffer.isUseOut)
	buffer.lck.Unlock()
	return buffer.isUseOut
}
 func (buffer *IDBuffer) Flush(tagChan chan string)  {
	 buffer.db.UpdateStep(buffer.bizTag)
	 tagChan <- "finish"
	 buffer.Debug("Do IDBuffer Write"  , <-tagChan)
 }

func NewIDBuffer(bizTag string, seederConfig config.SeederConfig) *IDBuffer {
	IdChan := make(chan map[string]uint64)
	typeIdMake := TypeIDMake{}
	dbGen := typeIdMake.Make(bizTag, seederConfig)
	go func(){
		maxId, cacheStep, _ := dbGen.GenerateSegment(bizTag)
		find := make(map[string]uint64)
		find["maxId"] = maxId
		find["cacheStep"] = cacheStep
		IdChan <- find
	}()
	row := make(map[string]uint64)
	row = <-IdChan
	buffer := &IDBuffer{bizTag:bizTag, currentId: row["maxId"], maxId: row["maxId"]  +  row["cacheStep"] , stats: &stats.Stats{}, lck:&sync.Mutex{}, db: dbGen, isUseOut:false}  //

	return buffer
}


