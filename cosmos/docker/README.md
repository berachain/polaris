# How to run dual node local network

1. make your changes in code
2. mage cosmos:dockerx base arm64 && mage cosmos:dockerx seed arm64
3. in terminal window 1: cd cosmos/docker && sh ./reset-temp.sh && docker-compose up
4. in terminal window 2: cd cosmos/docker, sh ./network-init.sh
5. in terminal window 2: docker exec -it polard-node0 bash -c ./scripts/seed-start.sh
6. in terminal window 3: docker exec -it polard-node1 bash -c ./scripts/seed-start.sh
7. in terminal window 4: docker exec -it polard-node2 bash -c ./scripts/seed-status.sh
8. in terminal window 5: docker exec -it polard-node3 bash -c ./scripts/seed-status.sh

note: added "-it" in steps 5-8, so that ctrl+c can kill the process
