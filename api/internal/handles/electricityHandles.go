package handles

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fernando-hub527/candieiro/internal/dtos"
	"github.com/Fernando-hub527/candieiro/internal/pkg/broker"
	rabbitmq "github.com/Fernando-hub527/candieiro/internal/pkg/broker"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/Fernando-hub527/candieiro/internal/pkg/utils"
	"github.com/Fernando-hub527/candieiro/internal/pkg/websocket"
	"github.com/Fernando-hub527/candieiro/internal/useCase/electricity"
	"github.com/Fernando-hub527/candieiro/internal/useCase/user"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ElectricityHandles struct {
	hub                *websocket.Hub
	userUseCase        user.IUserUseCase
	electricityUseCase electricity.IElectricityUseCase
}

func NewElectricityHandles(broker rabbitmq.IBroker, hub *websocket.Hub, userUseCase user.IUserUseCase, electricityUseCase electricity.IElectricityUseCase) *ElectricityHandles {
	elc := &ElectricityHandles{
		hub:                hub,
		userUseCase:        userUseCase,
		electricityUseCase: electricityUseCase,
	}
	// go elc.recordConsumption(broker, "electicity_record")
	go elc.updateConsumption(broker, "electicity_update")
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

	if err := elc.userUseCase.ValidAccess(context.Param("userName"), *planId, context.Request().Context()); err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}

	points, err := elc.electricityUseCase.ListPointsByPlant(*planId, context.Request().Context())
	if err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}
	context.JSON(200, points)
	return nil
}

func (elc *ElectricityHandles) ListConsumptionByInterval(context echo.Context) error {
	startTime, errS := utils.ValidTime(context.QueryParam("startTime"), elc.sendError, context)
	endTime, errE := utils.ValidTime(context.QueryParam("endTime"), elc.sendError, context)
	pointId, errP := utils.ValidObjectId(context.QueryParam("point"), elc.sendError, context)

	if errS != nil || errE != nil || errP != nil {
		return nil
	}

	point, err := elc.electricityUseCase.FindPointById(*pointId, context.Request().Context())
	if err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}

	if err := elc.userUseCase.ValidAccess(context.Param("userName"), point.PlanId, context.Request().Context()); err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}

	points, err := elc.electricityUseCase.ListConsumptionByIntervalAndPoint(*pointId, *startTime, *endTime, context.Request().Context())
	if err != nil {
		context.String(int(err.Status), err.ToString())
		return nil
	}
	context.JSON(200, points)
	return nil
}

func (elc *ElectricityHandles) recordConsumption(broker broker.IBroker, queue string) {
	chanMessager := broker.Consumer(queue)

	for msg := range chanMessager {
		var consuDTO dtos.ConsumptionRecordRequestDTO
		if err := json.Unmarshal(msg.GetMessager(), &consuDTO); err != nil {
			fmt.Println("Falha ao dessrializar json")
			msg.Reject()
		} else {
			elc.electricityUseCase.CreateConsumutionRecord(consuDTO)
			msg.Accept()
		}
	}
}

func (elc *ElectricityHandles) updateConsumption(broker broker.IBroker, queue string) {
	chanMessager := broker.Consumer(queue)
	fmt.Println(primitive.NewObjectID())
	for msg := range chanMessager {
		var consuDTO dtos.ConsumptionFluctuationRequestDTO
		if err := json.Unmarshal(msg.GetMessager(), &consuDTO); err != nil {
			fmt.Println("Falha ao dessrializar json")
			msg.Accept()
		} else {
			if _, err := elc.electricityUseCase.FindPointById(consuDTO.PointId, context.TODO()); err != nil {
				fmt.Println("Dados decebidos de ponto de consumo não registrado ", consuDTO.PointId)
			} else {
				message, _ := consuDTO.ParseToConsumptionFluctuation().ToJson()
				elc.hub.Messages <- map[string][]byte{consuDTO.PointId.String(): message}
			}
			msg.Accept()
		}
	}
}
