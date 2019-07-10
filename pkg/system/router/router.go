package router

import (
	v1 "github.com/AndreyLM/go-vue-server/pkg/controllers/v1/router"
	"github.com/AndreyLM/go-vue-server/pkg/types/routes"
	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
)

// Router - router
type Router struct {
	Router *mux.Router
}

// NewRouter - new router
func NewRouter() (r Router) {
	r.Router = mux.NewRouter().StrictSlash(true)
	return
}

// Init - initialization
func (r *Router) Init(DB *xorm.Engine) {
	r.Router.Use(Middleware)

	baseRoutes := GetRoutes(DB)
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	SubRoutes := v1.GetRoutes(DB)
	for name, pack := range SubRoutes {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, v1.Middleware)
	}
}

// AttachSubRouterWithMiddleware - attaches subroutes to main router
func (r *Router) AttachSubRouterWithMiddleware(path string, subroutes routes.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {
	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)

	for _, route := range subroutes {

		SubRouter.Name(route.Name).Methods(route.Method).Path(route.Pattern).Handler(route.HandlerFunc)
	}

	return
}
