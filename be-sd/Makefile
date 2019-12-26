GO=go
GORUN=$(GO) run
GOBUILD=$(GO) build
GOTEST=$(GO) test
GOCOVER=$(GO) tool cover

.PHONY: run/sd
run/sd:
	KN_BE_SD_PORT=9080 $(GORUN) cmd/sd/main.go

.PHONY: build
build:
	$(GOBUILD) -o kn-be-sd -v cmd/sd/main.go

.PHONY: image
image:
	docker build -f Dockerfile --tag kn-be-sd:latest .

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
	docker build -f Dockerfile -t kn-be-sd:latest .  # && \
	# docker-compose -f docker-compose.yml up --build --abort-on-container-exit --exit-code-from kn-be-sd
