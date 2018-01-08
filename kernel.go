package main

import (
	"fmt"
	"log"
	"seeder/bootstrap"
	"seeder/logger"
	"seeder/thrift/packages/generator"

	"context"
	"os"
	"os/signal"
	"syscall"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/takama/daemon"
	"seeder/zk"
)

type strapper bootstrap.Strapper

type Kernel struct {
	booted        bool
	bootstrappers []strapper
	SeederLogger.Logger
	applicaton    *bootstrap.Application
}
type Service struct {
	daemon.Daemon
}

func (service *Service) Manage() (string, error) {

	switch {
	case *installFlag:
		if *configFlag != "" {
			*configFlag = "-c=" + *configFlag
		}
		if *loggerFlag != "" {
			*loggerFlag = "-cc=" + *loggerFlag
		}
		return service.Install(*configFlag, *loggerFlag)
	case *removeFlag:
		return service.Remove()
	case *startFlag:
		return service.Start()
	case *stopFlag:
		return service.Stop()
	case *statusFlag:
		return service.Status()

	}

	kernel := NewKernel(true)
	kernel.SetApplication(NewApplication())
	kernel.BootstrapWith()
	kernel.Serve()
	// never happen, but need to complete code
	return "", nil
}

func NewKernel(debug bool) *Kernel {
	return new(Kernel)
}

func (s *Kernel) RegisterBootstrapper(b strapper) {
	s.bootstrappers = append(s.bootstrappers, b)
}
func (s *Kernel) SetApplication(app *bootstrap.Application) {
	s.applicaton = app
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

func (*IdGeneratorServiceImpl) Ping(ctx context.Context) (r string, err error) {
	return "idgen", nil

}

//GetId GetId
func (*IdGeneratorServiceImpl) GetId(ctx context.Context, params *generator.TGetIdParams) (r string, err error) {
	defer func() {
		if err := recover(); err != nil {
			applicaton.GetLogger().Error(err)
		}
	}()
	id, err := manager.GetId(params.GetTag(), params.GeneratorType)
	applicaton.GetLogger().Debug("request biz tag", params.GetTag())
	if err != nil {
		return "", NewSystemException(500, "SYSTEM_ERROR", "系统错误")
	}
	fmt.Println(id)
	return fmt.Sprintf("success %d", id), nil
}

func (kernal *Kernel) RegisterServer() {
	 soa := kernal.applicaton.GetServerSoa().(*zk.ServerSoa)
	 soa.Register()
}

// Parameters:
//  - Params
func (kernel *Kernel) Serve() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

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
	go func() {
		<-sigs
		fmt.Println("stop servering ")
		server.Stop()
		manager.Stop()
		fmt.Println("stop finish ")

	}()
	err = server.Listen()
	if err != nil {
		panic(err)
	}
	kernel.RegisterServer()
	fmt.Println("node", kernel.applicaton.GetConfig().Snow.Node)
	server.AcceptLoop()
}
