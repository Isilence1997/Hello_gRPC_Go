package common

import (
	"errors"
	"git.code.oa.com/atta/attaapi_go"
	"git.code.oa.com/video_app_short_video/hello_alice/config"
)

// atta上报api
var apiObj attaapi_go.AttaApi

func InitAtta() error{
	// 初始化很消耗资源，切忌频繁调用！建议单进程复用一个 atta api 对象
	result := apiObj.InitUDP()
	if result != attaapi_go.M_ATTA_REPORT_CODE_SUCCESS {
		return  errors.New("init atta failed")
	}
	return nil
}

func AttaSendFields(str1 string, str2 string) int {
	// 字段数组上报的字段顺序需要和http://atta.pcg.com中配置的AttaId字段顺序一致
	fieldValues := make([]string, 2)
	fieldValues[0] = str1
	fieldValues[1] = str2
	serviceConfig := config.GetConfig()
	attaConfig := serviceConfig.Atta
	// autoEscape（true 自动转义，false 不自动转义）
	return apiObj.SendFields(attaConfig.AttaId,attaConfig.Token,fieldValues,false)
}