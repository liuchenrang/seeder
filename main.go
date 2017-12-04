package main
import "fmt"
import "sync"
import (
	"seeder/generator/idgen"
	"runtime"
)

func main()  {
	 runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)
	inchan := make(chan int)
	lck := &sync.Mutex{}
	idGen := &idgen.DBGen{Counter:0, Fin: inchan,Lock: lck }
	i:=0
	go test(idGen)
	for i < 10000 {
		i = i + 1
	    go test(idGen)
	}
	fmt.Println("i", i)
	//select {
	//	case ct := <-inchan:
	//			fmt.Printf("1000000  ", ct)
	//	default:
	//			fmt.Printf("2000000 xx")
	//}
	 fmt.Printf("inchan -> %d", <-inchan)
	//var input string
	//fmt.Scanln(&input)
	//fmt.Printf("hh %s",input)
}
