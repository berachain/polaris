# syntax=docker/dockerfile:1
#
# Copyright (C) 2022, Berachain Foundation. All rights reserved.
# See the file LICENSE for licensing terms.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
# AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
# DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
# FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
# DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
# SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
# CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
# OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

#######################################################
###           Stage 0 - Build Arguments             ###
#######################################################

ARG GO_VERSION=1.20.4
ARG GOARCH=amd64
ARG GOOS=linux
ARG NAME=polaris-cosmos
ARG APP_NAME=polard
ARG DB_BACKEND=pebbledb
ARG CMD_PATH=./cosmos/simapp/polard
ARG FOUNDRY_DIR=contracts
ARG GO_WORK=""

#######################################################
###         Stage 1 - Build the Application         ###
#######################################################

FROM golang:${GO_VERSION}-alpine as builder

# Setup some alpine stuff that nobody really knows how or why it works.
# Like if ur reading this and u dunno just ask the devops guy or something.
RUN set -eux; \
    apk add git linux-headers ca-certificates build-base

# Set the working directory
WORKDIR /workdir

RUN echo $GO_WORK

# Copy go.mod and go.sum files (🔥 upgrade)
COPY ./go.work ./go.work.sum ./

# RUN for dir in $GO_WORK; do \
#         cp ./$dir/go.mod ./$dir/go.sum ./$dir/; \
#     done

COPY ./contracts/go.sum ./contracts/go.mod ./contracts/
COPY ./cosmos/go.sum ./cosmos/go.mod ./cosmos/
COPY ./eth/go.sum ./eth/go.mod ./eth/
COPY ./lib/go.sum ./lib/go.mod ./lib/
COPY ./magefiles/go.sum ./magefiles/go.mod ./magefiles/
COPY ./e2e/localnet/go.sum ./e2e/localnet/go.mod ./e2e/localnet/

# Download the go module dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build args
ARG NAME
ARG GOARCH
ARG GOOS
ARG APP_NAME
ARG DB_BACKEND
ARG CMD_PATH

# Build Executable
RUN VERSION=$(echo $(git describe --tags) | sed 's/^v//') && \
    COMMIT=$(git log -1 --format='%H') && \
    env GOOS=${GOOS} GOARCH=${GOARCH} && \
    env NAME=${NAME} DB_BACKEND=${DB_BACKEND} && \
    env APP_NAME=${APP_NAME} && \
    go build \
    -mod=readonly \
    -tags "netgo,ledger,muslc" \
    -ldflags "-X github.com/cosmos/cosmos-sdk/version.Name=$NAME \
    -X github.com/cosmos/cosmos-sdk/version.AppName=$APP_NAME \
    -X github.com/cosmos/cosmos-sdk/version.Version=$VERSION \
    -X github.com/cosmos/cosmos-sdk/version.Commit=$COMMIT \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags='netgo,ledger,muslc' \
    -X github.com/cosmos/cosmos-sdk/types.DBBackend=$DB_BACKEND \
    -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
    -trimpath \
    -o /workdir/bin/ \
    ${CMD_PATH}

#######################################################
###        Stage 2 - Prepare the Final Image        ###
#######################################################

FROM golang:${GO_VERSION}-alpine

# Build args
ARG APP_NAME

# Copy over built executable into a fresh container.
COPY --from=builder /workdir/bin/${APP_NAME} /bin/