run:
	go run -race cmd/api.go

build:
	go build -race -o bin/main cmd/api.go

lint:
	golangci-lint run ./...

test:
	go test -covermode=atomic -race -coverprofile=coverage.txt ./...

