package router

import (
	"github.com/gorilla/mux"
)

// Router - router
type Router struct {
	Router *mux.Router
}

// Init - initialization
func (r *Router) Init() {
	r.Router.Use(Middleware)
	baseRoutes := GetRoutes()
	for _, route := range baseRoutes {
		r.Router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
}

// NewRouter - new router
func NewRouter() (r Router) {
	r.Router = mux.NewRouter().StrictSlash(true)
	return
}
