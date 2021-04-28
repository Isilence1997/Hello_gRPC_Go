package main

import (
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
)

type greeterServiceImpl struct{}

func main() {

	s := trpc.NewServer()

	pb.RegisterGreeterService(s, &greeterServiceImpl{})

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
