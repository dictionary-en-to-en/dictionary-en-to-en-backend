package tools

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"reflect"
	"sync"
)

var poolSender = sync.Pool{New: func() interface{} {
	return new(Send)
}}

// SendResponse ....
type SendResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Error  *Error      `json:"error"`
}

// Error ....
type Error struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func (e Error) Error() string {
	//TODO implement me
	panic(e.Message)
}

func NewError(message string, code int, data interface{}) *Error {
	return &Error{
		Message: message,
		Code:    code,
		Data:    data,
	}
}

// Send ...
type Send struct {
	C            *fiber.Ctx
	Status       bool
	ResponseCode int
	ErrorMessage *Error
	Data         interface{}
}

// Send ...
func (provider *Send) Send() {
	response := new(SendResponse)

	provider.C.Response().Header.Set("Content-Type", "application/json")
	provider.C.Response().Header.SetStatusCode(provider.ResponseCode)

	response.Status = provider.Status
	response.Data = provider.Data
	response.Error = provider.ErrorMessage

	v, err := jsoniter.Marshal(response)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprint(provider.C, string(v))
	if err != nil {
		fmt.Println("Problem In Response - Internal Error")
	}

}

func (provider *Send) SendCore() {

	provider.C.Response().Header.Set("Content-Type", "application/json")
	provider.C.Response().Header.SetStatusCode(provider.ResponseCode)

	v, err := jsoniter.Marshal(provider.Data)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprint(provider.C, string(v))
	if err != nil {
		fmt.Println("Problem In Response - Internal Error")
	}

}

// Sender ...
func Sender(connection *fiber.Ctx, status bool, responseCode int, err *Error, data interface{}) {
	response := poolSender.Get().(*Send)
	defer releasePoolSender(response)

	response.Status = status
	response.ResponseCode = responseCode
	response.C = connection
	response.Data = data
	if err != nil {
		response.ErrorMessage = err
	}
	response.Send()
}

func releasePoolSender(response *Send) {
	Clear(response)
	poolSender.Put(response)
}

// Panic recovery
func SendError(c *fiber.Ctx) {
	if r := recover(); r != nil {
		// Customizing response status code and send data if you want
		if e, ok := r.(*Error); ok {
			Sender(c, false, e.Code, e, nil)
		} else if e, ok := r.(error); ok {
			// Send error message with 400 status code
			Sender(c, false, http.StatusBadRequest, &Error{
				Code:    http.StatusBadRequest,
				Message: e.Error(),
				Data:    nil,
			}, nil)

		} else if v, ok := r.(map[string]error); ok {
			// Send error message with 400 status code
			t := ""
			for eKey, errMessage := range v {
				t += eKey + errMessage.Error() + " , "
			}
			Sender(c, false, http.StatusBadRequest, &Error{
				Code:    http.StatusBadRequest,
				Message: t,
				Data:    nil,
			}, nil)
		}
	}
}

func Clear(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}
