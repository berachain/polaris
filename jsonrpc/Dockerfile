# syntax=docker/dockerfile:1
# SPDX-License-Identifier: BUSL-1.1
#
# Copyright (C) 2023, Berachain Foundation. All rights reserved.
# Use of this software is govered by the Business Source License included
# in the LICENSE file of this repository and at www.mariadb.com/bsl11.
#
# ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
# TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
# VERSIONS OF THE LICENSED WORK.
#
# THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
# LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
# LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
#
# TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
# AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
# EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
# TITLE.



ARG RUNNER_IMAGE
ARG GO_VERSION

#######################################################
###         Stage 1 - Build the Application         ###
#######################################################


FROM golang:${GO_VERSION}-alpine as builder

# Copy our source code into the container
WORKDIR /jsonrpc

# Install some stuff needed for the build
RUN set -eux
RUN apk update && apk add --no-cache musl-dev gcc build-base ca-certificates git;

# manage dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the build folder
COPY ./build ./build

# Install Mage
RUN go run build/mage/setup/main.go

# Copy the remaining files
COPY . ./

# For go-ethereum
ENV CGO_ENABLED=1

# Build jsonrpc binary
RUN  go build \
    -mod=readonly \
    -ldflags "-w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
    -trimpath \
    -o /jsonrpc/bin/ \
    ./...

#######################################################
###        Stage 2 - Prepare the Final Image        ###
#######################################################

FROM ${RUNNER_IMAGE}

COPY --from=builder /jsonrpc/bin/jsonrpcd /bin/jsonrpcd

ENV HOME /jsonrpc
WORKDIR $HOME

EXPOSE 8545
EXPOSE 8546

ENTRYPOINT ["jsonrpcd"]