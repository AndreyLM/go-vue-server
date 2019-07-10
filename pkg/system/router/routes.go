package router

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"

	"github.com/AndreyLM/go-vue-server/pkg/controllers/home"
	"github.com/AndreyLM/go-vue-server/pkg/types/routes"
)

// Middleware - middleware
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Global middleware...")
		next.ServeHTTP(w, r)
	})
}

// GetRoutes - get routes
func GetRoutes(DB *xorm.Engine) routes.Routes {
	home.Init(DB)

	return routes.Routes{
		routes.Route{Name: "Home", Method: "GET", Pattern: "/", HandlerFunc: home.Index},
	}
}
