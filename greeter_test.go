package main

import (
	"context"
	"testing"

	trpc "git.code.oa.com/trpc-go/trpc-go"
	_ "git.code.oa.com/trpc-go/trpc-go/http"
	_ "git.code.oa.com/trpc-go/trpc-selector-cl5"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	pb "git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter"
)

var greeterService = &greeterServiceImpl{}

//go:generate go mod tidy
//go:generate mockgen -destination=stub/git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter/greeter_mock.go -package=hello_alice_greeter -self_package=git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter GreeterClientProxy

func Test_Greeter_SayHello(t *testing.T) {

	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	greeterClientProxy := pb.NewMockGreeterClientProxy(ctrl)

	// 预期行为
	m := greeterClientProxy.EXPECT().SayHello(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	m.DoAndReturn(func(ctx context.Context, req interface{}, opts ...interface{}) (interface{}, error) {

		r, ok := req.(*pb.HelloRequest)
		if !ok {
			panic("invalid request")
		}

		rsp := &pb.HelloReply{}
		err := greeterService.SayHello(trpc.BackgroundContext(), r, rsp)
		return rsp, err
	})

	// 开始写单元测试逻辑
	req := &pb.HelloRequest{}

	rsp, err := greeterClientProxy.SayHello(trpc.BackgroundContext(), req)

	// 输出入参和返回 (检查t.Logf输出，运行 `go test -v`)
	t.Logf("Greeter_SayHello req: %v", req)
	t.Logf("Greeter_SayHello rsp: %v, err: %v", rsp, err)

	assert.Nil(t, err)
}
