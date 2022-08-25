all: clean gen init
	go build server.go
	go build client.go

init:
	go mod init kalra.com/goProjects || true
	go mod edit -replace kalra.com=/Users/developer/repos/goProjects || true
	go mod tidy

gen:
	mkdir -p bidiGRPC
	protoc --go_out=./bidiGRPC \
        --go_opt=paths=source_relative \
        --go-grpc_out=./bidiGRPC \
        --go-grpc_opt=paths=source_relative \
        contract.proto

clean:
	rm -rf bidiGRPC go.sum go.mod  contract.pb.go contract_grpc.pb.go server client
