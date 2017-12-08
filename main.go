package main

import (
	"flag"
	"seeder/bootstrap"
	"seeder/config"
	seederGenerator "seeder/generator"

	"seeder/logger"
	"github.com/alecthomas/log4go"
)

const VERSION = "1.0.0"

var (
	manager      seederGenerator.IDBufferSegmentManager
	seederConfig config.SeederConfig
	logger       SeederLogger.Logger
	applicaton   *bootstrap.Application
)

var helpstr =
"seeder " + VERSION + `
Options:
-start
        start service
-config
        config file
`

var debug = flag.Bool("d", false, "run in debug model")

var help = flag.Bool("h", true, "show tips")

var configFlag = flag.String("c", "./seeder.yaml", "config path")
var startFlag = flag.Bool("start", false, "start server")

func NewApplication() *bootstrap.Application{
	applicaton = bootstrap.NewApplication()
	seederConfig = config.NewSeederConfig(	*configFlag)
	applicaton.Set("globalSeederConfig", seederConfig)
	var level log4go.Level
	if seederConfig.Logger.Level == "DEBUG" {
		level = log4go.DEBUG
	}
	if seederConfig.Logger.Level == "CRITICAL" {
		level = log4go.CRITICAL
	}
	applicaton.Set("globalLogger", SeederLogger.NewLogger4g(level, seederConfig))
	manager = *seederGenerator.NewIDBufferSegmentManager(applicaton)
	return applicaton
}
func main() {
	flag.Parse()
	if *help && !*startFlag {
		println(helpstr)
		return
	}

	kernel := NewKernel(true)
	kernel.SetApplication(NewApplication())
	log := bootstrap.NewLogBootstrapper("/ab/abc")

	kernel.RegisterBootstrapper(log)

	kernel.BootstrapWith()

	kernel.Serve()
}
