package api

import (
	"baf/api/config"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func Run() {
	db, err := config.ConnectDB()
	if err !=nil {
		log.Fatal(err.Error())
	}

	r := config.CreateRouter()

	config.InitRouter(db, r).InitializeRoutes()
	config.Run(r)

}
