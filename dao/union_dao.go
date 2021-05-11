package dao

import (
	"fmt"

	"git.code.oa.com/trpc-go/trpc-go/client"
	union "git.code.oa.com/videocommlib/trpc-go-union"

	"git.code.oa.com/video_app_short_video/hello_alice/model"

	"git.code.oa.com/video_app_short_video/hello_alice/config"
)
var (
	// union 代理
	unionProxy union.Proxy
	serviceConfig  config.ServiceConfig
)

func initUnionProxy(){
	//初始化union proxy
	// 获取配置信息
	serviceConfig = config.GetConfig()
	unionConfig := serviceConfig.Union
	unionProxy = union.NewParamUnionProxy("union",8,unionConfig.AppId,unionConfig.AppKey,"")
}

func ReadUnion2071(vuid string)(unionRsp map[string]model.SocietyUserInfoUnion2071,err error){
	unionRsp = make(map[string]model.SocietyUserInfoUnion2071)
	initUnionProxy()
	//调用proxy，返回定义好的数据类型SocietyUserInfoUnion2071
	unionConfig := serviceConfig.Union
	err = unionProxy.GetUnion(uint32(2071),[]string{vuid},unionRsp,
		client.WithNamespace(unionConfig.Namespace),
		client.WithServiceName(unionConfig.ServiceName),// service name自己随便填，主要用于监控上报和寻找配置项
		client.WithTarget(unionConfig.Target),
		client.WithTimeout(800),
	)
	if err!=nil{
		return nil, err
	}
	_ , ok := unionRsp[vuid]
	if !ok {
		return nil,fmt.Errorf("vuid info not exists: %s", vuid)
	}
	return unionRsp,nil
}
