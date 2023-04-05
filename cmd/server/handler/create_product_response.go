package handlers

import "github.com/zabdielv/gin-exercises/internal/domain"

type CreateProductResponse struct {
	Estado string          `json:"estado"`
	Datos  domain.Producto `json:"producto"`
}

type CreateProductsResponse struct {
	Estado string            `json:"estado"`
	Datos  []domain.Producto `json:"productos"`
}

type CreateListResponse struct {
	Productos []domain.Producto `json:"productos"`
	Total     float64           `json:"total_price"`
}
