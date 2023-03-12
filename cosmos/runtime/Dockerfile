
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

ARG GO_VERSION

#######################################################
###         Stage 1 - Build Smart Contracts         ###
#######################################################

# Use the latest foundry image
FROM ghcr.io/foundry-rs/foundry as foundry

WORKDIR /polaris

# Build and test the source code
ARG PRECOMPILE_CONTRACTS_DIR
COPY ${PRECOMPILE_CONTRACTS_DIR} ${PRECOMPILE_CONTRACTS_DIR}
WORKDIR /polaris/${PRECOMPILE_CONTRACTS_DIR}

RUN forge build


# #############################dock##########################
# ###         Stage 2 - Build the Application         ###
# #######################################################

FROM golang:${GO_VERSION}-alpine as builder

# Copy our source code into the container
WORKDIR /polaris
COPY . .

# Setup some alpine stuff that nobody really knows why we need other
# than docker geeks cause let's be real, everyone else just googles this stuff
# or asks that one really smart guy on their devops team to fio.
RUN set -eux; apk add --no-cache ca-certificates build-base;
RUN apk add git



# Needed by github.com/zondax/hid
RUN apk add linux-headers

# Copy the forge output
ARG PRECOMPILE_CONTRACTS_DIR
# COPY --from=foundry /polaris/${PRECOMPILE_CONTRACTS_DIR}/out /polaris//${PRECOMPILE_CONTRACTS_DIR}/out

# Copy the go mod and sum files
COPY go.mod ./
COPY go.sum ./


# Build berad binary
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    VERSION=$(echo $(git describe --tags) | sed 's/^v//') && \
    COMMIT=$(git log -1 --format='%H') && \
    # GOARCH=amd64 \
    # GOOS=linux \
    go build \
    -mod=readonly \
    -tags "netgo,ledger,muslc" \
    -ldflags "-X github.com/cosmos/cosmos-sdk/version.Name="berachain" \
    -X github.com/cosmos/cosmos-sdk/version.AppName="berad" \
    -X github.com/cosmos/cosmos-sdk/version.Version=$VERSION \
    -X github.com/cosmos/cosmos-sdk/version.Commit=$COMMIT \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags='netgo,ledger,muslc' \
    -X github.com/cosmos/cosmos-sdk/types.DBBackend="pebbledb" \
    -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
    -trimpath \
    -o /polaris/bin/ \
    ./...

#######################################################
###        Stage 3 - Prepare the Final Image        ###
#######################################################

FROM golang:${GO_VERSION}-alpine

RUN apk add --no-cache bash
RUN apk add --no-cache jq

WORKDIR /polaris

COPY --from=builder /polaris/bin/polard /bin/
COPY --from=builder /polaris/init.sh /polaris/

ENV HOME /polaris
WORKDIR $HOME

# TODO: reduce ports and clarify 
EXPOSE 8545
EXPOSE 8546
EXPOSE 26654
EXPOSE 26653
EXPOSE 26656
EXPOSE 26657
EXPOSE 1317

CMD ["bash", "init.sh"]