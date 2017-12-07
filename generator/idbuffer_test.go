package generator_test

import (
	//"errors"
	//"fmt"
	"fmt"
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
	Application.Set("globalLogger", SeederLogger.NewLogger4g(3, seederConfig))

	idBuf := generator.NewIDBuffer("test", Application)
	idBuf.GetId()
	idBuf.Flush()

	i := 1
	var id uint64
	for i < 102 {
		id, _ = idBuf.GetId()
		fmt.Println("id", id)
		i++
	}
}

func TestNewIDBuffer(t *testing.T) {
	// Different allocations should not be equal.
	Application := bootstrap.NewApplication()
	config := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", config)

	idBuf := generator.NewIDBuffer("photo", Application)
	fmt.Println(idBuf)
	fmt.Println(idBuf.GetId())

}
