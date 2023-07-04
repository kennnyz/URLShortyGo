package delivery

import (
	"log"
	"net/http"
	"ozonTech/muhtarov/internal/service"
)

type Handler struct {
	// access to business logic
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/make-short-url", h.makeShortUrl)
	mux.HandleFunc("/get-long-url", h.getLongUrlByShort)

	return mux
}

func (h *Handler) makeShortUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		msg := "method not provide!"
		_, err := w.Write([]byte(msg))
		if err != nil {
			return
		}
	}

	longUrl := r.URL.Query().Get("url")

	ps, err := h.service.UrlShortyService.AddUrl(longUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	_, err = w.Write([]byte(ps.ShortUrl))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getLongUrlByShort(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		msg := "method not provide!"
		_, err := w.Write([]byte(msg))
		if err != nil {
			return
		}
	}

	shortUrl := r.URL.Query().Get("url")

	model, err := h.service.UrlShortyService.GetFullUrl(shortUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	_, err = w.Write([]byte(model.LongUrl))
}
