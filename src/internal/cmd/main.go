package main

import (
	"fmt"

	"github.com/Fernando-hub527/candieiro/internal/api"
	"github.com/Fernando-hub527/candieiro/internal/pkg/broker"
	"github.com/Fernando-hub527/candieiro/internal/pkg/mongodb"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	database, err := mongodb.SetUpDatabase("mongodb://root:example@localhost:27018/")
	if err != nil {
		panic("Falha ao se conectar ao banco \n")
	}
	broker, errBroker := broker.NewBroker("amqp://iot:iot@localhost:5673/")
	if errBroker != nil {
		panic("Falha ao conectar no broker")
	}
	startApi(broker, database)
}

func startApi(broker broker.IBroker, database *mongo.Database) {
	server := echo.New()
	api.SetRouts(broker, api.SetWebsocket(server), server, database)
	if err := server.Start(":1323"); err != nil {
		server.Logger.Fatal()
	}
	fmt.Println("server running")
}
