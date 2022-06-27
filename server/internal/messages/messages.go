package messages

import "encoding/json"

type RouteSection struct {
	Route string `json:"route"`
}

type Message[T any] struct {
	Route string `json:"route"`
	Data  T      `json:"data"`
}

func New[T any](raw []byte) (Message[T], error) {
	var m Message[T]
	if err := json.Unmarshal(raw, &m); err != nil {
		return Message[T]{}, err
	}
	return m, nil
}

func GetRoutePart(raw []byte) (RouteSection, error) {
	var m RouteSection
	if err := json.Unmarshal(raw, &m); err != nil {
		return RouteSection{}, err
	}
	return m, nil
}
