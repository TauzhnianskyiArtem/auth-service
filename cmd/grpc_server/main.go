package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"

	desc "github.com/TauzhnianskyiArtem/auth-service/pkg/user_v1"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 9090

type server struct {
	desc.UnimplementedUserV1Server
}

// Get ...
func (s *server) Get(ctx context.Context, req *desc.UserGetRequest) (*desc.UserGetResponse, error) {
	log.Printf("User id: %d", req.GetId())

	return &desc.UserGetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      1,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
