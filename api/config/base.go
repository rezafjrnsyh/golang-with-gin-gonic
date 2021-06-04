package config

import (
	"baf/api/article"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Server struct {
	DB	*sql.DB
	Router *gin.Engine
}

func ConnectDB(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (db *sql.DB) {
	db, err := sql.Open(DbDriver, DbUser+":"+DbPassword+"@tcp("+DbHost+":"+DbPort+")/"+DbName)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Print(err)
		_, _ = fmt.Scanln()
		log.Fatal(err)
	}
	log.Println("DataBase Successfully Connected")
	return db
}

func CreateRouter() *gin.Engine {
	return gin.Default()
}

func InitRouter(db *sql.DB, r *gin.Engine) *Server {
	return &Server{
		DB: db,
		Router: r,
	}
}

func (server *Server) InitializeRoutes()  {
	article.InitializeRoutesArticle(server.DB, server.Router)
}

func Run(r *gin.Engine, addr string) {
	fmt.Println("Listening to port 8800")
	log.Fatal(r.Run(addr))
}
