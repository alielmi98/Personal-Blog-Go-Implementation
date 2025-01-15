package handlers

import (
	"fmt"
	"net/http"

	"github.com/alielmi98/Personal-Blog-Go-Implementation/dto"
	"github.com/alielmi98/Personal-Blog-Go-Implementation/services"
)

type AdminHandler struct {
	service *services.ArticleService
}

// NewAdminHandler creates a new AdminHandler
func NewAdminHandler(service *services.ArticleService) *AdminHandler {
	return &AdminHandler{service: service}
}

// DashboardHandler handles the dashboard view
func (h *AdminHandler) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := h.service.GetAllArticles()
	if err != nil {
		http.Error(w, "Failed to load articles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the dashboard with loaded articles
	RenderTemplate(w, "templates/dashboard.tmpl", articles)
}

// CreateArticleHandler handles the creation of new articles
func (h *AdminHandler) CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		article := dto.ArticleDTO{
			Title:   r.FormValue("title"),
			Content: r.FormValue("content"),
		}
		if err := h.service.CreateArticle(article); err != nil {
			http.Error(w, "Failed to save article: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
		return
	}

	RenderTemplate(w, "templates/create_article.tmpl", nil)
}

// UpdateArticleHandler handles editing existing articles
func (h *AdminHandler) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/admin/articles/update/"):]
	var id int
	_, err := fmt.Sscanf(idStr, "%d", &id) // Assuming ID is an integer
	if err != nil {
		http.NotFound(w, r)
		return
	}

	article, exists := h.service.GetArticle(id)
	if !exists {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodPost {
		updatedArticle := dto.ArticleDTO{
			ID:      id,
			Title:   r.FormValue("title"),
			Content: r.FormValue("content"),
		}

		if err := h.service.UpdateArticle(updatedArticle); err != nil {
			http.Error(w, "Failed to update article: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
		return
	}

	RenderTemplate(w, "templates/update_article.tmpl", article)
}

// DeleteArticleHandler handles the deletion of articles
func (h *AdminHandler) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/admin/articles/delete/"):]

	var id int
	_, err := fmt.Sscanf(idStr, "%d", &id) // Assuming ID is an integer
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = h.service.DeleteArticle(id)
	if err != nil {
		http.Error(w, "Failed to delete article: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
}
