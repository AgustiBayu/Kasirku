package controllers

import (
	"encoding/json"
	"kasirku/helpers"
	"kasirku/models/domain"
	"kasirku/models/web"
	"kasirku/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TransactionControllerImpl struct {
	TransactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (c *TransactionControllerImpl) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request *domain.TransactionCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := c.TransactionService.Create(r.Context(), request)
	if err != nil {
		// A more sophisticated error handling could be implemented here
		// to return different status codes based on the error type.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(web.WebResponse{
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    response,
	})
}

func (c *TransactionControllerImpl) ShowPOSTerminal(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	helpers.RenderTemplate(w, "templates/transaction", "pos.html", nil)
}
