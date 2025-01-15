package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alielmi98/Personal-Blog-Go-Implementation/dto"
	"github.com/alielmi98/Personal-Blog-Go-Implementation/services"
)

type GuestHandler struct {
	service *services.ArticleService
}

// NewAdminHandler creates a new AdminHandler
func NewGuestHandler(service *services.ArticleService) *GuestHandler {
	return &GuestHandler{service: service}
}

func (h *GuestHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := h.service.GetAllArticles()
	if err != nil {
		http.Error(w, "Could not load articles", http.StatusInternalServerError)
		return
	}

	var truncatedArticles []dto.ArticleDTO

	for _, article := range articles {
		article.Content = truncate(article.Content, 100)
		truncatedArticles = append(truncatedArticles, article)
	}

	RenderTemplate(w, "templates/home.tmpl", truncatedArticles)
}

func (h *GuestHandler) ArticleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/article/"):])
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	article, isExist := h.service.GetArticle(id)

	if !isExist {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	RenderTemplate(w, "templates/article.tmpl", article)
}

func truncate(content string, length int) string {
	if len(content) > length {
		return content[:length] + "..."
	}
	return content
}
