package main

import (
	"git.code.oa.com/trpc-go/trpc-database/kafka"
	_ "go.uber.org/automaxprocs"

	_ "git.code.oa.com/trpc-go/trpc-config-tconf"
	_ "git.code.oa.com/trpc-go/trpc-filter/debuglog"
	_ "git.code.oa.com/trpc-go/trpc-filter/recovery"
	_ "git.code.oa.com/trpc-go/trpc-log-atta"
	_ "git.code.oa.com/trpc-go/trpc-metrics-m007"
	_ "git.code.oa.com/trpc-go/trpc-metrics-runtime"
	_ "git.code.oa.com/trpc-go/trpc-naming-polaris"
	_ "git.code.oa.com/trpc-go/trpc-opentracing-tjg"
	_ "git.code.oa.com/trpc-go/trpc-selector-cl5"

	"git.code.oa.com/trpc-go/trpc-go/log"

	trpc "git.code.oa.com/trpc-go/trpc-go"
	pb "git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter"
	"git.code.oa.com/video_app_short_video/hello_alice/dao"
)

type greeterServiceImpl struct{}

// 服务初始化
func ServiceInit() {
	// 初始化mysql
	err := dao.InitMysqlProxy()
	if err != nil {
		panic(err)
	}
	// 初始化redis
	err = dao.InitRedisProxy()
	if err!=nil{
		panic(err)
	}
	// 初始化wuji
	if err := dao.InitWujiProxy(); err != nil {
		panic(err)
	}
	// 初始化union
}
func main() {

	s := trpc.NewServer()
	ServiceInit()
	// 注册kafka消费handler,多个service的情况下 kafka.RegisterHandlerService(s.Service("name"), handle)
	// 没有指定name的情况，代表所有service共用同一个handler
	kafka.RegisterHandlerService(s.Service("trpc.video_app_short_video.hello_alice.consumer"), dao.ConsumeKafkaMsgHandler)
	pb.RegisterGreeterService(s, &greeterServiceImpl{})

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
