package handles

import (
	"github.com/Fernando-hub527/candieiro/internal/pkg/utils"
	"github.com/Fernando-hub527/candieiro/internal/pkg/websocket"
	"github.com/Fernando-hub527/candieiro/internal/useCase/electricity"
	"github.com/Fernando-hub527/candieiro/internal/useCase/user"
	"github.com/labstack/echo/v4"
	"github.com/rabbitmq/amqp091-go"
)

type ElectricityHandles struct {
	chSensors          chan amqp091.Delivery
	hub                *websocket.Hub
	userUseCase        user.IUserUseCase
	electricityUseCase electricity.IElectricityUseCase
}

func NewElectricityHandles(chSensors chan amqp091.Delivery, hub *websocket.Hub, userUseCase user.IUserUseCase, electricityUseCase electricity.IElectricityUseCase) *ElectricityHandles {
	return &ElectricityHandles{
		chSensors:          chSensors,
		hub:                hub,
		userUseCase:        userUseCase,
		electricityUseCase: electricityUseCase,
	}
}

func (elc *ElectricityHandles) ListPoints(context echo.Context) error {
	planId, err := utils.IsValidObjectId(context.QueryParam("plant"))
	if err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}

	if err := elc.userUseCase.ValidAccess(context.Param("userName"), planId); err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}

	points, err := elc.electricityUseCase.ListPointsByPlant(planId)
	if err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}
	context.JSON(200, points)
	return nil
}

func (elc *ElectricityHandles) ListConsumptionByInterval(context echo.Context) error {
	pointId, err := utils.IsValidObjectId(context.QueryParam("point"))

	if err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}

	point, err := elc.electricityUseCase.FindPointById(pointId)
	if err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}

	if err := elc.userUseCase.ValidAccess(context.Param("userName"), point.PlanId); err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}

	// points, err := elc.electricityUseCase.ListConsumptionByIntervalAndPoint()(plantId)
	// if err != nil {
	// 	context.String(int(err.Status), err.ToString())
	// 	return nil
	// }
	// context.JSON(200, points)
	return nil
}

func (elc *ElectricityHandles) ListShutdownSchedule(context echo.Context) error {
	context.String(200, "ok")
	return nil
}

func (elc *ElectricityHandles) FindSettingsByDevice(context echo.Context) error {
	context.String(200, "ok")
	return nil
}

func (elc *ElectricityHandles) AddShutdown(context echo.Context) error {
	context.String(200, "ok")
	return nil
}

func (elc *ElectricityHandles) RemoveShutdown(context echo.Context) error {
	context.String(200, "ok")
	return nil
}

func (elc *ElectricityHandles) UpdateSettings(context echo.Context) error {
	context.String(200, "ok")
	return nil
}

func (elc *ElectricityHandles) recordConsumption() {

}
