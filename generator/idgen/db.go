package idgen

import (
	"sync"
	//"sync/atomic"
	"sync/atomic"
)

type DBGen struct{
	Counter uint64
	Lock *sync.Mutex
	Fin chan<- int
}



func (dbgen *DBGen ) 	GetId(bizTag string , step int ) uint64 {
	//dbgen.Lock.Lock()
	pint := &dbgen.Counter
	atomic.AddUint64(pint, 1)
	if dbgen.Counter == 10001 {
		dbgen.Fin <- 10001
	}
	//defer dbgen.Lock.Unlock()
	return dbgen.Counter
}