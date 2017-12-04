package main

import (
	"fmt"
	"log"
	"seeder/bootstrap"
	"seeder/thrift/packages/generator"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type strapper bootstrap.Strapper

type Kernel struct {
	booted        bool
	bootstrappers []strapper
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

const (
	HOST = "localhost"
	PORT = "8080"
)

type IdGeneratorServiceImpl struct {
}

func (*IdGeneratorServiceImpl) Ping() (r string, user_exception *generator.UserException, system_exception *generator.SystemException, unknown_exception *generator.UnknownException, err error) {
	return "ping", nil, nil, nil, nil
}

func (*IdGeneratorServiceImpl) GetId(t *generator.TGetIdParams) (r string, user_exception *generator.UserException, system_exception *generator.SystemException, unknown_exception *generator.UnknownException, err error) {
	fmt.Printf("request tag: %v, type: %v", t.Tag, t.GeneratorType)

	return "abc", nil, nil, nil, nil
}

func (*Kernel) Serve() {

	handlers := &IdGeneratorServiceImpl{}

	processor := generator.NewIdGeneratorServiceProcessor(handlers)
	serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", HOST+":"+PORT . "\n")
	server.Serve()
}
