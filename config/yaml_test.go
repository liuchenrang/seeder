package config

import (
	"testing"
	"fmt"
)

func TestNewIDBuffer(t *testing.T) {
	// Different allocations should not be equal.
	seederConfig := NewSeederConfig("../seeder.yaml")
	fmt.Printf("--- t:\n%v\n\n", seederConfig)
}

