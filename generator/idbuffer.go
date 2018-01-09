package generator

import (
	error2 "seeder/error"
	"fmt"
	"seeder/bootstrap"
	"seeder/generator/idgen"
	"seeder/stats"
	"sync"
	"sync/atomic"
	"errors"
)

type IDBuffer struct {
	muGetId     sync.Mutex
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
	u := atomic.AddUint64(&this.currentId, this.step)
	out := this.IsUseOut()
	if out {
		return 0,  errors.New(error2.ID_USE_OUT)
	}
	this.stats.Dig()
	return u, nil
}
func (this *IDBuffer) GetStats() *stats.Stats {
	return this.stats
}
func (this *IDBuffer) IsUseOut() bool {

	this.muUseOut.Lock()
	defer this.muUseOut.Unlock()
	id := this.GetCurrentId()
	this.isUseOut = id > this.GetMaxId()  //大于MaxID 返回最后一次生成的Id ,  数据库中保存的是使用过的id
	this.application.GetLogger().Debug(" IDBufferIsUseOut currentId", id, "max ", this.GetMaxId(), "out", this.isUseOut, fmt.Sprintf("this %p", this))

	return this.isUseOut
}

func NewIDBuffer(bizTag string, application *bootstrap.Application) *IDBuffer {
	typeIdMake := TypeIDMake{}
	dbGen := typeIdMake.Make(bizTag, application)
	currentId, cacheStep, step, _ := dbGen.GenerateSegment(bizTag)
	if step <= 0 {
		panic(fmt.Sprintf("bizTag %s, step %d is error", bizTag, step))
	}
	this := &IDBuffer{
		bizTag:      bizTag, step: step, currentId: currentId, maxId: currentId + cacheStep, cacheStep: atomic.LoadUint64(&cacheStep), db: dbGen,
		application: application,
		stats:       &stats.Stats{},
	} //
	application.GetLogger().Debug(" InitNewIDBuffer currentId %d, max %d, out %t, this %p, tag %s", this.GetCurrentId(), this.GetMaxId(), this.isUseOut, this, bizTag)

	return this
}
func NewIDBuffer2(bizTag string, application *bootstrap.Application) *IDBuffer {
	typeIdMake := TypeIDMake{}
	dbGen := typeIdMake.Make(bizTag, application)
	this := &IDBuffer{
		bizTag:      bizTag, step: 5, currentId: 98555, maxId: 98555, cacheStep: 5, db: dbGen,
		application: application,
		stats:       &stats.Stats{},
	} //
	application.GetLogger().Debug(" InitNewIDBuffer currentId %d, max %d, out %t, this %p, tag %s", this.GetCurrentId(), this.GetMaxId(), this.isUseOut, this, bizTag)

	return this
}
