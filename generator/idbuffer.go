package generator

import (
	"errors"
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

	stats       stats.Stats
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

	return atomic.LoadUint64(&this.cacheStep)
}
func (this *IDBuffer) GetId() (id uint64, e error) {
	this.mu1.Lock()
	defer this.mu1.Unlock()
	out := this.IsUseOut()
	if out {
		return 0, errors.New("ID Use Out")
	}
	this.stats.Dig()
	atomic.AddUint64(&this.currentId, this.step)
	return this.currentId, nil
}
func (this *IDBuffer) GetStats() stats.Stats {
	return this.stats
}
func (this *IDBuffer) IsUseOut() bool {
	this.mu.Lock()
	defer this.mu.Unlock()
	if this.isUseOut {
		return this.isUseOut
	}


	cid := atomic.LoadUint64(&this.currentId)
	this.isUseOut = cid >= this.maxId
	this.application.GetLogger().Debug(" IDBuffer currentId", cid, "max ", this.maxId, "out", this.isUseOut)

	return this.isUseOut
}


func NewIDBuffer(bizTag string, application *bootstrap.Application) *IDBuffer {
	typeIdMake := TypeIDMake{}
	dbGen := typeIdMake.Make(bizTag, application)
	currentId, cacheStep, step, _ := dbGen.GenerateSegment(bizTag)

	this := &IDBuffer{
		bizTag: bizTag, step: step, currentId: currentId, maxId: currentId + cacheStep,cacheStep: atomic.LoadUint64(&cacheStep), db: dbGen, isUseOut: false,
		application: application,
	} //


	return this
}
