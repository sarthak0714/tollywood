# Tollywood 

Tollywood is a Go-based implementation of the actor model **[Actor Model](https://en.wikipedia.org/wiki/Actor_model)**. It offers a straightforward framework for building scalable and distributed systems using the actor pattern. 

## Features

- Lightweight, concurrent actors
- Message-passing based communication
- Remote actor management and communication via gRPC
- Easy to integrate with existing Go applications
- Scalable architecture suitable for distributed systems

## Installation

To install Tollywood, use the following command:

```
go get github.com/sarthak0714/tollywood
```

## Quick Start

Here's a simple example of how to use Tollywood:

```go
package main

import (
    "fmt"
    "github.com/sarthak0714/tollywood/internal/system"
    "github.com/sarthak0714/tollywood/proto"
)

func main() {
    // Create a new actor system
    actorSystem := system.NewActorSystem()

    // Spawn an actor
    actor, err := actorSystem.SpawnActor("myActor")
    if err != nil {
        fmt.Printf("Failed to spawn actor: %v\n", err)
        return
    }

    // Send a message to the actor
    message := &proto.Envelope{
        Sender:      "main",
        Target:      "myActor",
        MessageData: []byte("Hello, Actor!"),
    }
    err = actorSystem.SendMessage(message)
    if err != nil {
        fmt.Printf("Failed to send message: %v\n", err)
        return
    }

    // ... Handle actor responses and implement your logic
}
```

## Architecture

Tollywood consists of the following main components:

- `Actor`: The basic unit of computation that receives and processes messages.
- `ActorSystem`: Manages the lifecycle of actors and facilitates message passing.
- `RemoteServer`: Enables remote communication and management of the actor system using gRPC.

## Contributing

Contributions to Tollywood are welcome! Please feel free to submit a Pull Request.


