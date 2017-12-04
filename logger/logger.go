package SeederLogger

import (
	//"fmt"
	//"time"
	"time"
	"fmt"
)

type Logger struct{
	level int
	message string
}
func (logger Logger) Debug(a ...interface{}){
	now := time.Now()

	fmt.Println(now.Year(), now.Month(),now.Day(),now.Hour(), now.Minute(),now.Second(), a)
}
func (logger Logger) Info(a ...interface{}){
	now := time.Now()
	fmt.Println(now.Year(), now.Month(),now.Day(),now.Hour(), now.Minute(),now.Second(), a)
}
func New() Logger {
	return Logger{}
}