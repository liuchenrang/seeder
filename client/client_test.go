package client

import (
	"fmt"
	"log"
	"seeder/bootstrap"
	"seeder/config"
	"seeder/logger"
	thriftGenerator "seeder/thrift/packages/generator"
	"testing"

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
	for i < 3 {
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
func BenchmarkLoopsUts(b *testing.B) {
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

			fmt.Println("id", id)

		}
	}
	b.RunParallel(i)
}
// 测试并发效率
func BenchmarkLoopsTest2(b *testing.B) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	i := func(pb *testing.PB) {
		client := NewClient(Application)

		for pb.Next() {
			id, error := client.GetId(&thriftGenerator.TGetIdParams{Tag: "test2", GeneratorType: 1})
			if error != nil {
				log.Fatal(error)
			}

			fmt.Println("id", id)

		}
	}
	b.RunParallel(i)
}

// 测试并发效率
func BenchmarkLoopsTest5(b *testing.B) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	i := func(pb *testing.PB) {
		client := NewClient(Application)

		for pb.Next() {
			id, error := client.GetId(&thriftGenerator.TGetIdParams{Tag: "test5", GeneratorType: 1})
			if error != nil {
				log.Fatal(error)
			}

			fmt.Println("id", id)

		}
	}
	b.RunParallel(i)
}

