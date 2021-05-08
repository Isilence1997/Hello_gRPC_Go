package dao

import (
	"context"
	"fmt"

	"git.code.oa.com/trpc-go/trpc-database/redis"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/log"
)
var (
	// redis 客户端代理
	redisClientProxy redis.Client
)
//初始化请求接口
func initClientProxy() error{
	redisClientProxy = redis.NewClientProxy(
		"trpc.redis.redis.redis",
		client.WithNamespace("Production"),
		client.WithTarget("redis+polaris://:AzNIBb*PbIWQSJ,rwQ@sz4678.shortvideotest.redis.com"),
		)
	do, err := redisClientProxy.Do(context.Background(), "PING")
	if err != nil {
		log.Errorf("InitRedisProxy error, err:%v", err)
		return err
	}
	log.Infof("connect to redis successfully: %+v", do)
	return nil
}
// 对redis中的string类型进行操作
func AcessRedisString(ctx context.Context) (rsp string,err error) {
	err = initClientProxy()
	if err!=nil{
		return "", err
	}
	strResult, err := redis.String(redisClientProxy.Do(ctx, "SET", "key1", "hello"))
	if err != nil {
		log.Errorf("Set fail err=[%v]\n ", err)
		return "", err
	}
	rsp += fmt.Sprintf("Set result=[%v]\n ", strResult)
	intResult, err := redis.Int(redisClientProxy.Do(ctx, "EXISTS", "key1"))
	if err != nil {
		return "", err
	}
	rsp += fmt.Sprintf("Exists result=[%v]\n ",intResult)
	intResult, err = redis.Int(redisClientProxy.Do(ctx, "APPEND", "key1", " world"))
	if err != nil {
		log.Errorf("Append fail err=[%v]\n ", err)
		return "", err
	}
	rsp += fmt.Sprintf("Append result=[%v]\n ",intResult)
	strResult, err = redis.String(redisClientProxy.Do(ctx, "GET", "key1"))
	if err != nil {
		log.Errorf("Get fail err=[%v]\n ", err)
		return "", err
	}
	rsp += fmt.Sprintf("Get result=[%v]\n ",strResult)
	strResult, err = redis.String(redisClientProxy.Do(ctx, "GETRANGE", "key1", -3, -1))
	if err != nil {
		log.Errorf("GetRange fail err=[%v]\n ", err)
		return "", err
	}
	rsp += fmt.Sprintf("GetRange result=[%v]\n ",strResult)
	return rsp, nil
}
// 对redis中的hash类型进行操作
func AcessRedisHash(ctx context.Context) (rsp string,err error) {

	intResult, err := redis.Int(redisClientProxy.Do(ctx, "HSET", "myhash", "field1", "foo"))
	if err != nil {
		log.Errorf("HSet fail err=[%v]\n ", err)
		return "", err
	}
	rsp += fmt.Sprintf("HSet result=[%v]\n ",intResult)

	intResult, err = redis.Int(redisClientProxy.Do(ctx, "HSET", "myhash", "field2", "bar"))
	if err != nil {
		log.Errorf("HSet fail err=[%v]\n ", err)
		return "", err
	}
	rsp += fmt.Sprintf("HSet result=[%v]\n ",intResult)

	strResult, err := redis.String(redisClientProxy.Do(ctx, "HGET", "myhash", "field1"))
	if err != nil {
		log.Errorf("HGet fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("HGet result=[%v]\n ",strResult)

	return rsp, nil
}

func AcessRedisZset(ctx context.Context) (rsp string,err error) {
	intResult, err := redis.Int(redisClientProxy.Do(ctx, "ZADD", "myzset", 1, "one"))
	if err != nil {
		log.Errorf("ZAdd fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("ZAdd result=[%v]\n ",intResult)
	intResult, err = redis.Int(redisClientProxy.Do(ctx, "ZADD", "myzset", 2, "two", 3, "three"))
	if err != nil {
		log.Errorf("ZAdd fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("ZAdd result=[%v]\n ",intResult)
	strResult,err := redis.Strings(redisClientProxy.Do(ctx, "ZRANGE", "myzset", 0, -1, "WITHSCORES"))
	if err != nil {
		log.Errorf("ZRange fail err=[%v]\n ", err)
		return "", err
	}
	rsp += fmt.Sprintf("ZRange result=[%v]\n ",strResult)
	intResult,err = redis.Int(redisClientProxy.Do(ctx, "ZREM", "myzset", "one"))
	if err != nil {
		log.Errorf("ZRem fail err=[%v]\n", err)
		return "", err
	}
	rsp += fmt.Sprintf("ZRem result=[%v]\n ",intResult)
	strResult,err = redis.Strings(redisClientProxy.Do(ctx, "ZRANGEBYSCORE", "myzset", "(1", 3))
	if err != nil {
		log.Errorf("ZRangeByScore fail err=[%v]\n ", err)
		return "", err
	}
	rsp += fmt.Sprintf("ZRangeByScore result=[%v]\n ",strResult)
	return rsp, nil
}