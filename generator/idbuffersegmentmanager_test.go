package generator

import (
	"fmt"
	"seeder/bootstrap"
	"seeder/config"
	"seeder/logger"
	"testing"

	"github.com/alecthomas/log4go"
)

func TestSegManager(t *testing.T) {
	// Different allocations should not be equal.

	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))
	m := NewIDBufferSegmentManager(Application)

	i := 0
	runTime := 30
	for i <= runTime {
		i++
		id, _ := m.GetId("uts")
		fmt.Println("id====== ", id)
	}

}

func BenchmarkLoopsM(b *testing.B) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.CRITICAL, seederConfig))
	m := NewIDBufferSegmentManager(Application)
	i := func(pb *testing.PB) {

		for pb.Next() {

			i := 0
			runTime := 30
			for i <= runTime {
				i++
				id, _ := m.GetId("uts")
				fmt.Println("id====== ", id)
			}

		}
	}
	b.RunParallel(i)
}

func TestFmt(t *testing.T) {
	// Different allocations should not be equal.
	fmt.Println("xx")
}
