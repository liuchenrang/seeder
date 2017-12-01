package idgen

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)


func test(g IDGen){
	fmt.Println(g.generateSegment("biz"))
	fmt.Println(g.generateSegment("biz"))
}
func TestNewEqual(t *testing.T)  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	inchan := make(chan int)
	lck := &sync.Mutex{}
	idGen := &DBGen{ Fin: inchan,Lock: lck }
	test(idGen)

}
