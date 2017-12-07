package generator

import (
	"errors"
	"fmt"
	"seeder/bootstrap"
	"seeder/generator/idgen"
	"seeder/stats"
	"sync"
	"sync/atomic"
)

var mu sync.Mutex

type IDBuffer struct {
	currentId   uint64
	maxId       uint64
	step        uint64
	stats       *stats.Stats
	bizTag      string
	lck         *sync.Mutex
	isUseOut    bool
	db          idgen.IDGen
	application *bootstrap.Application
}

func (this *IDBuffer) GetCurrentId() (id uint64) {

	return this.currentId
}
func (this *IDBuffer) GetMaxId() (id uint64) {

	return this.maxId
}
func (this *IDBuffer) GetId() (id uint64, e error) {
	out := this.IsUseOut()
	if out {
		return 0, errors.New("ID Use Out")
	}
	this.stats.Dig()
	fmt.Println(" IDBuffer ", this.currentId)
	pint := &this.currentId
	atomic.AddUint64(pint, this.step)
	return this.currentId, nil
}
func (this *IDBuffer) GetStats() *stats.Stats {
	//this.application.Get("globalLogger").(log4go.Logger).Debug("this  nil ", this == nil)
	return this.stats
}
func (this *IDBuffer) IsUseOut() bool {
	if this.isUseOut {
		return this.isUseOut
	}
	this.lck.Lock()
	defer this.lck.Unlock()

	this.isUseOut = this.currentId >= this.maxId
	this.application.GetLogger().Debug(" IDBuffer currentId", this.currentId, "max ", this.maxId, "out", this.isUseOut)

	return this.isUseOut
}
func (this *IDBuffer) Flush() {
	this.db.UpdateStep(this.bizTag)
}

func NewIDBuffer(bizTag string, application *bootstrap.Application) *IDBuffer {
	mu.Lock()
	defer mu.Unlock()
	typeIdMake := TypeIDMake{}
	dbGen := typeIdMake.Make(bizTag, application)
	currentId, cacheStep, step, _ := dbGen.GenerateSegment(bizTag)
	this := &IDBuffer{
		bizTag: bizTag, step: step, currentId: currentId, maxId: currentId + cacheStep, stats: &stats.Stats{}, lck: &sync.Mutex{}, db: dbGen, isUseOut: false,
		application: application,
	} //
	return this
}
