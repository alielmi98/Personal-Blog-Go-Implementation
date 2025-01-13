package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/alielmi98/Personal-Blog-Go-Implementation/config"
	"github.com/alielmi98/Personal-Blog-Go-Implementation/handlers"
	middleware "github.com/alielmi98/Personal-Blog-Go-Implementation/middlewares"
	"github.com/alielmi98/Personal-Blog-Go-Implementation/services"
)

func main() {
	config.LoadConfig()

	dataFile := filepath.Join(".", "articles", "article.json")

	// Create an instance of the ArticleService
	articleService := services.NewArticleService(dataFile)
	err := articleService.LoadArticles()
	if err != nil {
		log.Fatal("Error loading articles:", err)
		return
	}

	// Create an instance of AdminHandler
	adminHandler := handlers.NewAdminHandler(articleService)
	guestHandler := handlers.NewGuestHandler(articleService)

	// Guest routers
	http.HandleFunc("/", http.HandlerFunc(guestHandler.HomeHandler))
	http.HandleFunc("/article/", http.HandlerFunc(guestHandler.ArticleHandler))

	// Admin routes with middleware
	http.Handle("/admin/dashboard", middleware.AuthMiddleware(http.HandlerFunc(adminHandler.DashboardHandler)))
	http.Handle("/admin/articles/create", middleware.AuthMiddleware(http.HandlerFunc(adminHandler.CreateArticleHandler)))
	http.Handle("/admin/articles/update/", middleware.AuthMiddleware(http.HandlerFunc(adminHandler.UpdateArticleHandler)))
	http.Handle("/admin/articles/delete/", middleware.AuthMiddleware(http.HandlerFunc(adminHandler.DeleteArticleHandler)))
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Starting server on :%s\n", config.AppConfig.Port)
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.Port, nil))
}
