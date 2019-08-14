package entity

type Message struct {
	Code    int
	Message string
	Result  string
}

//设置Message结果
func (message *Message) SetResult(result string) {
	message.Result = result
}
