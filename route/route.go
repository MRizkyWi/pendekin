package route

import (
	"pendekin/controller"
	"pendekin/service"

	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router, urlService service.UrlService) *mux.Router {
	urlController := controller.NewUrlController(urlService)

	// Url Routes
	router.HandleFunc("/urls", urlController.ShortenUrl).Methods("POST")
	router.HandleFunc("/urls/{shortUrl}", urlController.GetByShortUrl).Methods("GET")
	router.HandleFunc("/urls", urlController.UpdateStatus).Methods("PUT")

	return router
}
