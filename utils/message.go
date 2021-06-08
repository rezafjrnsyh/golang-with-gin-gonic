package utils

type ResponseMessage struct {
	Code	int
	Description	string
	Data	interface{}
}

func Response(code int, desc string, data interface{}) *ResponseMessage {
	return &ResponseMessage{Code: code,Description: desc, Data: data}
}
