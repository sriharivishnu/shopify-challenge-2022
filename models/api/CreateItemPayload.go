package models

type CreateItemPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
	City        string  `json:"city"`
}
