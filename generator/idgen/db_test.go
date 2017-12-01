package idgen

import (
	"fmt"
	"runtime"
	"testing"
)


func test(g IDGen){
	fmt.Println(g.GenerateSegment("uts"))
	fmt.Println(g.GenerateSegment("uts"))
}
func TestNewEqual(t *testing.T)  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	idGen := NewDBGen("uts")
	test(idGen)

}
