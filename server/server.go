package server

import "context"
import pb "github.com/suyashkumar/auth-grpc/protos"
import (
	"github.com/suyashkumar/auth"
	"github.com/suyashkumar/auth-grpc/config"
)

type server struct {
	auth auth.Auth
}

func (s *server) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}

func (s *server) GetToken(ctx context.Context, r *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	return nil, nil
}

func (s *server) Validate(ctx context.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return nil, nil
}

func NewServer() (pb.AuthServer, error) {
	a, err := auth.NewAuthenticator(
		config.Get(config.DBConnString),
		[]byte(config.Get(config.SigningKey)),
	)

	if err != nil {
		return nil, err
	}
	return &server{a}, nil
}
