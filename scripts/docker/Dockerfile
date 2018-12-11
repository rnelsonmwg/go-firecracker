FROM debian:buster-slim

ADD ./scripts/docker/sources.list /etc/apt/sources.list

RUN apt-get update && \
    apt-get install -y --no-install-recommends build-essential golang git ssh cmake unzip pkgconf curl \
            musl \
            musl-dev \
            musl-tools \
            iperf3 \
            iproute2 \
            jq \
            ca-certificates libssl1.1 openssl libcurl4-openssl-dev \
            libdw-dev \
            libiberty-dev \
            libssl-dev \
            lsof \
            zlib1g-dev && \
    apt-get clean -y && \
    rm -rf /var/lib/apt/lists/*

ARG CARGO_HOME=/usr/local/rust
ENV CARGO_HOME $CARGO_HOME
ENV RUSTUP_HOME $CARGO_HOME
ENV PATH="$PATH:$CARGO_HOME/bin"

RUN update-ca-certificates && \
    curl https:/sh.rustup.rs -sSf | sh -s -- -y && \
    rustup target add x86_64-unknown-linux-musl

ARG BUILD_UID=1000
ARG GOPATH=/go
ENV GOPATH $GOPATH
ENV PATH="$PATH:$GOPATH/bin"

RUN addgroup --gid $BUILD_UID user && \
    adduser --system --shell /bin/false --disabled-password \
        --home /firecracker \
        --uid $BUILD_UID --gid $BUILD_UID user && \
    rm -rf "$CARGO_HOME/registry" && \
    mkdir -p "/firecracker/build/cargo_registry" && \
    ln -s "/firecracker/build/cargo_registry" "$CARGO_HOME/registry" && \
    chmod o+w $CARGO_HOME/bin && \
    chown -R $BUILD_UID:$BUILD_UID /firecracker && \
    mkdir -p $GOPATH && \
    chown -R $BUILD_UID:$BUILD_UID $GOPATH

USER $BUILD_UID:$BUILD_UID
WORKDIR /firecracker

ARG FIRECRACKER_VERSION=0.11.0

RUN git init && \
    git remote add origin https://github.com/firecracker-microvm/firecracker.git && \
    git pull origin master && \
    git fetch --all --tags --prune && \
    git checkout tags/v${FIRECRACKER_VERSION} && \
    cargo build --target-dir "/firecracker/build" --release && \
    cp "/firecracker/build/x86_64-unknown-linux-musl/release"/firecracker $CARGO_HOME/bin && \
    cp "/firecracker/build/x86_64-unknown-linux-musl/release"/jailer $CARGO_HOME/bin

RUN go get -d -u github.com/golang/dep && \
    cd /go/src/github.com/golang/dep && \
    git checkout `git describe --abbrev=0 --tags` && \
    go install -ldflags="-X main.version=$(git describe --abbrev=0 --tags)" ./cmd/dep && \
    git checkout master

ARG GOLANGCI_VERSION=1.12.3

RUN go get -u github.com/mattn/goveralls && \
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
        sh -s -- -b $GOPATH/bin v${GOLANGCI_VERSION}

USER root

ADD ./scripts/docker/entrypoint.sh /entrypoint.sh
RUN chmod o+x /entrypoint.sh

USER $BUILD_UID:$BUILD_UID
WORKDIR /

ARG GIT_PROVIDER
ARG GIT_REPO
ARG COVERALLS_TOKEN
ENV GIT_PROVIDER=$GIT_PROVIDER
ENV GIT_REPO=$GIT_REPO
ENV COVERALLS_TOKEN=$COVERALLS_TOKEN

ENTRYPOINT ["/entrypoint.sh"]