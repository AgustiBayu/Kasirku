package controllers

import "net/http"

type AuthController interface {
	ShowLoginForm(w http.ResponseWriter, r *http.Request, params map[string]string)
	Login(w http.ResponseWriter, r *http.Request, params map[string]string)
}
