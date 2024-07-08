package benchmark

import (
	"fmt"
	"testing"

	"github.com/sarthak0714/tollywood/internal/system"
	"github.com/sarthak0714/tollywood/proto"
)

func BenchmarkActorMessagePassing(b *testing.B) {
	sys := system.NewActorSystem()

	// Spawn two actors
	sys.SpawnActor("actor1")
	sys.SpawnActor("actor2")

	b.ResetTimer()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sys.SendMessage(&proto.Envelope{
			Sender:      "actor1",
			Target:      "actor2",
			MessageData: []byte("Hello"),
		})
	}

	b.StopTimer()
	b.ReportMetric(float64(b.N)/b.Elapsed().Seconds(), "msgs/sec")
}

func BenchmarkActorSpawning(b *testing.B) {
	sys := system.NewActorSystem()

	b.ResetTimer()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actorID := fmt.Sprintf("actor%d", i)
		_, err := sys.SpawnActor(actorID)
		if err != nil {
			b.Fatalf("Failed to spawn actor: %v", err)
		}
	}

	b.StopTimer()
	b.ReportMetric(float64(b.N)/b.Elapsed().Seconds(), "actors/sec")
}
