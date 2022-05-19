package models

import "errors"

type CreateItemPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
	City        string  `json:"city"`
}

func (i CreateItemPayload) Validate() error {
	if i.Name == "" {
		return errors.New("name is required")
	}
	if i.City == "" {
		return errors.New("city is required")
	}

	return nil
}
