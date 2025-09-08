package main

import (
	"kasirku/app"
	"kasirku/controllers"
	"kasirku/helpers"
	"kasirku/models/domain"
	"kasirku/repositories"
	"kasirku/routes"
	"kasirku/services"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	helpers.InitJWT()
	db := app.DB()
	db.AutoMigrate(&domain.ProductCategory{}, &domain.Product{}, &domain.Transaction{}, &domain.TransactionDetail{}, &domain.User{})
	helpers.SeedUsers(db)
	validate := validator.New()

	// Repositories
	productCategoryRepo := repositories.NewProductCategoryRepository(db)
	productRepo := repositories.NewProductRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Services
	productCategoryService := services.NewProductCategoryService(productCategoryRepo, validate)
	productService := services.NewProductService(productRepo, productCategoryRepo, validate)
	transactionService := services.NewTransactionService(transactionRepo, productRepo, db, validate)
	authService := services.NewAuthService(userRepo)

	// Controllers
	productCategoryController := controllers.NewProductCategoryController(productCategoryService)
	productController := controllers.NewProductController(productService, productCategoryService)
	transactionController := controllers.NewTransactionController(transactionService)
	authController := controllers.NewAuthController(authService)

	router := routes.NewRouter(productCategoryController, productController, transactionController, authController)
	println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
