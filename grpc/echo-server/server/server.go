package server

import (
	"context"
	"fmt"
	"go0base/grpc/echo"
	"io"
	"log"

	"google.golang.org/grpc"
)

func NewEchoServer() echo.EchoServer {
	return &echoServer{}
}

type echoServer struct {
	echo.UnimplementedEchoServer
}

func (s *echoServer) UnaryEcho(ctx context.Context, in *echo.EchoMsg) (*echo.EchoMsg, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method UnaryEcho not implemented")
	fmt.Printf("server recv message :%+v \n", in)
	return in, nil
}

/*
// bili-code
func (s *echoServer) ClientStreamEcho(stream echo.Echo_ClientStreamEchoServer) error {
return status.Errorf(codes.Unimplemented,format:"method ClientStreamEcho not implemented")
}
*/
func (s *echoServer) ClientStreamEcho(stream grpc.ClientStreamingServer[echo.EchoMsg, echo.EchoResponse]) error {
	// return status.Errorf(codes.Unimplemented, "method ClientStreamEcho not implemented")

	for {
		in, err := stream.Recv()
		if err != nil && err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("server recv message :%+v \n", in)
	}

	err := stream.SendAndClose(&echo.EchoResponse{
		Ok: true,
	})
	return err
}
