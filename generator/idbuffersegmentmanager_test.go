package generator

import (
	"testing"
	"fmt"
	"seeder/config"
)


func TestManager(t *testing.T) {
	// Different allocations should not be equal.

	m := NewIDBufferSegmentManager(config.NewSeederConfig("../seeder.yaml"))

	wait := make(chan int)
	go func(){
		i := 0
		runTime := 100;
		for i <= runTime {
			i++
			id := m.GetId("test")
			if id <= 0 {
				logger.Debug("Do ChangeSlaveToMaster")
				m.segment.ChangeSlaveToMaster()
			}
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