package routes

import (
	"taxi-tracker-webapp/controller"
)
func Customer() []Route {
	routes := []Route{
		{
			Method:  "GET",
			Path:    "/",
			Handler: controller.Index,
		},
		{
			Method:  "GET",
			Path:    "",
			Handler: controller.Index,
		},
	}

	return routes
}
