package http

import (
	"log"
	"net/http"
	"ozonTech/muhtarov/internal/models"
)

func (h *Handler) getLongUrlByShort(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, models.MethodNotProvideErr.Error(), 400)
		return
	}

	shortUrl := r.URL.Query().Get("url")
	if shortUrl == "" {
		http.Error(w, models.NotValidUrlErr.Error(), 400)
		return
	}

	model, err := h.urlShorty.GetFullUrl(shortUrl)
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Println(err)
		return
	}

	_, err = w.Write([]byte(model.LongUrl))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
