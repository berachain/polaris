# How to run dual node local network

1. make your changes in code
2. mage cosmos:dockerx base arm64
3. mage cosmos:dockerx seed arm64
4. in terminal window 1: cd cosmos/docker docker-compose up
5. in terminal window 2: cd cosmos/docker, sh ./network-init.sh
6. export $
7. in terminal window 2: docker exec polard-node0 bash -c ./scripts/seed-start.sh
8. in terminal window 3: docker exec polard-node1 bash -c ./scripts/seed-start.sh