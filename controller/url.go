package controller

import (
	"encoding/json"
	"net/http"
	"pendekin/model"
	"pendekin/service"

	"github.com/gorilla/mux"
)

type UrlController struct {
	urlService service.UrlService
}

func NewUrlController(urlService service.UrlService) *UrlController {
	return &UrlController{
		urlService: urlService,
	}
}

func (u *UrlController) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var req model.NewUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	url, err := u.urlService.Save(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}

func (u *UrlController) GetByShortUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortUrl := params["shortUrl"]
	url, err := u.urlService.GetByShortUrl(shortUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.ActualUrl, http.StatusFound)
}

func (u *UrlController) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var req model.UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := u.urlService.UpdateStatus(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
