package openapi

import (
	"encoding/json"
	"log"
)

type Error struct {
	Code    int `json:"omitempty"`
	Message string
	Data    interface{} `json:"omitempty"`
}

//noinspection GoUnusedExportedFunction
func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

//noinspection GoUnusedExportedFunction
func NewErrorWithData(code int, message string, data interface{}) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (e *Error) Error() string {
	bs, err := json.Marshal(e)
	if err != nil {
		log.Print(e.Code, e.Message)
		panic(err)
	}
	return string(bs)
}
