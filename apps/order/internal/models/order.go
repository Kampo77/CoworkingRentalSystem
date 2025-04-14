package models

type Order struct {
	ID       uint           `json:"id"`
	UserID   uint           `json:"user_id"`
	Products []OrderProduct `json:"products"`
	Status   string         `json:"status"`
}

type OrderProduct struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
