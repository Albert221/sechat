package api

import (
	"github.com/Albert221/sechat/server/models"
	"errors"
)

var (
	RoomNotFoundError = errors.New("chatroom has not been found"),
)

type ChatRepository interface {
	Get(id string) (*models.ChatRoom, error)
	Exists(id string) bool
	Persist(room *models.ChatRoom) error
	Remove(room *models.ChatRoom) error
}
