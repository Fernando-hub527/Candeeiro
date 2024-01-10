package api

import (
	"github.com/Fernando-hub527/candieiro/internal/pkg/websocket"
	"github.com/labstack/echo/v4"
	"github.com/rabbitmq/amqp091-go"
)

func SetRouts(chSensors chan amqp091.Delivery, hub *websocket.Hub, server *echo.Echo) {
	// handlesElectricity := handles.NewElectricityHandles(chSensors, hub)

	// server.GET("api/v1/candieiro/points", handlesElectricity.ListPoints)
	// server.GET("api/v1/candieiro/point/consumption", handlesElectricity.ListConsumptionByInterval)
	// server.GET("api/v1/candieiro/point/shutdowns", handlesElectricity.ListShutdownSchedule)
	// server.GET("api/v1/candieiro/point/alert", handlesElectricity.FindSettingsByDevice)
	// server.POST("api/v1/candieiro/point/shutdown", handlesElectricity.AddShutdown)
	// server.DELETE("api/v1/candieiro/point/shutdown", handlesElectricity.RemoveShutdown)
	// server.PUT("api/v1/candieiro/point/alert", handlesElectricity.UpdateSettings)

}

func SetWebsocket(e *echo.Echo) *websocket.Hub {
	hub := websocket.NewHub()
	go hub.Run()
	e.GET("/ws", func(context echo.Context) error {
		return websocket.DefaultHandlesWs(hub, context)
	})
	return hub
}
