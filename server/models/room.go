package models

import (
	"github.com/Albert221/sechat/server/utils"
	"time"
	"bytes"
	"github.com/Albert221/sechat/server/api/updates"
)

type Room struct {
	Id              string
	CreatedAt       int64
	Clients         [2]*Client
	BothConnectChan chan bool
	BothConnected   bool
	Messages        []*Message
}

func NewChatRoom() Room {
	return Room{
		Id:              utils.RandomString(8),
		CreatedAt:       time.Now().Unix(),
		BothConnectChan: make(chan bool),
	}
}

func (cr *Room) SetFirstClient(client *Client) {
	cr.Clients[0] = client
}

func (cr *Room) SecondClientExists() bool {
	return cr.Clients[1] != nil
}

func (cr *Room) SetSecondClient(client *Client) {
	cr.Clients[1] = client

	cr.BothConnectChan <- true
	cr.BothConnected = true
}

// GetClientByPublicKey returns pointer to the client corresponding
// to the given encrypted public key or nil when given does not exist.
func (cr *Room) GetClientByPublicKey(pubKey []byte) *Client {
	for _, client := range cr.Clients {
		if client != nil && bytes.Equal(client.EncryptedPublicKey, pubKey) {
			return client
		}
	}

	return nil
}

// GetNeighborClient returns the other client than the one specified.
func (cr *Room) GetNeighborClient(client *Client) *Client {
	if bytes.Equal(client.EncryptedPublicKey, cr.Clients[0].EncryptedPublicKey) {
		return cr.Clients[1]
	} else {
		return cr.Clients[0]
	}
}

func (cr *Room) SendUpdate(update updates.Update) {
	for _, client := range cr.Clients {
		client.SendUpdate(update)
	}
}
