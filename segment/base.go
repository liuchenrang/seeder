package segment
type IDGen interface{
	getId(bizTag string , step int ) int
}

type DBGen struct{
   counter int 
   lock sync.Mutex
   fin chan<- int
}
func (gen *DBGen) getId(bizTag string , step int) (int){
   gen.lock.Lock()
   gen.counter = gen.counter+step
   if gen.counter == 10001 {
	   gen.fin<- gen.counter
   }
   defer gen.lock.Unlock()
   return gen.counter
}
