package models

import "fmt"

var (
	UrlNotFoundErr = fmt.Errorf("URL not found")
	NotValidUrlErr = fmt.Errorf("URL is not valid")
)
