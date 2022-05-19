package models

import "errors"

type CreateWarehousePayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Longitude   float32 `json:"longitude"`
	Latitude    float32 `json:"latitude"`
}

func (i CreateWarehousePayload) Validate() error {
	if i.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
