package main

import (
	"fmt"
	"log"
	"seeder/bootstrap"
	"seeder/logger"
	"seeder/thrift/packages/generator"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type strapper bootstrap.Strapper

type Kernel struct {
	booted        bool
	bootstrappers []strapper
	SeederLogger.Logger
	applicaton *bootstrap.Application
}

func NewKernel(debug bool) *Kernel {
	return new(Kernel)
}

func (s *Kernel) RegisterBootstrapper(b strapper) {
	s.bootstrappers = append(s.bootstrappers, b)
}
func (s *Kernel) SetApplication(app *bootstrap.Application) {
	s.applicaton  = app
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

func NewUserException(code generator.ErrorCode, errorName string, message string) *generator.UserException {

	uexp := generator.NewUserException()

	uexp.ErrorCode = code
	uexp.ErrorName = errorName
	uexp.Message = &message

	return uexp
}

func NewSystemException(code generator.ErrorCode, errorName string, message string) *generator.SystemException {

	sexp := generator.NewSystemException()

	sexp.ErrorCode = code
	sexp.ErrorName = errorName
	sexp.Message = &message

	return sexp
}

type IdGeneratorServiceImpl struct {
}

func (*IdGeneratorServiceImpl) Ping() (r string, err error) {
	return "idgen", nil

}

func (*IdGeneratorServiceImpl) GetId(params *generator.TGetIdParams) (r string, err error) {
	id, err := manager.GetId(params.GetTag())
	applicaton.GetLogger().Debug("request biz tag", params.GetTag())

	if err != nil {
		return "", NewSystemException(500, "SYSTEM_ERROR", "系统错误")
	}

	return fmt.Sprintf("%d", id), nil
}


func (kernel *Kernel) Serve() {
	handlers := &IdGeneratorServiceImpl{}
	processor := generator.NewIdGeneratorServiceProcessor(handlers)
	serverTransport, err := thrift.NewTServerSocket(seederConfig.Server.Host + ":" + fmt.Sprintf("%d", seederConfig.Server.Port))
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", seederConfig.Server.Host+":"+fmt.Sprintf("%d", seederConfig.Server.Port)+"\n")
	server.Serve()
}
