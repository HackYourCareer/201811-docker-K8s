#!/usr/bin/env sh

if [ -z "$1" ]; then
    echo "Missing directory arg!"
    exit 1
fi

export WORK_DIR="$1"
/mazegen.sh "$2" "$3"
#exit 0
