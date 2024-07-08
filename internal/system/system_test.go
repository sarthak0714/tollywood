package system

import (
	"testing"

	"github.com/sarthak0714/tollywood/proto"
)

func TestActorSystem(t *testing.T) {
	sys := NewActorSystem()

	// Test spawning an actor
	_, err := sys.SpawnActor("actor1")
	if err != nil {
		t.Fatalf("Failed to spawn actor: %v", err)
	}

	// Test getting an existing actor
	_, exists := sys.GetActor("actor1")
	if !exists {
		t.Fatalf("Actor not found after spawning")
	}

	// Test sending a message
	err = sys.SendMessage(&proto.Envelope{
		Sender:      "test",
		Target:      "actor1",
		MessageData: []byte("Hello"),
	})
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Test terminating an actor
	err = sys.TerminateActor("actor1")
	if err != nil {
		t.Fatalf("Failed to terminate actor: %v", err)
	}

	// Verify actor is removed
	_, exists = sys.GetActor("actor1")
	if exists {
		t.Fatalf("Actor still exists after termination")
	}
}
