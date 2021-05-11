package dao

import (
	"fmt"
	wuji "git.code.oa.com/open-wuji/go-sdk/wujiclient"
	"git.code.oa.com/video_app_short_video/hello_alice/model"
	"git.code.oa.com/video_app_short_video/hello_alice/config"
)

var (
	// 接入无极的client接口
	wujiClientProxy wuji.ClientInterface
)

//初始化wuji
func InitWujiProxy()(err error){
	serviceConfig := config.GetConfig()
	wujiConfig := serviceConfig.Wuji
	wujiClientProxy, err = wuji.NewClient(
		wuji.WithAppID(wujiConfig.AppID),
		wuji.WithSchemaID(wujiConfig.SchemaID),
		wuji.WithSchemaKey(wujiConfig.SchemaKey),
		wuji.WithRequestURL(wujiConfig.RequestURL),
		wuji.WithRequestTarget(wujiConfig.Target),  // L5寻址
		// wuji.WithRequestDirect(),   // 或者直接 dns 寻址
	)
	if err != nil {
		err = fmt.Errorf("initWujiClient error: %v",err)
	}
	return err
}

//获取无极表内容
func GetWujiContent(id string)(string,error){
	var obj model.WujiData
	err := wujiClientProxy.Get(id,&obj)//得到指定id的内容
	if err != nil {
		return "", err
	}
	wujiRsp := fmt.Sprintf("Get: id:%d name:%s ",obj.ID,obj.Name)
	raws := wujiClientProxy.GetKeys()//获取所有的key
	if raws == nil || len(raws) == 0{
		return "",fmt.Errorf("GetKeys error")
	}
	wujiRsp += fmt.Sprintf(" GetKeys(): %s",raws)
	return wujiRsp, nil
}
