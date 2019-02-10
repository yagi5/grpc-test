package main

import (
	"context"
	"fmt"
	"log"
	"net"

	userpb "github.com/yagi5/grpc-test/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Get(ctx context.Context, req *userpb.GetRequest) (*userpb.GetResponse, error) {
	if req.Id == "1" {
		return &userpb.GetResponse{User: &userpb.User{
			Id:    "1",
			Name:  "bob",
			Email: "email@example.com",
		}}, nil

	}
	return nil, fmt.Errorf("id invalid")
}

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	userpb.RegisterUsersServer(grpcServer, &server{})
	grpcServer.Serve(listenPort)
}
