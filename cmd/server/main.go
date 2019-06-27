package main

import (
	"flag"
	"log"

	"github.com/AndreyLM/go-vue-server/pkg/db"

	"github.com/AndreyLM/go-vue-server/pkg/system/app"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "Assigning the port")
	flag.Parse()
}

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	s := app.NewServer()

	s.Init(port, db)
	s.Start()
}
