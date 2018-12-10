#!/usr/bin/env bash

set +e

cd "$(dirname "$0")"

GIT_PROVIDER="github.com"
GIT_REPO="bitsnap/go-firecracker"

CLIENT_DOCKERFILE="build/client/Dockerfile.local"
[ "$1" == "ci" ] && CLIENT_DOCKERFILE="build/client/Dockerfile"

docker build -t go-firecracker/dependencies \
	--build-arg BUILD_UID=$(id -u) \
	-f build/dependencies/Dockerfile \
	..

docker build -t go-firecracker/client \
    --build-arg GIT_PROVIDER=${GIT_PROVIDER} \
    --build-arg GIT_REPO=${GIT_REPO} \
    --build-arg BUILD_UID=$(id -u) \
    --no-cache \
    -f ${CLIENT_DOCKERFILE} \
    ..

PIDS=()
docker build -t go-firecracker/lint \
    --build-arg BUILD_UID=$(id -u) \
    -f lint/Dockerfile \
    .. &
PIDS[0]=$!

docker build -t go-firecracker/test \
    --build-arg BUILD_UID=$(id -u) \
    --build-arg GIT_PROVIDER=${GIT_PROVIDER} \
    --build-arg GIT_REPO=${GIT_REPO} \
    -f test/Dockerfile \
    .. &
PIDS[1]=$!

for pid in ${PIDS[*]}; do
    wait ${pid}
done