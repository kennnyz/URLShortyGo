package delivery

import (
	"log"
	"net/http"
	"ozonTech/muhtarov/internal/models"
)

func (h *Handler) makeShortUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		msg := "method not provide!"
		_, err := w.Write([]byte(msg))
		if err != nil {
			return
		}
	}

	longUrl := r.URL.Query().Get("url")
	if longUrl == "" {
		http.Error(w, models.NotValidUrlErr.Error(), 400)
	}

	ps, err := h.service.UrlShortyService.AddUrl(longUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	_, err = w.Write([]byte(ps.ShortUrl))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
