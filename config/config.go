package config

import (
	"context"

	"git.code.oa.com/gRPC-go/gRPC-go/config"
	"git.code.oa.com/gRPC-go/gRPC-go/log"
	"gopkg.in/yaml.v3" // 解析yaml
)

type ServiceConfig struct {
	Mysql struct {
		ServiceName string `json:"service_name" yaml:"service_name"` // 服务名，监控上报使用
		Domain      string `json:"domain" yaml:"domain"`
		Port        uint16 `json:"port" yaml:"port"`
		DB          string `json:"db" yaml:"db"`
		User        string `json:"user" yaml:"user"`
		Password    string `json:"pwd" yaml:"pwd"`
	} `json:"mysql" yaml:"mysql"`
	Redis struct { // redis 配置
		ServiceName string `json:"service_name" yaml:"service_name"`
		ObjName     string `json:"obj_name" yaml:"obj_name"`
		Pwd         string `json:"pwd" yaml:"pwd"`
		Namespace   string `json:"namespace" yaml:"namespace"`
	} `json:"redis" yaml:"redis"`

	Union struct { // union 配置
		AppId       uint32 `json:"appid" yaml:"appid"`
		AppKey      string `json:"appkey" yaml:"appkey"`
		ServiceName string `json:"service_name" yaml:"service_name"`
		Namespace   string `json:"namespace" yaml:"namespace"`
		Target      string `json:"target" yaml:"target"`
	} `json:"union" yaml:"union"`

	Wuji struct { // wuji 配置
		AppID      string `yaml:"app_id"`
		SchemaID   string `yaml:"schema_id"`
		SchemaKey  string `yaml:"schema_key"`
		RequestURL string `yaml:"request_url"`
		Target     string `yaml:"target"`
	} `json:"wuji" yaml:"wuji"`

	Kafka struct {
		ServiceName string `json:"service_name" yaml:"service_name"`
		Topic       string `yaml:"topic"`
		ClientId    string `yaml:"client_id"`
		Target      string `yaml:"target"`
	} `json:"kafka" yaml:"kafka"`

	Atta struct { // atta 配置
		AttaId string `json:"atta_id" yaml:"atta_id"`
		Token  string `json:"atta_token" yaml:"atta_token"`
	} `json:"atta" yaml:"atta"`
}

var (
	serviceConfig ServiceConfig
)

// 初始化服务配置
func InitServiceConfig() {
	// 加载配置文件，使用tconf从123平台上加载
	confName := "greeter.yaml"
	serviceConfig = ServiceConfig{}
	err := config.GetYAML(confName, &serviceConfig) //yaml转化为struct
	if err != nil {
		log.Errorf("get yaml conf error, err:%v", err)
		panic(err)
	} else {
		log.Infof("yaml conf, conf:%+v", serviceConfig)
	}
	// 启动协程，监听配置文件变化，重新加载配置文件
	c, _ := config.Get("tconf").Watch(context.TODO(), confName) // Watch 监听配置项key的变更事件
	//加载本地配置文件
	//c, _ := config.Load("greeter.yaml")
	go func() { //goroutine
		select {
		case r := <-c: //通道c收到一个事件 r
			func() {
				err := yaml.Unmarshal([]byte(r.Value()), &serviceConfig) //解码到serviceConfig
				if err != nil {
					log.Errorf("watch conf yaml unmarshal error, value:%s, err:%v", r.Value(), err)
				} else {
					log.Infof("reload conf success, event:%s, value:%s", r.Event(), r.Value())
				}
			}()
		}
	}()
}

// 获取配置
func GetConfig() ServiceConfig {
	return serviceConfig
}
