fmt:
	go fmt ./...

tidy:
	go mod tidy

server: fmt tidy
	go run server.go

.PHONY: fmt tidy server