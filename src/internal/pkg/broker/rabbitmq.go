package broker

import (
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
)

type Broker struct {
	channel *amqp091.Channel
}

func NewBroker(url string) (*Broker, *errors.RequestError) {
	chAmqp, err := OpenChannel(url)
	if err != nil {
		return nil, errors.NewInternalErros("Unable to connect to broker")
	}
	return &Broker{
		channel: chAmqp,
	}, nil
}

func OpenChannel(url string) (*amqp091.Channel, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (broker *Broker) Consumer(queue string) (chan IBrokerMessager, error) {
	return nil, nil
}

func (broker *Broker) listenToQueues(channel chan amqp091.Delivery, channelRabbit *amqp091.Channel, queue string) {
}
