GO = go

.PHONY: tools
tools:
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.0
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11

all: tools