package handlers

import (
	"net/http"

	"github.com/alielmi98/Personal-Blog-Go-Implementation/config"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == config.AppConfig.AdminUser && password == config.AppConfig.AdminPass {
			http.SetCookie(w, &http.Cookie{
				Name:  "session",
				Value: "valid",
			})
			http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
			return
		}

		// Render login page with error message
		RenderTemplate(w, "templates/login.tmpl", map[string]interface{}{
			"Error": "Invalid username or password"})

		return
	}
	RenderTemplate(w, "templates/login.tmpl", nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Clear the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	// Redirect to the login page
	http.Redirect(w, r, "/login", http.StatusFound)
}
