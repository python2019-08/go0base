package main

import (
	"go0base/grpc/echo"
	"go0base/grpc/echo-server/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	echo.RegisterEchoServer(s, server.NewEchoServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
