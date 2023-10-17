package helper

//Response struct json
type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//Struct to give Empty Object
type EmptyObj struct {
}

//function Build Response
func BuildResponse(code string, message string, data interface{}) Response {
	res := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return res
}
