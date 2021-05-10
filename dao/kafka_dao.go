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
		client.WithTarget("kafka://9.134.106.151:9095?topic=FollewWritr&clientid=xxx&compression=xxx"),
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