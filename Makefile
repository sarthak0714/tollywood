build:
	go build -o bin/actorengine ./cmd/actorengine

test:
	go test -v ./...

run: build
	./bin/actorengine

clean:
	rm -rf bin

proto:
	protoc --go_out=. --go_opt=paths=source_relative pkg/pb/messages.proto

deps:
	go mod tidy
	go mod verify

