package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name"`
	Value     float64   `json:"value"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, value float64, stock int, createdAt time.Time) *Product {
	id, _ := uuid.NewUUID()
	return &Product{
		ID:        id,
		Name:      name,
		Value:     value,
		Stock:     stock,
		CreatedAt: createdAt,
	}
}
