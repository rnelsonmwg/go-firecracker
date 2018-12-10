#!/usr/bin/env bash

cd "$(dirname "$0")"

./build.sh > /dev/null

docker run -it --rm \
    --name go-firecracker-linting \
    go-firecracker/lint:latest