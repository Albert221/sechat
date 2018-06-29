package models

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

func (c *Client) SendUpdate(update Update) {
	if c.IsSessionOpened() {
		c.Session.Websocket.WriteJSON(update.UpdateStruct())
	}
}
