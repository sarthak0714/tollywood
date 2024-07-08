package actor

import (
	"testing"
	"time"

	"github.com/sarthak0714/tollywood/proto"
)

func TestActorLifecycle(t *testing.T) {
	a := NewActor("test")
	a.Start()

	// Send a message
	a.Send(&proto.Envelope{
		Sender:      "sender",
		Target:      "test",
		MessageData: []byte("Hello"),
	})

	// Allow some time for message processing
	time.Sleep(100 * time.Millisecond)

	a.Stop()
	// Verify that sending to a stopped actor doesn't panic
	a.Send(&proto.Envelope{
		Sender:      "sender",
		Target:      "test",
		MessageData: []byte("After stop"),
	})
}

func TestActorMailboxOverflow(t *testing.T) {
	a := NewActor("test")
	a.Start()

	// Fill the mailbox
	for i := 0; i < 101; i++ {
		a.Send(&proto.Envelope{
			Sender:      "sender",
			Target:      "test",
			MessageData: []byte("Message"),
		})
	}

	// The 101st message should not block
	a.Send(&proto.Envelope{
		Sender:      "sender",
		Target:      "test",
		MessageData: []byte("Overflow"),
	})

	a.Stop()
}
