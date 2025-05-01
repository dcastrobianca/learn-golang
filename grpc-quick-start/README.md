go mod edit -replace example.com/greetings=../go-official-tutorial/greetings
protoc --go_out=. --go-grpc_out=. grpc/v1/proto/greeter.proto