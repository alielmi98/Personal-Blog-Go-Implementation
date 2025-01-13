package handlers

import (
	"html/template"
	"net/http"

	"github.com/alielmi98/Personal-Blog-Go-Implementation/config"
	"github.com/alielmi98/Personal-Blog-Go-Implementation/dto"
)

// RenderTemplate renders HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	blogTitle := config.AppConfig.BlogTitle

	templateData := dto.TemplateData{
		BlogTitle: blogTitle,
		Data:      data,
	}

	t := template.Must(template.ParseFiles("templates/base.html", tmpl))
	if err := t.Execute(w, templateData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
