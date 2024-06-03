dep:
	@go mod tidy

build:
	@go build -o bin/task2 cmd/main.go

test:
	@go test -v ./tests/unit/...

run: build
	@./bin/task2

podman-image:
	@podman build -t task2 .

docker-image:
	@docker build -t task2 .

podman-up: podman-image
	@podman-compose up -d

podman-down:
	@podman-compose down

docker-up: docker-image
	@docker-compose up -d

docker-down:
	@docker-compose down