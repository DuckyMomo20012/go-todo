.PHONY: gen-proto
gen-proto:
	cd ./api/protobuf && buf generate

.PHONY: lint
lint:
	golangci-lint run

DOCKER_TAG=duckymomo20012/go-todo

.PHONY: docker-build
docker-build:
	docker build -t ghcr.io/$(DOCKER_TAG):latest -f ./docker/task/Dockerfile .
