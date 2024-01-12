package api

import (
	"github.com/Fernando-hub527/candieiro/internal/handles"
	"github.com/Fernando-hub527/candieiro/internal/pkg/websocket"
	consumutionrepository "github.com/Fernando-hub527/candieiro/internal/repository/consumutionRepository"
	pointrepository "github.com/Fernando-hub527/candieiro/internal/repository/pointRepository"
	userrepository "github.com/Fernando-hub527/candieiro/internal/repository/userRepository"
	"github.com/Fernando-hub527/candieiro/internal/useCase/electricity"
	"github.com/Fernando-hub527/candieiro/internal/useCase/user"
	"github.com/labstack/echo/v4"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetRouts(chBroker *amqp091.Channel, hub *websocket.Hub, server *echo.Echo, database *mongo.Database) {
	userUseCase := user.NewUserCase(userrepository.New(database))
	electricityUseCase := electricity.NewElectricityUseCase(pointrepository.New(database), consumutionrepository.New(database))
	handlesElectricity := handles.NewElectricityHandles(chBroker, hub, userUseCase, electricityUseCase)

	server.GET("/api/v1/candieiro/points", handlesElectricity.ListPoints)
	server.GET("/api/v1/candieiro/point/consumption", handlesElectricity.ListConsumptionByInterval)
}

func SetWebsocket(e *echo.Echo) *websocket.Hub {
	hub := websocket.NewHub()
	go hub.Run()
	e.GET("/ws", func(context echo.Context) error {
		return websocket.DefaultHandlesWs(hub, context)
	})
	return hub
}
