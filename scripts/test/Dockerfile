FROM go-firecracker/client:latest

ARG BUILD_UID=1000

USER root

ADD scripts/test/test.sh /test.sh
RUN chmod o+x /test.sh

USER $BUILD_UID:$BUILD_UID

ENTRYPOINT /test.sh