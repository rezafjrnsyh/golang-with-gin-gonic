package config

import (
	"baf/api/article"
	"baf/api/user"
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

func ConnectDB(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*sql.DB, error) {
	db, err := sql.Open(DbDriver, DbUser+":"+DbPassword+"@tcp("+DbHost+":"+DbPort+")/"+DbName)
	if err != nil {
		log.Fatal(err)
	}

	query := `CREATE TABLE IF NOT EXISTS article(id int primary key auto_increment, title text, content text)`
	//ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	//
	//defer cancelfunc()
	_, err = db.Exec(query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return nil, err
	}


	if err := db.Ping(); err != nil {
		log.Print(err)
		_, _ = fmt.Scanln()
		log.Fatal(err)
	}
	log.Println("DataBase Successfully Connected")
	return db, nil
}

func (server *Server) Close() {
	_ = server.DB.Close()
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
	article.CreateArticleController(server.DB, server.Router)
	user.CreateUserController(server.Router)
}

func Run(r *gin.Engine, addr string) {
	fmt.Println("Listening to port 8800")
	log.Fatal(r.Run(addr))
}
