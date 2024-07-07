package system

import (
	"fmt"
	"sync"

	"github.com/sarthak0714/tollywood/internal/actor"
	"github.com/sarthak0714/tollywood/proto"
)

type ActorSystem struct {
	actors map[string]*actor.Actor
	mu     sync.RWMutex
}

func NewActorSystem() *ActorSystem {
	return &ActorSystem{
		actors: make(map[string]*actor.Actor),
	}
}

func (s *ActorSystem) SpawnActor(id string) (*actor.Actor, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.actors[id]; exists {
		return nil, fmt.Errorf("actor with id %s already exists", id)
	}

	newActor := actor.NewActor(id)
	s.actors[id] = newActor
	newActor.Start()
	return newActor, nil
}

func (s *ActorSystem) GetActor(id string) (*actor.Actor, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	actor, exists := s.actors[id]
	return actor, exists
}

func (s *ActorSystem) TerminateActor(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	actor, exists := s.actors[id]
	if !exists {
		return fmt.Errorf("actor with id %s not found", id)
	}

	actor.Stop()
	delete(s.actors, id)
	return nil
}

func (s *ActorSystem) SendMessage(envelope *proto.Envelope) error {
	actor, exists := s.GetActor(envelope.Target)
	if !exists {
		return fmt.Errorf("actor with id %s not found", envelope.Target)
	}

	actor.Send(envelope)
	return nil
}
