package idgen

import (
	"sync"
)

type DBGen struct{
	counter uint64
	maxId uint64
	cacheStep uint64
	Lock *sync.Mutex
	Fin chan<- int
}



func (dbgen *DBGen ) generateSegment(bizTag string  ) (uint64, uint64, error) {
	dbgen.find(bizTag)
	return dbgen.maxId, dbgen.cacheStep, nil
}
func (dbgen *DBGen) flush(bizTag string){
	dbgen.find(bizTag)
}
func (dbgen *DBGen) find(bizTag string){
	dbgen.counter++
	dbgen.cacheStep = 500
	dbgen.maxId = 3000*dbgen.counter
}
func NewDBGen(bizTag string) *DBGen{
	return &DBGen{}
}