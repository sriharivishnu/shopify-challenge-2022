package models

type CreateWarehousePayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Longitude   float32 `json:"longitude"`
	Latitude    float32 `json:"latitude"`
}
