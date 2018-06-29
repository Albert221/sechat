package models

import (
	ws "github.com/gorilla/websocket"
)

type Client struct {
	EncryptedPublicKey []byte
	Session            ClientSession
}

type ClientSession struct {
	open           bool
	Websocket      *ws.Conn
}

func NewClient(pubKey []byte) Client {
	return Client{
		EncryptedPublicKey: pubKey,
	}
}

func (c *Client) IsSessionOpened() bool {
	return c.Session.open
}

func (c *Client) OpenSession(websocket *ws.Conn) {
	c.Session = ClientSession{
		open:           true,
		Websocket:      websocket,
	}

	websocket.SetCloseHandler(func(code int, text string) error {
		c.CloseSession()
		return nil
	})
}

func (c *Client) CloseSession() {
	// FIXME: or deleteme, not sure if this won't create infinite loop
	c.Session.Websocket.Close()
	c.Session = ClientSession{}
}
