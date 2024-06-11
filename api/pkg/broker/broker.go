package broker

type IBroker interface {
	Consumer(queue string) chan IBrokerMessager
	listenToQueues(channel chan IBrokerMessager, queue string) error
}

type IBrokerMessager interface {
	Reject() error
	Accept() error
	GetMessager() []byte
}

func deserializeMessage[T any]() {

}
