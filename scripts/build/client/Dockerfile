FROM go-firecracker/dependencies:latest

ARG GIT_PROVIDER
ARG GIT_REPO
ARG BUILD_UID=1000

USER $BUILD_UID:$BUILD_UID
WORKDIR /$GOPATH

RUN mkdir -p /go/src/$GIT_PROVIDER/$GIT_REPO && \
    cd /go/src/$GIT_PROVIDER/$GIT_REPO && \
    git clone git@$GIT_PROVIDER:$GIT_REPO.git . && \
    dep ensure

WORKDIR /go/src/$GIT_PROVIDER/$GIT_REPO
