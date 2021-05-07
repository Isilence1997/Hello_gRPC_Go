package main

import (
	"context"
	"fmt"

	"git.code.oa.com/video_app_short_video/hello_alice/logic"

	"git.code.oa.com/trpc-go/trpc-go/log"
	_ "git.code.oa.com/trpc-go/trpc-codec/videopacket"
	_ "git.code.oa.com/trpc-go/trpc-filter/validation"

	pb "git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter"
)

// SayHello 响应HelloRequest
func (s *greeterServiceImpl) SayHello(_ context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error {
	//填充响应内容
	rsp.Msg = "echo: " + req.GetMsg()
	return nil
}

// BathGetFansCount 批量获取粉丝数接口
func (s *greeterServiceImpl) BathGetFansCount(_ context.Context, _ *pb.HelloRequest, rsp *pb.HelloReply) error{
	ufrRsp,err := logic.BatchGetFansCount()
	if err != nil {
		log.Error(err)
		log.Info(err)
		return err
	}
	rsp.Msg = "[echo] " + fmt.Sprintf("FansCountMap : %v ",ufrRsp.FansCountMap)
	return nil
}

// GetUserInfo 调用jce服务接口，获得用户关注信息
func (s *greeterServiceImpl) GetUserInfo(_ context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error{
	ugcfiRsp,err := logic.GetUserInfo(req)
	if err != nil{
		log.Errorf("%+v",err)
		return err
	}
	log.Infof("%+v", ugcfiRsp)
	rsp.Msg = fmt.Sprintf("Result: %v",ugcfiRsp)
	return nil
}

//ReadUnion 读取union 2071表,获取用户头像和用户昵称
func (s *greeterServiceImpl) ReadUnion(_ context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error{
	unionRsp,err :=logic.ReadUnion2071(req)
	if err!=nil {
		log.Errorf("GetUnion2071 error, vuid:%s, err:%v",req.Msg,err)
		rsp.Msg = err.Error()
		return err
	}
	log.Infof("%v", unionRsp)
	rsp.Msg = fmt.Sprintf("UserHead: %v UserNick: %v", unionRsp[req.Msg].UserHead,unionRsp[req.Msg].UserNick)
	return nil
}

// AcessRedis 对redis进行操作
func (s *greeterServiceImpl) AcessRedis(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error {
	redisRsp, err := logic.AcessRedis(ctx)
	if err != nil {
		log.Errorf("%v",err)
		return err
	}
	rsp.Msg = redisRsp
	return nil
}