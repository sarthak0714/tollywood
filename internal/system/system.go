package system

import (
	"fmt"
	"sync"

	"github.com/sarthak0714/tollywood/internal/actor"
	"github.com/sarthak0714/tollywood/proto"
)

type ActorSystem struct {
	actors sync.Map
}

func NewActorSystem() *ActorSystem {
	return &ActorSystem{}
}

func (s *ActorSystem) SpawnActor(id string) (*actor.Actor, error) {
	if _, exists := s.actors.Load(id); exists {
		return nil, fmt.Errorf("actor with id %s already exists", id)
	}

	newActor := actor.NewActor(id)
	s.actors.Store(id, newActor)
	newActor.Start()
	return newActor, nil
}

func (s *ActorSystem) GetActor(id string) (*actor.Actor, bool) {
	if a, exists := s.actors.Load(id); exists {
		return a.(*actor.Actor), true
	}
	return nil, false
}

func (s *ActorSystem) TerminateActor(id string) error {
	if a, exists := s.actors.LoadAndDelete(id); exists {
		a.(*actor.Actor).Stop()
		return nil
	}
	return fmt.Errorf("actor with id %s not found", id)
}

func (s *ActorSystem) SendMessage(envelope *proto.Envelope) error {
	if a, exists := s.actors.Load(envelope.Target); exists {
		a.(*actor.Actor).Send(envelope)
		return nil
	}
	return fmt.Errorf("actor with id %s not found", envelope.Target)
}
