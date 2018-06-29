package models

import (
	"github.com/Albert221/sechat/server/utils"
	"time"
	"bytes"
)

type Room struct {
	Id          string
	CreatedAt   int64
	Clients     [2]Client
	BothConnect <-chan bool
	Messages    []*Message
}

func NewChatRoom() Room {
	return Room{
		Id:        utils.RandomString(8),
		CreatedAt: time.Now().Unix(),
	}
}

func (cr *Room) SetFirstClient(client Client) {
	cr.Clients[0] = client
}

func (cr *Room) SecondClientExists() bool {
	return len(cr.Clients[1].EncryptedPublicKey) > 0
}

func (cr *Room) SetSecondClient(client Client) {
	cr.Clients[1] = client

	// Send twice, for both clients
	cr.BothConnect <- true
	cr.BothConnect <- true
}

// GetClientByPublicKey returns pointer to the client corresponding
// to the given encrypted public key or nil when given does not exist.
func (cr *Room) GetClientByPublicKey(pubKey []byte) *Client {
	for _, client := range cr.Clients {
		if bytes.Equal(client.EncryptedPublicKey, pubKey) {
			return &client
		}
	}

	return nil
}

// GetNeighborClient returns the other client than the one specified.
func (cr *Room) GetNeighborClient(client *Client) *Client {
	if bytes.Equal(client.EncryptedPublicKey, cr.Clients[0].EncryptedPublicKey) {
		return &cr.Clients[1]
	} else {
		return &cr.Clients[0]
	}
}

func (cr *Room) SendUpdate(update Update) {
	for _, client := range cr.Clients {
		client.SendUpdate(update)
	}
}