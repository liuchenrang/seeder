package generator_test


import (
	//"errors"
	//"fmt"
	"testing"
	"seeder/generator"
	"time"
)
func TestNewEqual(t *testing.T) {
	// Different allocations should not be equal.
	segment := generator.NewIDBufferSegment("photo")
	id := segment.GetId()
	t.Log("id ", id)
	if id == 1 {
		t.Log("test ok")
	} else {
		t.Error("xxx")
	}
	id = segment.GetId()
	t.Log("id ", id)
	if id == 2 {
		t.Log("test ok")
	} else {
		t.Error("xxx")
	}

		go func(){
			t.Log("id",segment.GetId())
		}()

	time.Sleep(time.Second*3)
}

