# syntax=docker/dockerfile:1
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

ARG RUNNER_IMAGE
ARG GO_VERSION

#######################################################
###         Stage 1 - Build Smart Contracts         ###
#######################################################

# Use the latest foundry image
FROM ghcr.io/foundry-rs/foundry as foundry

# Copy our source code into the container
WORKDIR /polaris

# Recursively update all submodules.
COPY ./ ./
RUN git submodule update --recursive --remote 
RUN git submodule foreach --recursive git reset --hard HEAD

# Set working directory. 
WORKDIR /polaris/contracts

# Extra arguments are required for precompiles.
RUN forge build --extra-output-files bin --extra-output-files abi

#######################################################
###         Stage 2 - Build the Application         ###
#######################################################

FROM golang:${GO_VERSION}-alpine as builder

# Setup some alpine stuff that nobody really knows why we need other
# than docker geeks cause let's be real, everyone else just googles this stuff
# or asks that one really smart guy on their devops team to fio.
RUN set -eux; apk add --no-cache ca-certificates build-base;
RUN apk add git

# Needed by github.com/zondax/hid
RUN apk add linux-headers

# Copy our source code into the container
WORKDIR /polaris

# Copy the go mod and sum files
COPY ./go.work ./
COPY ./contracts/go.mod ./contracts/go.mod
COPY ./cosmos/go.mod ./cosmos/go.mod
COPY ./eth/go.mod ./eth/go.mod
COPY ./lib/go.mod ./lib/go.mod
COPY ./magefiles/go.mod ./magefiles/go.mod
COPY ./playground/go.mod ./playground/go.mod
COPY ./testing/go.mod ./testing/go.mod

# Download go dependencies
RUN go work sync
COPY ./magefiles ./magefiles
RUN go run magefiles/cmd/mage.go tidy

# Copy the remaining files
COPY ./ ./

# Generate Contract bindings
COPY --from=foundry /polaris/contracts/out /polaris/contracts/out
RUN go install github.com/ethereum/go-ethereum/cmd/abigen
RUN go install github.com/ethereum/go-ethereum/rlp/rlpgen
RUN go run magefiles/cmd/mage.go generate

# For go-ethereum
ENV CGO_ENABLED=1

# Build polard binary
RUN VERSION=$(echo $(git describe --tags) | sed 's/^v//') && \
    COMMIT=$(git log -1 --format='%H') && \
    go build \
    -mod=readonly \
    -tags "netgo,ledger,muslc" \
    -ldflags "-X github.com/cosmos/cosmos-sdk/version.Name="polaris" \
    -X github.com/cosmos/cosmos-sdk/version.AppName="polard" \
    -X github.com/cosmos/cosmos-sdk/version.Version=$VERSION \
    -X github.com/cosmos/cosmos-sdk/version.Commit=$COMMIT \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags='netgo,ledger,muslc' \
    -X github.com/cosmos/cosmos-sdk/types.DBBackend="pebbledb" \
    -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
    -trimpath \
    -o /polaris/bin/ \
    ./cosmos/...

#######################################################
###        Stage 3 - Prepare the Final Image        ###
#######################################################

FROM ${RUNNER_IMAGE}

COPY --from=builder /polaris/bin/polard /bin/polard
COPY --from=builder /polaris/cosmos/runtime /polaris/cosmos/runtime

ENV HOME /polaris
WORKDIR $HOME

EXPOSE 26654
EXPOSE 26656
EXPOSE 26657
EXPOSE 1317
EXPOSE 8545
EXPOSE 8546

ENTRYPOINT ["polard"]