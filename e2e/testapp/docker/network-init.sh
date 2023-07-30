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

NODE_COUNT="$1"
CONTAINER_PREFIX="polard-node"
HOMEDIR="/root/.polard"
CHAINID="polaris-2061"
: "${NODE_COUNT:?Argument (node count) not passed}"

# init nodes
for ((i=0; i<NODE_COUNT; i++))
do
  docker exec "$CONTAINER_PREFIX$i" bash -c "genbuild init $CONTAINER_PREFIX$i --home $HOMEDIR --chain-id $CHAINID"
done

# copy all gentx to container 0
for ((i=1; i<NODE_COUNT; i++))
do
  docker cp "$CONTAINER_PREFIX$i":$HOMEDIR/config/gentx ./temp
  docker cp ./temp/gentx "$CONTAINER_PREFIX"0:$HOMEDIR/config
done

# update genesis file using gentx
docker exec "$CONTAINER_PREFIX"0 bash -c "genbuild collect --home $HOMEDIR"

# copy genesis file to all containers
docker cp "$CONTAINER_PREFIX"0:$HOMEDIR/config/genesis.json ./temp/genesis.json
for ((i=1; i<NODE_COUNT; i++))
do
  docker cp ./temp/genesis.json "$CONTAINER_PREFIX$i":$HOMEDIR/config/genesis.json
done

# start
# docker exec -it $CONTAINER0 bash -c "$SCRIPTS/seed-start.sh"
# docker exec -it $CONTAINER1 bash -c "$SCRIPTS/seed-start.sh"
