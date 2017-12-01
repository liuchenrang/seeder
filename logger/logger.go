package logger

import (
	"fmt"
	"time"
)

type Logger struct{
	level int
	message string
}
func (logger Logger) Debug(message string){
	now := time.Now()
	fmt.Println(now.Year(), now.Month(),now.Day(),now.Hour(), now.Minute(),now.Second(), message)
}
func New() Logger {
	return Logger{}
}