package mappers

import (
	"github.com/suyashkumar/auth"
	pb "github.com/suyashkumar/auth-grpc/protos"
)

func ClaimsToProto(c auth.Claims) pb.Claims {
	return pb.Claims{
		UserUUID:    c.UserUUID,
		Email:       c.Email,
		Permissions: pb.Permissions(c.Permissions),
	}
}
