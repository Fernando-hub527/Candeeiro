package websocket

import "github.com/Fernando-hub527/candieiro/internal/pkg/utils"

type Hub struct {
	Romms map[string][]*Client

	registers chan map[*Client][]string

	unregister chan *Client

	messages chan map[string][]byte
}

/*
Dúvida: Aqui é retornado o endereço da struct, porém essa struct foi criada dentro da função,
Nesse caso, quando a função termina as variaveis da função são distruidas, logo, como é possível
retornar o endereço de uma variavel que quando a função acabar, teoricamente, não existe mais ?
*/
func NewHub() *Hub {
	return &Hub{
		Romms:      map[string][]*Client{},
		registers:  make(chan map[*Client][]string),
		unregister: make(chan *Client),
		messages:   make(chan map[string][]byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.registers:
			h.Romms = addClientsToRoom(client, h.Romms)
		case client := <-h.unregister:
			h.Romms = removeClientToRooms(client, h.Romms)
		case messages := <-h.messages:
			senMessage(messages, h.Romms)
		}
	}

}

func senMessage(messages map[string][]byte, rooms map[string][]*Client) {
	for room, message := range messages {
		for _, client := range rooms[room] {
			client.send <- message
		}
	}
}

func addClientsToRoom(client map[*Client][]string, rooms map[string][]*Client) map[string][]*Client {
	for client, clientRooms := range client {
		for _, room := range clientRooms {
			rooms[room] = append(rooms[room], client)
			return rooms
		}
	}
	return rooms
}

func removeClientToRooms(client *Client, rooms map[string][]*Client) map[string][]*Client {
	for _, clientRoom := range client.rooms {
		rooms[clientRoom] = utils.RemoveItemFromSlice2[*Client](rooms[clientRoom], client, func(itemA, itemB *Client) bool { return itemA.id == itemB.id })
		if len(rooms[clientRoom]) == 0 {
			delete(rooms, clientRoom)
		}
		close(client.send)
		return rooms
	}
	return rooms
}
