package entity

type Message struct {
	code    int
	message string
	result  string
}

//设置Message结果
func (message *Message) SetResult(result string) {
	message.result = result
}
