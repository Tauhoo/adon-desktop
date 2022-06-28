package messages

import "encoding/json"

type Type = int

const (
	ERROR   Type = 0
	SUCCESS Type = 1
)

type RouteSection struct {
	Route string `json:"route"`
}

type RequestMessage[T any] struct {
	Type  Type   `json:"type"`
	Route string `json:"route"`
	Data  T      `json:"data"`
}

type ResponseMessage[T any] struct {
	Type Type `json:"type"`
	Data T    `json:"data"`
}

func GetRoutePart(raw []byte) (RouteSection, error) {
	var m RouteSection
	if err := json.Unmarshal(raw, &m); err != nil {
		return RouteSection{}, err
	}
	return m, nil
}
