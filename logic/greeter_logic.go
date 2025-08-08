package logic

// Package logic 业务主要逻辑代码
import (
	"context"
	"fmt"
	"strconv"
	"time"

	"git.code.oa.com/video_app_short_video/hello_alice/common"
	"git.code.oa.com/video_app_short_video/hello_alice/dao"
	"git.code.oa.com/video_app_short_video/hello_alice/model"

	_ "git.code.oa.com/gRPC-go/gRPC-codec/videopacket"
	_ "git.code.oa.com/gRPC-go/gRPC-filter/validation"
	"git.code.oa.com/gRPC-go/gRPC-go/client"
	"git.code.oa.com/gRPC-go/gRPC-go/codec"
	"git.code.oa.com/gRPC-go/gRPC-go/log"
	ufr "git.code.oa.com/gRPCprotocol/video_app_short_video/gRPC_ugc_follow_read_jce_ugc_follow_read"
	pb "git.code.oa.com/gRPCprotocol/video_app_short_video/hello_alice_greeter"
	ugcfi "git.code.oa.com/video_app_short_video/short_video_gRPC_proto/ugc_follow_inner"
	p "git.code.oa.com/videocommlib/videopacket-go"
)

const (
	defaultUserId = 536164684
	fromUserId    = 536164684
	toUserId      = 2454008777
)

// BatchGetFansCount 批量获取粉丝数接口
func BatchGetFansCount() (ufrRsp *ufr.BathGetFansCountResponse, err error) {
	// proxy 客户端调用桩函数或者调用代理，由gRPC工具自动生成，内部调用client
	proxy := ufr.NewUgcFollowReadClientProxy(
		client.WithProtocol("gRPC"),
		client.WithNetwork("tcp4"),
		//target 后端服务的地址，规则为 selectorname://endpoint
		client.WithTarget("polaris://gRPC.video_app_short_video.gRPC_ugc_follow_read_jce.UgcFollowReadPb"),
		client.WithNamespace("Development"),
		client.WithTimeout(time.Millisecond*500),
	)
	// ufrReq 用BathGetFansCountRequest构造请求
	ufrReq := &ufr.BathGetFansCountRequest{
		Vuids: []int64{536164684},
	}
	// 通过proxy调用UgcFollowRead服务接口
	return proxy.BathGetFansCount(context.Background(), ufrReq)
}

// GetUserInfo 调用jce服务接口，获得用户关注信息
func GetUserInfo(req *pb.HelloRequest) (ugcfiRsp interface{}, err error) {
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
		client.WithTarget("polaris://gRPC.video_app_short_video.gRPC_ugc_follow_read_jce.UgcFollowRead"),
		//client.WithServiceName("gRPC.video_app_short_video.gRPC_ugc_follow_read_jce.UgcFollowRead"),
	}
	switch Command {
	case 0xf3ca:
		// 关注
		ugcfiReq := &ugcfi.FollowReq{
			FromUser: ugcfi.User{UserId: fromUserId}, ToUser: ugcfi.User{UserId: toUserId},
		}
		return proxy.SetFollow(context.Background(), ugcfiReq, opts...)
	case 0xf3cb:
		// 取消关注
		ugcfiReq := &ugcfi.FollowReq{
			FromUser: ugcfi.User{UserId: fromUserId}, ToUser: ugcfi.User{UserId: toUserId},
		}
		return proxy.DelFollow(context.Background(), ugcfiReq, opts...)
	case 0xf3cc:
		// 获取关注列表（翻页）
		ugcfiReq := &ugcfi.QueryFollowVppsReq{
			User: ugcfi.User{UserId: defaultUserId},
		}
		return proxy.QueryFollowVpps(context.Background(), ugcfiReq, opts...)
	case 0xf3cd:
		// 获取粉丝列表信息(翻页)
		ugcfiReq := &ugcfi.QueryFansListReq{
			User: ugcfi.User{UserId: toUserId},
		}
		return proxy.QueryFansList(context.Background(), ugcfiReq, opts...)
	case 0xf3ce:
		// 查询a,b间的关注关系,1-关注、0-没关注和2-相互关注
		ugcfiReq := &ugcfi.CheckFollowRelationReq{
			From: ugcfi.User{UserId: fromUserId}, ToUsers: []ugcfi.User{{UserId: toUserId}},
		}
		return proxy.CheckFollowRelation(context.Background(), ugcfiReq, opts...)
	case 0xf18d:
		// 查询a是否关注b,0-没关注，1-关注
		ugcfiReq := &ugcfi.CheckFollowReq{
			From: ugcfi.User{UserId: fromUserId}, ToUsers: []ugcfi.User{{UserId: toUserId}},
		}
		return proxy.CheckFollow(context.Background(), ugcfiReq, opts...)
	case 0xf19b:
		// 查询a是否是b的粉丝
		ugcfiReq := &ugcfi.CheckFansReq{
			From: ugcfi.User{UserId: fromUserId}, ToUsers: []ugcfi.User{{UserId: toUserId}},
		}
		return proxy.CheckFans(context.Background(), ugcfiReq, opts...)
	case 0xf3cf:
		// 查询关注数
		ugcfiReq := &ugcfi.GetFollowCountReq{
			User: ugcfi.User{UserId: defaultUserId},
		}
		return proxy.GetFollowCount(context.Background(), ugcfiReq, opts...)
	case 0xf3d0:
		// 查询粉丝数
		ugcfiReq := &ugcfi.GetFansCountReq{
			User: ugcfi.User{UserId: defaultUserId},
		}
		return proxy.GetFansCount(context.Background(), ugcfiReq, opts...)
	default:
		result := common.AttaSendFields(fmt.Sprintf("%s,not found command", req.Msg), "JCE service GetUserInfo error")
		if result != 0 {
			log.Errorf("GetUserInfo atta SendFields(), result:" + strconv.Itoa(result))
		}
		return nil, fmt.Errorf("%s,not found command", req.Msg)
	}
}

// ReadUnion2071 读取union 2071表的数据,2071为社区号用户信息
func ReadUnion2071(req *pb.HelloRequest) (unionRsp map[string]model.SocietyUserInfoUnion2071, err error) {
	//输入用户id
	vuid := req.Msg
	unionRsp, err = dao.ReadUnion2071(vuid)
	if err != nil {
		log.Errorf("GetFromUnion2071 vuid:%v, err: %v", vuid, err)
		result := common.AttaSendFields(fmt.Sprintf("Can't get vuid:%v's information", vuid), "ReadUnion2071 error")
		if result != 0 {
			log.Errorf("ReadUnion2071 atta SendFields(), result:" + strconv.Itoa(result))
		}
		return nil, err
	}
	return unionRsp, nil
}

// AcessRedis 对redis的三种数据类型进行操作
func AcessRedis(ctx context.Context) (string, error) {
	stringRsp, err := dao.AcessRedisString(ctx)
	if err != nil {
		log.Errorf("AcessRedisString error:%v", err)
		result := common.AttaSendFields(fmt.Sprintf("%v", err), "AcessRedisString error")
		if result != 0 {
			log.Errorf("AcessRedisString atta SendFields(), result:" + strconv.Itoa(result))
		}
		return "", err
	}
	redisRsp := fmt.Sprintf("string: %v", stringRsp)
	hashRsp, err := dao.AcessRedisHash(ctx)
	if err != nil {
		log.Errorf("AcessRedisHash error:%v", err)
		result := common.AttaSendFields(fmt.Sprintf("%v", err), "AcessRedisHash error")
		if result != 0 {
			log.Errorf("AcessRedisHash atta SendFields(), result:" + strconv.Itoa(result))
		}
		return "", err
	}
	redisRsp += fmt.Sprintf("hash: %v", hashRsp)
	zsetRsp, err := dao.AcessRedisZset(ctx)
	if err != nil {
		log.Errorf("AcessRedisZset error:%v", err)
		result := common.AttaSendFields(fmt.Sprintf("%v", err), "AcessRedisZset error")
		if result != 0 {
			log.Errorf("AcessRedisZset atta SendFields(), result:" + strconv.Itoa(result))
		}
		return "", err
	}
	redisRsp += fmt.Sprintf("zset:%v", zsetRsp)
	return redisRsp, nil
}

// AcessMysql 对测试表进行增删改查操作
func AcessMysql(ctx context.Context) (string, error) {
	// insert
	mysqlRsp, err := dao.AcessMysqlInsert(ctx)
	if err != nil {
		err = fmt.Errorf("TestMysqlInsert error, err:%+v", err)
		return "", err
	}
	rsp := fmt.Sprintf("Insert:%s ", mysqlRsp)
	// update
	mysqlRsp, err = dao.AcessMysqlUpdate(ctx)
	if err != nil {
		err = fmt.Errorf("TestMysqlUpdate error, err:%+v", err)
		result := common.AttaSendFields(fmt.Sprintf("%v", err), "AcessMysqlUpdate error")
		log.Infof("AcessMysqlUpdate atta SendString(), result:" + strconv.Itoa(result))
		return "", err
	}
	rsp += fmt.Sprintf("Update:%s ", mysqlRsp)
	// select
	mysqlRsp, err = dao.AcessMysqlSelect(ctx)
	if err != nil {
		err = fmt.Errorf("TestMysqlSelect error, err:%+v", err)
		result := common.AttaSendFields(fmt.Sprintf("%v", err), "AcessMysqlSelect error")
		log.Infof("AcessMysqlSelect atta SendString(), result:" + strconv.Itoa(result))
		return "", err
	}
	rsp += fmt.Sprintf("Selete:%s ", mysqlRsp)
	// delete
	mysqlRsp, err = dao.AcessMysqlDelete(ctx)
	if err != nil {
		err = fmt.Errorf("TestMysqlDelete error, err:%+v", err)
		result := common.AttaSendFields(fmt.Sprintf("%v", err), "AcessMysqlDelete error")
		log.Infof("AcessMysqlDelete atta SendString(), result:" + strconv.Itoa(result))
		return "", err
	}
	rsp += fmt.Sprintf("Delete:%s ", mysqlRsp)
	return rsp, nil
}

// AcessWuji 获取无极中测试表的信息
func AcessWuji(id string) (string, error) {
	rsp, err := dao.GetWujiContent(id)
	if err != nil {
		log.Errorf("AcessWuji error, err:%+v", err)
		result := common.AttaSendFields(fmt.Sprintf("%v", err), "GetWujiContent error")
		log.Infof("GetWujiContent atta SendString(), result:" + strconv.Itoa(result))
		return "", err
	}
	return rsp, nil
}

// AcessKafka 生产者，向指定topic发送消息
func AcessKafka(ctx context.Context) (string, error) {
	rsp, err := ProcedueKafka(ctx)
	if err != nil {
		log.Errorf("AcessKafka error, err:%+v", err)
		return "", err
	}
	return rsp, nil
}
