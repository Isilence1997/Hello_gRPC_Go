package dao

import (
	"context"
	"fmt"

	"git.code.oa.com/gRPC-go/gRPC-database/redis"
	"git.code.oa.com/gRPC-go/gRPC-go/client"
	"git.code.oa.com/gRPC-go/gRPC-go/log"

	"git.code.oa.com/video_app_short_video/hello_alice/config"
)
var (
	// redis 客户端代理
	redisClientProxy redis.Client
)

//初始化请求接口
func InitRedisProxy() error{
	serviceConfig := config.GetConfig()
	redisConfig := serviceConfig.Redis
	redisClientProxy = redis.NewClientProxy(
		redisConfig.ServiceName,
		client.WithNamespace(redisConfig.Namespace),
		//"redis+polaris://:pwd@zkname"
		client.WithTarget(fmt.Sprintf("redis+polaris://:%s@%s",redisConfig.Pwd,redisConfig.ObjName)),
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

// 对redis中的sorted set类型进行操作
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