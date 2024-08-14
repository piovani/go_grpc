package grpc

import (
	"fmt"
)

type ProductDto struct {
	Name      string `protobuf:"bytes,1,opt,name=name,proto3"`
	Value     string `protobuf:"bytes,2,opt,name=value,proto3"`
	Stock     string `protobuf:"bytes,3,opt,name=stock,proto3"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,proto3"`
}

func Execute() {
	fmt.Println("AQUI")
}

func CreateProduct() {
	var productDto ProductDto

	// if err := proto.Ummarshal(data, &productDto); err != nil {
	// 	panic(err)
	// }

	// product := entity.NewProduct(productDto.Name, productDto.Value, productDto.Stock, productDto.CreatedAt)

	fmt.Println(productDto)

}
