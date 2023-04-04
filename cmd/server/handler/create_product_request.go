package handlers

import "github.com/zabdielv/gin-exercises/internal/domain"

type CreateProductRequest struct {
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   int     `json:"code_value" `
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration" `
	Price        float64 `json:"price"`
}

func (request CreateProductRequest) ToDomain() domain.Producto {
	return domain.Producto{
		ID:           0,
		Name:         request.Name,
		Code_value:   request.Code_value,
		Is_published: request.Is_published,
		Expiration:   request.Expiration,
		Price:        request.Price,
	}
}
