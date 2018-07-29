package api

import (
	ws "github.com/gorilla/websocket"
)

type Connection struct {
	conn *ws.Conn
}

func NewConnection(conn *ws.Conn) Connection {
	return Connection{conn}
}

func (c *Connection) Read() ([]byte, error) {
	_, data, err := c.conn.ReadMessage()

	return data, err
}

func (c *Connection) Write(data []byte) error {
	return c.conn.WriteMessage(ws.BinaryMessage, data)
}

func (c *Connection) CloseGracefully() {
	c.conn.WriteMessage(ws.CloseMessage, []byte{})
}

func (c *Connection) Close() {
	c.conn.Close()
}
