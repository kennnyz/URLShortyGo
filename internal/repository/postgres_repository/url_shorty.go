package postgres_repository

import (
	"database/sql"
	"log"
	"ozonTech/muhtarov/internal/models"
)

type UrlRepoDB struct {
	db *sql.DB
}

func NewEmailRepo(db *sql.DB) *UrlRepoDB {
	return &UrlRepoDB{db: db}
}

func (e *UrlRepoDB) AddUrl(urlStruct models.UrlStruct) (models.UrlStruct, error) {

	// check if url exists
	checkQuery := "SELECT id, short_url, long_url FROM urls WHERE id = $1"
	err := e.db.QueryRow(checkQuery, urlStruct.Id).Scan(&urlStruct.Id, &urlStruct.ShortUrl, &urlStruct.LongUrl)
	if err == nil {
		return urlStruct, nil
	}

	insertQuery := "INSERT INTO urls (id, short_url, long_url) VALUES ($1, $2, $3) RETURNING id, short_url, long_url"
	err = e.db.QueryRow(insertQuery, urlStruct.Id, urlStruct.ShortUrl, urlStruct.LongUrl).Scan(&urlStruct.Id, &urlStruct.ShortUrl, &urlStruct.LongUrl)
	if err != nil {
		log.Println(urlStruct)
		return models.UrlStruct{}, err
	}

	return urlStruct, nil
}

func (e *UrlRepoDB) GetFullUrlByShort(shortUrl string) (models.UrlStruct, error) {
	var model models.UrlStruct
	checkQuery := "SELECT long_url FROM urls WHERE short_url = $1"

	err := e.db.QueryRow(checkQuery, shortUrl).Scan(&model.LongUrl)
	// if no rows in result return
	if err == sql.ErrNoRows {
		return models.UrlStruct{}, models.UrlNotFoundErr
	} else if err != nil {
		return models.UrlStruct{}, err
	}

	return model, nil
}
