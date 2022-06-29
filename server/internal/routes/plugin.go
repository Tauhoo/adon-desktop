package routes

import "github.com/Tauhoo/adon-desktop/internal/services"

var AddNewPlugin = func(service services.Service, transaction Transaction) {
	req := transaction.GetRequest().Data.(services.PluginBuildInfo)
	if err := service.AddNewPlugin(req); err != nil {
		transaction.SetError(err)
	} else {
		transaction.SetResponse(nil)
	}
}
