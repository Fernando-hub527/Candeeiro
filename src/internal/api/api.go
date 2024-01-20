package api

import (
	"github.com/Fernando-hub527/candieiro/internal/handles"
	"github.com/Fernando-hub527/candieiro/internal/pkg/broker"
	"github.com/Fernando-hub527/candieiro/internal/pkg/websocket"
	consumutionrepository "github.com/Fernando-hub527/candieiro/internal/repository/consumutionRepository"
	pointrepository "github.com/Fernando-hub527/candieiro/internal/repository/pointRepository"
	userrepository "github.com/Fernando-hub527/candieiro/internal/repository/userRepository"
	"github.com/Fernando-hub527/candieiro/internal/useCase/electricity"
	"github.com/Fernando-hub527/candieiro/internal/useCase/user"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetRouts(broker broker.IBroker, hub *websocket.Hub, server *echo.Echo, database *mongo.Database) {
	userUseCase := user.NewUserCase(userrepository.New(database))
	electricityUseCase := electricity.NewElectricityUseCase(pointrepository.New(database), consumutionrepository.New(database))

	handlesElectricity := handles.NewElectricityHandles(broker, hub, userUseCase, electricityUseCase)
	handlesUser := handles.NewUserHandles(userUseCase)

	server.POST("/api/v1/candieiro/login", handlesUser.Login)
	server.GET("/api/v1/candieiro/points", handlesElectricity.ListPoints)
	server.GET("/api/v1/candieiro/point/consumption", handlesElectricity.ListConsumptionByInterval)
}

func SetWebsocket(e *echo.Echo) *websocket.Hub {
	hub := websocket.NewHub()
	go hub.Run()
	e.GET("/ws", func(context echo.Context) error {
		websocket.DefaultHandlesWs(hub, context)
		return nil
	})
	return hub
}
