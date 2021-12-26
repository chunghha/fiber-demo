i.PHONY: clean test security build run

APP_NAME = fiber-demo
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

docker.run: docker.network swag docker.fiber

docker.network:
	docker network inspect demo-network >/dev/null 2>&1 || \
	docker network create -d bridge demo-network

docker.fiber.build:
	docker build -t fiber-demo .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name fiber-demo \
		--network demo-network \
		-p 3000:3000 \
		fiber-demo

docker.stop: docker.stop.fiber

docker.stop.fiber:
	docker stop fiber-demo

swag:
	swag init
