FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go test -v ./...
RUN make build
ENV PATH="/app:${PATH}"

#RUN kn-be-se

# run crawler with default settings
# it will try to crawl once and exit
# from docker-compose it must be running
# with count and interval settings
RUN kn-be-crawler
