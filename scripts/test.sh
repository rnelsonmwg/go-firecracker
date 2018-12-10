#!/usr/bin/env bash

cd "$(dirname "$0")"

[[ ! -f  "./env.coveralls" ]] && echo 'COVERALLS_TOKEN=""' >  ./env.coveralls

source ./env
source ./env.coveralls

cd ..

docker run -it --rm \
    --name go-firecracker-env \
    -v "$(pwd)"/:/go/src/${GIT_PROVIDER}/${GIT_REPO} \
    --device=/dev/kvm:/dev/kvm \
    --privileged \
    --security-opt seccomp=unconfined \
    --ulimit core=0 \
    go-firecracker:latest
