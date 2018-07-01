package updates

import (
	"github.com/oliveagle/jsonpath"
	"errors"
)

const (
	// inner
	TypeMyPublicKey = "myPublicKey"
	// outer
	TypeOtherPublicKey = "otherPublicKey"
	TypeError          = "error"
	// both
	TypeMessage = "message"
)

// Update is used for sending updates to clients
type Update interface {
	// UpdateStruct returns an associative array with two
	// required members: type and payload.
	UpdateStruct() map[string]interface{}
}

func ParseUpdate(v interface{}) (updateType string, payload interface{}, err error) {
	parsedUpdateType, err := jsonpath.JsonPathLookup(v, "$.type")
	if err != nil {
		return "", nil, err
	}
	updateType, correct := parsedUpdateType.(string)
	if !correct {
		return "", nil, errors.New("error")
	}

	payload, err = jsonpath.JsonPathLookup(v, "$.payload")
	if err != nil {
		return "", nil, err
	}

	return
}
