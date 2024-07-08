package remote

import (
	"context"
	"fmt"
	"net"

	"github.com/sarthak0714/tollywood/internal/system"
	"github.com/sarthak0714/tollywood/proto"
	"google.golang.org/grpc"
)

type RemoteServer struct {
	proto.UnimplementedRemoteServiceServer
	system *system.ActorSystem
}

func NewRemoteServer(system *system.ActorSystem) *RemoteServer {
	return &RemoteServer{system: system}
}

func (s *RemoteServer) Receive(stream proto.RemoteService_ReceiveServer) error {
	for {
		envelope, err := stream.Recv()
		if err != nil {
			return err
		}

		if err := s.system.SendMessage(envelope); err != nil {
			fmt.Printf("Error sending message: %v\n", err)
		}
	}
}

func (s *RemoteServer) Command(ctx context.Context, cmd *proto.RemoteCommand) (*proto.CommandResponse, error) {
	switch c := cmd.Command.(type) {
	case *proto.RemoteCommand_SpawnActor:
		_, err := s.system.SpawnActor(c.SpawnActor.ActorId)
		if err != nil {
			return &proto.CommandResponse{Success: false, Message: err.Error()}, nil
		}
		return &proto.CommandResponse{Success: true, Message: "Actor spawned"}, nil

	case *proto.RemoteCommand_TerminateActor:
		err := s.system.TerminateActor(c.TerminateActor.ActorId)
		if err != nil {
			return &proto.CommandResponse{Success: false, Message: err.Error()}, nil
		}
		return &proto.CommandResponse{Success: true, Message: "Actor terminated"}, nil

	default:
		return &proto.CommandResponse{Success: false, Message: "Unknown command"}, nil
	}
}

func StartRemoteServer(system *system.ActorSystem, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterRemoteServiceServer(s, NewRemoteServer(system))
	fmt.Printf("Starting gRPC server on %s\n", address)
	return s.Serve(lis)
}
