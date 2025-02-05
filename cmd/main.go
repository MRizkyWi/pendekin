package main

import (
	"log"
	"net/http"

	"pendekin/config"
	"pendekin/repository"
	"pendekin/route"
	"pendekin/service"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	urlRepo := repository.NewURLRepository(config.DB)
	urlService := service.NewUrlService(urlRepo)

	router := mux.NewRouter()
	router = route.InitializeRoutes(router, *urlService)

	log.Printf("Server running on %s \n", config.GetEnv("BASE_URL", ""))
	log.Fatal(http.ListenAndServe(":8000", router))
}
