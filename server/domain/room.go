package domain

import "github.com/google/uuid"

type Room struct {
	Id           uuid.UUID
	Clients      [2]*Client
	Messages     []*Message
	OnUpdate     func(room *Room)
	RegisterChan chan *Client
	OutboundChan chan *Message
	fresh        bool
}

func NewRoom() Room {
	return Room{
		Id:           uuid.New(),
		RegisterChan: make(chan *Client),
		OutboundChan: make(chan *Message),
		fresh:        true,
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.RegisterChan:
			if r.Clients[0] == nil {
				r.Clients[0] = client
			} else if r.Clients[1] == nil {
				r.Clients[1] = client
			} else {
				// TODO: throw error
				break
			}

			r.OnUpdate(r)
		case message := <-r.OutboundChan:
			r.Messages = append(r.Messages, message)
			for _, client := range r.Clients {
				if client == nil {
					continue
				}

				client.InboundChan <- message
			}

			r.OnUpdate(r)
		}
	}
}

func (r *Room) refresh() {
	if r.fresh {
		return
	}

	r.RegisterChan = make(chan *Client)
	r.OutboundChan = make(chan *Message)
	r.fresh = true
}