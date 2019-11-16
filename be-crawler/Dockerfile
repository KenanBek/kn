FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go test -v ./...
RUN go build -o kn cmd/main.go
ENV PATH="/app:${PATH}"
