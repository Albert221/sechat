package models

import (
	"time"
	"encoding/base64"
	"github.com/Albert221/sechat/server/api/updates"
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
		"type":    updates.TypeMessage,
		"payload": m,
	}
}
