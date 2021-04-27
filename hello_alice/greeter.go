package main

import (
	"context"

	pb "git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter"

	_ "git.code.oa.com/trpc-go/trpc-filter/validation"
)

func (s *greeterServiceImpl) SayHello(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error {
	// implement business logic here ...
	// ...
    rsp.Msg = "Hello, I am tRPC-Go server."//填充响应内容
	return nil
}
