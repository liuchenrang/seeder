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

}
