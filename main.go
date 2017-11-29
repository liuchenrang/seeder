package main
import "fmt"
import "sync"
import (
	"seeder/seed"
)
func test(gen seed.IDGen){
	fmt.Println("", gen.GetId("photo", 1))
}
func main()  {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)
	inchan := make(chan int)
	lck := &sync.Mutex{}
	idGen := &seed.DBGen{Counter:1, Fin: inchan,Lock: lck }
	i:=0
	for i < 10000 {
		i = i + 1
	    go test(idGen)
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
