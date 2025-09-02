package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	Create(W http.ResponseWriter, r *http.Request, ps httprouter.Params)
	FindAll(W http.ResponseWriter, r *http.Request, ps httprouter.Params)
	FindById(W http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Update(W http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Delete(W http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UploadThumbnail(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
