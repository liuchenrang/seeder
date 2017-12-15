package client

import (
	"fmt"
	"log"
	"seeder/bootstrap"
	"seeder/config"
	"seeder/logger"
	thriftGenerator "seeder/thrift/packages/generator"
	"testing"

	"github.com/liuchenrang/log4go"

	"net"
	"os"
	"git.apache.org/thrift.git/lib/go/thrift"
	"seeder/thrift/packages/inthrift"
)

func TestNewClient(t *testing.T) {

	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)

	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	client := NewClient(Application)

	idp, _ := client.Ping(nil)
	i := 0
	for i < 3 {
		id, error := client.GetId(nil,&thriftGenerator.TGetIdParams{Tag: "uts", GeneratorType: 1})
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
			id, error := client.GetId(nil, &thriftGenerator.TGetIdParams{Tag: "uts", GeneratorType: 1})
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
			id, error := client.GetId(nil, &thriftGenerator.TGetIdParams{Tag: "test2", GeneratorType: 1})
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
			id, error := client.GetId(nil, &thriftGenerator.TGetIdParams{Tag: "test5", GeneratorType: 1})
			if error != nil {
				log.Fatal(error)
			}

			fmt.Println("id", id)

		}
	}
	b.RunParallel(i)
}

func TestNewClient2(t *testing.T) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))
	client := NewInThriftClient(Application)
	id, error := client.Call(nil, "NABTestService", "getTestID","[]", "")
	fmt.Println("id",id,"e",error)


}
// 测试并发效率
func BenchmarkInThrift(b *testing.B) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	i := func(pb *testing.PB) {
		client := NewInThriftClient(Application)

		for pb.Next() {
			id, error := client.Call(nil, "NABTestService", "getTestID","[]", "")
			if error != nil {
				log.Fatal(error)
			}

			fmt.Println("id", id)

		}
	}
	b.RunParallel(i)
}

func NewInThriftClient(application *bootstrap.Application) *inthrift.ApiServiceClient {
	config := application.GetConfig()
	config.Server.Port = 9511
	config.Server.Host = "127.0.0.1"
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	fmt.Println("connect", config.Server.Host, config.Server.Port)
	// tsocket, err := thrift.NewTSocket(net.JoinHostPort("10.10.109.250", fmt.Sprintf("%d", config.Server.Port)))
	tsocket, err := thrift.NewTSocket(net.JoinHostPort(config.Server.Host, fmt.Sprintf("%d", config.Server.Port)))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	client := inthrift.NewApiServiceClientFactory(tsocket, protocolFactory)
	if err := tsocket.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+config.Server.Host+":"+fmt.Sprintf("%d", config.Server.Port), " ", err)
		os.Exit(1)
	}
	return client

}
