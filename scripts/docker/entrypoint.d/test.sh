#!/usr/bin/env bash

cd /go/src/${GIT_PROVIDER}/${GIT_REPO}

dep ensure

firecracker &
FPID=$!

[[ ! -d .test ]] && mkdir -p .test

[[ ! -f .test/hello-vmlinux.bin ]] && curl -fsSL -o .test/hello-vmlinux.bin https://s3.amazonaws.com/spec.ccfc.min/img/hello/kernel/hello-vmlinux.bin
[[ ! -f .test/hello-rootfs.ext4 ]] && curl -fsSL -o .test/hello-rootfs.ext4 https://s3.amazonaws.com/spec.ccfc.min/img/hello/fsfiles/hello-rootfs.ext4

go test -mutexprofile=.test/mutexprofile.out -cpuprofile=.test/cpuprofile.out -memprofile=.test/memprofile.out ./...
go test -cover -race -coverprofile=.test/coverage.out ./...
go test -trace=.test/trace.out ./...

[[ ! -z "$COVERALLS_TOKEN" ]] && [[ -f .test/coverage.out ]] && \
  goveralls -repotoken=${COVERALLS_TOKEN} -coverprofile=.test/coverage.out

rm -f /tmp/firecracker.socket
kill ${FPID} > /dev/null
