package karmatrain

// Much of this code is taken from the Gorilla Websocket chat example:
//   https://github.com/gorilla/websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	writeWait      = 5 * time.Second
	pongWait       = 15 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// yolo lol.
		return true
	},
}

type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

func NewWebsocket(response http.ResponseWriter, request *http.Request) (*connection, error) {
	c := &connection{}
	ws, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		log.Println(err)
		return c, err
	}
	c.ws = ws
	c.send = make(chan []byte, 256)
	go c.writePump()
	return c, nil
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// Write Kafka topic messages down to the websocket client.
func (c *connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
