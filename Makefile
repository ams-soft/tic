.PHONY: dev test fmt vet

dev:
	go run ./cmd/tic-demo/main.go

test:
	go test ./...

fmt:
	gofmt -w .

vet:
	go vet ./...
