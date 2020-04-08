package models

type Article struct {
	ID          string `json:id`
	Title       string `json:title`
	Description string `json:description`
	Content     string `json:content`
}
