package main

import (
	"fmt"
	"log"
	"seeder/bootstrap"
	"seeder/thrift/packages/generator"
	"git.apache.org/thrift.git/lib/go/thrift"
	generator2 "seeder/generator"
	"seeder/config"
	"seeder/logger"
)

type strapper bootstrap.Strapper

type Kernel struct {
	booted        bool
	bootstrappers []strapper
	SeederLogger.Logger
}

func NewKernel(debug bool) *Kernel {
	return new(Kernel)
}

func (s *Kernel) RegisterBootstrapper(b strapper) {
	s.bootstrappers = append(s.bootstrappers, b)
}

func (s *Kernel) BootstrapWith() {

	if s.booted {
		return
	}

	for _, v := range s.bootstrappers {
		v.Bootstrap()
	}

	s.booted = true
}


var (
	manager generator2.IDBufferSegmentManager
	seederConfig config.SeederConfig
	logger SeederLogger.Logger
)
type IdGeneratorServiceImpl struct {
}

func (*IdGeneratorServiceImpl)  Ping() (r string, err error){
	return "idgen", nil

}
// Parameters:
//  - Params
func (*IdGeneratorServiceImpl) GetId(params *generator.TGetIdParams) (r string, err error){

	id := manager.GetId(params.GetTag())
	fmt.Println("id", id)
	return fmt.Sprintf("%d", id), nil
}
func init()  {
	seederConfig = config.NewSeederConfig("./seeder.yaml")

	applicaton := bootstrap.NewApplication()
	applicaton.Set("globalSeederConfig", seederConfig)

	applicaton.Set("globalLogger", SeederLogger.NewLogger4g(3, seederConfig))

	manager = 	*generator2.NewIDBufferSegmentManager(applicaton)
	logger = SeederLogger.NewLogger(seederConfig)
}
func (kernel *Kernel) Serve() {
	handlers := &IdGeneratorServiceImpl{}
	processor := generator.NewIdGeneratorServiceProcessor(handlers)
	serverTransport, err := thrift.NewTServerSocket(seederConfig.Server.Host + ":" + fmt.Sprintf("%d",seederConfig.Server.Port))
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", seederConfig.Server.Host+":"+ fmt.Sprintf("%d",seederConfig.Server.Port) + "\n")
	server.Serve()
}
