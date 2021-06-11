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
	r := config.CreateRouter()

	db := config.ConnectDB()

	config.InitRouter(db, r).InitializeRoutes()
	config.Run(r)

}
