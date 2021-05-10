package logic

import (
	"context"
	"fmt"
	"git.code.oa.com/video_app_short_video/hello_alice/dao"
	"git.code.oa.com/video_app_short_video/hello_alice/model"
	"strconv"
	"time"

	_ "git.code.oa.com/trpc-go/trpc-codec/videopacket"
	_ "git.code.oa.com/trpc-go/trpc-filter/validation"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/codec"
	"git.code.oa.com/trpc-go/trpc-go/log"
	pb "git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter"
	ufr "git.code.oa.com/trpcprotocol/video_app_short_video/trpc_ugc_follow_read_jce_ugc_follow_read"
	ugcfi "git.code.oa.com/video_app_short_video/short_video_trpc_proto/ugc_follow_inner"
	union "git.code.oa.com/videocommlib/trpc-go-union" //trpc-go 操作union包
	p "git.code.oa.com/videocommlib/videopacket-go"
)
const (
	defaultUserId = 536164684
	fromUserId = 536164684
	toUserId = 2454008777
)
// BatchGetFansCount 批量获取粉丝数接口
func BatchGetFansCount() (ufrRsp *ufr.BathGetFansCountResponse, err error){
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
	return proxy.BathGetFansCount(context.Background(),ufrReq)
}

// GetUserInfo 调用jce服务接口，获得用户关注信息
func GetUserInfo(req *pb.HelloRequest)(ugcfiRsp interface{},err error) {
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
			FromUser: ugcfi.User{UserId: fromUserId },ToUser: ugcfi.User{UserId: toUserId},
		}
		return proxy.SetFollow(context.Background(),ugcfiReq,opts...)
	case 0xf3cb:
		// 取消关注
		ugcfiReq := &ugcfi.FollowReq{
			FromUser: ugcfi.User{UserId:fromUserId },ToUser: ugcfi.User{UserId: toUserId},
		}
		return proxy.DelFollow(context.Background(),ugcfiReq,opts...)
	case 0xf3cc:
		// 获取关注列表（翻页）
		ugcfiReq := &ugcfi.QueryFollowVppsReq{
			User:ugcfi.User{ UserId: defaultUserId },
		}
		return proxy.QueryFollowVpps(context.Background(),ugcfiReq,opts...)
	case 0xf3cd:
		// 获取粉丝列表信息(翻页)
		ugcfiReq := &ugcfi.QueryFansListReq{
			User:ugcfi.User{ UserId: toUserId },
		}
		return proxy.QueryFansList(context.Background(),ugcfiReq,opts...)
	case 0xf3ce:
		// 查询a,b间的关注关系,1-关注、0-没关注和2-相互关注
		ugcfiReq := &ugcfi.CheckFollowRelationReq{
			From:ugcfi.User{ UserId:fromUserId },ToUsers:[]ugcfi.User{{ UserId: toUserId}},
		}
		return proxy.CheckFollowRelation(context.Background(),ugcfiReq,opts...)
	case 0xf18d:
		// 查询a是否关注b,0-没关注，1-关注
		ugcfiReq := &ugcfi.CheckFollowReq{
			From:ugcfi.User{ UserId:fromUserId },ToUsers:[]ugcfi.User{{ UserId: toUserId}},
		}
		return proxy.CheckFollow(context.Background(),ugcfiReq,opts...)
	case 0xf19b:
		// 查询a是否是b的粉丝
		ugcfiReq := &ugcfi.CheckFansReq{
			From:ugcfi.User{ UserId:fromUserId },ToUsers:[]ugcfi.User{{ UserId: toUserId}},
		}
		return proxy.CheckFans(context.Background(),ugcfiReq,opts...)
	case 0xf3cf:
		// 查询关注数
		ugcfiReq := &ugcfi.GetFollowCountReq{
			User:ugcfi.User{ UserId: defaultUserId} ,
		}
		return proxy.GetFollowCount(context.Background(),ugcfiReq,opts...)
	case 0xf3d0:
		// 查询粉丝数
		ugcfiReq := &ugcfi.GetFansCountReq{
			User:ugcfi.User{ UserId:defaultUserId } ,
		}
		return proxy.GetFansCount(context.Background(),ugcfiReq,opts...)
	default:
		return nil,fmt.Errorf("%s,not found command",req.Msg)
	}
}

//ReadUnion 读取union 2071表的数据,2071为社区号用户信息
func ReadUnion2071(req *pb.HelloRequest)(unionRsp map[string]model.SocietyUserInfoUnion2071,err error){
	//输入用户id
	vuid := req.Msg
	unionRsp = make(map[string]model.SocietyUserInfoUnion2071)
	//初始化union proxy
	proxy := union.NewParamUnionProxy("union",8,20002564,"0993ef6bbd651722","")
	//调用proxy，返回定义好的数据类型SocietyUserInfoUnion2071
	err = proxy.GetUnion(uint32(2071),[]string{vuid},unionRsp,
	client.WithNamespace("Production"),
	client.WithServiceName("trpc.union.union.union"),// service name自己随便填，主要用于监控上报和寻找配置项
	client.WithTarget("polaris://243969:65536"),
//	client.WithTimeout(800),
	)
	if err!=nil{
		log.Errorf("GetFromUnion2071 vuid:%v, err: %v", vuid, err)
		return nil, err
	}
	_ , ok := unionRsp[vuid]
	if !ok {
		return nil,fmt.Errorf("vuid info not exists", vuid)
	}
	return unionRsp,nil
}
//
func AcessRedis(ctx context.Context)(string,error){
	stringRsp,err := dao.AcessRedisString(ctx)
	if err!=nil {
		log.Errorf("AcessRedisString error:%v", err)
		return "", err
	}
	redisRsp := fmt.Sprintf("string: %v",stringRsp)
	hashRsp,err := dao.AcessRedisHash(ctx)
	if err!=nil {
		log.Errorf("AcessRedisHash error:%v", err)
		return "", err
	}
	redisRsp += fmt.Sprintf("hash: %v",hashRsp)
	zsetRsp,err := dao.AcessRedisZset(ctx)
	if err!=nil {
		log.Errorf("AcessRedisZset error:%v", err)
		return "", err
	}
	redisRsp += fmt.Sprintf("zset:%v",zsetRsp)
	return redisRsp, nil
}

func AcessMysql(ctx context.Context)(string,error){
	//create table
	mysqlRsp, err := dao.AcessMysqlInit(ctx)
	if err != nil {
		err=fmt.Errorf("TestMysqlInit error, err:%+v", err)
		return "", err
	}
	rsp := fmt.Sprintf("Create:%s ",mysqlRsp)
	// insert
	mysqlRsp, err = dao.AcessMysqlInsert(ctx)
	if err != nil {
		err = fmt.Errorf("TestMysqlInsert error, err:%+v", err)
		return "", err
	}
	rsp += fmt.Sprintf("Insert:%s ",mysqlRsp)
	// update
	mysqlRsp,err = dao.AcessMysqlUpdate(ctx)
	if err != nil {
		err = fmt.Errorf("TestMysqlUpdate error, err:%+v", err)
		return "", err
	}
	rsp += fmt.Sprintf("Update:%s ",mysqlRsp)
	// select
	mysqlRsp,err = dao.AcessMysqlSelect(ctx)
	if err != nil {
		err = fmt.Errorf("TestMysqlSelect error, err:%+v", err)
		return "", err
	}
	rsp += fmt.Sprintf("Selete:%s ",mysqlRsp)
	// delete
	mysqlRsp,err = dao.AcessMysqlDelete(ctx)
	if err != nil {
		err = fmt.Errorf("TestMysqlDelete error, err:%+v", err)
		return "", err
	}
	rsp += fmt.Sprintf("Delete:%s ",mysqlRsp)
	return rsp, nil
}

func AcessWuji(id string) (string,error) {
	rsp,err := dao.GetWujiContent(id)
	if err != nil {
		log.Errorf("AcessWuji error, err:%+v", err)
		return "",err
	}
	return rsp,nil
}
func AcessKafka(ctx context.Context) (string,error) {
	rsp,err := dao.ProcedueKafka(ctx)
	if err != nil {
		log.Errorf("AcessKafka error, err:%+v", err)
		return "",err
	}
	return rsp,nil
}