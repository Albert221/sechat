package models

import "github.com/Albert221/sechat/server/api/updates"

type UpdateBus interface {
	Send(client *Client, update updates.Update)
}
