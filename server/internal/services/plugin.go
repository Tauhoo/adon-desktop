package services

import "github.com/Tauhoo/adon-desktop/internal/messages"

func (s service) GetPluginList(_ messages.RequestMessage[any]) messages.ResponseMessage[[]string] {
	result := []string{}

	for _, record := range s.pluginManager.GetPluginStorage().GetList() {
		result = append(result, record.Name)
	}

	return messages.ResponseMessage[[]string]{
		Type: messages.SUCCESS,
		Data: result,
	}
}
