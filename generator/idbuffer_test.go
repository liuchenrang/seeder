package generator_test

import (
	//"errors"
	//"fmt"
	"testing"
	"seeder/generator"
	"fmt"
	"seeder/config"
)
func TestIdBuffer(t *testing.T) {
	// Different allocations should not be equal.
	idBuf := generator.NewIDBuffer("photo", config.NewSeederConfig("../seeder.yaml"))
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
	idBuf := generator.NewIDBuffer("photo", config.NewSeederConfig("../seeder.yaml"))
	fmt.Println(idBuf)
	fmt.Println(idBuf.GetId())

}

