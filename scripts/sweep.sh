#!/usr/bin/env bash

ex=$(docker ps -qa --no-trunc --filter "status=exited")
[ ! -z "$ex" ] && docker rm ${ex}

ex=$(docker images --filter "dangling=true" -q --no-trunc)
[ ! -z "$ex" ] && docker rmi ${ex}

ex=$(docker images | grep "go-firecracker" | awk '/ / { print $3 }')
[ "$1" == "all" ] &&  [ ! -z "$ex" ] && docker ${ex}
