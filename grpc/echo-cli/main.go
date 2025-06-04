package main

import (
	"context"
	"fmt"
	"go0base/grpc/echo"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {

	// conn, err := grpc.Dial("localhost:50051",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	client := echo.NewEchoClient(conn)

	ctx := context.Background()

	in := &echo.EchoMsg{
		Name: "nick",
		Addr: &echo.Addr{
			Province: "湖南",
			City:     "长沙",
		},
		Birthday: timestamppb.New(time.Now()),
		Data:     []byte("今天天气不错"),
		Gender:   echo.Gender_MALE,
		Hobby:    []string{"羽毛球", "携带名", "看博客"},
	}
	//----  Unary request(一元请求)
	res, err := client.UnaryEcho(ctx, in)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)

	//---- stream request流式请求
	// metadata.NewOutgoingContext()
	// metadata.FromIncomingContext()
	stream, err := client.ClientStreamEcho(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 5; i++ {
		err = stream.Send(in)
		if err != nil {
			log.Fatalln(err)
		}
	}
	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)

}
