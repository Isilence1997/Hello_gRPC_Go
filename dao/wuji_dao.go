package dao

import (
	"database/sql"
	"fmt"
	wuji "git.code.oa.com/open-wuji/go-sdk/wujiclient"
	"git.code.oa.com/trpc-go/trpc-database/mysql"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/video_app_short_video/hello_alice/model"
)

var (
	// 接入无极的client接口
	filter wuji.FilterInterface
	wujiClientProxy wuji.ClientInterface
)

//初始化wuji
func initWujiProxy()(err error){
	wujiClientProxy, err := wuji.NewClient(
		wuji.WithAppID("hello_alice"),
		wuji.WithSchemaID("wuji_table_demo"),
		wuji.WithSchemaKey("b9b14f4585394e3094d603d8abe9887c"),
		wuji.WithRequestURL("http://nodeapi.webdev.com/x/api/wuji_public/object"),
		wuji.WithRequestTarget("polaris://64394561:131072"),  // L5寻址
		// wuji.WithRequestDirect(),   // 或者直接 dns 寻址
	)
	if err != nil {
		err = fmt.Errorf("initWujiClient error: %v",err)
		return err
	}
	return nil
}

func GetWujiContent(uid string)(string,error){
	err := initWujiProxy()
	if err!= nil {
		return "",err
	}
	var obj model.wujiData
	obj.ID=1
	wujiClientProxy.Get()
}
