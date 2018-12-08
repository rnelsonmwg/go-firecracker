#!/usr/bin/env bash

cd "$(dirname "$0")"

GIT_PROVIDER="github.com"
GIT_REPO="bitsnap/go-firecracker"

CLIENT_DOCKERFILE="build/client/Dockerfile.local"
[ "$1" == "ci" ] && CLIENT_DOCKERFILE="build/client/Dockerfile"

docker build -t go-firecracker/dependencies \
	--build-arg BUILD_UID=$(id -u) \
	--build-arg BUILD_GID=$(id -g) \
    --security-opt seccomp=unconfined \
	-f build/dependencies/Dockerfile \
	..

docker build -t go-firecracker/client \
    --build-arg GIT_PROVIDER=${GIT_PROVIDER} \
    --build-arg GIT_REPO=${GIT_REPO} \
    -f ${CLIENT_DOCKERFILE} \
    ..

PIDS=()
docker build -t go-firecracker/lint \
    -f lint/Dockerfile \
    .. &
PIDS[0]=$!

docker build -t go-firecracker/test \
    -f test/Dockerfile \
    .. &
PIDS[1]=$!

for pid in ${PIDS[*]}; do
    wait ${pid}
done