package example

import (
	"context"
	"fmt"
	"io"

	"github.com/heimonsy/grpc-api-gateway/proto"
	"google.golang.org/grpc"
)

type ExampleService struct{}

func (s *ExampleService) Register(gs *grpc.Server) {
	proto.RegisterExampleServer(gs, s)
}

func (s *ExampleService) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	return &proto.AddResponse{
		Result: req.A + req.B,
	}, nil
}

func (s *ExampleService) Connect(
	stream proto.Example_ConnectServer,
) error {
	for {
		cmd, err := stream.Recv()
		if err == io.EOF || err == context.Canceled {
			break
		}
		if err != nil {
			fmt.Println(err)
			return err
		}
		switch cmd.Type {
		case proto.Command_CLOSE:
			return nil
		default:
			err := stream.Send(&proto.CommandResponse{
				Id: cmd.Id,
			})
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	return nil
}
