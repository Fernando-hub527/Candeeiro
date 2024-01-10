package websocket

type Hub struct {
	clients map[*Client]bool

	register chan *Client

	unregister chan *Client

	messages chan []byte
}

/*
Dúvida: Aqui é retornado o endereço da struct, porém essa struct foi criada dentro da função,
Nesse caso, quando a função termina as variaveis da função são distruidas, logo, como é possível
retornar o endereço de uma variavel que quando a função acabar, teoricamente, não existe mais ?
*/
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		messages:   make(chan []byte),
	}
}

func (h *Hub) Run() {

}
