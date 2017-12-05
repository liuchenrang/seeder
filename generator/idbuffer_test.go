package generator_test

import (
	//"errors"
	//"fmt"
	"testing"
	"seeder/generator"
	"fmt"
	"seeder/config"
	"seeder/bootstrap"
	"seeder/logger"
)
func TestIdBuffer(t *testing.T) {
	// Different allocations should not be equal.
	Application := bootstrap.NewApplication()
	seederConfig :=  config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(3, seederConfig))

	idBuf := generator.NewIDBuffer("test", Application)
	idBuf.GetId()
	i := 1
	var id uint64;
	for  i < 50  {
		id = idBuf.GetId()
		getId := idBuf.GetId()
		if id+1 != getId {
			t.Error("id error")
			break;
		}
		fmt.Println("id", id)
		fmt.Println("id", getId)
		i++;

	}
}

func TestNewIDBuffer(t *testing.T) {
	// Different allocations should not be equal.
	Application := bootstrap.NewApplication()
	config :=  config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", config)

	idBuf := generator.NewIDBuffer("photo", Application)
	fmt.Println(idBuf)
	fmt.Println(idBuf.GetId())

}

