package main
import "fmt"
import "sync"
import "seeder/segment"
func test(gen base.IDGen){
	fmt.Println("", gen.GetId("photo", 1))
}
func main()  {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)
	inchan := make(chan int, 1)
	idGen := &base.DBGen{Counter:1, Fin: inchan,Lock: sync.Mutex{}}
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
