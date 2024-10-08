package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/piovani/go_grpc/entity"
)

type ProductDto struct {
	Name      string    `json:"name"`
	Value     float64   `json:"value"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}

func Execute() {
	r := chi.NewRouter()

	r.Post("/product", CreateProduct())

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println(err)
	}
}

func CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var productDto ProductDto

		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		product := entity.NewProduct(productDto.Name, productDto.Value, productDto.Stock, productDto.CreatedAt)

		productDtoOut := ProductDto{
			Name:      product.Name,
			Value:     product.Value,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(productDtoOut)
	}
}
