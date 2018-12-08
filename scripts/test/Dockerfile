FROM go-firecracker/client:latest

ADD scripts/test/test.sh /test.sh
RUN chmod o+x /test.sh

ENTRYPOINT /test.sh