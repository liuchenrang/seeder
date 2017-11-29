package main
import (
    "sync"
)
type IDGen interface{
	GetId(bizTag string , step int ) int
}

type DBGen struct{
   Counter int 
   Lock *sync.Mutex
   Fin chan<- int
}

func (dbgen DBGen ) 	GetId(bizTag string , step int ) int {
	return 0
}

type IDBuffer struct{

}
type TypeIDMake struct{

}

type TypeMake interface {
	 factory(makeType string) IDGen
}

func (typeMake TypeIDMake ) make() IDGen {
	return DBGen{}
}
func NewTypeIDMake() TypeIDMake {
	return TypeIDMake{}
}