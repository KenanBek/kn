FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go test -v ./...
RUN make build
ENV PATH="/app:${PATH}"
