# How to run dual node local network

1. make your changes in code
2. mage cosmos:dockerx base arm64 && mage cosmos:dockerx seed arm64
3. in terminal window 1: cd cosmos/docker && docker-compose up
4. in terminal window 2: cd cosmos/docker, sh ./network-init.sh
5. in terminal window 2: docker exec polard-node0 bash -c ./scripts/seed-start.sh
6. in terminal window 3: docker exec polard-node1 bash -c ./scripts/seed-start.sh

To kill a process:
docker exec polard-node1 bash -c ps
docker exec polard-node1 bash -c "kill -9 $PID"
