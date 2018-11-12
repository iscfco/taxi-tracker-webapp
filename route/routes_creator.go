package route

import (
	"github.com/gorilla/mux"
	r "taxi-tracker-webapp/route/routes"
)

func CreateRoutes(router *mux.Router) {
	var routes []r.Route
	routes = append(routes, r.Customer()...)

	for _, route := range routes {
		/*if route.Protected {
			router.Handle(route.Path,
				negroni.New(
					negroni.HandlerFunc(middleware.ValidateToken),
					negroni.WrapFunc(route.Handler),
				),
			).Methods(route.Method)
			continue
		}
		*/router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
