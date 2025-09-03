package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TransactionController interface {
	Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	ShowPOSTerminal(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
