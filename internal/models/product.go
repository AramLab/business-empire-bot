package models

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	CostPrice   float64 `json:"cost_price"`
	Category    string  `json:"category"`
	Rating      float64 `json:"rating" default:"5.0"`
}
