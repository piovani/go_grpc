package grpc

import (
	"context"
	"net"

	"github.com/piovani/go_grpc/ui/grpc/pb"
	grpc "google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedProductServer
}

func Execute() {
	listener, err := net.Listen("tcp", "localhost:7008")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *Server) CreateProduct(ctx context.Context, product *pb.ProductRequest) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{
		Name:      "test-name2",
		Value:     1.23,
		Stock:     10,
		CreatedAt: "2024",
	}, nil
	// var productDto ProductDto

	// if err := proto.Unmarshal(data, &productDto); err != nil {
	// 	// imprementar error
	// }

	// // if err := proto.Ummarshal(data, &productDto); err != nil {
	// // 	panic(err)
	// // }

	// // product := entity.NewProduct(productDto.Name, productDto.Value, productDto.Stock, productDto.CreatedAt)

	// fmt.Println(productDto)

}
