package routes

import "net/http"

type Route struct {
	Method    string
	Path      string
	Handler   func(w http.ResponseWriter, r *http.Request)
	Protected bool
}
