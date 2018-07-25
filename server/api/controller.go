package api

import (
	ws "github.com/gorilla/websocket"
	"net/http"
	"github.com/Albert221/sechat/server/models"
	"github.com/Albert221/sechat/server/api/updates"
	"io/ioutil"
	"github.com/gorilla/mux"
	"log"
	"encoding/base64"
)

type Controller struct {
	upgrader    ws.Upgrader
	repository  ChatRepository
}

func NewController(repository ChatRepository) Controller {
	return Controller{
		upgrader: ws.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			// FIXME: ONLY FOR DEBUG!
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		repository:   repository,
	}
}

func (c *Controller) NewEndpoint(w http.ResponseWriter, r *http.Request) {
	encryptedPubKey, _ := ioutil.ReadAll(r.Body)

	var room models.Room
	for {
		room = models.NewChatRoom()

		if !c.repository.Exists(room.Id) {
			break
		}
	}

	firstClient := room.NewClient(encryptedPubKey)
	room.SetFirstClient(&firstClient)

	c.repository.Persist(&room)

	w.Write([]byte(room.Id))
}

func (c *Controller) ChatEndpoint(w http.ResponseWriter, r *http.Request) {
	roomId := mux.Vars(r)["id"]

	room, err := c.repository.Get(roomId)
	if err == RoomNotFoundError {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c.handleWebsocket(&room, conn)
}

func (c *Controller) handleWebsocket(room *models.Room, conn *ws.Conn) {
	// Wait for the client's public key
	// ================================
	var client *models.Client
	var publicKey []byte
	for {
		var update interface{}
		conn.ReadJSON(&update)

		log.Println(update)

		updateType, payload, err := updates.ParseUpdate(update)
		if err != nil || updateType != updates.TypeMyPublicKey {
			continue
		}

		publicKey, err = base64.StdEncoding.DecodeString(payload.(string))
		if err != nil {
			continue
		}

		if client = room.GetClientByPublicKey(publicKey); client == nil {
			if room.SecondClientExists() {
				// There are both clients existing, but pub key doesn't match any
				conn.WriteJSON(
					updates.NewErrorUpdate(1, "forbidden").UpdateStruct())
				conn.Close()
				return
			} else {
				// There is no second client yet
				secondClient := room.NewClient(publicKey)
				room.SetSecondClient(&secondClient)
				client = &secondClient

				c.repository.Persist(room)
			}
		}

		break
	}

	// Check if session doesn't already exist and open it if it doesn't
	// =======
	if client.Socket != nil {
		conn.WriteJSON(
			updates.NewErrorUpdate(2, "session already opened"))
		conn.Close()
		return
	} else {
		client.Socket = conn
	}

	// Wait for both clients to connect
	// ================================
	if !room.BothConnected {
		<-room.BothConnectChan

		// Refresh room instance, it now has the neighbor pubkey
		newRoom, err := c.repository.Get(room.Id)
		if err != nil {
			return
		}

		room = &newRoom
		client = room.GetClientByPublicKey(publicKey)
		client.Socket = conn
	}

	neighbor := room.GetNeighborClient(client)

	// Send neighbour public key
	// ====================
	otherPublicKeyUpdate := updates.NewOtherPublicKeyUpdate(
		neighbor.EncryptedPublicKey)
	client.SendUpdate(&otherPublicKeyUpdate)

	// Listen for sent updates
	// ========================
	for {
		var update interface{}
		err := conn.ReadJSON(&update)
		if err != nil {
			return
		}

		updateType, payload, err := updates.ParseUpdate(update)
		if err != nil {
			continue
		}

		switch updateType {
		case updates.TypeMessage:
			message, err := base64.StdEncoding.DecodeString(payload.(string))
			if err != nil {
				continue
			}
			client.SendMessage(message)

			c.repository.Persist(room)
		}
	}
}
