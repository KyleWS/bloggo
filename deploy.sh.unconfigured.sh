#!/bin/bash
./build.sh
ssh root@206.189.79.199 '
docker network rm apinetwork
docker network create apinetwork
docker rm -f mongo-server
docker run -d --name mongo-server -p 27017:27017 --network apinetwork mongo
docker rm -f motiapp-server
docker pull kylews/motiapp-api
docker run -d \
-p 443:443 \
--network apinetwork \
-e DATABASE_ADDRESS=mongo-server:27017 \
-e ADDR=":443" \
--name motiapp-server \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=/etc/letsencrypt/live/motivationdashboard.app/fullchain.pem \
-e TLSKEY=/etc/letsencrypt/live/motivationdashboard.app/privkey.pem \
kylews/motiapp-api'
