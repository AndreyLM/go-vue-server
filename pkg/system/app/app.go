package app

import (
	"log"
	"net/http"

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
	log.Println("Starting server...")

	r := router.NewRouter()
	r.Init()
	http.ListenAndServe(s.port, r.Router)
}
