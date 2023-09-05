package main

import (
	"authorization/pkg/api"
	"authorization/pkg/authorization"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &authorization.GRPCServer{}
	api.RegisterAuthorizationServer(s, srv)

	listener, err := net.Listen("tcp", ":1200")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server in listening on: localhost:1200")
	s.Serve(listener)
}
