package models

import (
	"github.com/Albert221/sechat/server/utils"
	"time"
	"bytes"
)

type ChatRoom struct {
	Id          string
	CreatedAt   int64
	Clients     [2]Client
	BothConnect <-chan bool
	Messages    []*Message
}

func NewChatRoom() ChatRoom {
	return ChatRoom{
		Id:        utils.RandomString(8),
		CreatedAt: time.Now().Unix(),
	}
}

func (cr *ChatRoom) SetFirstClient(client Client) {
	cr.Clients[0] = client
}

func (cr *ChatRoom) SecondClientExists() bool {
	return len(cr.Clients[1].EncryptedPublicKey) > 0
}

func (cr *ChatRoom) SetSecondClient(client Client) {
	cr.Clients[1] = client

	// Send twice, for both clients
	cr.BothConnect <- true
	cr.BothConnect <- true
}

// GetClientByPublicKey returns pointer to the client corresponding
// to the given encrypted public key or nil when given does not exist.
func (cr *ChatRoom) GetClientByPublicKey(pubKey []byte) *Client {
	for _, client := range cr.Clients {
		if bytes.Equal(client.EncryptedPublicKey, pubKey) {
			return &client
		}
	}

	return nil
}

// GetNeighborClient returns the other client than the one specified.
func (cr *ChatRoom) GetNeighborClient(client *Client) *Client {
	if bytes.Equal(client.EncryptedPublicKey, cr.Clients[0].EncryptedPublicKey) {
		return &cr.Clients[1]
	} else {
		return &cr.Clients[0]
	}
}

func (cr *ChatRoom) Broadcast(v interface{}) {
	for _, client := range cr.Clients {
		if client.IsSessionOpened() {
			client.Session.Websocket.WriteJSON(v)
		}
	}
}