package main

import (
	"fmt"
	"application/route"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := gototenv.Load()
	if err != nil {
		log.Fatal(v...: "error laoding .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(stinrg(msg.Value))
	}

	// localRoute := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// localRoute.LoadPositions()
	// stringjson, _ := localRoute.ExportJsonPasitions()

	// fmt.Println(stringjson[0])
}
