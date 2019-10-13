#! /bin/bash

export TESTGROUND_BASEDIR=$(pwd)

go build
./testground -vv run --builder docker:go --runner local:docker --instances=2 dht/lookup-peers
docker stop testground-redis
docker rm testground-redis
docker network prune -f
sudo ifconfig
