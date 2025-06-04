# 1.golang kafka

Golang与Kafka结合常用于构建高并发、可扩展的消息系统，以下是核心知识点和实践方向：  


### **一、Golang操作Kafka的常用库**
- **Confluent Kafka Go Client**  
  官方推荐库，功能完整，支持原生Kafka协议，包含生产者、消费者、管理API（如创建Topic、查询Offset等）。  
  ```go
  import "github.com/confluentinc/confluent-kafka-go/kafka"
  ```
- **Sarama**  
  社区流行的开源库，纯Go实现，轻量级且文档完善，适合快速开发。  
  ```go
  import "github.com/Shopify/sarama"
  ```


### **二、核心功能实现示例（以Confluent库为例）**
#### **1. 生产者（Producer）**  
发送消息到Kafka Topic：  
```go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// 创建生产者配置
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Kafka服务器地址
	})
	if err != nil {
		log.Fatalf("创建生产者失败: %v", err)
	}
	defer producer.Close()

	// 异步消息发送回调（处理发送结果）
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("消息发送失败: %v", ev.TopicPartition)
				} else {
					log.Printf("消息发送成功: topic=%s, partition=%d, offset=%v",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	// 发送消息
	topic := "test-topic"
	message := "Hello Kafka from Golang!"
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
	if err != nil {
		log.Fatalf("发送消息出错: %v", err)
	}

	// 等待信号退出
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan
}
```

#### **2. 消费者（Consumer）**  
订阅Topic并消费消息：  
```go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// 创建消费者配置
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "my-group",       // 消费者组ID
		"auto.offset.reset": "earliest",       // 消费偏移量重置策略
	})
	if err != nil {
		log.Fatalf("创建消费者失败: %v", err)
	}
	defer consumer.Close()

	// 订阅Topic
	topic := "test-topic"
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("订阅Topic失败: %v", err)
	}

	// 消费消息循环
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			log.Printf("收到信号: %v，准备退出", sig)
			run = false
		default:
			// 轮询获取消息（超时100ms）
			msg, err := consumer.ReadMessage(-1)
			if err == nil {
				log.Printf("收到消息: topic=%s, partition=%d, offset=%v, value=%s",
					*msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset, string(msg.Value))
				// 手动提交偏移量（可选，默认自动提交）
				consumer.CommitMessage(msg)
			} else if err.(kafka.Error).Code() != kafka.ErrNoMessage {
				log.Printf("读取消息出错: %v", err)
				run = false
			}
		}
	}

	// 退出前取消订阅
	consumer.Unsubscribe()
}
```


### **三、进阶特性与最佳实践**
- **分区与负载均衡**  
  Kafka通过分区（Partition）实现数据分片和并行消费，Golang消费者组会自动分配分区到不同实例，确保负载均衡。  
- **消息顺序性**  
  若需保证消息顺序，可将同类型消息发送到同一分区（通过Key哈希实现），消费者按分区顺序消费。  
- **幂等性与 Exactly-Once 语义**  
  - 生产者开启幂等性（`enable.idempotence = true`）可避免重复消息。  
  - 结合Kafka事务（Transactions）实现跨分区的Exactly-Once语义，确保消息仅被处理一次。  
- **批量发送与压缩**  
  生产者可配置批量发送（`batch.size`、`linger.ms`）和消息压缩（`compression.type = gzip/snappy/zstd`），减少网络开销。  
- **异常处理与重试**  
  处理网络异常、Broker故障时，生产者需实现重试逻辑，消费者需处理Offset提交失败等场景。  


### **四、生产环境部署建议**
- **连接池管理**  
  避免频繁创建生产者和消费者，使用连接池复用资源。  
- **监控与告警**  
  通过Prometheus结合Kafka Exporter监控集群指标（如消息吞吐量、延迟、消费者滞后量等），Golang应用可自定义Metrics上报。  
- **日志与追踪**  
  集成Zap、Logrus等日志库，结合Jaeger等链路追踪工具定位消息处理瓶颈。  


通过Golang与Kafka的结合，可构建高可用的微服务消息系统，适用于日志收集、异步任务、微服务通信等场景。根据业务需求选择合适的库和配置，能有效提升系统的扩展性和稳定性。