GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCOVER=$(GOCMD) tool cover

BINARY_KN_NAME=kn
SOURCE_KN_NAME=cmd/main.go

.PHONY: build
build:
	$(GOBUILD) -o ${BINARY_KN_NAME} -v ${SOURCE_KN_NAME}

.PHONY: images
images:
	docker build -f Dockerfile --tag kn/be-se:latest

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: cover
test/cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out

.PHONY: test/e2e
test/e2e:
	docker build -f Dockerfile -t kn/be-se:latest .
	docker-compose -f docker-compose.yml up --build --abort-on-container-exit --exit-code-from kn
