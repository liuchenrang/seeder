package main
import "fmt"
type IDGen interface{
	 getId(bizTag string , step int ) int
}

type DBGen struct{
	counter int 
}
func (gen *DBGen) getId(bizTag string , step int) (int){
	gen.counter = gen.counter+step
	return gen.counter
}
func test(gen IDGen){
	fmt.Println("%d", gen.getId("photo", 3))
}

func main()  {
	idGen := &DBGen{counter:1}
	i:=0
	for i < 100 {
		i++
		go test(idGen)
	}
}
