package entity

type Message struct {
	Code    int
	Message string
	Result  interface{}
}

func BadResponse(msg string) Message {
	return Message{Code: MESSAGE_BAD, Message: msg}
}

// 只有正常请求才能有结果
func OKResponse(message interface{}) Message {
	result := Message{Code: MESSAGE_OK, Message: "OK", Result: message}
	return result
}

func ErrorResponse() Message {
	return Message{Code: MESSAGE_ERROR, Message: "error"}
}
