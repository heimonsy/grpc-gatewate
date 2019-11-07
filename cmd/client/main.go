package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/heimonsy/grpc-api-gateway/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8810", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := proto.NewExampleClient(conn)

	for i := 0; i < 10; i++ {
		resp, err := client.Add(context.Background(), &proto.AddRequest{A: 2, B: 2})
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Add Result:", resp.Result)
	}

	stream, err := client.Connect(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	var id int32
	id++
	err = stream.Send(&proto.Command{Type: proto.Command_PING, Id: id})
	if err != nil {
		log.Fatalln(err)
	}
	cmdResp, err := stream.Recv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("CMD:", cmdResp.Id)

	stream.Send(&proto.Command{Type: proto.Command_CLOSE})
	<-time.After(time.Millisecond * 50)
}
