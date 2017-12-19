package main

import (
	"flag"
	"seeder/bootstrap"
	"seeder/config"
	seederGenerator "seeder/generator"

	"seeder/logger"
	"github.com/takama/daemon"
	"os"
	"fmt"
	"log"
)

const (
	 VERSION = "1.0.0"
	 project_name = "seeder"
	 description = "id generator"
)


var (
 	dependencies = []string{"seeder.service"}
 	stdlog, errlog *log.Logger

	manager      seederGenerator.IDBufferSegmentManager
	seederConfig config.SeederConfig
	logger       SeederLogger.Logger
	applicaton   *bootstrap.Application
//install | remove | start | stop | status

	configFlag = flag.String("c", "./seeder.yaml", "config path")
	loggerFlag = flag.String("cc", "./log4go.xml", "log config path")
	versionFlat = flag.String("version", VERSION, "")

	removeFlag = flag.Bool("remove", false, "-remove")
	startFlag = flag.Bool("start", false, "-start")
	stopFlag = flag.Bool("stop", false, "-stop")
	statusFlag = flag.Bool("status", false, "-status")
	installFlag = flag.Bool("install", false, "-install muset set -c and -cc ")
)

func NewApplication() *bootstrap.Application {
	applicaton = bootstrap.NewApplication()
	seederConfig = config.NewSeederConfig(*configFlag)
	applicaton.Set("globalSeederConfig", seederConfig)

	applicaton.Set("globalLogger", SeederLogger.NewLogger4g(0, seederConfig))
	manager = *seederGenerator.NewIDBufferSegmentManager(applicaton)
	manager.StartHotPreLoad()
	return applicaton
}
func init()  {
	stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
}
func main() {
	flag.Parse()
	srv, err := daemon.New(project_name, description, dependencies...)
	if err != nil {
		stdlog.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		stdlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
