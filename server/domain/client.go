package domain

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Client struct {
	Id               uuid.UUID
	EncodedPublicKey string
	InboundChan      chan *Message
	connection       Connection
	room             *Room
}

func NewClient(room *Room, encodedPublicKey string, connection Connection) Client {
	return Client{
		Id:               uuid.New(),
		EncodedPublicKey: encodedPublicKey,
		InboundChan:      make(chan *Message, 255),
		room:             room,
		connection:       connection,
	}
}

func (c *Client) RunReading() {
	defer func() {
		c.connection.Close()
	}()

	for {
		data, err := c.connection.Read()
		if err != nil {
			break
		}

		var message Message
		err = json.Unmarshal(data, &message)
		if err != nil {
			continue
		}

		c.room.OutboundChan <- &message
	}
}

func (c *Client) RunWriting() {
	defer func() {
		c.connection.Close()
	}()

	for {
		select {
		case message, ok := <-c.InboundChan:
			if !ok {
				c.connection.CloseGracefully()
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				continue
			}

			c.connection.Write(data)
		}
	}
}
