package rabbitmq

import "github.com/rabbitmq/amqp091-go"

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

func Consumer(channel chan amqp091.Delivery, channelRabbit *amqp091.Channel, queue string) {

}
