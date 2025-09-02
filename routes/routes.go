package routes

import (
	"kasirku/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(productCategoryController controllers.ProductCategoryController, productController controllers.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.ServeFiles("/image/*filepath", http.Dir("image"))

	router.GET("/categories", productCategoryController.FindAll)
	router.GET("/categories/add", productCategoryController.Create)
	router.POST("/categories/add", productCategoryController.Create)
	router.GET("/categories/edit/:id", productCategoryController.FindById)
	router.POST("/categories/edit/:id", productCategoryController.Update)
	router.GET("/categories/delete/:id", productCategoryController.Delete)
	router.GET("/product", productController.FindAll)
	router.GET("/product/add", productController.Create)
	router.POST("/product/add", productController.Create)
	router.GET("/product/edit/:productId", productController.FindById)
	router.POST("/product/edit/:productId", productController.Update)
	router.GET("/product/delete/:productId", productController.Delete)
	router.POST("/product/upload-thumbnail/:productId", productController.UploadThumbnail)

	return router
}
