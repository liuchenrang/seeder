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


type IDBuffer struct {
	mu1 sync.Mutex
	currentId   uint64
	maxId       uint64
	step        uint64
	cacheStep   uint64

	stats       *stats.Stats
	bizTag      string

	mu sync.Mutex
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
func (this *IDBuffer) GetCacheStep() (id uint64) {

	return this.cacheStep
}
func (this *IDBuffer) GetId() (id uint64, e error) {
	this.mu1.Lock()
	defer this.mu1.Unlock()
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
	return this.stats
}
func (this *IDBuffer) IsUseOut() bool {
	if this.isUseOut {
		return this.isUseOut
	}
	this.mu.Lock()
	defer this.mu.Unlock()

	this.isUseOut = this.currentId >= this.maxId
	this.application.GetLogger().Debug(" IDBuffer currentId", this.currentId, "max ", this.maxId, "out", this.isUseOut)

	return this.isUseOut
}


func NewIDBuffer(bizTag string, application *bootstrap.Application) *IDBuffer {
	typeIdMake := TypeIDMake{}
	dbGen := typeIdMake.Make(bizTag, application)
	currentId, cacheStep, step, _ := dbGen.GenerateSegment(bizTag)

	this := &IDBuffer{
		bizTag: bizTag, step: step, currentId: currentId, maxId: currentId + cacheStep,cacheStep: cacheStep, stats: &stats.Stats{}, db: dbGen, isUseOut: false,
		application: application,
	} //
	return this
}
