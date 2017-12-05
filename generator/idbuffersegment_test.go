package generator_test

import (
	//"errors"
	//"fmt"
	"testing"
	"seeder/generator"
	"seeder/logger"
	"fmt"
	"seeder/config"
)

func TestNewEqual(t *testing.T) {
	// Different allocations should not be equal.
	segment := generator.NewIDBufferSegment("test",config.NewSeederConfig("../seeder.yaml"))
	segment.CreateMasterIDBuffer("test")
	segment.CreateSlaveIDBuffer("test")
	id := segment.GetId()
	logger := SeederLogger.New()
	var i uint64
	for i < 40 {
		id = segment.GetId()
		nextId := segment.GetId()
		logger.Debug("id ", id, "nextId", nextId)
		fmt.Printf("xxxx")
		if id+1 != nextId {
			t.Error("id error")
			break;
		}
		fmt.Println()
		i++;
	}
}

func TestStats(t *testing.T) {
	// Different allocations should not be equal.
	segment := generator.NewIDBufferSegment("test" , config.NewSeederConfig("../seeder.yaml"))
	segment.CreateMasterIDBuffer("test")
	segment.ChangeSlaveToMaster()
	segment.GetMasterIdBuffer().GetStats()
}
