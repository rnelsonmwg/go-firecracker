#!/usr/bin/env bash

GIT_PROVIDER="github.com"
GIT_REPO="bitsnap/go-firecracker"

cd "$(dirname "$0")"

./build.sh > /dev/null

cd ..
mkdir -p .test

docker run -it --rm \
    --name go-firecracker-testing \
    -v "$(pwd)"/.test:/go/src/${GIT_PROVIDER}/${GIT_REPO}/.test \
    go-firecracker/test:latest
