# This simulation runs JSON-RPC API tests.
FROM golang:1-alpine as builder
RUN apk add --update git ca-certificates gcc musl-dev linux-headers

# We do not support the full execution-apis repo, so we only copy the tests we need.
# RUN git clone --depth 1 https://github.com/ethereum/execution-apis.git /execution-apis

# To run local tests, copy the directory into the same as the simulator and
# uncomment the line below.
ADD tests /execution-apis/tests

# Build the simulator executable.
ADD . /source
WORKDIR /source
RUN go build -v .

# Build the simulator run container.
FROM alpine:latest
ADD . /source
WORKDIR /source
COPY --from=builder /source/rpc-compat .
COPY --from=builder /execution-apis/tests ./tests

ENTRYPOINT ["./rpc-compat"]