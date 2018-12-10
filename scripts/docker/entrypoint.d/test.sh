#!/usr/bin/env bash

cd /go/src/${GIT_PROVIDER}/${GIT_REPO}

dep ensure

[[ ! -d .test ]] && mkdir -p .test
go test -mutexprofile=.test/mutexprofile.out -cpuprofile=.test/cpuprofile.out -memprofile=.test/memprofile.out ./...
go test -cover -race -coverprofile=.test/coverage.out ./...
go test -trace=.test/trace.out ./...

[[ ! -z "$COVERALLS_TOKEN" ]] && goveralls -repotoken=${COVERALLS_TOKEN} -coverprofile=.test/coverage.out
