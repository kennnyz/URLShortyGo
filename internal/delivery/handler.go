package delivery

import (
	"net/http"
	"ozonTech/muhtarov/internal/models"
)

type URLShorty interface {
	AddUrl(mail string) (models.UrlStruct, error)
	GetFullUrl(keyword string) (models.UrlStruct, error) // check if user exists
}

type Handler struct {
	// access to business logic
	urlShorty URLShorty
}

func NewHandler(urlShortyService URLShorty) *Handler {
	return &Handler{
		urlShorty: urlShortyService,
	}
}

func (h *Handler) Init() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/make-short-url", h.makeShortUrl)
	mux.HandleFunc("/get-long-url", h.getLongUrlByShort)

	return mux
}
