package domain

import "sync"

type Pool struct {
	used         map[string]*Room
	lock         sync.Mutex
	storage      Storage
	onRoomUpdate func(room *Room)
}

func NewPool(storage Storage) Pool {
	pool := Pool{
		storage: storage,
		used:    make(map[string]*Room),
		lock:    sync.Mutex{},
	}
	pool.onRoomUpdate = func(room *Room) {
		pool.storage.Persist(room)
	}

	return pool
}

func (p *Pool) Create() *Room {
	room := NewRoom()
	room.OnUpdate = p.onRoomUpdate

	p.lock.Lock()
	p.used[room.Id.String()] = &room
	p.lock.Unlock()

	p.storage.Persist(&room)

	return &room
}

func (p *Pool) Get(id string) (*Room, error) {
	if room, ok := p.used[id]; ok {
		return room, nil
	}

	room, err := p.storage.Get(id)
	if err != nil {
		return nil, err
	}

	room.refresh()

	p.lock.Lock()
	p.used[room.Id.String()] = &room
	p.lock.Unlock()

	room.OnUpdate = p.onRoomUpdate

	return &room, nil
}

func (p *Pool) Free(room *Room) {
	p.lock.Lock()
	delete(p.used, room.Id.String())
	p.lock.Unlock()
}

func (p *Pool) Remove(room *Room) error {
	p.Free(room)

	return p.storage.Remove(room)
}
