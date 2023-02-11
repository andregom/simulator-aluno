package main

import (
	// "fmt"
	// "application/route"
	"infra/kafka"
	"log"

	gototenv "github.com/joho/godotenv"
)

func init() {
	err := gototenv.Load()
	if err != nil {
		log.Fatal("error laoding .env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "readtest", producer)

	for {
		_ = 1
	}

	// localRoute := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// localRoute.LoadPositions()
	// stringjson, _ := localRoute.ExportJsonPasitions()

	// fmt.Println(stringjson[0])
}
