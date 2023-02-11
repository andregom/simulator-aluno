package main

import (
	"fmt"
	"log"

	gototenv "github.com/joho/godotenv"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	kafka "simulator/infra/kafka"
	kafka2 "simulator/application/kafka"
)

func init() {
	err := gototenv.Load()
	if err != nil {
		log.Fatal("error laoding .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
