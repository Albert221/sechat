package models

import (
	"github.com/Albert221/sechat/server/api/updates"
	"log"
)

type Client struct {
	EncryptedPublicKey []byte
	Session            ClientSession
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
	if c.IsSessionOpened() {
		log.Println("sending update lol")
		c.Session.Websocket.WriteJSON(update.UpdateStruct())
	}
}
