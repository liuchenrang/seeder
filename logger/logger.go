package SeederLogger

import (
	//"time"
	"time"
	"fmt"
	//l4g "github.com/alecthomas/log4go"
	"seeder/config"
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
	//globalLogger = l4g.NewDefaultLogger(l4g.DEBUG)
	//flw := l4g.NewFileLogWriter(seederCconfig.Logger.Path + "/" + seederCconfig.Logger.File, false)
	//flw.SetFormat("[%D %T] [%L] (%S) %M")
	//flw.SetRotate(true)
	//flw.SetRotateSize(1024*1024*100)
	//flw.SetRotateLines(100)
	//flw.SetRotateDaily(true)
	//globalLogger.AddFilter("file", l4g.DEBUG, flw)
	return Logger{}
}

