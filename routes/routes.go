package routes

import (
	"kasirku/controllers"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(productCategoryController controllers.ProductCategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/categories", productCategoryController.FindAll)
	router.GET("/categories/create", productCategoryController.Create)
	router.POST("/categories/create", productCategoryController.Create)
	router.GET("/categories/update/:id", productCategoryController.FindById)
	router.POST("/categories/update/:id", productCategoryController.Update)
	router.GET("/categories/delete/:id", productCategoryController.Delete)

	return router
}