package models

type UrlStruct struct {
	Id       int64  `json:"id"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}
