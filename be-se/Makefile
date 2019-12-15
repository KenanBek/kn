GO=go
GORUN=$(GO) run
GOBUILD=$(GO) build
GOTEST=$(GO) test
GOCOVER=$(GO) tool cover

.PHONY: run/se
run/se:
	$(GORUN) cmd/se/main.go

.PHONY: run/crawler
run/crawler:
	$(GORUN) cmd/crawler/main.go

.PHONY: run/mongo
run/mongo:
	docker run --name kn-mongo -d -p 27017:27017 mongo:latest

.PHONY: run/mongo/destroy
run/mongo/destroy:
	docker stop kn-mongo
	docker rm kn-mongo

.PHONY: build
build:
	$(GOBUILD) -o kn-be-se -v cmd/se/main.go
	$(GOBUILD) -o kn-be-crawler -v cmd/crawler/main.go

.PHONY: image
image:
	docker build -f Dockerfile --tag kn-be-se:latest .

.PHONY: lint
lint:
	golint -set_exit_status=1 `go list ./...`

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: mocks
mocks:
	$(GO) generate -v ./...

.PHONY: test/cover
test/cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out

.PHONY: test/e2e
test/e2e:
	docker build -f Dockerfile -t kn-be-se:latest . && \
	docker-compose -f docker-compose.yml up --build --abort-on-container-exit --exit-code-from kn-be-se
