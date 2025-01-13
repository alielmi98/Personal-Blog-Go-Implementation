package dto

import "time"

type ArticleDTO struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type TemplateData struct {
	BlogTitle string
	Data      interface{}
}
