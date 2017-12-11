package main

import (
	"flag"
	"seeder/bootstrap"
	"seeder/config"
	seederGenerator "seeder/generator"

	"seeder/logger"
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

	applicaton.Set("globalLogger", SeederLogger.NewLogger4g(0, seederConfig))
	manager = *seederGenerator.NewIDBufferSegmentManager(applicaton)
	go manager.SegmentManager()

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
