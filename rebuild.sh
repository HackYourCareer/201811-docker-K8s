#!/usr/bin/env bash

set -e

TAG=0.0.4
pushd step2
env GOOS=linux ./build.sh
docker build -t tomek:latest .
popd

pushd step3
docker build -t tsmsap/mazegen:latest .
popd

pushd step6
env GOOS=linux ./build_worker.sh
docker build -t tsmsap/mazegen-worker:"${TAG}" .
popd

pushd step7
env GOOS=linux ./build_controller.sh
docker build -t tsmsap/mazegen-controller:"${TAG}" .
popd

