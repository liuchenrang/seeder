package idgen

import (
	"fmt"
	"runtime"
	"testing"
	"seeder/config"
)


func test(g IDGen){
	fmt.Println(g.GenerateSegment("uts"))
	fmt.Println(g.GenerateSegment("uts"))
}
func TestNewEqual(t *testing.T)  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	seederConfig := config.NewSeederConfig("../../seeder.yaml")
	idGen := NewDBGen("uts", seederConfig)
	test(idGen)
	fmt.Println(seederConfig)
}
