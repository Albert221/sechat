package models

const (
	// outer
	TypeOtherPublicKey = "otherPublicKey"
	// both
	TypeMessage = "message"
)

// Update is used for sending updates to clients
type Update interface {
	// UpdateStruct returns an associative array with two
	// required members: type and payload.
	UpdateStruct() map[string]interface{}
}
