package generator_test

import (
	//"errors"
	//"fmt"
	"testing"
	"seeder/generator"
	"seeder/logger"
)

func TestNewEqual(t *testing.T) {
	// Different allocations should not be equal.
	segment := generator.NewIDBufferSegment("photo")
	id := segment.GetId()
	logger := logger.New()
	var i uint64
	for i < 2000 {
		id = segment.GetId()
		if id+1 != segment.GetId() {
			t.Error("id error")
			break;
		}
		logger.Debug("id ")
		i++;
	}
}
