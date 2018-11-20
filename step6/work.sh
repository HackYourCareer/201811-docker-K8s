#!/usr/bin/env sh

set -e

if [ -z "$1" ]; then
    echo "Missing directory arg!"
    exit 1
fi

export WORK_DIR="$1"
#Suppress output to make logs visible
/mazegen.sh "$2" "$3" > /dev/null
