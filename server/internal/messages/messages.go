package messages

import (
	"encoding/json"

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

type RequestMessage struct {
	Route string `json:"route"`
	Data  any    `json:"data"`
}

type ResponseMessage struct {
	Code string `json:"code"`
	Data any    `json:"data"`
}

func GetRoutePart(raw []byte) (RouteSection, error) {
	var m RouteSection
	if err := json.Unmarshal(raw, &m); err != nil {
		return RouteSection{}, err
	}
	return m, nil
}

func NewResponseMessage(data any) ResponseMessage {
	return ResponseMessage{
		Code: "SUCCESS",
		Data: data,
	}
}

func NewResponseEmptyMessage() ResponseMessage {
	return ResponseMessage{
		Code: "SUCCESS",
		Data: nil,
	}
}

func NewResponseErrorMessage(err errors.Error) ResponseMessage {
	return ResponseMessage{
		Code: string(err.GetCode()),
		Data: err.GetData(),
	}
}
