#!/usr/bin/env bash

set -e
TAG=0.0.4

echo "################################################################################"
echo "# pushing: tsmsap/mazegen-worker:${TAG}"
echo "################################################################################"
docker push tsmsap/mazegen-worker:"${TAG}"

echo "################################################################################"
echo "# pushing tsmsap/mazegen-controller:${TAG}"
echo "################################################################################"
docker push tsmsap/mazegen-controller:"${TAG}"

