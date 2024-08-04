package services

import (
	"context"
	"time"

	"github.com/Fernando-hub527/candieiro/hexa/internal/core/consumer/domain"
	"github.com/Fernando-hub527/candieiro/hexa/internal/core/consumer/ports"
)

type consumerService struct {
	consumerRepository *ports.ConsumerRepository
}

func NewConsumerService(consumerRepository *ports.ConsumerRepository) *consumerService {
	return &consumerService{
		consumerRepository: consumerRepository,
	}
}

// Function responsible for listing consumer records by interval and plant
func (c *consumerService) ListPlantConsumptionByInterval(ctx context.Context, startTime, endTime time.Time, consumerId int) (*[]domain.Consumer, error) {
	return nil, nil
}

// Function responsible for listing consumer records by interval
func (c *consumerService) ListConsumptionByInterval(ctx context.Context, startTime, endTime time.Time, consumerId int) (*domain.Consumer, error) {
	return nil, nil
}

// Function responsible for updating devices that collect consumer information
func (c *consumerService) UpdateConsumerDevice(ctx context.Context, serialDevice uint64, consumerId uint32) (*domain.Consumer, error) {
	return nil, nil
}

// Function responsible for adding a client to the list of listeners for consumption updates
func (c *consumerService) ListenConsumption(ctx context.Context, consumerId int64) (*domain.Record, error) {
	return nil, nil
}
