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

CONTAINER0="polard-node0"
CONTAINER1="polard-node1"
CONTAINER2="polard-node2"
CONTAINER3="polard-node3"

HOMEDIR="/root/.polard"
SCRIPTS="/scripts"

rm -rf ./temp
mkdir ./temp
mkdir ./temp/seed0
mkdir ./temp/seed1
mkdir ./temp/seed2
mkdir ./temp/seed3
touch ./temp/genesis.json

# init step 1 
docker exec $CONTAINER0 bash -c "$SCRIPTS/seed0-init-step1.sh"
docker exec $CONTAINER1 bash -c "$SCRIPTS/seed1-init-step1.sh seed-1"
docker exec $CONTAINER2 bash -c "$SCRIPTS/seed1-init-step1.sh seed-2"
docker exec $CONTAINER3 bash -c "$SCRIPTS/seed1-init-step1.sh seed-3"

# copy genesis.json from seed-0 to seed-1
docker cp $CONTAINER0:$HOMEDIR/config/genesis.json ./temp/genesis.json
docker cp ./temp/genesis.json $CONTAINER1:$HOMEDIR/config/genesis.json

# init step 2
docker exec $CONTAINER1 bash -c "$SCRIPTS/seed1-init-step2.sh seed-1"

# copy genesis.json from seed-1 to seed-2
docker cp $CONTAINER1:$HOMEDIR/config/genesis.json ./temp/genesis.json
docker cp ./temp/genesis.json $CONTAINER2:$HOMEDIR/config/genesis.json

# init step 2
docker exec $CONTAINER2 bash -c "$SCRIPTS/seed1-init-step2.sh seed-2"

# copy genesis.json from seed-2 to seed-3
docker cp $CONTAINER2:$HOMEDIR/config/genesis.json ./temp/genesis.json
docker cp ./temp/genesis.json $CONTAINER3:$HOMEDIR/config/genesis.json

# init step 2
docker exec $CONTAINER3 bash -c "$SCRIPTS/seed1-init-step2.sh seed-3"


# copy genesis.json from seed-3 to seed-0
docker cp $CONTAINER3:$HOMEDIR/config/genesis.json ./temp/genesis.json
docker cp ./temp/genesis.json $CONTAINER0:$HOMEDIR/config/genesis.json

# copy gentx
docker cp $CONTAINER1:$HOMEDIR/config/gentx ./temp
docker cp $CONTAINER2:$HOMEDIR/config/gentx ./temp
docker cp $CONTAINER3:$HOMEDIR/config/gentx ./temp
docker cp ./temp/gentx $CONTAINER0:$HOMEDIR/config 

# init step 2
docker exec $CONTAINER0 bash -c "$SCRIPTS/seed0-init-step2.sh"

# copy genesis.json from seed-0 to seed-1,2,3
docker cp $CONTAINER0:$HOMEDIR/config/genesis.json ./temp/genesis.json
docker cp ./temp/genesis.json $CONTAINER1:$HOMEDIR/config/genesis.json
docker cp ./temp/genesis.json $CONTAINER2:$HOMEDIR/config/genesis.json
docker cp ./temp/genesis.json $CONTAINER3:$HOMEDIR/config/genesis.json

# start
# docker exec -it $CONTAINER0 bash -c "$SCRIPTS/seed-start.sh"
# docker exec -it $CONTAINER1 bash -c "$SCRIPTS/seed-start.sh"
# docker exec -it $CONTAINER2 bash -c "$SCRIPTS/seed-start.sh"
# docker exec -it $CONTAINER3 bash -c "$SCRIPTS/seed-start.sh"

# docker exec -it polard-node0 bash -c "/scripts/seed-start.sh"
# docker exec -it polard-node1 bash -c "/scripts/seed-start.sh"
# docker exec -it polard-node2 bash -c "/scripts/seed-start.sh"
# docker exec -it polard-node3 bash -c "/scripts/seed-start.sh"
