package main

<<<<<<< HEAD
import "fmt"

const (
	SEEDER_CONFIG_PATH = "/usr/local/Cellar/go/gopath/src/seeder/seeder.yaml"
)
func main()  {
	fmt.Println(SEEDER_CONFIG_PATH)
	//select {
	//	case ct := <-inchan:
	//			fmt.Printf("1000000  ", ct)
	//	default:
	//			fmt.Printf("2000000 xx")
	//}
	//var input string
	//fmt.Scanln(&input)
	//fmt.Printf("hh %s",input)
import (
	"flag"
	"seeder/bootstrap"
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
