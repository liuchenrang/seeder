package generator

import (
	"testing"
	"fmt"
	"seeder/config"
	"seeder/bootstrap"
	"seeder/logger"
	"github.com/alecthomas/log4go"
)


func TestManager(t *testing.T) {
	// Different allocations should not be equal.

	Application := bootstrap.NewApplication()
	seederConfig :=  config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)

	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	m := NewIDBufferSegmentManager(Application)

	m.GetId("uts")
	return
	wait := make(chan int)
	go func(){
		i := 0
		runTime := 30;
		for i <= runTime {
			i++
			id := m.GetId("uts")

			fmt.Println("id ", id)


		}
		if i == runTime {
			wait<- runTime
		}

		wait<- runTime
	}()
	<-wait

}

func TestFmt(t *testing.T) {
	// Different allocations should not be equal.
	fmt.Println("xx")
}