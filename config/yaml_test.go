package config

import (
	"testing"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"time"
)

func TestNewIDBuffer(t *testing.T) {
	 //time.LoadLocation("Asia/Chongqing")
	// Different allocations should not be equal.
	seederConfig := NewSeederConfig("../seeder.yaml")
	fmt.Printf("--- t:\n%+v\n\n", seederConfig)
	log := l4g.NewDefaultLogger(l4g.DEBUG)
	defer log.Close()
	log.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter())
	log.Info("The time is now: %s", time.Now().Format("2006-01-02 15:04:05"))
}

