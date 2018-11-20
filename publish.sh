#!/usr/bin/env bash

set -e
TAG=0.0.4

echo "pushing: tsmsap/mazegen-worker:${TAG}"
docker push tsmsap/mazegen-worker:"${TAG}"

echo "pushing tsmsap/mazegen-controller:${TAG}"
docker push tsmsap/mazegen-controller:"${TAG}"

