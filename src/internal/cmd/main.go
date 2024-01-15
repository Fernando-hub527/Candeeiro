package main

import (
	"fmt"

	"github.com/Fernando-hub527/candieiro/internal/api"
	rabbitmq "github.com/Fernando-hub527/candieiro/internal/pkg/broker"
	"github.com/Fernando-hub527/candieiro/internal/pkg/mongodb"
	"github.com/labstack/echo/v4"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	database, err := mongodb.SetUpDatabase("mongodb://root:example@localhost:27018/")
	if err != nil {
		panic("Falha ao se conectar ao banco \n")
	}
	chRabbit := startBroker("amqp://iot:iot@localhost:5673/")
	startApi(chRabbit, database)
}

func startBroker(url string) *amqp091.Channel {
	ch, err := rabbitmq.OpenChannel(url)
	if err != nil {
		panic(err)
	}
	return ch
}

func startApi(chSensors *amqp091.Channel, database *mongo.Database) {
	server := echo.New()
	api.SetRouts(chSensors, api.SetWebsocket(server), server, database)
	if err := server.Start(":1323"); err != nil {
		server.Logger.Fatal()
	}
	fmt.Println("server running")
}
