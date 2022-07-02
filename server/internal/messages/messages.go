package messages

import (
	"github.com/Tauhoo/adon-desktop/internal/errors"
)

type Type = int

const (
	ERROR   Type = 0
	SUCCESS Type = 1
)

type RouteSection struct {
	Route string `json:"route"`
}

type RequestMessage[T any] struct {
	Route string `json:"route"`
	Data  T      `json:"data"`
}

type ResponseMessage[T any] struct {
	Code string `json:"code"`
	Data T      `json:"data"`
}

func NewResponseMessage[T any](data T) ResponseMessage[T] {
	return ResponseMessage[T]{
		Code: "SUCCESS",
		Data: data,
	}
}

func NewRequestMessage[T any](route string, data T) RequestMessage[T] {
	return RequestMessage[T]{
		Route: route,
		Data:  data,
	}
}

func NewResponseEmptyMessage() ResponseMessage[interface{}] {
	return ResponseMessage[interface{}]{
		Code: "SUCCESS",
		Data: nil,
	}
}

func NewResponseErrorMessage(err errors.Error) ResponseMessage[interface{}] {
	return ResponseMessage[interface{}]{
		Code: string(err.GetCode()),
		Data: err.GetData(),
	}
}
