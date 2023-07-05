package delivery

import (
	"log"
	"net/http"
	"ozonTech/muhtarov/internal/models"
)

func (h *Handler) getLongUrlByShort(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		msg := "method not provide!"
		_, err := w.Write([]byte(msg))
		if err != nil {
			return
		}
	}

	shortUrl := r.URL.Query().Get("url")
	if shortUrl == "" {
		http.Error(w, models.NotValidUrlErr.Error(), 400)
	}

	model, err := h.service.UrlShortyService.GetFullUrl(shortUrl)
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
