package config

import (
	"baf/api/app/controller"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"net/url"
)

type Server struct {
	DB	*sql.DB
	Router *gin.Engine
}

func ConnectDB() *sql.DB {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err := sql.Open(`mysql`, dsn)
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
		return nil
	}


	if err := db.Ping(); err != nil {
		log.Print(err)
		_, _ = fmt.Scanln()
		log.Fatal(err)
	}
	log.Println("DataBase Successfully Connected")
	return db
}

func (server *Server) Close() {
	_ = server.DB.Close()
}

func CreateRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}

func InitRouter(db *sql.DB, r *gin.Engine) *Server {
	return &Server{
		DB: db,
		Router: r,
	}
}

func (server *Server) InitializeRoutes()  {
	r := server.Router.Group("/v1")
	controller.CreateArticleController(server.DB, r)
	controller.CreateUserController(server.Router)
}

func Run(r *gin.Engine, addr string) {
	fmt.Println("Listening to port 8800")
	log.Fatal(r.Run(addr))
}
