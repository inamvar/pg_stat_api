package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"pg_stats/database"
	"pg_stats/handlers"
)

type ApiServer struct {
	App *fiber.App
}


func (s *ApiServer) init(){

	_,err := database.CreateConnection()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	s.App = fiber.New()
	s.App.Get("/", handlers.StatsHandler)
}

func (s *ApiServer) Start() {
    s.init()
	s.App.Listen(":3000")
}
func (s *ApiServer) ShutDown() {
	s.App.Shutdown()
	database.Db.Close()
}
