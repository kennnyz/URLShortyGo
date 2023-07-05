package delivery

import (
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
