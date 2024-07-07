package actor

import (
	"context"
	"fmt"
	"sync"

	"github.com/sarthak0714/tollywood/proto"
)

type Actor struct {
	Id      string
	mailbox chan *proto.Envelope
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
}

func NewActor(id string) *Actor {
	ctx, cancel := context.WithCancel(context.Background())
	return &Actor{
		Id:      id,
		mailbox: make(chan *proto.Envelope, 100),
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (a *Actor) Start() {
	a.wg.Add(1)
	go a.processMessages()
}

func (a *Actor) Stop() {
	a.cancel()
	a.wg.Wait()
}

func (a *Actor) Send(envelope *proto.Envelope) {
	select {
	case a.mailbox <- envelope:
	default:
		fmt.Printf("Mailbox full for actor %s\n", a.Id)
	}
}

func (a *Actor) processMessages() {
	defer a.wg.Done()
	for {
		select {
		case msg := <-a.mailbox:
			fmt.Printf("Actor %s received message from %s: %s\n", a.Id, msg.Sender, string(msg.MessageData))
		case <-a.ctx.Done():
			return
		}
	}
}
