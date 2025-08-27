package main

import (
	"kasirku/app"
	"kasirku/controllers"
	"kasirku/models/domain"
	"kasirku/repositories"
	"kasirku/routes"
	"kasirku/services"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.DB()
	db.AutoMigrate(&domain.ProductCategory{})

	productCategoryRepo := repositories.NewProductCategoryRepository(db)
	validate := validator.New()
	productCategoryService := services.NewProductCategoryService(productCategoryRepo, validate)
	productCategoryController := controllers.NewProductCategoryController(productCategoryService)
	router := routes.NewRouter(productCategoryController)

	println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
