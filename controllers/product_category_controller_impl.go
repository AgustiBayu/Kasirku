package controllers

import (
	"context"
	"html/template"
	"kasirku/models/domain"
	"kasirku/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductCategoryControllerImpl struct {
	ProductCategoryService services.ProductCategoryService
}

func NewProductCategoryController(productCategoryService services.ProductCategoryService) ProductCategoryController {
	return &ProductCategoryControllerImpl{
		ProductCategoryService: productCategoryService,
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplParsed := template.Must(template.ParseFiles("templates/layout.html", tmpl))
	tmplParsed.ExecuteTemplate(w, "layout.html", data)
}

func (c *ProductCategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == http.MethodPost {
		category := r.FormValue("category")
		req := &domain.ProductCategoryCreateRequest{Category: category}
		_ = c.ProductCategoryService.Create(context.Background(), req)

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
		return
	}
	renderTemplate(w, "templates/category_form.html", nil)
}

func (c *ProductCategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	categories, _ := c.ProductCategoryService.FindAll(context.Background())
	data := map[string]interface{}{
		"Categories": categories,
	}
	renderTemplate(w, "templates/category_list.html", data)
}

func (c *ProductCategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	category, _ := c.ProductCategoryService.FindById(context.Background(), id)

	data := map[string]interface{}{
		"Category": category,
	}
	renderTemplate(w, "templates/category_form.html", data)
}

func (c *ProductCategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// mirip dengan Create, tapi update data
}

func (c *ProductCategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	_ = c.ProductCategoryService.Delete(context.Background(), id)

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
