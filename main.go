package main

import (
	"fmt"
	// "application/route"
	"log"
	kafka "infra/kafka"
	kafka2 "application/kafka"

	gototenv "github.com/joho/godotenv"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
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

	// localRoute := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// localRoute.LoadPositions()
	// stringjson, _ := localRoute.ExportJsonPasitions()

	// fmt.Println(stringjson[0])
}
