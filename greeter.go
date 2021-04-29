package main

import (
	"context"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/log"
	"time"
	"fmt"

	pb "git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter"

	_ "git.code.oa.com/trpc-go/trpc-filter/validation"
	ufr "git.code.oa.com/trpcprotocol/video_app_short_video/trpc_ugc_follow_read_jce_ugc_follow_read"
)

// SayHello 响应HelloRequest
func (s *greeterServiceImpl) SayHello(_ context.Context, _ *pb.HelloRequest, rsp *pb.HelloReply) error {
	//填充响应内容
	rsp.Msg = "hello,I am tRPC-go sever."
	return nil
}
// GetUserInfo 显示客户端的请求内容
func (s *greeterServiceImpl) GetUserInfo(_ context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error {
	//填充响应内容
	rsp.Msg = "echo: " + req.GetMsg()
	return nil
}
// BathGetFansCount 批量获取粉丝数接口
func (s *greeterServiceImpl) BathGetFansCount(_ context.Context, _ *pb.HelloRequest, rsp *pb.HelloReply) error{
	// proxy 客户端调用桩函数或者调用代理，由trpc工具自动生成，内部调用client
	proxy := ufr.NewUgcFollowReadClientProxy(
		client.WithProtocol("trpc"),
		client.WithNetwork("tcp4"),
		//target 后端服务的地址，规则为 selectorname://endpoint
		client.WithTarget("polaris://trpc.video_app_short_video.trpc_ugc_follow_read_jce.UgcFollowReadPb"),
		client.WithNamespace("Development"),
		client.WithTimeout(time.Millisecond*500),
		)
	// ufrReq 用BathGetFansCountRequest构造请求
	ufrReq := &ufr.BathGetFansCountRequest{
		Vuids: []int64{536164684},
	}
	// 通过proxy调用UgcFollowRead服务接口
	ufrRsp,err := proxy.BathGetFansCount(context.Background(),ufrReq)
	if err != nil {
		log.Error(err)
		log.Info(err)
		return err
	}
	rsp.Msg = "[echo] " + fmt.Sprintf("%#v", ufrRsp)
	return nil
}