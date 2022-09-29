package models

type Product struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
	Price     string `json:"price"`
	Quantity  uint8  `json:"quantity"`
}
