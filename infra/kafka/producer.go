package kafka

import (
	"os"

	ckafka "github.com/conffluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.configMap{
		"bootstrapservers": os.Getenv(key: "KafkaBootsrapServers")
	}
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Erro())
	}
	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PatitionAny},
		Value: 			[]byte(msg),
	}
	err := producer.Produce(message, deliveryChan: nil)
	if err != nil {
		return err
	}
	return nil
}