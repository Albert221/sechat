package models

type Message struct {
	Content []byte
	Author  *Client
	SentAt  int64
}