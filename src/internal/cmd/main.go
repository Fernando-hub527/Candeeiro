package main

import (
	"github.com/Fernando-hub527/candieiro/internal/api"
	"github.com/Fernando-hub527/candieiro/internal/pkg/mongodb"
	"github.com/Fernando-hub527/candieiro/internal/pkg/rabbitmq"
	"github.com/labstack/echo/v4"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	mongodb.SetUpDatabase("mongodb://root:example@localhost:27018/")
	chRabbit := startBroker("amqp://iot:iot@localhost:5673/")
	startApi(chRabbit)
}

func startBroker(url string) *amqp091.Channel {
	ch, err := rabbitmq.OpenChannel(url)
	if err != nil {
		panic(err)
	}
	return ch
}

func startApi(chSensors *amqp091.Channel) {
	server := echo.New()
	api.SetRouts(chSensors, api.SetWebsocket(server), server)
	server.Logger.Fatal(server.Start(":1323"))
}
