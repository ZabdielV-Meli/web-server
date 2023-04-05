package handlers

type CreateProductResponse struct {
	Estado string `json:"estado"`
	Datos  any    `json:"datos"`
}

type CreateListResponse struct {
	Productos any     `json:"productos"`
	Total     float64 `json:"total_price"`
}
