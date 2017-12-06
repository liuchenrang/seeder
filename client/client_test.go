package client

import (
	"testing"
	thriftGenerator "seeder/thrift/packages/generator"
	"fmt"
	"log"
	"seeder/bootstrap"
	"seeder/config"
	"seeder/logger"
	"github.com/alecthomas/log4go"
)

func TestNewClient(t *testing.T) {

	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)

	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	client := NewClient(Application)

	idp, _ := client.Ping()
	i := 0
	for i < 5 {
		id, error := client.GetId(&thriftGenerator.TGetIdParams{Tag: "uts", GeneratorType: 1})
		if error != nil {
			log.Fatal(error)
		}
		fmt.Println("ping ", idp)
		fmt.Println("id", id)
		i++
	}

}

func BenchmarkLoops(b *testing.B) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)

	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	client := NewClient(Application)

	idp, _ := client.Ping()

	for i := 0; i < b.N; i++ {
		id, error := client.GetId(&thriftGenerator.TGetIdParams{Tag: "uts", GeneratorType: 1})
		if error != nil {
			log.Fatal(error)
		}
		fmt.Println("ping ", idp)
		fmt.Println("id", id)
		i++
	}

}

// 测试并发效率
func BenchmarkLoopsParallel(b *testing.B) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	i := func(pb *testing.PB) {
		client := NewClient(Application)

		for pb.Next() {
			id, error := client.GetId(&thriftGenerator.TGetIdParams{Tag: "uts", GeneratorType: 1})
			if error != nil {
				log.Fatal(error)
			}

			fmt.Println("ping ", id)

		}
	}
	b.RunParallel(i)
}
