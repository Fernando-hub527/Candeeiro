package broker

import "github.com/rabbitmq/amqp091-go"

type IBroker interface {
	Consumer(queue string) chan IBrokerMessager
	listenToQueues(channel chan IBrokerMessager, channelRabbit *amqp091.Channel, queue string) error
}

type IBrokerMessager interface {
	Reject() error
	Accept() error
	GetMessager() []byte
}
