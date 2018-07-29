package api

import (
	"github.com/Albert221/sechat/server/domain"
	"github.com/gorilla/mux"
	ws "github.com/gorilla/websocket"
	"net/http"
)

type Controller struct {
	pool     *domain.Pool
	upgrader ws.Upgrader
	router   *mux.Router
}

func NewController(pool *domain.Pool, router *mux.Router) Controller {
	return Controller{
		pool: pool,
		upgrader: ws.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			// FIXME: ONLY FOR DEBUG!
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		router: router,
	}
}

func (c *Controller) New(w http.ResponseWriter, r *http.Request) {
	room := c.pool.Create()

	go room.Run()

	url, _ := c.router.Get("room").URL("id", room.Id.String())
	w.Write([]byte(url.String()))
}

func (c *Controller) Room(w http.ResponseWriter, r *http.Request) {
	roomId := mux.Vars(r)["id"]

	room, err := c.pool.Get(roomId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// TODO: Authentication

	conn, err := c.upgrader.Upgrade(w, r, nil)

	encodedPublicKey := "abc123"
	connection := NewConnection(conn)

	// In case of new client
	client := domain.NewClient(room, encodedPublicKey, &connection)
	room.RegisterChan <- &client

	go client.RunReading()
	go client.RunWriting()
}
