.PHONY: init
init:
	$(MAKE) download-deps

	go env -w GOPRIVATE="github.com/DuckyMomo20012/*"

.PHONY: download-deps
download-deps:
	@# Ref: https://github.com/golang/go/issues/25922#issuecomment-1038394599
	@# Ref: https://marcofranssen.nl/manage-go-tools-via-go-modules
	cat ./internal/tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

	brew bundle --file=./internal/tools/Brewfile;

.PHONY: gen-proto
gen-proto:
	cd ./api/protobuf && buf generate

.PHONY: lint
lint:
	golangci-lint run

DOCKER_TAG=duckymomo20012/go-todo

.PHONY: docker-build
docker-build:
	docker build -t ghcr.io/$(DOCKER_TAG):latest -f ./docker/go-todo/Dockerfile .

.PHONY: docker-push
docker-push:
	docker push ghcr.io/$(DOCKER_TAG):latest
