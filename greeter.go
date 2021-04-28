package main

import (
	"context"

	pb "git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter"

	_ "git.code.oa.com/trpc-go/trpc-filter/validation"
	_ "git.code.oa.com/trpcprotocol/video_app_short_video/trpc_ugc_follow_read_jce_ugc_follow_read"
)
// SayHello 响应HelloRequest
func (s *greeterServiceImpl) SayHello(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error {
	// implement business logic here ...
	// ...
	rsp.Msg = "hello,I am tRPC-go sever."//填充响应内容
	return nil
}
// GetUserInfo 显示客户端的请求内容
func (s *greeterServiceImpl) GetUserInfo(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error {
	// implement business logic here ...
	// ...
	rsp.Msg = "echo: " + req.GetMsg()//填充响应内容
	return nil
}

//func (s *greeterServiceImpl) BathGetFansCount(ctx context.Context, req *pb., rsp *pb.HelloReply)