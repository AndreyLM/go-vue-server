package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"

	"github.com/AndreyLM/go-vue-server/pkg/system/router"
	"github.com/go-xorm/xorm"
)

// Server - app server
type Server struct {
	port string
	Db   *xorm.Engine
}

// NewServer - creates server
func NewServer() Server {
	return Server{}
}

// Init - init all vals
func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("Init server...")
	s.port = ":" + port
	s.Db = db
}

// Start - start server
func (s *Server) Start() {
	log.Println("Starting server on port " + s.port + "...")

	r := router.NewRouter()
	r.Init(s.Db)

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{"*"}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(newServer.ListenAndServe())
}
