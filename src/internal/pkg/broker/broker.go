package broker

type IBroker interface {
	Consumer(queue string) (chan IBrokerMessager, error)
}

type IBrokerMessager interface {
	Reject() error
	Accept() error
	GetMessager() []byte
}
