package api

import (
	"github.com/Albert221/sechat/server/models"
	"errors"
	"log"
)

var (
	RoomNotFoundError = errors.New("chatroom has not been found")
)

type ChatRepository interface {
	Get(id string) (models.Room, error)
	Exists(id string) bool
	Persist(room *models.Room) error
	Remove(room *models.Room) error
}

type InMemoryChatRepository struct {
	rooms []models.Room
}

func NewInMemoryChatRepository() InMemoryChatRepository {
	return InMemoryChatRepository{}
}

func (r *InMemoryChatRepository) Get(id string) (models.Room, error) {
	for i, room := range r.rooms {
		if room.Id == id {
			return r.rooms[i], nil
		}
	}

	return models.Room{}, RoomNotFoundError
}

func (r *InMemoryChatRepository) Exists(id string) bool {
	for _, room := range r.rooms {
		if room.Id == id {
			return true
		}
	}

	return false
}

func (r *InMemoryChatRepository) Persist(room *models.Room) error {
	// FIXME: debug
	log.Println("saving")

	for i, iRoom := range r.rooms {
		if iRoom.Id == room.Id {
			r.rooms[i] = *room
			return nil
		}
	}

	r.rooms = append(r.rooms, *room)
	return nil
}

func (r *InMemoryChatRepository) Remove(room *models.Room) error {
	for i, iRoom := range r.rooms {
		if iRoom.Id == room.Id {
			r.rooms = append(r.rooms[:i], r.rooms[i+1:]...)
			return nil
		}
	}

	return RoomNotFoundError
}