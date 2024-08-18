package grpc

import (
	"context"
	"net"
	"time"

	"github.com/piovani/go_grpc/entity"
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

func (s *Server) CreateProduct(ctx context.Context, productRequest *pb.ProductRequest) (*pb.ProductResponse, error) {
	created_at, err := time.Parse("2006-01-02 15:04:05", productRequest.CreatedAt)
	if err != nil {
		panic(err)
	}

	prooduct := entity.NewProduct(
		productRequest.Name,
		float64(productRequest.Value),
		int(productRequest.Stock),
		created_at,
	)

	return &pb.ProductResponse{
		Name:      prooduct.Name,
		Value:     float32(prooduct.Value),
		Stock:     int32(prooduct.Stock),
		CreatedAt: prooduct.CreatedAt.String(),
	}, nil
}
