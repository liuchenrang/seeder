package client

import (
	"fmt"
	"net"
	"os"
	"seeder/bootstrap"
	"seeder/thrift/packages/generator"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func NewClient(application *bootstrap.Application) *generator.IdGeneratorServiceClient {
	config := application.GetConfig()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	fmt.Println("connect", config.Server.Host, config.Server.Port)
	// tsocket, err := thrift.NewTSocket(net.JoinHostPort("10.10.109.250", fmt.Sprintf("%d", config.Server.Port)))
	tsocket, err := thrift.NewTSocket(net.JoinHostPort(config.Server.Host, fmt.Sprintf("%d", config.Server.Port)))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	transport := thrift.NewTFramedTransport(tsocket)
	client := generator.NewIdGeneratorServiceClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+config.Server.Host+":"+fmt.Sprintf("%d", config.Server.Port), " ", err)
		os.Exit(1)
	}
	return client

}
