package dao

import (
	"fmt"
	wuji "git.code.oa.com/open-wuji/go-sdk/wujiclient"
	"git.code.oa.com/video_app_short_video/hello_alice/model"
)

var (
	// 接入无极的client接口
	wujiClientProxy wuji.ClientInterface
)

//初始化wuji
func InitWujiProxy()(err error){
	wujiClientProxy, err = wuji.NewClient(
		wuji.WithAppID("hello_alice"),
		wuji.WithSchemaID("wuji_table_demo"),
		wuji.WithSchemaKey("b9b14f4585394e3094d603d8abe9887c"),
		wuji.WithRequestURL("http://nodeapi.webdev.com/x/api/wuji_public/object"),
		wuji.WithRequestTarget("polaris://64394561:131072"),  // L5寻址
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
		err = fmt.Errorf("Get error: %v",err)
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
