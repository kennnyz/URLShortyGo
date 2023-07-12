package models

import "fmt"

var (
	UrlNotFoundErr      = fmt.Errorf("URL not found")
	NotValidUrlErr      = fmt.Errorf("URL is not valid")
	MethodNotProvideErr = fmt.Errorf("method not provide")
)
