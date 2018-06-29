package models

import (
	"time"
	"encoding/base64"
)

type Message struct {
	Content    []byte  `json:"content"`
	Author     *Client `json:"-"`
	AuthorName string  `json:"author"`
	SentAt     int64   `json:"sentAt"`
}

func NewMessage(content []byte, author *Client) Message {
	return Message{
		Content:    content,
		Author:     author,
		SentAt:     time.Now().Unix(),
		AuthorName: base64.StdEncoding.EncodeToString(author.EncryptedPublicKey),
	}
}

func (m *Message) UpdateStruct() map[string]interface{} {
	return map[string]interface{}{
		"type":    TypeMessage,
		"payload": m,
	}
}
