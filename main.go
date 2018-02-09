package main

import (
	"google.golang.org/grpc"
	pb "github.com/suyashkumar/auth-grpc/protos"
	"github.com/suyashkumar/auth-grpc/server"
	"net"
	"log"
)

func main() {

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, server.NewServer())

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to open listener on PORT. %+v", err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve %+v", err)
	}

}
