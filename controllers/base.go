package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	DB	*sql.DB
	Router *gin.Engine
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	// membuat instance / init router
	server.Router = gin.Default()
	// membuat routes
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8800")
	log.Fatal(server.Router.Run(addr))
}
