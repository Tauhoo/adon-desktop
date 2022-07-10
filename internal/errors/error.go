package errors

import "encoding/json"

type Code string

type Error interface {
	GetCode() Code
	GetData() any
	json.Marshaler
}

type adonError struct {
	Code Code `json:"code"`
	Data any  `json:"data"`
}

func (e adonError) MarshalJSON() ([]byte, error) {
	return json.Marshal(e)
}

func (e adonError) GetCode() Code {
	return e.Code
}

func (e adonError) GetData() any {
	return e.Data
}

func New(code Code, data any) Error {
	return adonError{
		Code: code,
		Data: data,
	}
}

func NewWithoutData(code Code) Error {
	return adonError{
		Code: code,
		Data: nil,
	}
}
