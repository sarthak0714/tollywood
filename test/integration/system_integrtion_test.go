package integration

import (
	"fmt"
	"testing"
	"time"

	"github.com/sarthak0714/tollywood/internal/system"
	"github.com/sarthak0714/tollywood/proto"
)

func TestActorSystemIntegration(t *testing.T) {
	sys := system.NewActorSystem()

	// Spawn multiple actors
	for i := 0; i < 10; i++ {
		_, err := sys.SpawnActor(fmt.Sprintf("actor%d", i))
		if err != nil {
			t.Fatalf("Failed to spawn actor: %v", err)
		}
	}

	// Send messages between actors
	for i := 0; i < 10; i++ {
		err := sys.SendMessage(&proto.Envelope{
			Sender:      fmt.Sprintf("actor%d", i),
			Target:      fmt.Sprintf("actor%d", (i+1)%10),
			MessageData: []byte("Hello"),
		})
		if err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	// Allow time for message processing
	time.Sleep(100 * time.Millisecond)

	// Terminate all actors
	for i := 0; i < 10; i++ {
		err := sys.TerminateActor(fmt.Sprintf("actor%d", i))
		if err != nil {
			t.Fatalf("Failed to terminate actor: %v", err)
		}
	}
}
