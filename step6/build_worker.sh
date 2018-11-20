#!/usr/bin/env bash

BUILD_PATH="${GOPATH}"/src/mazehyc/web/worker

if [ -n "${GOOS}" ]; then
    env GOOS="${GOOS}" go build -o ./maze-worker "${BUILD_PATH}"/cmd/main.go
else
    go build -o ./maze-worker "${BUILD_PATH}"/cmd/main.go
fi
