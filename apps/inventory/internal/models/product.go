package models

type Product struct {
	ID          string  `json:"id" bson:"_id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"category_id"`
	Price       float64 `json:"price"`
	Stock       uint    `json:"stock"`
}
