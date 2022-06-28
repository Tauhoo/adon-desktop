package routes

import "github.com/Tauhoo/adon-desktop/internal/services"

func getRouteHandlerList(service services.Service) []RouteHandler {
	return []RouteHandler{
		routeHandler[any, []string]{
			route:   "get_plugin_list",
			handler: service.GetPluginList,
		},
	}
}
