# syntax=docker/dockerfile:1
# Copyright (C) 2023, Berachain Foundation. All rights reserved.
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