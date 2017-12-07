package generator_test

import (
	//"errors"
	//"fmt"
	"fmt"
	"seeder/bootstrap"
	"seeder/config"
	"seeder/generator"
	"seeder/logger"
	"testing"
)

func TestMasterChange(t *testing.T) {
	// Different allocations should not be equal.
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(3, seederConfig))
	segment := generator.NewIDBufferSegment("uts", Application)
	segment.CreateMasterIDBuffer("uts")
	segment.CreateSlaveIDBuffer("uts")
	var id uint64
	logger := SeederLogger.NewLogger(seederConfig)
	var i uint64
	for i < 40 {
		if segment.IsMasterUserOut() {
			segment.ChangeSlaveToMaster()
		}
		id = segment.GetId()
		nextId := segment.GetId()
		logger.Debug("id ", id, "nextId", nextId)
		fmt.Printf("xxxx")
		if id+1 != nextId {
			t.Error("id error")
			break
		}
		fmt.Println()
		i++
	}
}

func TestStats(t *testing.T) {
	// Different allocations should not be equal.
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)

	segment := generator.NewIDBufferSegment("test", Application)
	segment.CreateMasterIDBuffer("test")
	segment.ChangeSlaveToMaster()
	segment.GetMasterIdBuffer().GetStats()
}
