package main
import "fmt"
import "sync"
// import "runtime"
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
func test(gen IDGen){
	fmt.Println("", gen.getId("photo", 1))
}
func main()  {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)
	inchan := make(chan int, 1)
	idGen := &DBGen{counter:1, fin: inchan,lock: sync.Mutex{}}
	i:=1
	for i <= 10000 {
	    go test(idGen)
		i = i + 1
	}
	select {
	case ct := <-inchan:
			fmt.Printf(" selct ", ct)
	default:
			fmt.Printf("xx")
	}
	// fmt.Printf("%d", <-inchan)
	var input string
	fmt.Scanln(&input)
	fmt.Printf("hh %s",input)
}
