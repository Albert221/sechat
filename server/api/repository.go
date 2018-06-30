package api

import (
	"github.com/Albert221/sechat/server/models"
	"errors"
)

var (
	RoomNotFoundError = errors.New("chatroom has not been found")
)

type ChatRepository interface {
	Get(id string) (*models.Room, error)
	Exists(id string) bool
	Persist(room *models.Room) error
	Remove(room *models.Room) error
}
