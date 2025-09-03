package controllers

import (
	"context"
	"encoding/json"
	"kasirku/helpers"
	"kasirku/models/domain"
	"kasirku/models/web"
	"kasirku/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService         services.ProductService
	ProductCategoryService services.ProductCategoryService
}

func NewProductController(productService services.ProductService, productCategoryService services.ProductCategoryService) ProductController {
	return &ProductControllerImpl{
		ProductService:         productService,
		ProductCategoryService: productCategoryService,
	}
}

func (c *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		slug := r.FormValue("slug")
		priceStr := r.FormValue("price")
		exp := r.FormValue("exp")
		categoryIdStr := r.FormValue("category_id")
		barcode := r.FormValue("barcode")

		price, _ := strconv.Atoi(priceStr)
		categoryId, _ := strconv.Atoi(categoryIdStr)
		req := &domain.ProductCreateRequest{Name: name, Slug: slug,
			Price: uint(price), Exp: exp, CategoryID: uint(categoryId), Barcode: barcode}
		if err := c.ProductService.Create(context.Background(), req); err != nil {
			http.Error(w, "Gagal menyimpan data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/product", http.StatusSeeOther)
		return
	}
	categories, _ := c.ProductCategoryService.FindAll(context.Background())

	// Siapkan data untuk template
	data := map[string]interface{}{
		"Categories": categories,
	}
	helpers.RenderTemplate(w, "templates/product", "product_form_add.html", data)
}

func (c *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	products, _ := c.ProductService.FindAll(context.Background())

	helpers.RenderTemplate(w, "templates/product", "product_list.html", products)
}

func (c *ProductControllerImpl) FindAllJson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	products, err := c.ProductService.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    products,
	})
}

func (c *ProductControllerImpl) FindByBarcode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	barcode := ps.ByName("barcode")

	product, err := c.ProductService.FindByBarcode(context.Background(), barcode)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusNotFound,
			Message: "Product not found",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    product,
	})
}

func (c *ProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("productId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := c.ProductService.FindById(context.Background(), id)
	if err != nil {
		http.Error(w, "Failed to find product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := c.ProductCategoryService.FindAll(context.Background())
	if err != nil {
		http.Error(w, "Failed to find categories: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Product":    product,
		"Categories": categories,
	}
	helpers.RenderTemplate(w, "templates/product", "product_form_edit.html", data)
}

func (c *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == http.MethodPost {
		idStr := ps.ByName("productId")
		id, _ := strconv.Atoi(idStr)
		name := r.FormValue("name")
		slug := r.FormValue("slug")
		barcode := r.FormValue("barcode")
		priceStr := r.FormValue("price")
		exp := r.FormValue("exp")
		categoryIdStr := r.FormValue("category_id")

		price, _ := strconv.Atoi(priceStr)
		categoryId, _ := strconv.Atoi(categoryIdStr)

		req := &domain.ProductUpdateRequest{ID: uint(id), Name: name, Slug: slug, Barcode: barcode,
			Price: uint(price), Exp: exp, CategoryID: uint(categoryId)}

		if err := c.ProductService.Update(context.Background(), req); err != nil {
			http.Error(w, "Gagal memperbarui data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/product", http.StatusSeeOther)
		return
	}

	idStr := ps.ByName("productId")
	id, _ := strconv.Atoi(idStr)

	product, _ := c.ProductService.FindById(context.Background(), id)
	categories, _ := c.ProductCategoryService.FindAll(context.Background())
	data := map[string]interface{}{
		"Product":    product,
		"Categories": categories,
	}
	helpers.RenderTemplate(w, "templates/product", "product_form_edit.html", data)
}

func (c *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("productId")
	id, _ := strconv.Atoi(idStr)

	c.ProductService.Delete(context.Background(), id)

	http.Redirect(w, r, "/product", http.StatusSeeOther)
}

func (c *ProductControllerImpl) UploadThumbnail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("productId")
	id, _ := strconv.Atoi(idStr)

	file, handler, err := r.FormFile("thumbnail")
	if err != nil {
		http.Error(w, "Gagal mengunggah file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := c.ProductService.UploadThumbnail(context.Background(), uint(id), file, handler); err != nil {
		http.Error(w, "Gagal menyimpan thumbnail: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/product/edit/"+idStr, http.StatusSeeOther)
}
