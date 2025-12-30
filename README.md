# grpc-golang

## Getting Started
```
$ go run ./cmd/serve
$ grpcurl -plaintext \
  -d '{"name":"Guest"}' \
  localhost:50051 \
  greeter.Greeter/SayHello
```

## gRPC
```
protoc \
  --go_out=../internal/pb --go_opt=paths=source_relative \
  --go-grpc_out=../internal/pb --go-grpc_opt=paths=source_relative \
  $(find . -name '*.proto' -print)
```