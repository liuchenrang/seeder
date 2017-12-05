package generator

import (
	"testing"
	"fmt"
	"seeder/config"
	"seeder/bootstrap"
)


func TestManager(t *testing.T) {
	// Different allocations should not be equal.

	Application := bootstrap.NewApplication()
	seederConfig :=  config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)

	m := NewIDBufferSegmentManager(Application)
	wait := make(chan int)
	go func(){
		i := 0
		runTime := 10;
		for i <= runTime {
			i++
			id := m.GetId("uts")
			if id <= 0 {
				m.getSegmentByBizTag("uts").ChangeSlaveToMaster()
			}
			fmt.Println("id ", id)

			id = m.GetId("eyas")

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