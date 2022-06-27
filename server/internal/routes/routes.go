package routes

import (
	"encoding/json"

	"github.com/Tauhoo/adon-desktop/internal/messages"
	"github.com/Tauhoo/adon-desktop/internal/services"
	"github.com/asticode/go-astilectron"
)

type RouteHandler interface {
	GetRoute() string
	Handle(m *astilectron.EventMessage) (v interface{})
}

type routeHandler[I any, O any] struct {
	route   string
	handler func(messages.Message[I]) messages.Message[O]
}

func (rh routeHandler[I, O]) GetRoute() string {
	return rh.route
}

func (rh routeHandler[I, O]) Handle(m *astilectron.EventMessage) (v interface{}) {
	var input messages.Message[I]
	if err := m.Unmarshal(&input); err != nil {
		return err
	}

	output := rh.handler(input)

	if raw, err := json.Marshal(output); err != nil {
		return err
	} else {
		return raw
	}
}

func Regist(service services.Service, w *astilectron.Window) {
	routeHandlerList := []RouteHandler{}
	w.OnMessage(func(m *astilectron.EventMessage) (v interface{}) {
		var routeSection messages.RouteSection
		if err := m.Unmarshal(routeSection); err != nil {
			return err
		}

		for _, rh := range routeHandlerList {
			if rh.GetRoute() == routeSection.Route {
				return rh.Handle(m)
			}
		}

		return nil
	})
}
