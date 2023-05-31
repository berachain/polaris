# This simulation runs JSON-RPC API tests.

ARG GO_VERSION=1.20.4

FROM golang:${GO_VERSION}-alpine as builder
RUN apk add --update gcc musl-dev linux-headers

# Build the simulator executable.
ADD . /source
WORKDIR /source
RUN go mod tidy
RUN go build .

# Build the simulator run container.
FROM alpine:latest
ADD . /source
WORKDIR /source
COPY --from=builder /source/rpc .
ENTRYPOINT ["./rpc"]
