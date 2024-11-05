package models

type Company struct {
	Name     string  `json:"name"`
	Industry string  `json:"industry"`
	Budget   float64 `json:"budget"`
}
