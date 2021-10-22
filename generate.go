package learning

//go:generate protoc --proto_path=. --go_out=. --go_opt=module=learning --go-grpc_out=. --go-grpc_opt=module=learning api/hello/v1/hello.proto
