package actor

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/sarthak0714/tollywood/proto"
)

type ActorState int32

const (
	Running ActorState = iota
	Stopped
)

type Actor struct {
	ID       string
	mailbox  chan *proto.Envelope
	ctx      context.Context
	state    ActorState
	cancelch chan struct{}
}

func NewActor(id string) *Actor {
	ctx := context.WithoutCancel(context.Background())
	return &Actor{
		ID:       id,
		mailbox:  make(chan *proto.Envelope, 100),
		ctx:      ctx,
		state:    Running,
		cancelch: make(chan struct{}, 1),
	}
}

func (a *Actor) Start() {
	go a.processMessages()
}

func (a *Actor) Stop() {
	if atomic.CompareAndSwapInt32((*int32)(&a.state), int32(Running), int32(Stopped)) {
		a.cancelch <- struct{}{}

	}
}

func (a *Actor) Send(envelope *proto.Envelope) {
	select {
	case a.mailbox <- envelope:
	default:
		fmt.Printf("Mailbox full for actor %s\n", a.ID)
	}
}

func (a *Actor) processMessages() {
	for {
		select {
		case msg := <-a.mailbox:
			fmt.Printf("Actor %s received message from %s: %s\n", a.ID, msg.Sender, string(msg.MessageData))
		case <-a.ctx.Done():
			return
		}
	}
}
