package router

import (
	"net/http"

	"github.com/AndreyLM/go-vue-server/pkg/controllers/home"
	"github.com/AndreyLM/go-vue-server/pkg/types/routes"
)

// Middleware - middleware
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes - get routes
func GetRoutes() routes.Routes {
	return routes.Routes{
		routes.Route{Name: "Home", Method: "GET", Pattern: "/", HandlerFunc: home.Index},
	}
}
