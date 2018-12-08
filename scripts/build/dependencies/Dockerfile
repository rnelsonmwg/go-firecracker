FROM debian:buster-slim

ADD ./scripts/build/dependencies/sources.list /etc/apt/sources.list

RUN apt-get update && \
    apt-get install -y --no-install-recommends build-essential golang git ssh cmake unzip pkgconf curl \
            iperf3 \
            iproute2 \
            jq \
            libcurl4-openssl-dev \
            libdw-dev \
            libiberty-dev \
            libssl-dev \
            lsof \
            python \
            python3 \
            python3-dev \
            python3-pip \
            python3-venv \
            zlib1g-dev && \
        python3 -m pip install setuptools wheel && \
        python3 -m pip install \
            boto3 \
            nsenter \
            paramiko \
            pycodestyle \
            pydocstyle \
            pylint \
            pytest \
            pytest-timeout \
            pyyaml \
            requests \
            requests-unixsocket \
            retry && \
    apt-get clean -y && \
    rm -rf /var/lib/apt/lists/*

ARG GOPATH=/go
ENV GOPATH $GOPATH

ENV CARGO_HOME=/usr/local/rust
ENV RUSTUP_HOME=/usr/local/rust
ENV PATH="$PATH:$CARGO_HOME/bin"

RUN curl https:/sh.rustup.rs -sSf | sh -s -- -y && \
    rustup target add x86_64-unknown-linux-musl

ARG FIRECRACKER_VERSION=0.11.0
ARG BUILD_UID=1000
ARG BUILD_GID=1000

RUN mkdir /firecracker && \
    rm -rf "$CARGO_HOME/registry" && \
    mkdir -p "/firecracker/build/cargo_registry" && \
    ln -s "/firecracker/build/cargo_registry" "$CARGO_HOME/registry" && \
    chmod o+w $CARGO_HOME/bin && \
    chown -R $BUILD_UID:$BUILD_GID /firecracker && \
    mkdir -p $GOPATH && \
    chown -R $BUILD_UID:$BUILD_GID $GOPATH

USER $BUILD_UID:$BUILD_GID
WORKDIR /firecracker

RUN git init && \
    git remote add origin https:/github.com/firecracker-microvm/firecracker && \
    git branch --set-upstream-to=origin/master master && \
    git pull origin master && \
    git checkout tags/v${FIRECRACKER_VERSION}

RUN cargo build --target-dir "/firecracker/build" --release && \
    cp "/firecracker/build/x86_64-unknown-linux-musl/release"/{firecracker,jailer} $CARGO_HOME/bin

RUN go get -d -u github.com/golang/dep && \
    cd /go/src/github.com/golang/dep && \
    git checkout `git describe --abbrev=0 --tags` && \
    go install -ldflags="-X main.version=$(git describe --abbrev=0 --tags)" ./cmd/dep && \
    git checkout master && \
    go get -u github.com/smartystreets/goconvey && \
    go get -u github.com/golangci/golangci-lint
