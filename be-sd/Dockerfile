FROM golang:latest as builder
WORKDIR ..
RUN mkdir kn-be-sd
WORKDIR kn-be-sd
COPY . ./
ARG version=dev
RUN make test
RUN make build

FROM alpine
COPY --from=builder /kn-be-sd/kn-be-sd .
ENV PATH="/kn-be-sd:${PATH}"
