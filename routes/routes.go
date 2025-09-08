package routes

import (
	"kasirku/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Middleware wrapper
func protected(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// This is a simplified middleware check. 
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		// In a real app, you would validate the token here.
		h(w, r, ps)
	}
}

func NewRouter(productCategoryController controllers.ProductCategoryController, productController controllers.ProductController, transactionController controllers.TransactionController, authController controllers.AuthController) *httprouter.Router {
	router := httprouter.New()

	router.ServeFiles("/images/*filepath", http.Dir("images"))

	// Auth routes
	router.GET("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authController.ShowLoginForm(w, r, nil)
	})
	router.POST("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authController.Login(w, r, nil)
	})

	// API Routes
	router.GET("/api/products", protected(productController.FindAllJson))
	router.GET("/api/products/low-stock", protected(productController.FindLowStock))
	router.GET("/api/products/barcode/:barcode", protected(productController.FindByBarcode))

	// Transaction Route
	router.GET("/pos", protected(transactionController.ShowPOSTerminal))
	router.POST("/transactions", protected(transactionController.Create))

	// Product Category Routes
	router.GET("/categories", protected(productCategoryController.FindAll))
	router.GET("/categories/add", protected(productCategoryController.Create))
	router.POST("/categories/add", protected(productCategoryController.Create))
	router.GET("/categories/edit/:id", protected(productCategoryController.FindById))
	router.POST("/categories/edit/:id", protected(productCategoryController.Update))
	router.GET("/categories/delete/:id", protected(productCategoryController.Delete))

	// Product Routes
	router.GET("/product", protected(productController.FindAll))
	router.GET("/product/add", protected(productController.Create))
	router.POST("/product/add", protected(productController.Create))
	router.GET("/product/edit/:productId", protected(productController.FindById))
	router.POST("/product/edit/:productId", protected(productController.Update))
	router.GET("/product/delete/:productId", protected(productController.Delete))
	router.POST("/product/upload-thumbnail/:productId", protected(productController.UploadThumbnail))

	return router
}
