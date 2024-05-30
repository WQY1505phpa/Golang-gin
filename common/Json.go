package common

type Response struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func JsonRes(code string, message string, data map[string]interface{}) *Response {
	res := &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return res
}
