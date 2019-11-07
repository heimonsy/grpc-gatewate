package main

import (
	"log"
	"net"

	"github.com/heimonsy/grpc-api-gateway/service/example"
	"google.golang.org/grpc"
)

func main() {
	ln, err := net.Listen("tcp", ":8811")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	(&example.ExampleService{}).Register(grpcServer)

	err = grpcServer.Serve(ln)
	if err != nil {
		log.Fatalln(err)
	}
}
