package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID
	Name      string
	Value     float64
	Stock     int
	CreatedAt time.Time
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
