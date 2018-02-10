package server

import "context"
import pb "github.com/suyashkumar/auth-grpc/protos"
import (
	"github.com/sirupsen/logrus"
	"github.com/suyashkumar/auth"
	"github.com/suyashkumar/auth-grpc/config"
	"github.com/suyashkumar/auth-grpc/mappers"
)

type server struct {
	auth auth.Auth
}

func (s *server) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	u := auth.User{
		Email:              r.Email,
		FirstName:          r.FirstName,
		LastName:           r.LastName,
		MaxPermissionLevel: auth.PERMISSIONS_USER, //TODO: more robust permissioning
	}
	err := s.auth.Register(&u, r.Password)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{}, nil
}

func (s *server) GetToken(ctx context.Context, r *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	//TODO: add validation layer
	token, err := s.auth.GetToken(r.Email, r.Password, auth.Permissions(r.RequestedPermissions))
	if err != nil {
		logrus.WithError(err).Error("Error in GetToken()")
		return nil, err // consider wrapping in a service-level error
	}
	return &pb.GetTokenResponse{token}, nil
}

func (s *server) Validate(ctx context.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	c, err := s.auth.Validate(r.Token)
	if err != nil {
		logrus.WithError(err).Error("Error in Validate()")
	}

	pc := mappers.ClaimsToProto(*c)
	return &pb.ValidateResponse{
		Claims: &pc,
	}, nil
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
