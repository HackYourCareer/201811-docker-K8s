#!/usr/bin/env bash

set -e

TAG=0.0.4

echo "########################################"
echo "# Build base image                     #"
echo "########################################"
pushd step2
docker build -t tomek:latest .
popd

echo "########################################"
echo "# Build mazegen image (maze + convert) #"
echo "########################################"
pushd step3
docker build -t tsmsap/mazegen:latest .
popd

echo "########################################"
echo "# Build worker image                   #"
echo "########################################"
pushd step6
env GOOS=linux ./build_worker.sh
docker build -t tsmsap/mazegen-worker:"${TAG}" .
popd

echo "########################################"
echo "# Build controller image               #"
echo "########################################"
pushd step7
env GOOS=linux ./build_controller.sh
docker build -t tsmsap/mazegen-controller:"${TAG}" .
popd

