GO=go
GORUN=$(GO) run
GOBUILD=$(GO) build
GOTEST=$(GO) test
GOCOVER=$(GO) tool cover

BINARY_KN_NAME=kn
SOURCE_KN_NAME=cmd/kn/kn.go

.PHONY: run/se
run/se:
	$(GORUN) cmd/se/main.go

.PHONY: run/crawler
run/crawler:
	$(GORUN) cmd/crawler/main.go

.PHONY: build
build:
	$(GOBUILD) -o kn-se -v cmd/se/main.go
	$(GOBUILD) -o kn-se-crawler -v cmd/crawler/main.go

.PHONY: images
image:
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
	docker build -f Dockerfile -t kn-be-se:latest . && \
	docker-compose -f docker-compose.yml up --build --abort-on-container-exit --exit-code-from kn-be-se
