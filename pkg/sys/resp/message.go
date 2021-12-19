package resp

type Message struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BadResponse(msg string) Message {
	return Message{Code: MESSAGE_BAD, Message: msg}
}

func OkResponse() Message {
	result := Message{Code: MESSAGE_OK, Message: "ok"}
	return result
}

func OkResponseWithRet(message interface{}) Message {
	return Message{Code: MESSAGE_OK, Message: "ok", Data: message}
}

func ErrorResponse() Message {
	return Message{Code: MESSAGE_ERROR, Message: "error"}
}
