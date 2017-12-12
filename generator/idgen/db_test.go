package idgen

import (
	"fmt"
	"github.com/alecthomas/log4go"
	"runtime"
	"seeder/bootstrap"
	"seeder/config"
	"seeder/logger"
	"testing"
)

func test(g IDGen) {
	fmt.Println(g.GenerateSegment("uts"))
	fmt.Println(g.GenerateSegment("uts"))
}
func TestNewEqual(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	idGen := NewDBGen("uts", Application)
	test(idGen)
	fmt.Println(seederConfig)
	Application.GetLogger().Close()
}
