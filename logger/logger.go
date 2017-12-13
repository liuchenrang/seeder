package SeederLogger

import (
	//"time"
	"fmt"
	"github.com/liuchenrang/log4go"
	"os"
	"seeder/config"
	"time"
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

func NewLogger4g(level log4go.Level, seederConfig config.SeederConfig) log4go.Logger {
	log := log4go.NewDefaultLogger(level)
	logConfigFile := "/usr/local/Cellar/go/gopath/src/seeder/log4go.xml"
	if _, err := os.Stat(logConfigFile); err == nil {
		log.LoadConfiguration(logConfigFile)
		fmt.Println("load log4go config")
	}

	return log
}
