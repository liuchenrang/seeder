package idgen

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)


func test(g IDGen){
	fmt.Println(g.GenerateSegment("biz"))
	fmt.Println(g.GenerateSegment("biz"))
}
func TestNewEqual(t *testing.T)  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	inchan := make(chan int)
	lck := &sync.Mutex{}
	idGen := &DBGen{ Fin: inchan,Lock: lck }
	test(idGen)

}
