package generator

import (
	"testing"
	"fmt"
)


func TestManager(t *testing.T) {
	// Different allocations should not be equal.
	m := NewIDBufferSegmentManager("uts")
	wait := make(chan int)
	go func(){
		i := 0
		runTime := 1000;
		for i <= runTime {
			i++
			id := m.GetId("uts")
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
