package models

import (
	"github.com/Albert221/sechat/server/api/updates"
	"log"
	ws "github.com/gorilla/websocket"
)

type Client struct {
	EncryptedPublicKey []byte
	Socket             *ws.Conn
	room               *Room
}

func (cr *Room) NewClient(pubKey []byte) Client {
	return Client{
		EncryptedPublicKey: pubKey,
		room:               cr,
	}
}

func (c *Client) SendMessage(messageContent []byte) {
	message := NewMessage(messageContent, c)

	c.room.Messages = append(c.room.Messages, &message)
	c.room.SendUpdate(&message)
}

func (c *Client) SendUpdate(update updates.Update) {
	log.Println("sending update lol")
	c.Socket.WriteJSON(update.UpdateStruct())
}
