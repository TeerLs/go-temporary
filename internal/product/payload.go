package product

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

type UpdateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}