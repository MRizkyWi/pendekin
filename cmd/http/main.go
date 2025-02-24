package main

import (
	"log"
	"net/http"

	"pendekin/cache"
	"pendekin/config"
	"pendekin/repository"
	"pendekin/route"
	"pendekin/service"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	cache.ConnectRedis()

	urlRepo := repository.NewURLRepository(config.DB)
	urlService := service.NewUrlService(urlRepo, cache.GetRedisClient())

	router := mux.NewRouter()
	router = route.InitializeRoutes(router, *urlService)

	log.Printf("Server running on %s \n", config.GetEnv("BASE_URL", ""))
	log.Fatal(http.ListenAndServe(":8000", router))
}
