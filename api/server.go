package api

import (
	"baf/api/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	// digunakan untuk mengatur bahwa config.json akan menjadi file config
	viper.SetConfigFile(`config.json`)
	// membaca isi config
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func Run() {
	// fungsi untuk koneksi ke database
	db, err := config.ConnectDB()

	// apabila ada error program langsung berhenti
	if err !=nil {
		log.Fatal(err.Error())
	}

	// menginitialisasi router / membuat router
	r := config.CreateRouter()

	// menggunakan router dan menginitialisasi routes
	config.InitRouter(db, r).InitializeRoutes()

	// menjalankan program
	errun := config.Run(r)
	if errun != nil {
		log.Fatal(errun.Error())
	}
}
