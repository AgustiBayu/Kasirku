package main

import (
	"kasirku/app"
	"kasirku/controllers"
	"kasirku/models/domain"
	"kasirku/repositories"
	"kasirku/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Koneksi ke DB
	db := app.DB()

	// Auto migrate table
	db.AutoMigrate(&domain.ProductCategory{})

	// Init repository, service, controller
	productCategoryRepo := repositories.NewProductCategoryRepository(db)
	validate := validator.New()
	productCategoryService := services.NewProductCategoryService(productCategoryRepo, validate)
	productCategoryController := controllers.NewProductCategoryController(productCategoryService)

	// Router (pakai RESTful endpoint)
	router := httprouter.New()
	router.GET("/categories", productCategoryController.FindAll)        // List
	router.GET("/categories/:id", productCategoryController.FindById)   // Detail
	router.POST("/categories/create", productCategoryController.Create) // Create
	router.PUT("/categories/:id", productCategoryController.Update)     // Update
	router.DELETE("/categories/:id", productCategoryController.Delete)  // Delete

	println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
