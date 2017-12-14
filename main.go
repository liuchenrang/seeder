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

var debug = flag.Bool("d", false, "run in debug model")
var configFlag = flag.String("c", "./seeder.yaml", "config path")
var versionFlat = flag.String("version", VERSION, "")
var startFlag = flag.Bool("start", false, "start server")

func NewApplication() *bootstrap.Application {
	applicaton = bootstrap.NewApplication()
	seederConfig = config.NewSeederConfig(*configFlag)
	applicaton.Set("globalSeederConfig", seederConfig)

	applicaton.Set("globalLogger", SeederLogger.NewLogger4g(0, seederConfig))
	manager = *seederGenerator.NewIDBufferSegmentManager(applicaton)
	manager.StartHotPreLoad()
	return applicaton
}

func main() {
	flag.Parse()
	if !*startFlag {
		println("seeder version ", VERSION)
		println("usage: seeder -start ")
		return
	}

	kernel := NewKernel(true)
	kernel.SetApplication(NewApplication())
	kernel.BootstrapWith()
	kernel.Serve()
}
