package domain

type Producto struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   int     `json:"code_value" `
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration" `
	Price        float64 `json:"price"`
}
