GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

BINARY_KN_NAME=kn
SOURCE_KN_NAME=cmd/main.go

.PHONY: build
build:
	$(GOBUILD) -o ${BINARY_KN_NAME} -v ${SOURCE_KN_NAME}

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: test/e2e
test/e2e:
	docker build -f Dockerfile -t kn/be:latest . 
	docker-compose -f docker-compose.yml up --build --abort-on-container-exit --exit-code-from kn
