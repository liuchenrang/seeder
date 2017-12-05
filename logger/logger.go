package SeederLogger

import (
	//"time"
	"time"
	"fmt"
	"seeder/config"
	"github.com/alecthomas/log4go"
)
var (
	seederCconfig config.SeederConfig
	//globalLogger l4g.Logger
)


func init()  {

}
type Logger struct{
	level int
	message string
	log4go.Logger
}

func (logger Logger) Debug(a ...interface{}){
	now := time.Now()
	fmt.Sprintf("%s-%s-%s %s:%s:%s",now.Year(), now.Month(),now.Day(),now.Hour(), now.Minute(),now.Second())
	fmt.Println(a)
}
func (logger Logger) Info(a ...interface{}){
	now := time.Now()
	fmt.Println(now.Year(), now.Month(),now.Day(),now.Hour(), now.Minute(),now.Second(), a)
}
func NewLogger(seederConfig config.SeederConfig) Logger {
	lg := Logger{}

	return lg
}

func NewLogger4g(level log4go.Level, seederConfig config.SeederConfig) log4go.Logger{
	log := log4go.NewDefaultLogger(level)
	s := seederConfig.Logger.Path + "/" + seederConfig.Logger.File
	fmt.Println("log path ", s)
	flw := log4go.NewFileLogWriter(s, false)
	flw.SetFormat("[%D %T] [%L] (%S) %M")
	flw.SetRotate(true)
	flw.SetRotateSize(1024*1024*100)
	flw.SetRotateLines(100)
	flw.SetRotateDaily(true)
	log.AddFilter("file", level, flw)
	return log
}