# SPDX-License-Identifier: MIT
#
# Copyright (c) 2023 Berachain Foundation
#
# Permission is hereby granted, free of charge, to any person
# obtaining a copy of this software and associated documentation
# files (the "Software"), to deal in the Software without
# restriction, including without limitation the rights to use,
# copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the
# Software is furnished to do so, subject to the following
# conditions:
#
# The above copyright notice and this permission notice shall be
# included in all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
# EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
# OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
# NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
# HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
# WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
# FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
# OTHER DEALINGS IN THE SOFTWARE.

# This simulation runs JSON-RPC API tests.

# ARG GO_VERSION=1.20.4
# ARG GOARCH=arm64
# ARG GOOS=linux

# FROM golang:${GO_VERSION}-alpine as builder
# RUN set -eux; \
#     apk add git linux-headers ca-certificates build-base
# RUN apk add --update gcc musl-dev linux-headers git

# # Build the simulator executable.
# ADD . /source
# WORKDIR /source
# RUN go mod tidy

# ARG GOARCH
# ARG GOOS

# RUN env GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=1 && \
#     go build \
#     -mod=readonly \
#     -tags "netgo,ledger,cgo,muslc" \
#     -ldflags "-X github.com/cosmos/cosmos-sdk/version.Name=$NAME \
#     -X github.com/cosmos/cosmos-sdk/version.AppName=$APP_NAME \
#     -X github.com/cosmos/cosmos-sdk/version.Version=$VERSION \
#     -X github.com/cosmos/cosmos-sdk/version.Commit=$COMMIT \
#     -X github.com/cosmos/cosmos-sdk/version.BuildTags='netgo,ledger,muslc' \
#     -X github.com/cosmos/cosmos-sdk/types.DBBackend=$DB_BACKEND \
#     -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
#     .

# # Build the simulator run container.
# FROM alpine:latest
# ADD . /source
# WORKDIR /source
# COPY --from=builder /source/rpc .
# ENTRYPOINT ["./rpc"]

# This simulation runs JSON-RPC API tests.
FROM golang:1-alpine as builder
RUN apk add --update gcc musl-dev linux-headers

# Build the simulator executable.
ADD . /source
WORKDIR /source
RUN ls -al
RUN go mod tidy
RUN go build -v .

# Build the simulator run container.
FROM alpine:latest
ADD . /source
WORKDIR /source
COPY --from=builder /source/rpc .
ENTRYPOINT ["./rpc"]
