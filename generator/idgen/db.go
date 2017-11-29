package idgen

import "sync"

type DBGen struct{
	Counter int
	Lock *sync.Mutex
	Fin chan<- int
}



func (dbgen DBGen ) 	GetId(bizTag string , step int ) int {
	return 0
}