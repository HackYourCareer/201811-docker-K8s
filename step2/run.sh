#!/usr/bin/env sh

set -e

if [ -z "${WORK_DIR}" ]; then
    WORK_DIR="./images"
fi

mkdir -p "${WORK_DIR}"
FILE_NAME="${WORK_DIR}"/maze.txt

if [ -n "$2" ]; then
    ./maze "$1" "$2" > "${FILE_NAME}"
else
    ./maze "$1" > "${FILE_NAME}"
fi

cat "${FILE_NAME}"
