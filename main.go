package main

import (
	"flag"
	"seeder/bootstrap"
	"runtime"
)
const VERSION = "1.0.0"

var helpstr = "[" + VERSION + `] seeder command [arguments]
Options:
    -debug
         debug mode
The commands are:
    seeder run 
`

var debug = flag.Bool("debug", false, "run in debug model")

var help = flag.Bool("help", false, "show tips")

func main() {
	runtime.GOMAXPROCS(1)
	flag.Parse()

	if *help {
		println(helpstr)
		return
	}

	kernel := NewKernel(true)

	log := bootstrap.NewLogBootstrapper("/ab/abc")

	kernel.RegisterBootstrapper(log)

	kernel.BootstrapWith()

	kernel.Serve()
}
