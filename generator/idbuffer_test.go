package generator_test

import (
	//"errors"
	//"fmt"
	"testing"
	"seeder/generator"
)
func TestIdBuffer(t *testing.T) {
	// Different allocations should not be equal.
	idBuf := generator.NewIDBuffer("photo")
	idBuf.GetId()
	i := 0
	var id uint64;
	for  i < 1000  {
		i++;
		id = idBuf.GetId()
		if id != idBuf.GetId() {
			t.Error("id error")
			break;
		}
	}
}

