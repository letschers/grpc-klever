build_proto:
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     proto/crypto.proto

run_server:
	go run main.go

run_client:
	go run client/client.go