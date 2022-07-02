package api

import (
	"encoding/json"

	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/asticode/go-astilectron"
)

type API interface {
	PluginAdded(pluginName string) errors.Error
}

type api struct {
	window *astilectron.Window
}

func (a api) send(message interface{}, callbacks ...astilectron.CallbackMessage) errors.Error {
	if result, err := json.Marshal(message); err != nil {
		return errors.NewWithoutData(MarshalRequestFailCode)
	} else {
		a.window.SendMessage(string(result), callbacks...)
		return nil
	}
}

func New(window *astilectron.Window) API {
	return api{window: window}
}
