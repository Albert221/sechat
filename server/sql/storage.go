package sql

import (
	"github.com/Albert221/sechat/server/domain"
	"errors"
)

// TODO: Use SQL here
type Storage struct {
	rooms map[string]domain.Room
}

func NewStorage() Storage {
	return Storage{
		rooms: make(map[string]domain.Room),
	}
}

func (s *Storage) Get(id string) (domain.Room, error) {
	if room, ok := s.rooms[id]; ok {
		return room, nil
	}

	return domain.Room{}, errors.New("can't find room with given id")
}

func (s *Storage) Persist(room *domain.Room) error {
	s.rooms[room.Id.String()] = *room

	return nil
}

func (s *Storage) Remove(room *domain.Room) error {
	delete(s.rooms, room.Id.String())

	return nil
}
