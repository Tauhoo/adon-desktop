package routes

import (
	"encoding/json"

	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/messages"
	"github.com/Tauhoo/adon-desktop/internal/services"
	"github.com/asticode/go-astilectron"
)

type Transaction interface {
	GetRequest() messages.RequestMessage
	GetResponse() messages.ResponseMessage

	SetResponse(v any)
	SetError(v errors.Error)
}

type transaction struct {
	response messages.ResponseMessage
	request  messages.RequestMessage
}

func (t transaction) GetRequest() messages.RequestMessage {
	return t.request
}

func (t transaction) SetResponse(v any) {
	t.response = messages.NewResponseMessage(v)
}

func (t transaction) SetError(v errors.Error) {
	t.response = messages.NewResponseErrorMessage(v)
}

func (t transaction) GetResponse() messages.ResponseMessage {
	return t.response
}

func NewTansaction(m *astilectron.EventMessage) (Transaction, errors.Error) {
	var req messages.RequestMessage
	err := m.Unmarshal(&req)
	if err != nil {
		return nil, errors.New(UnmashalFailErrCode, err.Error())
	}
	return transaction{
		request: req,
		response: messages.NewResponseErrorMessage(
			errors.NewWithoutData(NoResponseErrCode),
		),
	}, nil
}

type Handler = func(service services.Service, transaction Transaction)

type Router interface {
	Route(tx Transaction) []byte
}

type router struct {
	handlers map[string]Handler
	service  services.Service
}

func (r router) Route(tx Transaction) []byte {
	msg := tx.GetRequest()
	handler, ok := handlers[msg.Route]
	if !ok {
		return []byte(`{"code":"ROUTE_NOT_FOUND"}`)
	}

	handler(r.service, tx)
	raw, err := json.Marshal(tx.GetResponse())
	if err != nil {
		return []byte(`{"code":"MASHAL_REPONSE_FAIL"}`)
	}
	return raw
}

func NewRouter(service services.Service) Router {
	return router{
		handlers: handlers,
		service:  service,
	}
}

func Regist(service services.Service, w *astilectron.Window) {
	r := NewRouter(service)
	w.OnMessage(func(m *astilectron.EventMessage) (v interface{}) {
		tx, err := NewTansaction(m)
		if err != nil {
			return `{"code":"CREATE_TRANSACTION_FAIL"}`
		}
		return r.Route(tx)
	})
}
