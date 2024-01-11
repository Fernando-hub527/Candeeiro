package handles

import (
	"fmt"

	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/Fernando-hub527/candieiro/internal/pkg/utils"
	"github.com/Fernando-hub527/candieiro/internal/pkg/websocket"
	"github.com/Fernando-hub527/candieiro/internal/useCase/electricity"
	"github.com/Fernando-hub527/candieiro/internal/useCase/user"
	"github.com/labstack/echo/v4"
	"github.com/rabbitmq/amqp091-go"
)

type ElectricityHandles struct {
	hub                *websocket.Hub
	userUseCase        user.IUserUseCase
	electricityUseCase electricity.IElectricityUseCase
}

func NewElectricityHandles(chBroker *amqp091.Channel, hub *websocket.Hub, userUseCase user.IUserUseCase, electricityUseCase electricity.IElectricityUseCase) *ElectricityHandles {
	elc := &ElectricityHandles{
		hub:                hub,
		userUseCase:        userUseCase,
		electricityUseCase: electricityUseCase,
	}
	go elc.recordConsumption(chBroker, "consumptio/electicity/record")
	go elc.updateConsumption(chBroker, "consumptio/electicity/update")
	return elc
}

func (elc *ElectricityHandles) sendError(context echo.Context, err errors.RequestError) error {
	return context.String(int(err.Status), err.ToString())
}

func (elc *ElectricityHandles) ListPoints(context echo.Context) error {
	planId, err := utils.ValidObjectId(context.QueryParam("plant"), elc.sendError, context)

	if err != nil {
		return nil
	}

	if err := elc.userUseCase.ValidAccess(context.Param("userName"), *planId); err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}

	points, err := elc.electricityUseCase.ListPointsByPlant(*planId)
	if err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}
	context.JSON(200, points)
	return nil
}

func (elc *ElectricityHandles) ListConsumptionByInterval(context echo.Context) error {
	startTime, errS := utils.ValidTime(context.QueryParam("startMoment"), elc.sendError, context)
	endTime, errE := utils.ValidTime(context.QueryParam("startMoment"), elc.sendError, context)
	pointId, errP := utils.ValidObjectId(context.QueryParam("point"), elc.sendError, context)

	if errS != nil || errE != nil || errP != nil {
		return nil
	}

	point, err := elc.electricityUseCase.FindPointById(*pointId)
	if err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}

	if err := elc.userUseCase.ValidAccess(context.Param("userName"), point.PlanId); err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}

	points, err := elc.electricityUseCase.ListConsumptionByIntervalAndPoint(*pointId, *startTime, *endTime)
	if err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}
	context.JSON(200, points)
	return nil
}

func (elc *ElectricityHandles) recordConsumption(chBroker *amqp091.Channel, queue string) {
	for {
		fmt.Println("Iniciando consumo da fila ", queue)
	}
}

func (elc *ElectricityHandles) updateConsumption(chBroker *amqp091.Channel, queue string) {
	for {
		fmt.Println("Iniciando consumo da fila ", queue)
	}
}
