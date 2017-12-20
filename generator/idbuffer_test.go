package generator_test

import (
	//"errors"
	//"fmt"
	"fmt"

	"github.com/liuchenrang/log4go"

	"seeder/bootstrap"
	"seeder/config"
	"seeder/generator"
	"seeder/logger"
	"testing"
)

func TestIdBuffer(t *testing.T) {
	// Different allocations should not be equal.
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	idBuf := generator.NewIDBuffer("uts", Application)

	i := 1
	var id uint64
	for i < 100 {
		id, _ = idBuf.GetId()

		fmt.Println("id", id)
		i++
	}
}

func TestNewIDBuffer(t *testing.T) {
	// Different allocations should not be equal.
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	idBuf := generator.NewIDBuffer("photo", Application)
	fmt.Println(idBuf)
	fmt.Println(idBuf.GetId())

}
