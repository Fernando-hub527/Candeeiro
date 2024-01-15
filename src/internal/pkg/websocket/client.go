package websocket

import (
	"time"

	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/Fernando-hub527/candieiro/internal/pkg/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Client struct {
	id    uint64
	hub   *Hub
	rooms []string
	conn  *websocket.Conn
	send  chan []byte
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			for i := 0; i < len(c.send); i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func DefaultHandlesWs(hub *Hub, context echo.Context) {

	id, errId := utils.ValidObjectId(context.QueryParam("plantId"), func(ctx echo.Context, err errors.RequestError) error {
		return context.String(int(err.Status), err.ToString())
	}, context)

	if errId != nil {
		return
	}

	conn, err := upgrader.Upgrade(context.Response(), context.Request(), nil)
	if err != nil {
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.registers <- map[*Client][]string{client: {id.String()}}

	go client.writePump()
}
