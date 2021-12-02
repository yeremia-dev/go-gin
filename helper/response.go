package helper

import "strings"

//Response struct is used to build response data
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmptyObject is used if there is response that needs an empty object
type EmptyObject struct{}

//BuildResponse function is used to build success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}

	return res
}

//BuildErrorResponse is used to build a response if there is any error
func BuildErrorResponse(message string, err string, data interface{}) Response {

	splittedErrorMessage := strings.Split(err, "\n")

	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedErrorMessage,
		Data:    data,
	}

	return res
}
