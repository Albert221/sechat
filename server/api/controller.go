package api

import (
	ws "github.com/gorilla/websocket"
	"net/http"
	"github.com/Albert221/sechat/server/models"
	"io/ioutil"
	"github.com/gorilla/mux"
	"log"
	"encoding/base64"
)

type Controller struct {
	upgrader   ws.Upgrader
	repository ChatRepository
}

const (
	// inner

	// outer
	TypeOtherClientPublicKey = "other_client_public_key"

	// both
	TypeMessage = "message"
)

type websocketMessage struct {
	Type    string `json:"type"`
	Sender  string `json:"sender,omitempty"`
	Payload string `json:"payload"`
}

func NewController() Controller {
	return Controller{
		upgrader: ws.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (c *Controller) NewEndpoint(w http.ResponseWriter, r *http.Request) {
	encryptedPubKey, _ := ioutil.ReadAll(r.Body)

	var room models.ChatRoom
	for {
		room = models.NewChatRoom()

		if !c.repository.Exists(room.Id) {
			break
		}
	}

	room.SetFirstClient(
		models.NewClient(encryptedPubKey))

	c.repository.Persist(&room)

	w.Write([]byte(room.Id))
}

func (c *Controller) ChatEndpoint(w http.ResponseWriter, r *http.Request) {
	roomId := mux.Vars(r)["id"]
	encryptedPubKey, _ := ioutil.ReadAll(r.Body)

	room, err := c.repository.Get(roomId)
	if err == RoomNotFoundError {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var client *models.Client
	if client = room.GetClientByPublicKey(encryptedPubKey); client == nil {
		if room.SecondClientExists() {
			// There are both clients existing, but pub key doesn't match any
			w.WriteHeader(http.StatusForbidden)
			return
		} else {
			// There is no second client yet
			room.SetSecondClient(
				models.NewClient(encryptedPubKey))
			client = &room.Clients[1]

			c.repository.Persist(room)
		}
	}

	// Return 429 when session is already opened
	if client.IsSessionOpened() {
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}

	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client.OpenSession(conn)
	c.handleWebsocket(room, client)
}

func (c *Controller) handleWebsocket(room *models.ChatRoom, client *models.Client) {
	// Wait for both clients to connect
	// ================================
	<-room.BothConnect

	neighbor := room.GetNeighborClient(client)

	// Send neighbour public key
	// ====================
	c.sendPublicKeyOfSourceTo(neighbor, client)

	// Listen for sent messages
	// ========================
	for {
		var message websocketMessage
		client.Session.Websocket.ReadJSON(&message)

		switch message.Type {
		case TypeMessage:
			room.Broadcast(message)
		}
	}
}

func (c *Controller) sendPublicKeyOfSourceTo(source, receiver *models.Client) {
	base64Payload := base64.StdEncoding.EncodeToString(source.EncryptedPublicKey)

	message := &websocketMessage{
		Type:    TypeOtherClientPublicKey,
		Payload: base64Payload,
	}

	receiver.Session.Websocket.WriteJSON(message)
}
