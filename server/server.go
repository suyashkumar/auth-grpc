package server

import "context"
import pb "github.com/suyashkumar/auth-grpc/protos"

type server struct{}

func (s *server) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}

func (s *server) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (s *server) Validate(ctx context.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return nil, nil
}

func NewServer() pb.AuthServer {
	return &server{}
}