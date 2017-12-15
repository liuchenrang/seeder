package generator

import (
	"errors"
	"seeder/bootstrap"
	"seeder/generator/idgen"
	"seeder/stats"
	"sync"
	"sync/atomic"
	"fmt"
)

type IDBuffer struct {
	muGetId sync.Mutex
	muCurrentId sync.Mutex
	currentId   uint64
	maxId       uint64
	step        uint64
	cacheStep   uint64

	stats  *stats.Stats
	bizTag string

	mu sync.Mutex

	muUseOut    sync.Mutex
	isUseOut    bool
	db          idgen.IDGen
	application *bootstrap.Application
}

func (this *IDBuffer) GetCurrentId() (id uint64) {
	this.muCurrentId.Lock()
	defer this.muCurrentId.Unlock()
	return atomic.LoadUint64(&this.currentId)
}
func (this *IDBuffer) GetMaxId() (id uint64) {
	return atomic.LoadUint64(&this.maxId)
}
func (this *IDBuffer) GetCacheStep() (id uint64) {

	return atomic.LoadUint64(&this.cacheStep)
}
func (this *IDBuffer) GetId() (id uint64, e error) {

	this.muGetId.Lock()
	defer this.muGetId.Unlock()

	out := this.IsUseOut()
	if out {
		return 0, errors.New("ID Use Out")
	}
	this.stats.Dig()

	return atomic.AddUint64(&this.currentId, this.step),nil
}
func (this *IDBuffer) GetStats() *stats.Stats {
	return this.stats
}
func (this *IDBuffer) IsUseOut() bool {

	this.muUseOut.Lock()
	defer this.muUseOut.Unlock()
	id := this.GetCurrentId()
	this.isUseOut = id  > this.GetMaxId()
	this.application.GetLogger().Debug(" IDBufferIsUseOut currentId", id, "max ", this.GetMaxId(), "out", this.isUseOut, fmt.Sprintf("this %p",this) )

	return this.isUseOut
}

func NewIDBuffer(bizTag string, application *bootstrap.Application) *IDBuffer {
	typeIdMake := TypeIDMake{}
	dbGen := typeIdMake.Make(bizTag, application)
	currentId, cacheStep, step, _ := dbGen.GenerateSegment(bizTag)

	this := &IDBuffer{
		bizTag: bizTag, step: step, currentId: currentId, maxId: currentId + cacheStep, cacheStep: atomic.LoadUint64(&cacheStep), db: dbGen,
		application: application,
		stats: &stats.Stats{},
	} //
	application.GetLogger().Error(" InitNewIDBuffer currentId", this.GetCurrentId(), "max ", this.GetMaxId(), "out", this.isUseOut, fmt.Sprintf("this %p",this) )

	return this
}
