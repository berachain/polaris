CONTAINER1="polaris-node0"
CONTAINER2="polaris-node1"

HOMEDIR="/pv/.polard"
SCRIPTS="/scripts"

rm -rf ./temp
mkdir ./temp
mkdir ./temp/seed0
mkdir ./temp/seed1
touch ./temp/genesis.json

#reset pods
# docker exec $CONTAINER1 bash -c "$SCRIPTS/reset.sh"
# docker exec $CONTAINER2 bash -c "$SCRIPTS/reset.sh"

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
