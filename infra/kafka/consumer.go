import ckafka "github.com/confluentinc/qconfluent-kafka-go/qkafka"

import (
	"fmt"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan * ckafka.Message) * KafkaConsumer {
	return {
		MsgChan: msgChan,
	}
}

func(k *KafkaConsumer) Consume() {
	configMap := &ckafka.configMap{
		"bootstrap.servers": os.Getenv(key: "KafkaBootstrapServers"),
		"group.id": os.Getenv(key: "KafkaConsumerGroupId"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kfafka message:" + err.Error())
	}
	topics := []string{os.Getenv(key: "KafkaREadTopic")}
	c.SubscribeTopics(topics, rebalanceCB:nil)
	fmt.Println(a...: "Kafka consumer has benn started")
	for {
		msg, err := c.ReadMessage(timeout: -1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}