.PHONY: proto

proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/messages.proto

test:
	@go test ./... -v


bench:
	@go test -bench=. ./test/benchmark 




