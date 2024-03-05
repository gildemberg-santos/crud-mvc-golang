package config

import (
	"log"
	"net/http"

	db "github.com/gildemberg-santos/crud-mvc-golang/model"
	"github.com/gildemberg-santos/crud-mvc-golang/routes"
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
