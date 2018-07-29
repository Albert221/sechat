package domain

type Storage interface {
	Get(id string) (Room, error)
	Persist(room *Room) error
	Remove(room *Room) error
}
