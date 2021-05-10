package main

import (
	"context"
	"testing"

	"git.code.oa.com/trpc-go/trpc-go"
	_ "git.code.oa.com/trpc-go/trpc-go/http"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	pb "git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter"
)

var greeterService = &greeterServiceImpl{}

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

func Test_Greeter_BathGetFansCount(t *testing.T) {

	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	greeterClientProxy := pb.NewMockGreeterClientProxy(ctrl)

	// 预期行为
	m := greeterClientProxy.EXPECT().BathGetFansCount(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	m.DoAndReturn(func(ctx context.Context, req interface{}, opts ...interface{}) (interface{}, error) {

		r, ok := req.(*pb.HelloRequest)
		if !ok {
			panic("invalid request")
		}

		rsp := &pb.HelloReply{}
		err := greeterService.BathGetFansCount(trpc.BackgroundContext(), r, rsp)
		return rsp, err
	})

	// 开始写单元测试逻辑
	req := &pb.HelloRequest{}

	rsp, err := greeterClientProxy.BathGetFansCount(trpc.BackgroundContext(), req)

	// 输出入参和返回 (检查t.Logf输出，运行 `go test -v`)
	t.Logf("Greeter_BathGetFansCount req: %v", req)
	t.Logf("Greeter_BathGetFansCount rsp: %v, err: %v", rsp, err)

	assert.Nil(t, err)
}

func Test_Greeter_GetUserInfo(t *testing.T) {

	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	greeterClientProxy := pb.NewMockGreeterClientProxy(ctrl)

	// 预期行为
	m := greeterClientProxy.EXPECT().GetUserInfo(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	m.DoAndReturn(func(ctx context.Context, req interface{}, opts ...interface{}) (interface{}, error) {

		r, ok := req.(*pb.HelloRequest)
		if !ok {
			panic("invalid request")
		}

		rsp := &pb.HelloReply{}
		err := greeterService.GetUserInfo(trpc.BackgroundContext(), r, rsp)
		return rsp, err
	})

	// 开始写单元测试逻辑
	req := &pb.HelloRequest{}

	rsp, err := greeterClientProxy.GetUserInfo(trpc.BackgroundContext(), req)

	// 输出入参和返回 (检查t.Logf输出，运行 `go test -v`)
	t.Logf("Greeter_GetUserInfo req: %v", req)
	t.Logf("Greeter_GetUserInfo rsp: %v, err: %v", rsp, err)

	assert.Nil(t, err)
}

func Test_Greeter_ReadUnion(t *testing.T) {

	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	greeterClientProxy := pb.NewMockGreeterClientProxy(ctrl)

	// 预期行为
	m := greeterClientProxy.EXPECT().ReadUnion(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	m.DoAndReturn(func(ctx context.Context, req interface{}, opts ...interface{}) (interface{}, error) {

		r, ok := req.(*pb.HelloRequest)
		if !ok {
			panic("invalid request")
		}

		rsp := &pb.HelloReply{}
		err := greeterService.ReadUnion(trpc.BackgroundContext(), r, rsp)
		return rsp, err
	})

	// 开始写单元测试逻辑
	req := &pb.HelloRequest{}

	rsp, err := greeterClientProxy.ReadUnion(trpc.BackgroundContext(), req)

	// 输出入参和返回 (检查t.Logf输出，运行 `go test -v`)
	t.Logf("Greeter_ReadUnion req: %v", req)
	t.Logf("Greeter_ReadUnion rsp: %v, err: %v", rsp, err)

	assert.Nil(t, err)
}

func Test_Greeter_AcessRedis(t *testing.T) {

	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	greeterClientProxy := pb.NewMockGreeterClientProxy(ctrl)

	// 预期行为
	m := greeterClientProxy.EXPECT().AcessRedis(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	m.DoAndReturn(func(ctx context.Context, req interface{}, opts ...interface{}) (interface{}, error) {

		r, ok := req.(*pb.HelloRequest)
		if !ok {
			panic("invalid request")
		}

		rsp := &pb.HelloReply{}
		err := greeterService.AcessRedis(trpc.BackgroundContext(), r, rsp)
		return rsp, err
	})

	// 开始写单元测试逻辑
	req := &pb.HelloRequest{}

	rsp, err := greeterClientProxy.AcessRedis(trpc.BackgroundContext(), req)

	// 输出入参和返回 (检查t.Logf输出，运行 `go test -v`)
	t.Logf("Greeter_AcessRedis req: %v", req)
	t.Logf("Greeter_AcessRedis rsp: %v, err: %v", rsp, err)

	assert.Nil(t, err)
}

func Test_Greeter_AcessMysql(t *testing.T) {

	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	greeterClientProxy := pb.NewMockGreeterClientProxy(ctrl)

	// 预期行为
	m := greeterClientProxy.EXPECT().AcessMysql(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	m.DoAndReturn(func(ctx context.Context, req interface{}, opts ...interface{}) (interface{}, error) {

		r, ok := req.(*pb.HelloRequest)
		if !ok {
			panic("invalid request")
		}

		rsp := &pb.HelloReply{}
		err := greeterService.AcessMysql(trpc.BackgroundContext(), r, rsp)
		return rsp, err
	})

	// 开始写单元测试逻辑
	req := &pb.HelloRequest{}

	rsp, err := greeterClientProxy.AcessMysql(trpc.BackgroundContext(), req)

	// 输出入参和返回 (检查t.Logf输出，运行 `go test -v`)
	t.Logf("Greeter_AcessMysql req: %v", req)
	t.Logf("Greeter_AcessMysql rsp: %v, err: %v", rsp, err)

	assert.Nil(t, err)
}

func Test_Greeter_AcessWuji(t *testing.T) {

	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	greeterClientProxy := pb.NewMockGreeterClientProxy(ctrl)

	// 预期行为
	m := greeterClientProxy.EXPECT().AcessWuji(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	m.DoAndReturn(func(ctx context.Context, req interface{}, opts ...interface{}) (interface{}, error) {

		r, ok := req.(*pb.HelloRequest)
		if !ok {
			panic("invalid request")
		}

		rsp := &pb.HelloReply{}
		err := greeterService.AcessWuji(trpc.BackgroundContext(), r, rsp)
		return rsp, err
	})

	// 开始写单元测试逻辑
	req := &pb.HelloRequest{}

	rsp, err := greeterClientProxy.AcessWuji(trpc.BackgroundContext(), req)

	// 输出入参和返回 (检查t.Logf输出，运行 `go test -v`)
	t.Logf("Greeter_AcessWuji req: %v", req)
	t.Logf("Greeter_AcessWuji rsp: %v, err: %v", rsp, err)

	assert.Nil(t, err)
}

func Test_Greeter_AcessKafka(t *testing.T) {

	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	greeterClientProxy := pb.NewMockGreeterClientProxy(ctrl)

	// 预期行为
	m := greeterClientProxy.EXPECT().AcessKafka(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	m.DoAndReturn(func(ctx context.Context, req interface{}, opts ...interface{}) (interface{}, error) {

		r, ok := req.(*pb.HelloRequest)
		if !ok {
			panic("invalid request")
		}

		rsp := &pb.HelloReply{}
		err := greeterService.AcessKafka(trpc.BackgroundContext(), r, rsp)
		return rsp, err
	})

	// 开始写单元测试逻辑
	req := &pb.HelloRequest{}

	rsp, err := greeterClientProxy.AcessKafka(trpc.BackgroundContext(), req)

	// 输出入参和返回 (检查t.Logf输出，运行 `go test -v`)
	t.Logf("Greeter_AcessKafka req: %v", req)
	t.Logf("Greeter_AcessKafka rsp: %v, err: %v", rsp, err)

	assert.Nil(t, err)
}
