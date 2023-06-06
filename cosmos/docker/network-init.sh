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

CONTAINER1="polard-node0"
CONTAINER2="polard-node1"

HOMEDIR="/pv/.polard"
SCRIPTS="/scripts"

rm -rf ./temp
mkdir ./temp
mkdir ./temp/seed0
mkdir ./temp/seed1
touch ./temp/genesis.json

# init step 1 
docker exec $CONTAINER1 bash -c "$SCRIPTS/seed0-init-step1.sh"
docker exec $CONTAINER2 bash -c "$SCRIPTS/seed1-init-step1.sh"

# copy genesis.json from seed-0 to seed-1
docker cp $CONTAINER1:$HOMEDIR/config/genesis.json ./temp/genesis.json
docker cp ./temp/genesis.json $CONTAINER2:$HOMEDIR/config/genesis.json

# init step 2
docker exec $CONTAINER2 bash -c "$SCRIPTS/seed1-init-step2.sh"

# copy genesis.json from seed-1 to seed-0
docker cp $CONTAINER2:$HOMEDIR/config/genesis.json ./temp/genesis.json
docker cp ./temp/genesis.json $CONTAINER1:$HOMEDIR/config/genesis.json

# copy gentx
docker cp $CONTAINER2:$HOMEDIR/config/gentx ./temp/gentx
docker cp ./temp/gentx $CONTAINER1:$HOMEDIR/config 

# init step 2
docker exec $CONTAINER1 bash -c "$SCRIPTS/seed0-init-step2.sh"

# copy genesis.json from seed-0 to seed-1
docker cp $CONTAINER1:$HOMEDIR/config/genesis.json ./temp/genesis.json
docker cp ./temp/genesis.json $CONTAINER2:$HOMEDIR/config/genesis.json

# start
# docker exec $CONTAINER1 bash -c "$SCRIPTS/seed-start.sh"
# docker exec $CONTAINER2 bash -c "$SCRIPTS/seed-start.sh"
