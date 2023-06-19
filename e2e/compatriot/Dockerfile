
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

ARG GOVERSION
ARG GOARCH=arm64
ARG GOOS=darwin
ARG POLARD_VERSION=test-compatriot
ARG COMPATRIOT_PATH=./e2e/compatriot

# Use the latest foundry image
FROM ghcr.io/foundry-rs/foundry:nightly as foundry

# Set working directory
WORKDIR /workdir

# Required for forge install.
COPY .git/ .git/

# Install dependecies for solidity contracts.
RUN forge install --no-commit

# Use the latest foundry image
FROM polard-base:${POLARD_VERSION} as polard-compatriot

RUN apk add --no-cache bash jq

WORKDIR /

# RUN curl -L https://foundry.paradigm.xyz | bash
# RUN source /root/.profile
# RUN foundryup

# Copy polard requirements
COPY ./e2e/compatriot/config /config

# Copy compatriot scripts
COPY ./e2e/compatriot/scripts /scripts
COPY ./e2e/compatriot/compatriot.sh /compatriot.sh

RUN ls /config

RUN chmod +x scripts/*
RUN chmod +x *.sh

# Set the home directory
ENV HOME /
WORKDIR $HOME

EXPOSE 8545

CMD ["bash", "compatriot.sh"]
