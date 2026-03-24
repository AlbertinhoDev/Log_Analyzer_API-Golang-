APP_NAME=log-analyzer-api

.PHONY: run test docker-build docker-run

run:
	go run ./cmd/server

test:
	go test ./...

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run -p 8080:8080 $(APP_NAME)
