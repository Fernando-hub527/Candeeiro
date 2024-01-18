package broker

import (
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
)

type Broker struct {
	channel *amqp091.Channel
}

func NewBrokerRabbit(url string) (*Broker, *errors.RequestError) {
	chAmqp, err := openChannel(url)
	if err != nil {
		return nil, errors.NewInternalErros("Unable to connect to broker")
	}
	return &Broker{
		channel: chAmqp,
	}, nil
}

func openChannel(url string) (*amqp091.Channel, error) {
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

func (broker *Broker) Consumer(queue string) chan IBrokerMessager {
	chanMessage := make(chan IBrokerMessager)
	go broker.listenToQueues(chanMessage, broker.channel, queue)
	return chanMessage
}

func (broker *Broker) listenToQueues(channel chan IBrokerMessager, channelRabbit *amqp091.Channel, queue string) error {
	msgs, err := channelRabbit.Consume(
		queue,
		"electricity_consumption",
		false, false, false, false, nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		channel <- NewBrokerMessage(msg)
	}
	return nil
}

type BrokerMessagerRabbit struct {
	msg amqp091.Delivery
}

func NewBrokerMessage(msg amqp091.Delivery) *BrokerMessagerRabbit {
	return &BrokerMessagerRabbit{
		msg: msg,
	}
}

func (broker *BrokerMessagerRabbit) Reject() error {
	return broker.msg.Reject(false)
}

func (broker *BrokerMessagerRabbit) Accept() error {
	return broker.msg.Ack(false)
}

func (broker *BrokerMessagerRabbit) GetMessager() []byte {
	return broker.msg.Body
}
