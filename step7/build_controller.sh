#!/usr/bin/env bash

BUILD_PATH="${GOPATH}"/src/mazehyc/web/controller

if [ -n "${GOOS}" ]; then
    env GOOS="${GOOS}" go build -o ./maze-controller "${BUILD_PATH}"/cmd/main.go
else
    go build -o ./maze-controller "${BUILD_PATH}"/cmd/main.go
fi
