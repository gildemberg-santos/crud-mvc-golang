package config

import (
	"log"
	"net/http"

	db "github.com/gildemberg-santos/picpay-desafio-backend/model"
	"github.com/gildemberg-santos/picpay-desafio-backend/routes"
)

func Start() {
	initDB()
	initWebServer()
}

func initDB() {
	log.Println("Initializing database")
	db.NewDB(&db.User{})
}

func initWebServer() {
	log.Println("Initializing web server")
	log.Fatal(http.ListenAndServe(":8080", routes.NewRouter()))
}
