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

# Use the latest foundry image
FROM polard-base:test-hive as polaris-hive

RUN apk add --no-cache bash jq

WORKDIR /

# Add a default genesis file + Inject the startup script
ADD build_eth_genesis.sh /build_eth_genesis.sh
ADD version.txt /version.txt
ADD hive-init.sh /hive-init.sh
ADD config /config

# Build the ethereum style genesis from the cosmos genesis.
RUN ./build_eth_genesis.sh

RUN export HIVE_CHAIN_ID=2061
RUN export HIVE_NETWORK_ID=2061


# Export the usual networking ports to allow outside access to the node
EXPOSE 8545 8546 8547 8551 30303 30303/udp

CMD ["bash", "hive-init.sh"]
