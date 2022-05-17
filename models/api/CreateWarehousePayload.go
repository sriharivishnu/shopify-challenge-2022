package models

type CreateWarehousePayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Lon         float32 `json:"lon"`
	Lat         float32 `json:"lat"`
}
