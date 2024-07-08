package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sarthak0714/tollywood/internal/remote"
	"github.com/sarthak0714/tollywood/internal/system"
	"github.com/sarthak0714/tollywood/proto"
	"google.golang.org/grpc"
)

func main() {
	actorSystem := system.NewActorSystem()

	// Start the remote server
	go func() {
		if err := remote.StartRemoteServer(actorSystem, ":8080"); err != nil {
			log.Fatalf("Failed to start remote server: %v", err)
		}
	}()

	// Wait for the server to start
	time.Sleep(time.Second)

	// Connect to the remote server
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to remote server: %v", err)
	}
	defer conn.Close()

	client := proto.NewRemoteServiceClient(conn)

	// Spawn an actor
	spawnResp, err := client.Command(context.Background(), &proto.RemoteCommand{
		Command: &proto.RemoteCommand_SpawnActor{
			SpawnActor: &proto.SpawnActor{ActorId: "actor1"},
		},
	})
	if err != nil {
		log.Fatalf("Failed to spawn actor: %v", err)
	}
	fmt.Println(spawnResp.Message)

	// Send a message to the actor
	stream, err := client.Receive(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}

	err = stream.Send(&proto.Envelope{
		Sender:      "main",
		Target:      "actor1",
		MessageData: []byte("Hello, Actor!"),
	})
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Wait for a moment to allow message processing
	time.Sleep(time.Second)

	// Terminate the actor
	terminateResp, err := client.Command(context.Background(), &proto.RemoteCommand{
		Command: &proto.RemoteCommand_TerminateActor{
			TerminateActor: &proto.TerminateActor{ActorId: "actor1"},
		},
	})
	if err != nil {
		log.Fatalf("Failed to terminate actor: %v", err)
	}
	fmt.Println(terminateResp.Message)
}
