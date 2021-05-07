package dao

import (
	"context"
	"fmt"
	"git.code.oa.com/trpc-go/trpc-database/goredis"
	"git.code.oa.com/trpc-go/trpc-go/log"
)

// 对redis中的string类型进行操作
func AcessRedisString(ctx context.Context) (string,error) {
	var rsp string
	//初始化请求接口
	cli, err := goredis.NewClientProxy("trpc.gamecenter.test.redis",nil)
		//client.WithTarget("redis+polaris://:AzNIBb*PbIWQSJ,rwQ@sz4678.shortvideotest.redis.com")//请求服务地址格式：redis://<user>:<password>@<host>:<port>/<db_number>
	if err != nil {
		log.Errorf("InitRedisProxy fail err=[%v]\n", err)
		return "",err
	}
	strResult, err := cli.Set(ctx,"key1","hello",0).Result()
	if err != nil {
		log.Errorf("Set fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("Set result=[%v]\n", strResult)
	intResult, err := cli.Exists(ctx,"key1").Result()
	if err != nil {
		return "", err
	}
	rsp += fmt.Sprintf("Exists result=[%v]\n",intResult)
	intResult, err = cli.Append(ctx,"key1", " world").Result()
	if err != nil {
		log.Errorf("Append fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("Append result=[%v]\n",intResult)
	strResult, err = cli.Get(ctx,  "key1").Result()
	if err != nil {
		log.Errorf("Get fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("Get result=[%v]\n",strResult)
	strResult, err = cli.GetRange(ctx, "key1", -3, -1).Result()
	if err != nil {
		log.Errorf("GetRange fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("GetRange result=[%v]\n",strResult)
	return rsp, nil
}
// 对redis中的hash类型进行操作
func AcessRedisHash(ctx context.Context) (rsp string,err error) {
	//初始化请求接口
	cli, err := goredis.NewClientProxy("trpc.gamecenter.test.redis",nil)
	//client.WithTarget("redis+polaris://:AzNIBb*PbIWQSJ,rwQ@sz4678.shortvideotest.redis.com"请求服务地址格式：redis://<user>:<password>@<host>:<port>/<db_number>
	if err != nil {
		log.Errorf("InitRedisProxy fail err=[%v]\n", err)
		return "",err
	}

	intResult, err := cli.HSet(ctx, "myhash", "field1","foo").Result()
	if err != nil {
		log.Errorf("HSet fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("HSet result=[%v]\n",intResult)

	intResult, err = cli.HSet(ctx, "myhash", "field2", "bar").Result()
	if err != nil {
		log.Errorf("HSet fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("HSet result=[%v]\n",intResult)

	strResult, err := cli.HGet(ctx, "myhash", "field1").Result()
	if err != nil {
		log.Errorf("HGet fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("HGet result=[%v]\n",strResult)

	return rsp, nil
}

func AcessRedisZset(ctx context.Context) (rsp string,err error) {
	//初始化请求接口
	cli, err := goredis.NewClientProxy("trpc.gamecenter.test.redis",nil)
	//client.WithTarget("redis+polaris://:AzNIBb*PbIWQSJ,rwQ@sz4678.shortvideotest.redis.com"请求服务地址格式：redis://<user>:<password>@<host>:<port>/<db_number>
	if err != nil {
		log.Errorf("InitRedisProxy fail err=[%v]\n", err)
		return "",err
	}
	intResult, err := cli.ZAdd(ctx, "myzset", ).Result()
	if err != nil {
		log.Errorf("ZAdd fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("ZAdd result=[%v]\n",intResult)
	return rsp, nil
}