FROM golang:alpine as builder
WORKDIR ..
RUN mkdir kn-be-sd
WORKDIR kn-be-sd
COPY . ./
ARG version=dev
RUN make test
RUN make build

FROM alpine
COPY --from=builder /pricing-surge/pricing-surge .
COPY --from=builder /pricing-surge/scripts ./scripts
COPY --from=builder /pricing-surge/migrations ./migrations
RUN chmod +x ./scripts/migrate_and_run.sh
ENV PATH="/kn-be-sd:${PATH}"
