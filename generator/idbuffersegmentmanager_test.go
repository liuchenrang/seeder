package generator

import (
	"testing"
)


func TestNewEqual(t *testing.T) {
	// Different allocations should not be equal.
	NewIDBufferSegmentManager("uts")
}
