package updates

type ErrorUpdate struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorUpdate(code int, message string) ErrorUpdate {
	return ErrorUpdate{
		Code:    code,
		Message: message,
	}
}

func (e ErrorUpdate) UpdateStruct() map[string]interface{} {
	return map[string]interface{}{
		"type":    TypeError,
		"payload": e,
	}
}
