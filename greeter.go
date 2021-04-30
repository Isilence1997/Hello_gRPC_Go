package main

import (
	"context"
	"fmt"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/codec"
	"git.code.oa.com/trpc-go/trpc-go/log"
	"strconv"
	"time"

	_ "git.code.oa.com/trpc-go/trpc-codec/videopacket"
	_ "git.code.oa.com/trpc-go/trpc-filter/validation"
	pb "git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter"
	ufr "git.code.oa.com/trpcprotocol/video_app_short_video/trpc_ugc_follow_read_jce_ugc_follow_read"
	ugcfi "git.code.oa.com/video_app_short_video/short_video_trpc_proto/ugc_follow_inner"
	p "git.code.oa.com/videocommlib/videopacket-go"
)

// SayHello 响应HelloRequest
func (s *greeterServiceImpl) SayHello(_ context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error {
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
	rsp.Msg = "[echo] " + fmt.Sprintf("FansCountMap : %v ",ufrRsp.FansCountMap)
	return nil
}

// GetUserInfo 调用jce服务接口，获得用户关注信息
func (s *greeterServiceImpl) GetUserInfo(_ context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error{
	packet := p.NewVideoPacket()
	Command, _ := strconv.ParseInt(req.Msg, 0, 64) //convert strings to int64
	packet.CommHeader.BasicInfo.Command = int16(Command)
	proxy := ugcfi.NewUgcFollowInnerServiceProxy("")
	opts := []client.Option{
		client.WithReqHead(packet),
		client.WithProtocol("videopacket"),
		client.WithNetwork("tcp4"),
		client.WithNamespace("Development"),
		client.WithSerializationType(codec.SerializationTypeJCE),
		client.WithTarget("polaris://trpc.video_app_short_video.trpc_ugc_follow_read_jce.UgcFollowRead"),
		//client.WithServiceName("trpc.video_app_short_video.trpc_ugc_follow_read_jce.UgcFollowRead"),
	}
	switch Command {
	case 0xf3ca:
		// 关注
		ugcfiReq := &ugcfi.FollowReq{
			FromUser: ugcfi.User{UserId: 536164684 },ToUser: ugcfi.User{UserId: 2454008777},
		}
		ugcfiRsp,err := proxy.SetFollow(context.Background(),ugcfiReq,opts...)
		if err != nil{
			log.Errorf("%+v",err)
			return err
		}
		log.Infof("%+v", ugcfiRsp)
		rsp.Msg = fmt.Sprintf("Result: %v",ugcfiRsp.Result)
	case 0xf3cb:
		// 取消关注
		ugcfiReq := &ugcfi.FollowReq{
			FromUser: ugcfi.User{UserId:536164684 },ToUser: ugcfi.User{UserId:2454008777},
		}
		ugcfiRsp,err := proxy.DelFollow(context.Background(),ugcfiReq,opts...)
		if err != nil{
			log.Errorf("%+v",err)
			return err
		}
		log.Infof("%+v", ugcfiRsp)
		rsp.Msg = fmt.Sprintf("Result: %v",ugcfiRsp.Result)
	case 0xf3cc:
		// 获取关注列表（翻页）
		ugcfiReq := &ugcfi.QueryFollowVppsReq{
			User:ugcfi.User{ UserId:536164684 },
		}
		ugcfiRsp,err := proxy.QueryFollowVpps(context.Background(),ugcfiReq,opts...)
		if err != nil{
			log.Errorf("%+v",err)
			return err
		}
		log.Infof("%+v", ugcfiRsp)
		rsp.Msg = fmt.Sprintf("%v",ugcfiRsp)
	case 0xf3cd:
		// 获取粉丝列表信息(翻页)
		ugcfiReq := &ugcfi.QueryFansListReq{
			User:ugcfi.User{ UserId:2454008777 },
		}
		ugcfiRsp,err := proxy.QueryFansList(context.Background(),ugcfiReq,opts...)
		if err != nil{
			log.Errorf("%+v",err)
			return err
		}
		log.Info(ugcfiRsp)
		rsp.Msg = fmt.Sprintf("%v",ugcfiRsp.VecFollowedUser)
	case 0xf3ce:
		// 查询a,b间的关注关系,1-关注、0-没关注和2-相互关注
		ugcfiReq := &ugcfi.CheckFollowRelationReq{
			From:ugcfi.User{ UserId:536164684 },ToUsers:[]ugcfi.User{{ UserId:742676956}},
		}
		ugcfiRsp,err := proxy.CheckFollowRelation(context.Background(),ugcfiReq,opts...)
		if err != nil{
			log.Errorf("%+v",err)
			return err
		}
		log.Info(ugcfiRsp)
		rsp.Msg = fmt.Sprintf("%v",ugcfiRsp)
	case 0xf18d:
		// 查询a是否关注b,0-没关注，1-关注
		ugcfiReq := &ugcfi.CheckFollowReq{
			From:ugcfi.User{ UserId:536164684 },ToUsers:[]ugcfi.User{{ UserId: 2454008777}},
		}
		ugcfiRsp,err := proxy.CheckFollow(context.Background(),ugcfiReq,opts...)
		if err != nil{
			log.Errorf("%+v",err)
			return err
		}
		log.Info(ugcfiRsp)
		rsp.Msg = fmt.Sprintf("%v",ugcfiRsp)
	case 0xf19b:
		// 查询a是否是b的粉丝
		ugcfiReq := &ugcfi.CheckFansReq{
			From:ugcfi.User{ UserId:536164684 },ToUsers:[]ugcfi.User{{ UserId: 2454008777}},
		}
		ugcfiRsp,err := proxy.CheckFans(context.Background(),ugcfiReq,opts...)
		if err != nil{
			log.Errorf("%+v",err)
			return err
		}
		log.Info(ugcfiRsp)
		rsp.Msg = fmt.Sprintf("%v",ugcfiRsp)
	case 0xf3cf:
		// 查询关注数
		ugcfiReq := &ugcfi.GetFollowCountReq{
			User:ugcfi.User{ UserId: 536164684} ,
		}
		ugcfiRsp,err := proxy.GetFollowCount(context.Background(),ugcfiReq,opts...)
		if err != nil{
			log.Errorf("%+v",err)
			return err
		}
		log.Info(ugcfiRsp)
		rsp.Msg = fmt.Sprintf("%v",ugcfiRsp.IFollowCount)
	case 0xf3d0:
		// 查询粉丝数
		ugcfiReq := &ugcfi.GetFansCountReq{
			User:ugcfi.User{ UserId:536164684 } ,
		}
		ugcfiRsp,err := proxy.GetFansCount(context.Background(),ugcfiReq,opts...)
		if err != nil{
			log.Errorf("%+v",err)
			return err
		}
		log.Info(ugcfiRsp)
		rsp.Msg = fmt.Sprintf("%v",ugcfiRsp.LFansCount)
	default:
		return fmt.Errorf("not found command")
	}
	return nil
}


