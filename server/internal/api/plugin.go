package api

import (
	"github.com/Tauhoo/adon"
	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/messages"
)

func (a api) PluginAdded(pluginName string) errors.Error {
	return a.send(messages.NewRequestMessage("route/plugin-added", pluginName))
}

func (a api) PluginDeleted(pluginName string) errors.Error {
	return a.send(messages.NewRequestMessage("route/plugin-deleted", pluginName))
}

type ExecutionStateChangeEvent struct {
	State string `json:"state"`
	Info  any    `json:"info"`
}

func (a api) ExecutionStateChange(pluginName, function string, state adon.ExecuteState, info any) errors.Error {
	return a.send(messages.NewRequestMessage("route/execute-state-change", ExecutionStateChangeEvent{
		State: state.String(),
		Info:  info,
	}))
}
