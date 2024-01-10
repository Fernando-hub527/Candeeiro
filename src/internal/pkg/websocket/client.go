package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func (*Client) writePump() {

}

func (*Client) readPump() {

}

func DefaultHandlesWs(hub *Hub, context echo.Context) error {
	return nil
}
