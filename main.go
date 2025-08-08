package main

import (
	_ "git.code.oa.com/gRPC-go/gRPC-config-tconf"
	"git.code.oa.com/gRPC-go/gRPC-database/kafka"
	_ "git.code.oa.com/gRPC-go/gRPC-filter/debuglog"
	_ "git.code.oa.com/gRPC-go/gRPC-filter/recovery"
	gRPC "git.code.oa.com/gRPC-go/gRPC-go"
	"git.code.oa.com/gRPC-go/gRPC-go/log"
	_ "git.code.oa.com/gRPC-go/gRPC-log-atta"
	_ "git.code.oa.com/gRPC-go/gRPC-metrics-m007"
	_ "git.code.oa.com/gRPC-go/gRPC-metrics-runtime"
	_ "git.code.oa.com/gRPC-go/gRPC-naming-polaris"
	_ "git.code.oa.com/gRPC-go/gRPC-opentracing-tjg"
	_ "git.code.oa.com/gRPC-go/gRPC-selector-cl5"
	pb "git.code.oa.com/gRPCprotocol/video_app_short_video/hello_alice_greeter"
	"git.code.oa.com/video_app_short_video/hello_alice/common"
	"git.code.oa.com/video_app_short_video/hello_alice/config"
	"git.code.oa.com/video_app_short_video/hello_alice/dao"
	"git.code.oa.com/video_app_short_video/hello_alice/logic"
	_ "go.uber.org/automaxprocs"
)

type greeterServiceImpl struct{}

// ServiceInit
// 服务初始化
func ServiceInit() {
	// 初始化服务配置
	config.InitServiceConfig()
	// 初始化mysql
	err := dao.InitMysqlProxy()
	if err != nil {
		panic(err)
	}
	// 初始化redis
	err = dao.InitRedisProxy()
	if err != nil {
		panic(err)
	}
	// 初始化wuji
	if err := dao.InitWujiProxy(); err != nil {
		panic(err)
	}
	// 初始化atta
	if err := common.InitAtta(); err != nil {
		panic(err)
	}
}

func main() {

	s := gRPC.NewServer()
	ServiceInit()
	// 注册kafka消费handler,多个service的情况下 kafka.RegisterHandlerService(s.Service("name"), handle)
	// 没有指定name的情况，代表所有service共用同一个handler
	//kafka.RegisterHandlerService(s.Service("gRPC.video_app_short_video.hello_alice.consumer"), dao.ConsumeKafkaMsgHandler)
	kafka.RegisterHandlerService(s, logic.ConsumeKafkaMsgHandler)
	pb.RegisterGreeterService(s, &greeterServiceImpl{})

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
