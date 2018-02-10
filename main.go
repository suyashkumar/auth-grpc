package main

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/suyashkumar/auth-grpc/config"
	"github.com/suyashkumar/auth-grpc/log"
	pb "github.com/suyashkumar/auth-grpc/protos"
	"github.com/suyashkumar/auth-grpc/server"
	"google.golang.org/grpc"
)

func main() {

	log.Configure()

	s := grpc.NewServer()
	as, err := server.NewServer()
	if err != nil {
		logrus.WithError(err).Error("Error initializing AuthServer implementation")
	}
	pb.RegisterAuthServer(s, as)

	l, err := net.Listen("tcp", config.Get(config.Port))
	if err != nil {
		logrus.Error(err)
	}

	if err := s.Serve(l); err != nil {
		logrus.WithError(err).Error("Error starting up server")
	}

}
