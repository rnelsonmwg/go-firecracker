FROM go-firecracker/dependencies:latest

ARG GIT_PROVIDER
ARG GIT_REPO

RUN mkdir -p /go/src/$GIT_PROVIDER/$GIT_REPO && \
    cd /go/src/$GIT_PROVIDER/$GIT_REPO && \
    git clone git@$GIT_PROVIDER:$GIT_REPO.git .

WORKDIR /go/src/$GIT_PROVIDER/$GIT_REPO

RUN dep ensure