package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"git.code.oa.com/trpc-go/trpc-database/kafka"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/log"
	"git.code.oa.com/video_app_short_video/hello_alice/model"
	"time"
)

const (
	fromUserId = 536164684
	toUserId = 2454008777
)

// 生产消息
func ProcedueKafka(ctx context.Context) (rsp string,err error) {
	var msg model.FollowWriteKafkaMsg
	msg.IsFake=0
	msg.FollowAction=1
	msg.FromVuid= fromUserId
	msg.ToVuid =toUserId

	kafkaProxy := kafka.NewClientProxy(
		"trpc.video_app_short_video.hello_alice.producer",
		//clientid=生产者ID
		//compression: none 等于sarama.CompressionNone(默认值);gzip 等于sarama.CompressionGZIP;snappy 等于sarama.CompressionSnappy;lz4 等于sarama.CompressionLZ4;zstd 等于sarama.CompressionZSTD
		client.WithTarget("kafka://cdmqszentry01.data.mig:10005,cdmqszentry02.data.mig:10069,cdmqszentry05.data.mig:10033,cdmqszentry06.data.mig:10021?topic=U_TOPIC_FOLLOW&clientid=p_FOLLOW"),
	)
	for i:=0;i<3;i++{ //生产3次
		msg.Timestamp = time.Now().Unix()
		value, _ := json.Marshal(msg)//序列化msg
		err := kafkaProxy.Produce(ctx, []byte("hello"), value)//key和value都是字节流
		if err != nil {
			log.Errorf("kafka Produce error, err:%+v", err)
			return "",err
		}
		time.Sleep(time.Millisecond * 50)
		rsp += fmt.Sprintf("toVuid:%d, fromVuid:%d, timestamp:%+v ",
			msg.ToVuid, msg.FromVuid, msg.Timestamp)
	}
	return rsp,nil
}

//消费消息
/*
 service:                                                                                #业务服务提供的service，可以有多个
    - name: trpc.${app}.${server}.consumer                                                      #service的路由名称 如果使用的是123平台，需要使用trpc.${app}.${server}.consumer
      address:cdmqszentry01.data.mig:10005,cdmqszentry02.data.mig:10069,cdmqszentry05.data.mig:10033,cdmqszentry06.data.mig:10021?topic=U_TOPIC_FOLLOW&clientid=p_FOLLOW         #kafka consumer broker address，version如果不设置则为1.1.1.0，部分ckafka需要指定0.10.2.0
      protocol: kafka                                                                     #应用层协议
      timeout: 1000
*/
func ConsumeKafkaMsgHandler(_ context.Context, key, value []byte, topic string,
	partition int32, offset int64) error {
	kafakRsp := fmt.Sprintf("ConsumeKafkaMsgHandler, [key]%v, [value]%v, [topic]%v, [partition]%v, [offset]%v", string(key),
		string(value), topic, partition, offset)
	// 创建对象，反序列化
	var followWriteKafkaMsg model.FollowWriteKafkaMsg
	err := json.Unmarshal(value, &followWriteKafkaMsg)
	if err != nil {
		err = fmt.Errorf("ConsumeKafkaMsgHandler json.Unmarshal error, value:%s, topic:%s,"+
			" partition:%d, offset:%d, err:%v", string(value), topic, partition, offset, err)
		return err
	}
	log.Infof("%s %v",kafakRsp, followWriteKafkaMsg)
	return nil
}