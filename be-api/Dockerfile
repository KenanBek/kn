FROM golang:latest as builder
RUN mkdir /app
WORKDIR /app
COPY . ./
RUN make test
ARG version=dev
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags "-X main.version=$version" -o kn-be-api -v ./cmd/api/main.go

FROM alpine
COPY --from=builder /app/kn-be-api /
