#!/usr/bin/env bash

cd "$(dirname "$0")"

[[ ! -f  "./env.coveralls" ]] && echo 'COVERALLS_TOKEN=""' >  ./env.coveralls

source ./env
source ./env.coveralls

docker build -t go-firecracker \
    --build-arg BUILD_UID=${BUILD_UID} \
    --build-arg GIT_PROVIDER=${GIT_PROVIDER} \
    --build-arg GIT_REPO=${GIT_REPO} \
    --build-arg COVERALLS_TOKEN=${COVERALLS_TOKEN} \
	-f docker/Dockerfile \
	..
