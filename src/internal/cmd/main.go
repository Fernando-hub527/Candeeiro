package main

import (
	"fmt"

	"github.com/Fernando-hub527/candieiro/internal/api"
	"github.com/Fernando-hub527/candieiro/internal/pkg/rabbitmq"
	"github.com/labstack/echo/v4"
	"github.com/rabbitmq/amqp091-go"
)

func a() {

}

func main() {
	validators := map[string]func(){
		"time": a,
	}
	b := validators["times"]
	fmt.Println(b == nil)
	// chanelSensors := make(chan amqp091.Delivery)

	// mongodb.SetUpDatabase("mongodb://root:example@localhost:27018/")
	// startBroker(chanelSensors, "amqp://iot:iot@localhost:5673/")
	// startApi(chanelSensors)
}

func startBroker(chSensors chan amqp091.Delivery, url string) {
	ch, err := rabbitmq.OpenChannel(url)
	if err != nil {
		panic(err)
	}
	go rabbitmq.Consumer(chSensors, ch, "sensors")
}

func startApi(chSensors chan amqp091.Delivery) {
	server := echo.New()
	api.SetRouts(chSensors, api.SetWebsocket(server), server)
	server.Logger.Fatal(server.Start(":1323"))
}
