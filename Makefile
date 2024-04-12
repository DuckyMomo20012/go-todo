.PHONY: openapi
openapi: openapi_http

.PHONY: openapi_http
openapi_http:
	oapi-codegen -generate fiber -o internal/tasks/ports/openapi_api.gen.go -package ports api/openapi/tasks.yaml
	oapi-codegen -generate types -o internal/tasks/ports/openapi_types.gen.go -package ports api/openapi/tasks.yaml

.PHONY: gen-proto
gen-proto:
	cd ./api/protobuf && buf generate

.PHONY: lint
lint:
	golangci-lint run

DOCKER_TAG=duckymomo20012/go-todo

.PHONY: docker-build
docker-build:
	docker build -t ghcr.io/$(DOCKER_TAG):latest -f ./docker/tasks/Dockerfile .
