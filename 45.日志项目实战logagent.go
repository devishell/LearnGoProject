package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"math/rand"
	"time"
)

func main() {
	//用到的框架工具

	//消息队列的通信模型
	//点对点queue 发布订阅topic

	//1.kafka   --LinkedIn用scala编写2011开源捐给apache
	//分布式数据流平台 高吞吐 低延迟 高容错
	//Producer ->
	//Kafka Cluster集群(Broker服务器节点(Topic消息主题/partition主题分区:优先连[leader一个]若leader挂了则连[follower有多个备份]))
	//leader和follower(备胎)在不同的机器,ACK消息链,消费组是一个整体,
	//->Consumer
	//kafka作用：1.MQ 还有传统的ActiveMQ RabbitMQ等 2.最初用来进行追踪网络活动 3.Metrics传输监控数据 4.日志聚合，收集多个服务器的日志
	//下载 http://kafka.apache.org/downloads
	//解压查看server.properties 设置log位置默认log.dirs=/tmp/kafka-logs
	//查看zookeeper.properties 设置默认dataDir=/tmp/zookeeper
	//先启动zookeeper-server-start.sh ../config/zookeeper.properties
	//启动kafka-server-start.sh ../config/server.properties

	//2.zookeeper  -- 雅虎仿照谷歌开发，致力于开发和维护开源服务器，该服务器可实现高度可靠的分布式协调。
	//分布式服务 跨进程跨程序公共资源锁 kafka默认用zookeeper做集群管理
	//下载 https://zookeeper.apache.org/releases.html

	//go使用 go get github.com/Shopify/sarama win只能使用1.19因为后续版本需要gcc
	config := sarama.NewConfig()
	//生产者配置
	config.Producer.RequiredAcks = sarama.WaitForAll          // ACK
	config.Producer.Partitioner = sarama.NewRandomPartitioner //分区
	config.Producer.Return.Successes = true                   //确认
	//连接kafka
	syncProducer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed ,err:", err)
		return
	}
	defer syncProducer.Close()
	//封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "shopping"
	msg.Value = sarama.StringEncoder("2020.10.19 开业"+string(rand.Uint64()))
	//发送消息
	partition, offset, err := syncProducer.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg err:", err)
		return
	}
	fmt.Println("partition:",partition,"offset", offset)
	//执行后终端cd /tmp/kafka-logs/shopping-0/ 后 cat 00000000000000000000.log 里面信息 IG19?u??u????????????."2020.10.19 开业I?hu?u???????????."2020.10.19 开业%
	//可以启动kafka-concole-consumer.sh消费该消息终端输入 kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic shopping --from-beginning
	//会打印：2020.10.19 开业 ,运行几次打印几次

	//tailIf从文件一直读取最新的内容常用来读取日志文件打印 https://github.com/hpcloud/tail
	//模拟命令 tail -f /tmp/kafka-logs/shopping-0/00000000000000000000.log
	t, err := tail.TailFile("debug.log", tail.Config{
		ReOpen: true, //超过一定大小日志文件变换后会跟随打开文件
		Follow: true, //一直读最新的行
		Location: &tail.SeekInfo{Offset: 0,Whence: 2}, //Whence: 2 从文件末尾读
		MustExist: false,
		Poll: false,//轮询方式
	})
	if err != nil {
		fmt.Println("tailIf err:",err)
		return
	}
	//手动修改debug.log的内容按快捷键保存会更快刷新到文件中,查看打印
	for line := range t.Lines {
		fmt.Println("line:",line.Text)
		time.Sleep(time.Second)
	}
}
