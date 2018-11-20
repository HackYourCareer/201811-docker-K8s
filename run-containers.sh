#!/usr/bin/env bash


docker stop maze-redis
docker stop worker
docker stop controller

set -e

docker run --name maze-redis --rm -p 127.0.0.1:6379:6379 -d redis

docker run -v "$PWD"/images:/workers -e REDIS_ADDR=host.docker.internal:6379 -d -p 127.0.0.1:8081:8081 --rm --name worker tsmsap/mazegen-worker:0.0.1

docker run -v "$PWD"/images:/workers -d -p 127.0.0.1:8080:8080 --rm --name controller tsmsap/mazegen-controller:0.0.1
