package generator

import (
	"testing"
	"fmt"
)


func TestNewEqual(t *testing.T) {
	// Different allocations should not be equal.
	m := NewIDBufferSegmentManager("uts")
	//m.GetId("uts")
	fmt.Println(m.GetId("uts"))
}
