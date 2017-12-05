package SeederLogger

import (
	"testing"
	//"time"
	//l4g "github.com/alecthomas/log4go"
	"seeder/config"
)

func TestLogger(t *testing.T) {
	seederConfig := config.NewSeederConfig("../seeder.yaml")

	lg := NewLogger(seederConfig)
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")
	lg.Debug("xxx")

	//filename := "xxx.log"
	//// Get a new logger instance
	//log := l4g.NewDefaultLogger(l4g.DEBUG)
	//
	//// Create a default logger that is logging messages of FINE or higher
	////log.AddFilter("file", l4g.FINE, l4g.NewFileLogWriter(filename, false))
	//log.AddFilter("file", l4g.FINE, l4g.NewConsoleLogWriter())
	//log.Close()
	//
	///* Can also specify manually via the following: (these are the defaults) */
	//flw := l4g.NewFileLogWriter(filename, false)
	//flw.SetFormat("[%D %T] [%L] (%S) %M")
	//flw.SetRotate(false)
	//flw.SetRotateSize(0)
	//flw.SetRotateLines(0)
	//flw.SetRotateDaily(false)
	//log.AddFilter("file", l4g.FINE, flw)
	//// Log some experimental messages
	//log.Finest("Everything is created now (notice that I will not be printing to the file)")
	//log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	//log.Critical("Time to close out!")
	//
	//// Close the log
	//log.Close()
}
