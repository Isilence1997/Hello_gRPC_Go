package dao

import (
	"fmt"

	"git.code.oa.com/gRPC-go/gRPC-go/client"
	union "git.code.oa.com/videocommlib/gRPC-go-union"

	"git.code.oa.com/video_app_short_video/hello_alice/model"

	"git.code.oa.com/video_app_short_video/hello_alice/config"
)
var (
	// union 代理
	unionProxy union.Proxy
	serviceConfig  config.ServiceConfig
)

// 初始化union proxy
func initUnionProxy(){
	// 获取配置信息
	serviceConfig = config.GetConfig()
	unionConfig := serviceConfig.Union
	unionProxy = union.NewParamUnionProxy("union",8,unionConfig.AppId,unionConfig.AppKey,"")
}

// ReadUnion2071 读取union 2071表，获取社区用户信息
func ReadUnion2071(vuid string)(unionRsp map[string]model.SocietyUserInfoUnion2071,err error){
	unionRsp = make(map[string]model.SocietyUserInfoUnion2071)
	initUnionProxy()
	unionConfig := serviceConfig.Union
	//调用proxy，获取2071表中的社区用户信息
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
