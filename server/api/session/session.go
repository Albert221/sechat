package session

import (
	ws "github.com/gorilla/websocket"
	"github.com/Albert221/sechat/server/utils"
	"errors"
)

type SocketStore struct {
	sockets map[string]*ws.Conn
}

var SocketDoesNotExist = errors.New("socket with given id does not exist")

func NewSocketStore() SocketStore {
	return SocketStore{}
}

func (s *SocketStore) Add(socket *ws.Conn) (id string) {
	for id == "" || s.Exists(id) {
		id = utils.RandomString(5)
	}

	s.sockets[id] = socket
	return
}

func (s *SocketStore) Get(id string) (*ws.Conn, error) {
	if !s.Exists(id) {
		return nil, SocketDoesNotExist
	}

	return s.sockets[id], nil
}

func (s *SocketStore) Exists(id string) bool {
	_, exists := s.sockets[id]

	return exists
}

func (s *SocketStore) Remove(id string) error {
	if !s.Exists(id) {
		return SocketDoesNotExist
	}

	delete(s.sockets, id)

	return nil
}