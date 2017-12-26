package SeederLogger

import (
	//"time"
	"fmt"
	"github.com/liuchenrang/log4go"
	"os"
	"seeder/config"
	"time"
	"flag"
)

var (
	seederCconfig config.SeederConfig
	//globalLogger l4g.Logger

	Author string
)

func init() {
	Author = "xinghuo"
}

type Logger struct {
	level   int
	message string
	log4go.Logger
}

func (logger Logger) Debug(a ...interface{}) {
	now := time.Now()
	fmt.Sprintf("%s-%s-%s %s:%s:%s", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(a)
}
func (logger Logger) Info(a ...interface{}) {
	now := time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), a)
}
func NewLogger(seederConfig config.SeederConfig) Logger {
	lg := Logger{}

	return lg
}
func init()  {

}
func NewLogger4g(level log4go.Level, seederConfig config.SeederConfig) log4go.Logger {
	log := log4go.NewDefaultLogger(level)
	loggerFlag  := flag.String("cc", "./log4go.xml", "log config path")

	if _, err := os.Stat(*loggerFlag); err == nil {
		log.LoadConfiguration(*loggerFlag)
		fmt.Println("load log4go config")
	}

	return log
}
func NewLogger4gWithConfig(level log4go.Level, seederConfig config.SeederConfig, loggerFlag *string) log4go.Logger {
	log := log4go.NewDefaultLogger(0)

	if _, err := os.Stat(*loggerFlag); err == nil {
		log.LoadConfiguration(*loggerFlag)
		fmt.Println("load log4go config")
	}

	return log
}
