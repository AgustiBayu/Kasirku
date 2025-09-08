package controllers

import (
	"html/template"
	"kasirku/models/domain"
	"kasirku/services"
	"net/http"
)

type AuthControllerImpl struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &AuthControllerImpl{authService: authService}
}

func (c *AuthControllerImpl) ShowLoginForm(w http.ResponseWriter, r *http.Request, params map[string]string) {
	tmpl, err := template.ParseFiles("templates/auth/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func (c *AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request, params map[string]string) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	request := domain.LoginRequest{
		Username: username,
		Password: password,
	}

	token, err := c.authService.Login(request)
	if err != nil {
		// TODO: show error message on login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
