package router

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"

	"github.com/AndreyLM/go-vue-server/pkg/controllers/v1/status"
	"github.com/AndreyLM/go-vue-server/pkg/types/routes"
)

var db *xorm.Engine

// Middleware - subrouter middleware
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		log.Println("Inside v1 middleware")
		log.Println(token)
		next.ServeHTTP(w, r)
	})
}

// GetRoutes - gets subroutes
func GetRoutes(DB *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {
	db = DB
	status.Init(db)
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": routes.SubRoutePackage{
			Routes: routes.Routes{
				routes.Route{Name: "Status", Method: "GET", Pattern: "/status", HandlerFunc: status.Index},
			},
			Middleware: Middleware,
		},
	}

	return
}
