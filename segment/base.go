package base 
import (
    "sync"
)
type IDGen interface{
	GetId(bizTag string , step int ) int
}

type DBGen struct{
   Counter int 
   Lock sync.Mutex
   Fin chan<- int
}
func (gen *DBGen) GetId(bizTag string , step int) (int){
   gen.Lock.Lock()
   gen.Counter = gen.Counter+step
   if gen.Counter == 10001 {
	   gen.Fin<- gen.Counter
   }
   defer gen.Lock.Unlock()
   return gen.Counter
}
