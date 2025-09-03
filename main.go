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
	db.AutoMigrate(&domain.ProductCategory{}, &domain.Product{}, &domain.Transaction{}, &domain.TransactionDetail{})
	validate := validator.New()

	// Repositories
	productCategoryRepo := repositories.NewProductCategoryRepository(db)
	productRepo := repositories.NewProductRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	// Services
	productCategoryService := services.NewProductCategoryService(productCategoryRepo, validate)
	productService := services.NewProductService(productRepo, productCategoryRepo, validate)
	transactionService := services.NewTransactionService(transactionRepo, productRepo, db, validate)

	// Controllers
	productCategoryController := controllers.NewProductCategoryController(productCategoryService)
	productController := controllers.NewProductController(productService, productCategoryService)
	transactionController := controllers.NewTransactionController(transactionService)

	router := routes.NewRouter(productCategoryController, productController, transactionController)
	router.ServeFiles("/images/*filepath", http.Dir("images"))
	println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
