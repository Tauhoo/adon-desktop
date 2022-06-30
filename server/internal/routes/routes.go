package routes

import (
	"encoding/json"

	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/logs"
	"github.com/Tauhoo/adon-desktop/internal/messages"
	"github.com/Tauhoo/adon-desktop/internal/services"
	"github.com/asticode/go-astilectron"
)

// type Transaction interface {
// 	GetRequest() (messages.RequestMessage, errors.Error)
// 	GetResponse() messages.ResponseMessage

// 	SetResponse(v any)
// 	SetError(v errors.Error)
// }

// type transaction struct {
// 	m *astilectron.EventMessage
// 	response messages.ResponseMessage
// 	request  messages.RequestMessage
// }

// func (t transaction) GetRequest() messages.RequestMessage {
// 	return t.request
// }

// func (t transaction) SetResponse(v any) {
// 	t.response = messages.NewResponseMessage(v)
// }

// func (t transaction) SetError(v errors.Error) {
// 	t.response = messages.NewResponseErrorMessage(v)
// }

// func (t transaction) GetResponse() messages.ResponseMessage {
// 	return t.response
// }

func ReadEventMessage[T any](m *astilectron.EventMessage) (messages.RequestMessage[T], errors.Error) {
	var req messages.RequestMessage[T]
	var reqString string
	err := m.Unmarshal(&reqString)
	if err != nil {
		logs.ErrorLogger.Printf("fail unmarshal astilectron message to string - message: %#v, error: %s\n", m, err.Error())
		return messages.RequestMessage[T]{}, errors.New(UnmarshalFailErrCode, err.Error())
	}

	err = json.Unmarshal([]byte(reqString), &req)
	if err != nil {
		logs.ErrorLogger.Printf("fail unmarshal astilectron message to string - message: %#v, error: %s\n", reqString, err.Error())
		return messages.RequestMessage[T]{}, errors.New(UnmarshalFailErrCode, err.Error())
	}

	return req, nil
}

func ReadEventMessageRoute(m *astilectron.EventMessage) (messages.RouteSection, errors.Error) {
	var req messages.RouteSection
	var reqString string
	err := m.Unmarshal(&reqString)
	if err != nil {
		logs.ErrorLogger.Printf("fail unmarshal astilectron message to string - message: %#v, error: %s\n", m, err.Error())
		return messages.RouteSection{}, errors.New(UnmarshalFailErrCode, err.Error())
	}

	err = json.Unmarshal([]byte(reqString), &req)
	if err != nil {
		logs.ErrorLogger.Printf("fail unmarshal astilectron message to string - message: %#v, error: %s\n", reqString, err.Error())
		return messages.RouteSection{}, errors.New(UnmarshalFailErrCode, err.Error())
	}

	return req, nil
}

type Handler = func(service services.Service, m *astilectron.EventMessage) any

type Router interface {
	Route(m *astilectron.EventMessage) string
}

type router struct {
	handlers map[string]Handler
	service  services.Service
}

func (r router) Route(m *astilectron.EventMessage) string {
	routeSection, err := ReadEventMessageRoute(m)
	if err != nil {
		return `{"code":"CANNOT_READ_ROUTE_SECTION"}`
	}

	logs.InfoLogger.Printf("start route transaction - route: %s\n", routeSection.Route)

	handler, ok := handlers[routeSection.Route]
	if !ok {
		logs.ErrorLogger.Printf("route not found - route: %s\n", routeSection.Route)
		return `{"code":"ROUTE_NOT_FOUND"}`
	}

	res := handler(r.service, m)

	raw, rawerr := json.Marshal(res)
	if rawerr != nil {
		logs.ErrorLogger.Printf("marshal response fail - route: %s\n", routeSection.Route)
		return `{"code":"MARSHAL_REPONSE_FAIL"}`
	}
	return string(raw)
}

func NewRouter(service services.Service) Router {
	return router{
		handlers: handlers,
		service:  service,
	}
}

func Regist(service services.Service, w *astilectron.Window) {
	r := NewRouter(service)
	logs.InfoLogger.Printf("regist handlers\n")

	w.OnMessage(func(m *astilectron.EventMessage) (v interface{}) {
		return r.Route(m)
	})

	for route, _ := range handlers {
		logs.InfoLogger.Printf(route)
	}
}
