package product

import (
	"net/http"
	"strconv"
	"temporary/pkg/db"
	"temporary/pkg/res"
	"temporary/req"
)

type ProductHandler struct {
	repo *ProductRepo
}

type ProductHandlerDeps struct {
	DB *db.DB
}

func NewProductHandler(router *http.ServeMux, deps *ProductHandlerDeps) *ProductHandler {
	repo := NewProductRepo(deps.DB.DB)
	handler := &ProductHandler{repo: repo}

	router.HandleFunc("/products", handler.GetAll())
	router.HandleFunc("/products/{id}", handler.GetByID())
	router.HandleFunc("/products/create", handler.Create())
	router.HandleFunc("/products/update/{id}", handler.Update())
	router.HandleFunc("/products/delete/{id}", handler.Delete())

	return handler
}

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CreateProductRequest](&w, r)
		if err != nil {
			return
		}

		product := &Product{
			Name:        body.Name,
			Description: body.Description,
		}		

		result, err := h.repo.Create(product)
		if err != nil {
			res.WriteJSON(w, map[string]string{
				"status": "error",
				"error":  err.Error(),
			}, http.StatusInternalServerError)
			return
		}

		res.WriteJSON(w, result, http.StatusOK)
	}
}

func (h *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			res.WriteJSON(w, map[string]string{
				"status": "error",
				"error":  err.Error(),
			}, http.StatusBadRequest)
			return
		}

		product, err := h.repo.GetByID(uint(id))
		if err != nil {
			res.WriteJSON(w, map[string]string{
				"status": "error",
				"error":  err.Error(),
			}, http.StatusInternalServerError)
			return
		}
		res.WriteJSON(w, product, http.StatusOK)
	}
}

func (h *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.repo.GetAll()
		if err != nil {
			res.WriteJSON(w, map[string]string{
				"status": "error",
				"error":  err.Error(),
			}, http.StatusInternalServerError)
			return
		}
		res.WriteJSON(w, products, http.StatusOK)
	}
}

func (h *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			res.WriteJSON(w, map[string]string{
				"status": "error",
				"error":  err.Error(),
			}, http.StatusBadRequest)
			return
		}
		body, err := req.HandleBody[UpdateProductRequest](&w, r)
		if err != nil {
			return
		}

		product := &Product{
			Name:        body.Name,
			Description: body.Description,
		}

		err = h.repo.Update(product, uint(id))
		if err != nil {
			res.WriteJSON(w, map[string]string{
				"status": "error",
				"error":  err.Error(),
			}, http.StatusInternalServerError)
			return
		}

		res.WriteJSON(w, product, http.StatusOK)
	}
}

func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			res.WriteJSON(w, map[string]string{
				"status": "error",
				"error":  err.Error(),
			}, http.StatusBadRequest)
			return
		}

		err = h.repo.Delete(uint(id))
		if err != nil {
			res.WriteJSON(w, map[string]string{
				"status": "error",
				"error":  err.Error(),
			}, http.StatusInternalServerError)
			return
		}

		res.WriteJSON(w, map[string]string{
			"status": "success",
		}, http.StatusOK)
	}
}