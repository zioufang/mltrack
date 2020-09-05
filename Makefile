build:
	go build -o bin/mltrack cmd/cli/main.go

server:
	go run cmd/api/main.go
