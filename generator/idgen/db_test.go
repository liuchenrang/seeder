package idgen

import (
	"fmt"
	"github.com/liuchenrang/log4go"

	"runtime"
	"seeder/bootstrap"
	"seeder/config"
	"seeder/logger"
	"testing"
	"sort"
)


func TestNewEqual(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	idGen := NewDBGen("uts", Application)
	fmt.Println(idGen.GenerateSegment("uts5"))

	fmt.Println(seederConfig)
	Application.GetLogger().Close()
}

func BenchmarkNewDBGen(b *testing.B) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.CRITICAL, seederConfig))
	sgGroup := make([]int, 1)

	i := func(pb *testing.PB) {
		m := NewDBGen("uts", Application)

		for pb.Next() {

			id, _, _, _ := m.GenerateSegment("uts")
			fmt.Println("id====== ", id)
		}
	}
	b.RunParallel(i)
	sort.Ints(sgGroup)

}

