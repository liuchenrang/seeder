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
	segment := generator.NewIDBufferSegment("uts",config.NewSeederConfig("../seeder.yaml"))
	segment.CreateMasterIDBuffer("uts")
	id := segment.GetId()
	logger := SeederLogger.New()
	var i uint64
	for i < 2000 {
		id = segment.GetId()
		if id+1 != segment.GetId() {
			t.Error("id error")
			break;
		}
		fmt.Println()
		logger.Debug("id ", id)
		i++;
	}
}

func TestStats(t *testing.T) {
	// Different allocations should not be equal.
	segment := generator.NewIDBufferSegment("uts" , config.NewSeederConfig("../seeder.yaml"))
	segment.CreateMasterIDBuffer("uts")
	segment.ChangeSlaveToMaster()
	segment.GetMasterIdBuffer().GetStats()
}
