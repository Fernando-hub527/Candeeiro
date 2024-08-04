package consumer

import (
	"context"
	"time"

	"github.com/Fernando-hub527/candieiro/hexa/internal/core/consumer/domain"
)

type mongodbRepository struct {
}

func NewMongodbRepository() *mongodbRepository {
	return &mongodbRepository{}
}

func (m *mongodbRepository) CreateConsumptionRecord(ctx context.Context, record *domain.Record) (*domain.Record, error) {
	return nil, nil
}

func (m *mongodbRepository) ListConsumptionByIntervalAndConsumer(ctx context.Context, startTime, endTime time.Time, consumer uint32) (*[]domain.Record, error) {
	return nil, nil
}

func (m *mongodbRepository) ListConsumptionByIntervalAndPlant(ctx context.Context, startTime, endTime time.Time, plant uint32) (*[]domain.Record, error) {
	return nil, nil
}

func (m *mongodbRepository) CreateConsumer(ctx context.Context, consumer *domain.Consumer) (*domain.Consumer, error) {
	return nil, nil
}

func (m *mongodbRepository) ListConsumersByPlant(ctx context.Context, plant uint32) (*[]domain.Consumer, error) {
	return nil, nil
}

func (m *mongodbRepository) findConsumerById(ctx context.Context, id uint32) (*domain.Consumer, error) {
	return nil, nil
}
